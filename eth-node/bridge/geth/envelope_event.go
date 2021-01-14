package gethbridge

import (
	"github.com/status-im/status-go/eth-node/types"
	"github.com/status-im/status-go/waku"
	wakucommon "github.com/status-im/status-go/waku/common"
	"github.com/status-im/status-go/whisper"
)

// NewWhisperEnvelopeEventWrapper returns a types.EnvelopeEvent object that mimics Geth's EnvelopeEvent
func NewWhisperEnvelopeEventWrapper(envelopeEvent *whisper.EnvelopeEvent) *types.EnvelopeEvent {
	if envelopeEvent == nil {
		panic("envelopeEvent should not be nil")
	}

	wrappedData := envelopeEvent.Data
	switch data := envelopeEvent.Data.(type) {
	case []whisper.EnvelopeError:
		wrappedData := make([]types.EnvelopeError, len(data))
		for index := range data {
			wrappedData[index] = *NewWhisperEnvelopeErrorWrapper(&data[index])
		}
	case *whisper.MailServerResponse:
		wrappedData = NewWhisperMailServerResponseWrapper(data)
	case whisper.SyncEventResponse:
		wrappedData = NewGethSyncEventResponseWrapper(data)
	}
	return &types.EnvelopeEvent{
		Event: types.EventType(envelopeEvent.Event),
		Hash:  types.Hash(envelopeEvent.Hash),
		Batch: types.Hash(envelopeEvent.Batch),
		Peer:  types.EnodeID(envelopeEvent.Peer),
		Data:  wrappedData,
	}
}

// NewWakuEnvelopeEventWrapper returns a types.EnvelopeEvent object that mimics Geth's EnvelopeEvent
func NewWakuEnvelopeEventWrapper(envelopeEvent *wakucommon.EnvelopeEvent) *types.EnvelopeEvent {
	if envelopeEvent == nil {
		panic("envelopeEvent should not be nil")
	}

	wrappedData := envelopeEvent.Data
	switch data := envelopeEvent.Data.(type) {
	case []wakucommon.EnvelopeError:
		wrappedData := make([]types.EnvelopeError, len(data))
		for index := range data {
			wrappedData[index] = *NewWakuEnvelopeErrorWrapper(&data[index])
		}
	case *waku.MailServerResponse:
		wrappedData = NewWakuMailServerResponseWrapper(data)
	}
	return &types.EnvelopeEvent{
		Event: types.EventType(envelopeEvent.Event),
		Hash:  types.Hash(envelopeEvent.Hash),
		Batch: types.Hash(envelopeEvent.Batch),
		Peer:  types.EnodeID(envelopeEvent.Peer),
		Data:  wrappedData,
	}
}
