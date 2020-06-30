package protocol

import (
	"crypto/ecdsa"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	gethbridge "github.com/status-im/status-go/eth-node/bridge/geth"
	"github.com/status-im/status-go/eth-node/crypto"
	"github.com/status-im/status-go/eth-node/types"
	"github.com/status-im/status-go/protocol/tt"
	"github.com/status-im/status-go/whisper/v6"
)

func TestMessengerMuteSuite(t *testing.T) {
	suite.Run(t, new(MessengerMuteSuite))
}

type MessengerMuteSuite struct {
	suite.Suite
	m          *Messenger        // main instance of Messenger
	privateKey *ecdsa.PrivateKey // private key for the main instance of Messenger

	// If one wants to send messages between different instances of Messenger,
	// a single Whisper service should be shared.
	shh types.Whisper

	tmpFiles []*os.File // files to clean up
	logger   *zap.Logger
}

func (s *MessengerMuteSuite) SetupTest() {
	s.logger = tt.MustCreateTestLogger()

	config := whisper.DefaultConfig
	config.MinimumAcceptedPOW = 0
	shh := whisper.New(&config)
	s.shh = gethbridge.NewGethWhisperWrapper(shh)
	s.Require().NoError(shh.Start(nil))

	s.m = s.newMessenger(s.shh)
	s.privateKey = s.m.identity
}

func (s *MessengerMuteSuite) newMessengerWithKey(shh types.Whisper, privateKey *ecdsa.PrivateKey) *Messenger {
	tmpFile, err := ioutil.TempFile("", "")
	s.Require().NoError(err)

	options := []Option{
		WithCustomLogger(s.logger),
		WithMessagesPersistenceEnabled(),
		WithDatabaseConfig(tmpFile.Name(), "some-key"),
		WithDatasync(),
	}
	installationID := uuid.New().String()
	m, err := NewMessenger(
		privateKey,
		&testNode{shh: shh},
		installationID,
		options...,
	)
	s.Require().NoError(err)

	err = m.Init()
	s.Require().NoError(err)

	s.tmpFiles = append(s.tmpFiles, tmpFile)

	return m
}

func (s *MessengerMuteSuite) newMessenger(shh types.Whisper) *Messenger {
	privateKey, err := crypto.GenerateKey()
	s.Require().NoError(err)

	return s.newMessengerWithKey(s.shh, privateKey)
}

func (s *MessengerMuteSuite) TestSetMute() {
	key, err := crypto.GenerateKey()
	s.Require().NoError(err)

	theirMessenger := s.newMessengerWithKey(s.shh, key)

	chatID := "status"

	chat := CreatePublicChat(chatID, s.m.transport)

	err = s.m.SaveChat(&chat)
	s.Require().NoError(err)

	err = s.m.Join(chat)
	s.Require().NoError(err)

	err = theirMessenger.SaveChat(&chat)
	s.Require().NoError(err)

	s.Require().NoError(s.m.MuteChat(chatID))

	s.Require().Len(s.m.Chats(), 1)
	s.Require().True(s.m.Chats()[0].Muted)

	s.Require().NoError(s.m.UnmuteChat(chatID))
	s.Require().False(s.m.Chats()[0].Muted)
}
