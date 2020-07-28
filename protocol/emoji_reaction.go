package protocol

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"

	"github.com/status-im/status-go/eth-node/crypto"
	"github.com/status-im/status-go/eth-node/types"
	"github.com/status-im/status-go/protocol/protobuf"
)

// EmojiReaction represents an emoji reaction from a user in the application layer, used for persistence, querying and
// signaling
type EmojiReaction struct {
	protobuf.EmojiReaction

	// From is a public key of the author of the emoji reaction.
	From string `json:"from,omitempty"`

	// SigPubKey is the ecdsa encoded public key of the emoji reaction author
	SigPubKey *ecdsa.PublicKey `json:"-"`

	// LocalChatID is the chatID of the local chat (one-to-one are not symmetric)
	LocalChatID string `json:"localChatId"`
}

// ID is the Keccak256() contatenation of From-MessageID-EmojiType
func (e EmojiReaction) ID() string {
	return types.EncodeHex(crypto.Keccak256([]byte(fmt.Sprintf("%s%s%d", e.From, e.MessageId, e.Type))))
}

// GetSigPubKey returns an ecdsa encoded public key
// this function is required to implement the ChatEntity interface
func (e EmojiReaction) GetSigPubKey() *ecdsa.PublicKey {
	return e.SigPubKey
}

// GetProtoBuf returns the struct's embedded protobuf struct
// this function is required to implement the ChatEntity interface
func (e EmojiReaction) GetProtobuf() proto.Message {
	return &e.EmojiReaction
}

// SetMessageType a setter for the MessageType field
// this function is required to implement the ChatEntity interface
func (e *EmojiReaction) SetMessageType(messageType protobuf.MessageType) {
	e.MessageType = messageType
}

func (e EmojiReaction) MarshalJSON() ([]byte, error) {
	type EmojiAlias EmojiReaction
	item := struct {
		ID          string                      `json:"id"`
		Clock       uint64                      `json:"clock,omitempty"`
		ChatID      string                      `json:"chatId,omitempty"`
		LocalChatID string                      `json:"localChatId,omitempty"`
		From        string                      `json:"from"`
		MessageID   string                      `json:"messageId,omitempty"`
		MessageType protobuf.MessageType        `json:"messageType,omitempty"`
		EmojiID     protobuf.EmojiReaction_Type `json:"emojiId,omitempty"`
	}{

		ID:          e.ID(),
		Clock:       e.Clock,
		ChatID:      e.ChatId,
		LocalChatID: e.LocalChatID,
		From:        e.From,
		MessageID:   e.MessageId,
		MessageType: e.MessageType,
		EmojiID:     e.Type,
	}

	return json.Marshal(item)
}
