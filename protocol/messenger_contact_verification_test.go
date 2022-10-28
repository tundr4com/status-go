package protocol

import (
	"context"
	"crypto/ecdsa"
	"testing"

	gethbridge "github.com/status-im/status-go/eth-node/bridge/geth"
	"github.com/status-im/status-go/eth-node/crypto"
	"github.com/status-im/status-go/protocol/common"
	"github.com/status-im/status-go/protocol/requests"
	"github.com/status-im/status-go/protocol/tt"
	"github.com/status-im/status-go/protocol/verification"
	"github.com/status-im/status-go/waku"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/status-im/status-go/eth-node/types"
)

func TestMessengerVerificationRequests(t *testing.T) {
	suite.Run(t, new(MessengerVerificationRequests))
}

type MessengerVerificationRequests struct {
	suite.Suite
	m          *Messenger        // main instance of Messenger
	privateKey *ecdsa.PrivateKey // private key for the main instance of Messenger

	// If one wants to send messages between different instances of Messenger,
	// a single Waku service should be shared.
	shh types.Waku

	logger *zap.Logger
}

func (s *MessengerVerificationRequests) SetupTest() {
	s.logger = tt.MustCreateTestLogger()
	config := waku.DefaultConfig
	config.MinimumAcceptedPoW = 0
	shh := waku.New(&config, s.logger)
	s.shh = gethbridge.NewGethWakuWrapper(shh)
	s.Require().NoError(shh.Start())
	s.m = s.newMessenger(s.shh)
	s.privateKey = s.m.identity
	// We start the messenger in order to receive installations
	_, err := s.m.Start()
	s.Require().NoError(err)
}

func (s *MessengerVerificationRequests) mutualContact(theirMessenger *Messenger) {
	messageText := "hello!"

	contactID := types.EncodeHex(crypto.FromECDSAPub(&theirMessenger.identity.PublicKey))
	request := &requests.SendContactRequest{
		ID:      types.Hex2Bytes(contactID),
		Message: messageText,
	}

	// Send contact request
	resp, err := s.m.SendContactRequest(context.Background(), request)
	s.Require().NoError(err)

	s.Require().NotNil(resp)
	s.Require().Len(resp.Messages(), 1)
	s.Require().Equal(common.ContactRequestStatePending, resp.Messages()[0].ContactRequestState)

	// Make sure it's not returned as coming from us
	contactRequests, _, err := s.m.PendingContactRequests("", 10)
	s.Require().NoError(err)
	s.Require().Len(contactRequests, 0)

	// Make sure contact is added on the sender side
	contacts := s.m.AddedContacts()
	s.Require().Len(contacts, 1)
	s.Require().Equal(ContactRequestStateSent, contacts[0].ContactRequestState)

	// Wait for the message to reach its destination
	resp, err = WaitOnMessengerResponse(
		theirMessenger,
		func(r *MessengerResponse) bool {
			return len(r.Contacts) > 0 && len(r.Messages()) > 0 && len(r.ActivityCenterNotifications()) > 0
		},
		"no messages",
	)

	// Check contact request has been received
	s.Require().NoError(err)

	// Check activity center notification is of the right type
	s.Require().Len(resp.ActivityCenterNotifications(), 1)
	s.Require().Equal(ActivityCenterNotificationTypeContactRequest, resp.ActivityCenterNotifications()[0].Type)
	s.Require().NotNil(resp.ActivityCenterNotifications()[0].Message)
	s.Require().Equal(common.ContactRequestStatePending, resp.ActivityCenterNotifications()[0].Message.ContactRequestState)

	// Check the contact state is correctly set
	s.Require().Len(resp.Contacts, 1)
	s.Require().Equal(ContactRequestStateReceived, resp.Contacts[0].ContactRequestState)

	// Make sure it's the pending contact requests
	contactRequests, _, err = theirMessenger.PendingContactRequests("", 10)
	s.Require().NoError(err)
	s.Require().Len(contactRequests, 1)
	s.Require().Equal(contactRequests[0].ContactRequestState, common.ContactRequestStatePending)

	// Accept contact request, receiver side
	resp, err = theirMessenger.AcceptContactRequest(context.Background(), &requests.AcceptContactRequest{ID: types.Hex2Bytes(contactRequests[0].ID)})
	s.Require().NoError(err)

	// Make sure the message is updated
	s.Require().NotNil(resp)
	s.Require().Len(resp.Messages(), 1)
	s.Require().Equal(resp.Messages()[0].ID, contactRequests[0].ID)
	s.Require().Equal(common.ContactRequestStateAccepted, resp.Messages()[0].ContactRequestState)

	s.Require().Len(resp.ActivityCenterNotifications(), 1)
	s.Require().Equal(resp.ActivityCenterNotifications()[0].ID.String(), contactRequests[0].ID)
	s.Require().NotNil(resp.ActivityCenterNotifications()[0].Message)
	s.Require().Equal(common.ContactRequestStateAccepted, resp.ActivityCenterNotifications()[0].Message.ContactRequestState)

	// Check the contact state is correctly set
	s.Require().Len(resp.Contacts, 1)
	s.Require().Equal(ContactRequestStateMutual, resp.Contacts[0].ContactRequestState)

	// Make sure the sender is added to our contacts
	contacts = theirMessenger.AddedContacts()
	s.Require().Len(contacts, 1)

	// Make sure we consider them a mutual contact, receiver side
	mutualContacts := theirMessenger.MutualContacts()
	s.Require().Len(mutualContacts, 1)

	// Wait for the message to reach its destination
	resp, err = WaitOnMessengerResponse(
		s.m,
		func(r *MessengerResponse) bool {
			return len(r.Contacts) > 0 && len(r.Messages()) > 0 && len(r.ActivityCenterNotifications()) > 0
		},
		"no messages",
	)
	s.Require().NoError(err)

	// Check activity center notification is of the right type
	s.Require().Equal(ActivityCenterNotificationTypeContactRequest, resp.ActivityCenterNotifications()[0].Type)
	s.Require().NotNil(resp.ActivityCenterNotifications()[0].Message)
	s.Require().Equal(common.ContactRequestStateAccepted, resp.ActivityCenterNotifications()[0].Message.ContactRequestState)

	// Make sure the message is updated, sender s2de
	s.Require().NotNil(resp)
	s.Require().Len(resp.Messages(), 1)
	s.Require().Equal(resp.Messages()[0].ID, contactRequests[0].ID)
	s.Require().Equal(common.ContactRequestStateAccepted, resp.Messages()[0].ContactRequestState)

	// Make sure we consider them a mutual contact, sender side
	mutualContacts = s.m.MutualContacts()
	s.Require().Len(mutualContacts, 1)

	// Check the contact state is correctly set
	s.Require().Len(resp.Contacts, 1)
	s.Require().Equal(ContactRequestStateMutual, resp.Contacts[0].ContactRequestState)

}

func (s *MessengerVerificationRequests) TestAcceptVerificationRequests() {
	theirMessenger := s.newMessenger(s.shh)
	_, err := theirMessenger.Start()
	s.Require().NoError(err)

	s.mutualContact(theirMessenger)

	theirPk := types.EncodeHex(crypto.FromECDSAPub(&theirMessenger.identity.PublicKey))
	challenge := "challenge"

	resp, err := s.m.SendContactVerificationRequest(context.Background(), theirPk, challenge)
	s.Require().NoError(err)
	s.Require().Len(resp.VerificationRequests, 1)
	verificationRequestID := resp.VerificationRequests[0].ID

	s.Require().Len(resp.Messages(), 1)
        s.Require().NotEmpty(resp.Messages()[0].OutgoingStatus)
	s.Require().Equal(challenge, resp.Messages()[0].Text)
	s.Require().Equal(common.ContactVerificationStatePending, resp.Messages()[0].ContactVerificationState)

	// Wait for the message to reach its destination
	resp, err = WaitOnMessengerResponse(
		theirMessenger,
		func(r *MessengerResponse) bool {
			return len(r.VerificationRequests) > 0 && len(r.ActivityCenterNotifications()) > 0
		},
		"no messages",
	)
	s.Require().NoError(err)
	s.Require().Len(resp.VerificationRequests, 1)
	s.Require().Equal(resp.VerificationRequests[0].ID, verificationRequestID)
	s.Require().Equal(resp.ActivityCenterNotifications()[0].Type, ActivityCenterNotificationTypeContactVerification)
	s.Require().Equal(resp.ActivityCenterNotifications()[0].ContactVerificationStatus, verification.RequestStatusPENDING)

	s.Require().NotNil(resp.ActivityCenterNotifications()[0].Message)
	s.Require().Equal(challenge, resp.ActivityCenterNotifications()[0].Message.Text)
	s.Require().Equal(common.ContactVerificationStatePending, resp.ActivityCenterNotifications()[0].Message.ContactVerificationState)
	s.Require().Len(resp.Messages(), 1)
        s.Require().Empty(resp.Messages()[0].OutgoingStatus)
	s.Require().Equal(challenge, resp.Messages()[0].Text)
	s.Require().Equal(common.ContactVerificationStatePending, resp.Messages()[0].ContactVerificationState)

	resp, err = theirMessenger.AcceptContactVerificationRequest(context.Background(), verificationRequestID, "hello back")

	s.Require().NoError(err)

	s.Require().NoError(err)
	s.Require().Len(resp.VerificationRequests, 1)
	s.Require().Equal(resp.VerificationRequests[0].ID, verificationRequestID)
	s.Require().Equal(resp.VerificationRequests[0].RequestStatus, verification.RequestStatusACCEPTED)
	s.Require().NotEmpty(resp.VerificationRequests[0].RepliedAt)

	s.Require().Len(resp.ActivityCenterNotifications(), 1)
	s.Require().Equal(resp.ActivityCenterNotifications()[0].ID.String(), verificationRequestID)

	s.Require().Equal(resp.ActivityCenterNotifications()[0].ContactVerificationStatus, verification.RequestStatusACCEPTED)
	s.Require().Equal(common.ContactVerificationStateAccepted, resp.ActivityCenterNotifications()[0].Message.ContactVerificationState)
	s.Require().Len(resp.Messages(), 1)
	s.Require().Equal(common.ContactVerificationStateAccepted, resp.Messages()[0].ContactVerificationState)

	s.Require().NotNil(resp.ActivityCenterNotifications()[0].ReplyMessage)
	s.Require().NotEmpty(resp.ActivityCenterNotifications()[0].ReplyMessage.OutgoingStatus)
	s.Require().Equal("hello back", resp.ActivityCenterNotifications()[0].ReplyMessage.Text)

	// Wait for the message to reach its destination
	resp, err = WaitOnMessengerResponse(
		s.m,
		func(r *MessengerResponse) bool {
			return len(r.VerificationRequests) > 0
		},
		"no messages",
	)
	s.Require().NoError(err)
	s.Require().Len(resp.VerificationRequests, 1)
	s.Require().Equal(resp.VerificationRequests[0].ID, verificationRequestID)

        messages := resp.Messages()
	s.Require().Len(messages, 2)
        var originalMessage *common.Message
        var replyMessage *common.Message

        if messages[0].ID == resp.VerificationRequests[0].ID {
          originalMessage = messages[0]
          replyMessage = messages[1]
        } else {
          originalMessage = messages[1]
          replyMessage = messages[0]
        }

	s.Require().Equal(common.ContactVerificationStateAccepted, originalMessage.ContactVerificationState)
        // Not sure this is the right state for this, it should probably be pendingtrustedverification or something
	s.Require().Equal(common.ContactVerificationStateAccepted, replyMessage.ContactVerificationState)

	s.Require().Len(resp.ActivityCenterNotifications(), 1)
	s.Require().Equal(resp.ActivityCenterNotifications()[0].ID.String(), verificationRequestID)
	s.Require().Equal(resp.ActivityCenterNotifications()[0].ContactVerificationStatus, verification.RequestStatusACCEPTED)
	s.Require().Equal(common.ContactVerificationStateAccepted, resp.ActivityCenterNotifications()[0].Message.ContactVerificationState)

	s.Require().NotNil(resp.ActivityCenterNotifications()[0].ReplyMessage)
	s.Require().Empty(resp.ActivityCenterNotifications()[0].ReplyMessage.OutgoingStatus)
	s.Require().Equal("hello back", resp.ActivityCenterNotifications()[0].ReplyMessage.Text)
}

func (s *MessengerVerificationRequests) TestDeclineVerificationRequests() {
	theirMessenger := s.newMessenger(s.shh)
	_, err := theirMessenger.Start()
	s.Require().NoError(err)

	s.mutualContact(theirMessenger)

	theirPk := types.EncodeHex(crypto.FromECDSAPub(&theirMessenger.identity.PublicKey))
	challenge := "challenge"

	resp, err := s.m.SendContactVerificationRequest(context.Background(), theirPk, challenge)
	s.Require().NoError(err)
	s.Require().Len(resp.VerificationRequests, 1)
	verificationRequestID := resp.VerificationRequests[0].ID

	s.Require().Len(resp.Messages(), 1)
	s.Require().Equal(challenge, resp.Messages()[0].Text)
	s.Require().Equal(common.ContactVerificationStatePending, resp.Messages()[0].ContactVerificationState)

	// Wait for the message to reach its destination
	resp, err = WaitOnMessengerResponse(
		theirMessenger,
		func(r *MessengerResponse) bool {
			return len(r.VerificationRequests) > 0 && len(r.ActivityCenterNotifications()) > 0
		},
		"no messages",
	)
	s.Require().NoError(err)
	s.Require().Len(resp.VerificationRequests, 1)
	s.Require().Equal(resp.VerificationRequests[0].ID, verificationRequestID)
	s.Require().Equal(resp.ActivityCenterNotifications()[0].Type, ActivityCenterNotificationTypeContactVerification)
	s.Require().Equal(resp.ActivityCenterNotifications()[0].ContactVerificationStatus, verification.RequestStatusPENDING)

	s.Require().NotNil(resp.ActivityCenterNotifications()[0].Message)
	s.Require().Equal(challenge, resp.ActivityCenterNotifications()[0].Message.Text)
	s.Require().Equal(common.ContactVerificationStatePending, resp.ActivityCenterNotifications()[0].Message.ContactVerificationState)
	s.Require().Len(resp.Messages(), 1)
	s.Require().Equal(challenge, resp.Messages()[0].Text)
	s.Require().Equal(common.ContactVerificationStatePending, resp.Messages()[0].ContactVerificationState)

	// Make sure it's stored and retrieved correctly
	notifications, err := theirMessenger.UnreadActivityCenterNotifications("", 4, ActivityCenterNotificationTypeContactVerification)
	s.Require().NoError(err)
	s.Require().Len(notifications.Notifications, 1)
	s.Require().Equal(notifications.Notifications[0].ContactVerificationStatus, verification.RequestStatusPENDING)
	s.Require().Equal(common.ContactVerificationStatePending, notifications.Notifications[0].Message.ContactVerificationState)

	resp, err = theirMessenger.DeclineContactVerificationRequest(context.Background(), verificationRequestID)

	s.Require().NoError(err)

	s.Require().NotNil(resp)

	s.Require().Len(resp.VerificationRequests, 1)
	s.Require().Equal(resp.VerificationRequests[0].ID, verificationRequestID)
	s.Require().Equal(resp.VerificationRequests[0].RequestStatus, verification.RequestStatusDECLINED)
	s.Require().NotEmpty(resp.VerificationRequests[0].RepliedAt)

	s.Require().Len(resp.ActivityCenterNotifications(), 1)
	s.Require().Equal(resp.ActivityCenterNotifications()[0].ID.String(), verificationRequestID)

	s.Require().Equal(resp.ActivityCenterNotifications()[0].ContactVerificationStatus, verification.RequestStatusDECLINED)
	s.Require().Equal(common.ContactVerificationStateDeclined, resp.ActivityCenterNotifications()[0].Message.ContactVerificationState)
	s.Require().Len(resp.Messages(), 1)
	s.Require().Equal(common.ContactVerificationStateDeclined, resp.Messages()[0].ContactVerificationState)

	// Make sure it's stored and retrieved correctly
	notifications, err = theirMessenger.UnreadActivityCenterNotifications("", 4, ActivityCenterNotificationTypeContactVerification)
	s.Require().NoError(err)
	s.Require().Len(notifications.Notifications, 1)
	s.Require().Equal(notifications.Notifications[0].ContactVerificationStatus, verification.RequestStatusDECLINED)
	s.Require().Equal(common.ContactVerificationStateDeclined, notifications.Notifications[0].Message.ContactVerificationState)

	// Wait for the message to reach its destination
	resp, err = WaitOnMessengerResponse(
		s.m,
		func(r *MessengerResponse) bool {
			return len(r.VerificationRequests) > 0
		},
		"no messages",
	)
	s.Require().NoError(err)
	s.Require().Len(resp.VerificationRequests, 1)
	s.Require().Equal(resp.VerificationRequests[0].ID, verificationRequestID)

	s.Require().Len(resp.Messages(), 1)
	s.Require().Equal(common.ContactVerificationStateDeclined, resp.Messages()[0].ContactVerificationState)

	s.Require().Len(resp.ActivityCenterNotifications(), 1)
	s.Require().Equal(resp.ActivityCenterNotifications()[0].ID.String(), verificationRequestID)
	s.Require().Equal(resp.ActivityCenterNotifications()[0].ContactVerificationStatus, verification.RequestStatusDECLINED)
	s.Require().Equal(common.ContactVerificationStateDeclined, resp.ActivityCenterNotifications()[0].Message.ContactVerificationState)
}

func (s *MessengerVerificationRequests) TearDownTest() {
	s.Require().NoError(s.m.Shutdown())
}

func (s *MessengerVerificationRequests) newMessenger(shh types.Waku) *Messenger {
	privateKey, err := crypto.GenerateKey()
	s.Require().NoError(err)
	messenger, err := newMessengerWithKey(s.shh, privateKey, s.logger, nil)
	s.Require().NoError(err)
	return messenger
}
