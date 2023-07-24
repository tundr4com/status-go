package activity

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"

	"github.com/status-im/status-go/services/wallet/async"
	w_common "github.com/status-im/status-go/services/wallet/common"
	"github.com/status-im/status-go/services/wallet/token"
	"github.com/status-im/status-go/services/wallet/walletevent"
)

const (
	// FilterResponse json is sent as a message in the EventActivityFilteringDone event
	EventActivityFilteringDone          walletevent.EventType = "wallet-activity-filtering-done"
	EventActivityGetRecipientsDone      walletevent.EventType = "wallet-activity-get-recipients-result"
	EventActivityGetOldestTimestampDone walletevent.EventType = "wallet-activity-get-oldest-timestamp-result"
)

const (
	filterTaskID                 = int32(1)
	filterTaskPolicy             = async.ReplacementPolicyCancelOld
	getRecipientsTaskID          = int32(2)
	getRecipientsTaskPolicy      = async.ReplacementPolicyIgnoreNew
	getOldestTimestampTaskID     = int32(3)
	getOldestTimestampTaskPolicy = async.ReplacementPolicyCancelOld
)

func makeTaskType(requestID int32, originalID int32, policy async.ReplacementPolicy) async.TaskType {
	return async.TaskType{
		ID:     int64(requestID)<<32 | int64(originalID),
		Policy: policy,
	}
}

type Service struct {
	db           *sql.DB
	tokenManager *token.Manager
	eventFeed    *event.Feed

	scheduler *async.Scheduler
}

func NewService(db *sql.DB, tokenManager *token.Manager, eventFeed *event.Feed) *Service {
	return &Service{
		db:           db,
		tokenManager: tokenManager,
		eventFeed:    eventFeed,
		scheduler:    async.NewScheduler(),
	}
}

type ErrorCode = int

const (
	ErrorCodeSuccess ErrorCode = iota + 1
	ErrorCodeTaskCanceled
	ErrorCodeFailed
)

type FilterResponse struct {
	Activities []Entry `json:"activities"`
	Offset     int     `json:"offset"`
	// Used to indicate that there might be more entries that were not returned
	// based on a simple heuristic
	HasMore   bool      `json:"hasMore"`
	ErrorCode ErrorCode `json:"errorCode"`
}

// FilterActivityAsync allows only one filter task to run at a time
// and it cancels the current one if a new one is started
// All calls will trigger an EventActivityFilteringDone event with the result of the filtering
func (s *Service) FilterActivityAsync(requestID int32, addresses []common.Address, chainIDs []w_common.ChainID, filter Filter, offset int, limit int) {
	s.scheduler.Enqueue(makeTaskType(requestID, filterTaskID, filterTaskPolicy), func(ctx context.Context) (interface{}, error) {
		activities, err := getActivityEntries(ctx, s.getDeps(), addresses, chainIDs, filter, offset, limit)
		return activities, err
	}, func(result interface{}, taskType async.TaskType, err error) {
		res := FilterResponse{
			ErrorCode: ErrorCodeFailed,
		}

		if errors.Is(err, context.Canceled) || errors.Is(err, async.ErrTaskOverwritten) {
			res.ErrorCode = ErrorCodeTaskCanceled
		} else if err == nil {
			activities := result.([]Entry)
			res.Activities = activities
			res.Offset = offset
			res.HasMore = len(activities) == limit
			res.ErrorCode = ErrorCodeSuccess
		}

		s.sendResponseEvent(&requestID, EventActivityFilteringDone, res, err)
	})
}

type GetRecipientsResponse struct {
	Addresses []common.Address `json:"addresses"`
	Offset    int              `json:"offset"`
	// Used to indicate that there might be more entries that were not returned
	// based on a simple heuristic
	HasMore   bool      `json:"hasMore"`
	ErrorCode ErrorCode `json:"errorCode"`
}

// GetRecipientsAsync returns true if a task is already running or scheduled due to a previous call; meaning that
// this call won't receive an answer but client should rely on the answer from the previous call.
// If no task is already scheduled false will be returned
func (s *Service) GetRecipientsAsync(requestID int32, offset int, limit int) bool {
	return s.scheduler.Enqueue(makeTaskType(requestID, getRecipientsTaskID, getRecipientsTaskPolicy), func(ctx context.Context) (interface{}, error) {
		var err error
		result := &GetRecipientsResponse{
			Offset:    offset,
			ErrorCode: ErrorCodeSuccess,
		}
		result.Addresses, result.HasMore, err = GetRecipients(ctx, s.db, offset, limit)
		return result, err
	}, func(result interface{}, taskType async.TaskType, err error) {
		res := result.(*GetRecipientsResponse)
		if errors.Is(err, context.Canceled) || errors.Is(err, async.ErrTaskOverwritten) {
			res.ErrorCode = ErrorCodeTaskCanceled
		} else if err != nil {
			res.ErrorCode = ErrorCodeFailed
		}

		s.sendResponseEvent(&requestID, EventActivityGetRecipientsDone, result, err)
	})
}

type GetOldestTimestampResponse struct {
	Timestamp int64     `json:"timestamp"`
	ErrorCode ErrorCode `json:"errorCode"`
}

func (s *Service) GetOldestTimestampAsync(requestID int32, addresses []common.Address) {
	s.scheduler.Enqueue(makeTaskType(requestID, getOldestTimestampTaskID, getOldestTimestampTaskPolicy), func(ctx context.Context) (interface{}, error) {
		timestamp, err := GetOldestTimestamp(ctx, s.db, addresses)
		return timestamp, err
	}, func(result interface{}, taskType async.TaskType, err error) {
		res := GetOldestTimestampResponse{
			ErrorCode: ErrorCodeFailed,
		}

		if errors.Is(err, context.Canceled) || errors.Is(err, async.ErrTaskOverwritten) {
			res.ErrorCode = ErrorCodeTaskCanceled
		} else if err == nil {
			res.Timestamp = result.(int64)
			res.ErrorCode = ErrorCodeSuccess
		}

		s.sendResponseEvent(&requestID, EventActivityGetOldestTimestampDone, res, err)
	})
}

func (s *Service) Stop() {
	s.scheduler.Stop()
}

func (s *Service) getDeps() FilterDependencies {
	return FilterDependencies{
		db: s.db,
		tokenSymbol: func(t Token) string {
			info := s.tokenManager.LookupTokenIdentity(uint64(t.ChainID), t.Address, t.TokenType == Native)
			if info == nil {
				return ""
			}
			return info.Symbol
		},
		tokenFromSymbol: func(chainID *w_common.ChainID, symbol string) *Token {
			var cID *uint64
			if chainID != nil {
				cID = new(uint64)
				*cID = uint64(*chainID)
			}
			t, detectedNative := s.tokenManager.LookupToken(cID, symbol)
			if t == nil {
				return nil
			}
			tokenType := Native
			if !detectedNative {
				tokenType = Erc20
			}
			return &Token{
				TokenType: tokenType,
				ChainID:   w_common.ChainID(t.ChainID),
				Address:   t.Address,
			}
		},
	}
}

func (s *Service) sendResponseEvent(requestID *int32, eventType walletevent.EventType, payloadObj interface{}, resErr error) {
	payload, err := json.Marshal(payloadObj)
	if err != nil {
		log.Error("Error marshaling response: %v; result error: %w", err, resErr)
	} else {
		err = resErr
	}

	log.Debug("wallet.api.activity.Service RESPONSE", "requestID", requestID, "eventType", eventType, "error", err, "payload.len", len(payload))

	event := walletevent.Event{
		Type:    eventType,
		Message: string(payload),
	}

	if requestID != nil {
		event.RequestID = new(int)
		*event.RequestID = int(*requestID)
	}

	s.eventFeed.Send(event)
}
