package communities

import (
	"crypto/ecdsa"
	"database/sql"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/status-im/status-go/appdatabase"
	"github.com/status-im/status-go/eth-node/crypto"
	"github.com/status-im/status-go/eth-node/types"
	"github.com/status-im/status-go/protocol/common"
	"github.com/status-im/status-go/protocol/communities/token"
	"github.com/status-im/status-go/protocol/protobuf"
	"github.com/status-im/status-go/protocol/sqlite"
	"github.com/status-im/status-go/services/wallet/bigint"
	"github.com/status-im/status-go/t/helpers"
)

func TestPersistenceSuite(t *testing.T) {
	suite.Run(t, new(PersistenceSuite))
}

type PersistenceSuite struct {
	suite.Suite

	db *Persistence
}

func (s *PersistenceSuite) SetupTest() {
	s.db = nil

	db, err := helpers.SetupTestMemorySQLDB(appdatabase.DbInitializer{})
	s.Require().NoError(err, "creating sqlite db instance")

	err = sqlite.Migrate(db)
	s.Require().NoError(err, "protocol migrate")

	s.db = &Persistence{db: db, timesource: &TimeSourceStub{}}
}

func (s *PersistenceSuite) TestSaveCommunity() {
	id, err := crypto.GenerateKey()
	s.Require().NoError(err)

	// there is one community inserted by default
	communities, err := s.db.AllCommunities(&id.PublicKey, "")
	s.Require().NoError(err)
	s.Require().Len(communities, 1)

	community := Community{
		config: &Config{
			PrivateKey:           id,
			ControlNode:          &id.PublicKey,
			ControlDevice:        true,
			ID:                   &id.PublicKey,
			Joined:               true,
			Spectated:            true,
			Verified:             true,
			Muted:                true,
			MuteTill:             time.Time{},
			CommunityDescription: &protobuf.CommunityDescription{},
		},
	}
	s.Require().NoError(s.db.SaveCommunity(&community))

	communities, err = s.db.AllCommunities(&id.PublicKey, "")
	s.Require().NoError(err)
	s.Require().Len(communities, 2)
	s.Equal(types.HexBytes(crypto.CompressPubkey(&id.PublicKey)), communities[1].ID())
	s.Equal(true, communities[1].Joined())
	s.Equal(true, communities[1].Spectated())
	s.Equal(true, communities[1].Verified())
	s.Equal(true, communities[1].Muted())
	s.Equal(time.Time{}, communities[1].MuteTill())
}

func (s *PersistenceSuite) TestShouldHandleSyncCommunity() {
	sc := &protobuf.SyncInstallationCommunity{
		Id:          []byte("0x123456"),
		Description: []byte("this is a description"),
		Joined:      true,
		Verified:    true,
		Clock:       uint64(time.Now().Unix()),
	}

	// check an empty db to see if a community should be synced
	should, err := s.db.ShouldHandleSyncCommunity(sc)
	s.Require().NoError(err, "SaveSyncCommunity")
	s.True(should)

	// add a new community to the db
	err = s.db.saveRawCommunityRow(fromSyncCommunityProtobuf(sc))
	s.Require().NoError(err, "saveRawCommunityRow")

	rcrs, err := s.db.getAllCommunitiesRaw()
	s.Require().NoError(err, "should have no error from getAllCommunitiesRaw")
	s.Len(rcrs, 2, "length of all communities raw should be 2")

	// check again to see is the community should be synced
	sc.Clock--
	should, err = s.db.ShouldHandleSyncCommunity(sc)
	s.Require().NoError(err, "SaveSyncCommunity")
	s.False(should)

	// check again to see is the community should be synced
	sc.Clock++
	sc.Clock++
	should, err = s.db.ShouldHandleSyncCommunity(sc)
	s.Require().NoError(err, "SaveSyncCommunity")
	s.True(should)
}

func (s *PersistenceSuite) TestSetSyncClock() {
	sc := &protobuf.SyncInstallationCommunity{
		Id:          []byte("0x123456"),
		Description: []byte("this is a description"),
		Joined:      true,
		Verified:    true,
	}

	// add a new community to the db
	err := s.db.saveRawCommunityRow(fromSyncCommunityProtobuf(sc))
	s.Require().NoError(err, "saveRawCommunityRow")

	// retrieve row from db synced_at must be zero
	rcr, err := s.db.getRawCommunityRow(sc.Id)
	s.Require().NoError(err, "getRawCommunityRow")
	s.Require().Zero(rcr.SyncedAt, "synced_at must be zero value")

	// Set the synced_at value
	clock := uint64(time.Now().Unix())
	err = s.db.SetSyncClock(sc.Id, clock)
	s.Require().NoError(err, "SetSyncClock")

	// Retrieve row from db and check clock matches synced_at value
	rcr, err = s.db.getRawCommunityRow(sc.Id)
	s.Require().NoError(err, "getRawCommunityRow")
	s.Require().Equal(clock, rcr.SyncedAt, "synced_at must equal the value of the clock")

	// Set Synced At with an older clock value
	olderClock := clock - uint64(256)
	err = s.db.SetSyncClock(sc.Id, olderClock)
	s.Require().NoError(err, "SetSyncClock")

	// Retrieve row from db and check olderClock matches synced_at value
	rcr, err = s.db.getRawCommunityRow(sc.Id)
	s.Require().NoError(err, "getRawCommunityRow")
	s.Require().NotEqual(olderClock, rcr.SyncedAt, "synced_at must not equal the value of the olderClock value")

	// Set Synced At with a newer clock value
	newerClock := clock + uint64(512)
	err = s.db.SetSyncClock(sc.Id, newerClock)
	s.Require().NoError(err, "SetSyncClock")

	// Retrieve row from db and check olderClock matches synced_at value
	rcr, err = s.db.getRawCommunityRow(sc.Id)
	s.Require().NoError(err, "getRawCommunityRow")
	s.Equal(newerClock, rcr.SyncedAt, "synced_at must equal the value of the newerClock value")
}

func (s *PersistenceSuite) TestSetPrivateKey() {
	sc := &protobuf.SyncInstallationCommunity{
		Id:          []byte("0x123456"),
		Description: []byte("this is a description"),
		Joined:      true,
		Verified:    true,
	}

	// add a new community to the db with no private key
	err := s.db.saveRawCommunityRow(fromSyncCommunityProtobuf(sc))
	s.Require().NoError(err, "saveRawCommunityRow")

	// retrieve row from db, private key must be zero
	rcr, err := s.db.getRawCommunityRow(sc.Id)
	s.Require().NoError(err, "getRawCommunityRow")
	s.Zero(rcr.PrivateKey, "private key must be zero value")

	// Set private key
	pk, err := crypto.GenerateKey()
	s.Require().NoError(err, "crypto.GenerateKey")
	err = s.db.SetPrivateKey(sc.Id, pk)
	s.Require().NoError(err, "SetPrivateKey")

	// retrieve row from db again, private key must match the given key
	rcr, err = s.db.getRawCommunityRow(sc.Id)
	s.Require().NoError(err, "getRawCommunityRow")
	s.Equal(crypto.FromECDSA(pk), rcr.PrivateKey, "private key must match given key")
}

func (s *PersistenceSuite) TestJoinedAndPendingCommunitiesWithRequests() {
	identity, err := crypto.GenerateKey()
	s.Require().NoError(err, "crypto.GenerateKey shouldn't give any error")

	clock := uint64(time.Now().Unix())

	// Add a new community that we have joined
	com := s.makeNewCommunity(identity)
	com.Join()
	sc, err := com.ToSyncInstallationCommunityProtobuf(clock, nil, nil)
	s.Require().NoError(err, "Community.ToSyncInstallationCommunityProtobuf shouldn't give any error")
	err = s.db.saveRawCommunityRow(fromSyncCommunityProtobuf(sc))
	s.Require().NoError(err, "saveRawCommunityRow")

	// Add a new community that we have requested to join, but not yet joined
	com2 := s.makeNewCommunity(identity)
	err = s.db.SaveCommunity(com2)
	s.Require().NoError(err, "SaveCommunity shouldn't give any error")

	rtj := &RequestToJoin{
		ID:          types.HexBytes{1, 2, 3, 4, 5, 6, 7, 8},
		PublicKey:   common.PubkeyToHex(&identity.PublicKey),
		Clock:       clock,
		CommunityID: com2.ID(),
		State:       RequestToJoinStatePending,
	}
	err = s.db.SaveRequestToJoin(rtj)
	s.Require().NoError(err, "SaveRequestToJoin shouldn't give any error")

	comms, err := s.db.JoinedAndPendingCommunitiesWithRequests(&identity.PublicKey, "")
	s.Require().NoError(err, "JoinedAndPendingCommunitiesWithRequests shouldn't give any error")
	s.Len(comms, 2, "Should have 2 communities")

	for _, comm := range comms {
		switch comm.IDString() {
		case com.IDString():
			s.Len(comm.RequestsToJoin(), 0, "Should have no RequestsToJoin")
		case com2.IDString():
			rtjs := comm.RequestsToJoin()
			s.Len(rtjs, 1, "Should have one RequestsToJoin")
			s.Equal(rtjs[0], rtj, "RequestToJoin should match the Request stored in the db")
		}
	}
}

func (s *PersistenceSuite) TestSaveRequestToLeave() {
	rtl := &RequestToLeave{
		ID:          []byte("0x123456"),
		PublicKey:   "0xffffff",
		Clock:       2,
		CommunityID: []byte("0x654321"),
	}

	err := s.db.SaveRequestToLeave(rtl)
	s.Require().NoError(err)

	// older clocks should not be saved
	rtl.Clock = 1
	err = s.db.SaveRequestToLeave(rtl)
	s.Error(err)
}

func (s *PersistenceSuite) makeNewCommunity(identity *ecdsa.PrivateKey) *Community {
	comPrivKey, err := crypto.GenerateKey()
	s.Require().NoError(err, "crypto.GenerateKey shouldn't give any error")

	com, err := New(Config{
		MemberIdentity: &identity.PublicKey,
		PrivateKey:     comPrivKey,
		ControlNode:    &comPrivKey.PublicKey,
		ControlDevice:  true,
		ID:             &comPrivKey.PublicKey,
	}, &TimeSourceStub{})
	s.NoError(err, "New shouldn't give any error")

	md, err := com.MarshaledDescription()
	s.Require().NoError(err, "Community.MarshaledDescription shouldn't give any error")
	com.config.CommunityDescriptionProtocolMessage = md

	return com
}

func (s *PersistenceSuite) TestGetSyncedRawCommunity() {
	sc := &protobuf.SyncInstallationCommunity{
		Id:          []byte("0x123456"),
		Description: []byte("this is a description"),
		Joined:      true,
		Verified:    true,
		Spectated:   true,
	}

	// add a new community to the db
	err := s.db.saveRawCommunityRowWithoutSyncedAt(fromSyncCommunityProtobuf(sc))
	s.Require().NoError(err, "saveRawCommunityRow")

	// retrieve row from db synced_at must be zero
	rcr, err := s.db.getRawCommunityRow(sc.Id)
	s.Require().NoError(err, "getRawCommunityRow")
	s.Zero(rcr.SyncedAt, "synced_at must be zero value")

	// retrieve synced row from db, should fail
	src, err := s.db.getSyncedRawCommunity(sc.Id)
	s.EqualError(err, sql.ErrNoRows.Error())
	s.Nil(src)

	// Set the synced_at value
	clock := uint64(time.Now().Unix())
	err = s.db.SetSyncClock(sc.Id, clock)
	s.Require().NoError(err, "SetSyncClock")

	// retrieve row from db synced_at must not be zero
	rcr, err = s.db.getRawCommunityRow(sc.Id)
	s.Require().NoError(err, "getRawCommunityRow")
	s.NotZero(rcr.SyncedAt, "synced_at must be zero value")

	// retrieve synced row from db, should succeed
	src, err = s.db.getSyncedRawCommunity(sc.Id)
	s.Require().NoError(err)
	s.NotNil(src)
	s.Equal(clock, src.SyncedAt)
}

func (s *PersistenceSuite) TestGetCommunitiesSettings() {
	settings := []CommunitySettings{
		{CommunityID: "0x01", HistoryArchiveSupportEnabled: false},
		{CommunityID: "0x02", HistoryArchiveSupportEnabled: true},
		{CommunityID: "0x03", HistoryArchiveSupportEnabled: false},
	}

	for i := range settings {
		stg := settings[i]
		err := s.db.SaveCommunitySettings(stg)
		s.Require().NoError(err)
	}

	rst, err := s.db.GetCommunitiesSettings()
	s.Require().NoError(err)
	s.Equal(settings, rst)
}

func (s *PersistenceSuite) TestSaveCommunitySettings() {
	settings := CommunitySettings{CommunityID: "0x01", HistoryArchiveSupportEnabled: false}
	err := s.db.SaveCommunitySettings(settings)
	s.Require().NoError(err)
	rst, err := s.db.GetCommunitiesSettings()
	s.Require().NoError(err)
	s.Equal(1, len(rst))
}

func (s *PersistenceSuite) TestDeleteCommunitySettings() {
	settings := CommunitySettings{CommunityID: "0x01", HistoryArchiveSupportEnabled: false}

	err := s.db.SaveCommunitySettings(settings)
	s.Require().NoError(err)

	rst, err := s.db.GetCommunitiesSettings()
	s.Require().NoError(err)
	s.Equal(1, len(rst))
	s.Require().NoError(s.db.DeleteCommunitySettings(types.HexBytes{0x01}))
	rst2, err := s.db.GetCommunitiesSettings()
	s.Require().NoError(err)
	s.Equal(0, len(rst2))
}

func (s *PersistenceSuite) TestUpdateCommunitySettings() {
	settings := []CommunitySettings{
		{CommunityID: "0x01", HistoryArchiveSupportEnabled: true},
		{CommunityID: "0x02", HistoryArchiveSupportEnabled: false},
	}

	s.Require().NoError(s.db.SaveCommunitySettings(settings[0]))
	s.Require().NoError(s.db.SaveCommunitySettings(settings[1]))

	settings[0].HistoryArchiveSupportEnabled = true
	settings[1].HistoryArchiveSupportEnabled = false

	s.Require().NoError(s.db.UpdateCommunitySettings(settings[0]))
	s.Require().NoError(s.db.UpdateCommunitySettings(settings[1]))

	rst, err := s.db.GetCommunitiesSettings()
	s.Require().NoError(err)
	s.Equal(settings, rst)
}

func (s *PersistenceSuite) TestGetCommunityToken() {
	tokens, err := s.db.GetCommunityTokens("123")
	s.Require().NoError(err)
	s.Require().Len(tokens, 0)

	tokenERC721 := token.CommunityToken{
		CommunityID:        "123",
		TokenType:          protobuf.CommunityTokenType_ERC721,
		Address:            "0x123",
		Name:               "StatusToken",
		Symbol:             "STT",
		Description:        "desc",
		Supply:             &bigint.BigInt{Int: big.NewInt(123)},
		InfiniteSupply:     false,
		Transferable:       true,
		RemoteSelfDestruct: true,
		ChainID:            1,
		DeployState:        token.InProgress,
		Base64Image:        "ABCD",
	}

	err = s.db.AddCommunityToken(&tokenERC721)
	s.Require().NoError(err)

	token, err := s.db.GetCommunityToken("123", 1, "0x123")
	s.Require().NoError(err)
	s.Require().Equal(&tokenERC721, token)
}

func (s *PersistenceSuite) TestGetCommunityTokens() {
	tokens, err := s.db.GetCommunityTokens("123")
	s.Require().NoError(err)
	s.Require().Len(tokens, 0)

	tokenERC721 := token.CommunityToken{
		CommunityID:        "123",
		TokenType:          protobuf.CommunityTokenType_ERC721,
		Address:            "0x123",
		Name:               "StatusToken",
		Symbol:             "STT",
		Description:        "desc",
		Supply:             &bigint.BigInt{Int: big.NewInt(123)},
		InfiniteSupply:     false,
		Transferable:       true,
		RemoteSelfDestruct: true,
		ChainID:            1,
		DeployState:        token.InProgress,
		Base64Image:        "ABCD",
		Deployer:           "0xDep1",
		PrivilegesLevel:    token.OwnerLevel,
	}

	tokenERC20 := token.CommunityToken{
		CommunityID:        "345",
		TokenType:          protobuf.CommunityTokenType_ERC20,
		Address:            "0x345",
		Name:               "StatusToken",
		Symbol:             "STT",
		Description:        "desc",
		Supply:             &bigint.BigInt{Int: big.NewInt(345)},
		InfiniteSupply:     false,
		Transferable:       true,
		RemoteSelfDestruct: true,
		ChainID:            2,
		DeployState:        token.Failed,
		Base64Image:        "QWERTY",
		Decimals:           21,
		Deployer:           "0xDep2",
		PrivilegesLevel:    token.CommunityLevel,
	}

	err = s.db.AddCommunityToken(&tokenERC721)
	s.Require().NoError(err)
	err = s.db.AddCommunityToken(&tokenERC20)
	s.Require().NoError(err)

	tokens, err = s.db.GetCommunityTokens("123")
	s.Require().NoError(err)
	s.Require().Len(tokens, 1)
	s.Require().Equal(tokenERC721, *tokens[0])

	err = s.db.UpdateCommunityTokenState(1, "0x123", token.Deployed)
	s.Require().NoError(err)
	tokens, err = s.db.GetCommunityTokens("123")
	s.Require().NoError(err)
	s.Require().Len(tokens, 1)
	s.Require().Equal(token.Deployed, tokens[0].DeployState)

	tokens, err = s.db.GetCommunityTokens("345")
	s.Require().NoError(err)
	s.Require().Len(tokens, 1)
	s.Require().Equal(tokenERC20, *tokens[0])

	err = s.db.UpdateCommunityTokenAddress(1, "0x123", "0x123-newAddr")
	s.Require().NoError(err)
	tokens, err = s.db.GetCommunityTokens("123")
	s.Require().NoError(err)
	s.Require().Len(tokens, 1)
	s.Require().Equal("0x123-newAddr", tokens[0].Address)
}

func (s *PersistenceSuite) TestSaveCheckChannelPermissionResponse() {

	viewAndPostPermissionResults := make(map[string]*PermissionTokenCriteriaResult)
	viewAndPostPermissionResults["one"] = &PermissionTokenCriteriaResult{
		Criteria: []bool{true, true, true, true},
	}
	viewAndPostPermissionResults["two"] = &PermissionTokenCriteriaResult{
		Criteria: []bool{false},
	}
	chatID := "some-chat-id"
	communityID := "some-community-id"

	checkChannelPermissionResponse := &CheckChannelPermissionsResponse{
		ViewOnlyPermissions: &CheckChannelViewOnlyPermissionsResult{
			Satisfied:   true,
			Permissions: make(map[string]*PermissionTokenCriteriaResult),
		},
		ViewAndPostPermissions: &CheckChannelViewAndPostPermissionsResult{
			Satisfied:   true,
			Permissions: viewAndPostPermissionResults,
		},
	}

	err := s.db.SaveCheckChannelPermissionResponse(communityID, chatID, checkChannelPermissionResponse)
	s.Require().NoError(err)

	responses, err := s.db.GetCheckChannelPermissionResponses(communityID)
	s.Require().NoError(err)
	s.Require().Len(responses, 1)
	s.Require().NotNil(responses[chatID])
	s.Require().True(responses[chatID].ViewOnlyPermissions.Satisfied)
	s.Require().Len(responses[chatID].ViewOnlyPermissions.Permissions, 0)
	s.Require().True(responses[chatID].ViewAndPostPermissions.Satisfied)
	s.Require().Len(responses[chatID].ViewAndPostPermissions.Permissions, 2)
	s.Require().Equal(responses[chatID].ViewAndPostPermissions.Permissions["one"].Criteria, []bool{true, true, true, true})
	s.Require().Equal(responses[chatID].ViewAndPostPermissions.Permissions["two"].Criteria, []bool{false})
}

func (s *PersistenceSuite) TestGetCommunityRequestsToJoinWithRevealedAddresses() {
	identity, err := crypto.GenerateKey()
	s.Require().NoError(err, "crypto.GenerateKey shouldn't give any error")

	clock := uint64(time.Now().Unix())
	communityID := types.HexBytes{7, 7, 7, 7, 7, 7, 7, 7}
	revealedAddresses := []string{"address1", "address2", "address3"}
	chainIds := []uint64{1, 2}

	// No data in database
	rtjResult, err := s.db.GetCommunityRequestsToJoinWithRevealedAddresses(communityID)
	s.Require().NoError(err, "GetCommunityRequestsToJoinWithRevealedAddresses shouldn't give any error")
	s.Require().Len(rtjResult, 0)

	// RTJ with 2 revealed Addresses
	expectedRtj1 := &RequestToJoin{
		ID:          types.HexBytes{1, 2, 3, 4, 5, 6, 7, 8},
		PublicKey:   common.PubkeyToHex(&identity.PublicKey),
		Clock:       clock,
		CommunityID: communityID,
		State:       RequestToJoinStateAccepted,
		RevealedAccounts: []*protobuf.RevealedAccount{
			{
				Address: revealedAddresses[0],
			},
			{
				Address: revealedAddresses[1],
			},
		},
	}
	err = s.db.SaveRequestToJoin(expectedRtj1)
	s.Require().NoError(err, "SaveRequestToJoin shouldn't give any error")

	err = s.db.SaveRequestToJoinRevealedAddresses(expectedRtj1.ID, expectedRtj1.RevealedAccounts)
	s.Require().NoError(err, "SaveRequestToJoinRevealedAddresses shouldn't give any error")

	rtjResult, err = s.db.GetCommunityRequestsToJoinWithRevealedAddresses(communityID)
	s.Require().NoError(err, "GetCommunityRequestsToJoinWithRevealedAddresses shouldn't give any error")
	s.Require().Len(rtjResult, 1)
	s.Require().Equal(expectedRtj1.ID, rtjResult[0].ID)
	s.Require().Equal(expectedRtj1.PublicKey, rtjResult[0].PublicKey)
	s.Require().Equal(expectedRtj1.Clock, rtjResult[0].Clock)
	s.Require().Equal(expectedRtj1.CommunityID, rtjResult[0].CommunityID)
	s.Require().Len(rtjResult[0].RevealedAccounts, 2)

	for index, account := range rtjResult[0].RevealedAccounts {
		s.Require().Equal(revealedAddresses[index], account.Address)
	}

	// RTJ with 1 revealed Address, ChainIds, IsAirdropAddress and Signature
	signature := []byte("test")
	expectedRtj2 := &RequestToJoin{
		ID:          types.HexBytes{8, 7, 6, 5, 4, 3, 2, 1},
		PublicKey:   common.PubkeyToHex(&identity.PublicKey),
		Clock:       clock,
		CommunityID: communityID,
		State:       RequestToJoinStateAccepted,
		RevealedAccounts: []*protobuf.RevealedAccount{
			{
				Address:          revealedAddresses[2],
				ChainIds:         chainIds,
				IsAirdropAddress: true,
				Signature:        signature,
			},
		},
	}
	err = s.db.SaveRequestToJoin(expectedRtj2)
	s.Require().NoError(err, "SaveRequestToJoin shouldn't give any error")

	err = s.db.SaveRequestToJoinRevealedAddresses(expectedRtj2.ID, expectedRtj2.RevealedAccounts)
	s.Require().NoError(err, "SaveRequestToJoinRevealedAddresses shouldn't give any error")

	rtjResult, err = s.db.GetCommunityRequestsToJoinWithRevealedAddresses(communityID)
	s.Require().NoError(err, "GetCommunityRequestsToJoinWithRevealedAddresses shouldn't give any error")
	s.Require().Len(rtjResult, 2)

	s.Require().Len(rtjResult[1].RevealedAccounts, 1)
	s.Require().Equal(revealedAddresses[2], rtjResult[1].RevealedAccounts[0].Address)
	s.Require().Equal(chainIds, rtjResult[1].RevealedAccounts[0].ChainIds)
	s.Require().Equal(true, rtjResult[1].RevealedAccounts[0].IsAirdropAddress)
	s.Require().Equal(rtjResult[1].RevealedAccounts[0].Signature, signature)

	// RTJ without RevealedAccounts
	expectedRtjWithoutRevealedAccounts := &RequestToJoin{
		ID:          types.HexBytes{1, 6, 6, 6, 6, 6, 6, 6},
		PublicKey:   common.PubkeyToHex(&identity.PublicKey),
		Clock:       clock,
		CommunityID: communityID,
		State:       RequestToJoinStateAccepted,
	}
	err = s.db.SaveRequestToJoin(expectedRtjWithoutRevealedAccounts)
	s.Require().NoError(err, "SaveRequestToJoin shouldn't give any error")

	rtjResult, err = s.db.GetCommunityRequestsToJoinWithRevealedAddresses(communityID)
	s.Require().NoError(err, "GetCommunityRequestsToJoinWithRevealedAddresses shouldn't give any error")
	s.Require().Len(rtjResult, 3)

	s.Require().Len(rtjResult[2].RevealedAccounts, 0)

	// RTJ with RevealedAccount but with empty Address
	expectedRtjWithEmptyAddress := &RequestToJoin{
		ID:          types.HexBytes{2, 6, 6, 6, 6, 6, 6, 6},
		PublicKey:   common.PubkeyToHex(&identity.PublicKey),
		Clock:       clock,
		CommunityID: communityID,
		State:       RequestToJoinStateAccepted,
		RevealedAccounts: []*protobuf.RevealedAccount{
			{
				Address: "",
			},
		},
	}
	err = s.db.SaveRequestToJoin(expectedRtjWithEmptyAddress)
	s.Require().NoError(err, "SaveRequestToJoin shouldn't give any error")

	rtjResult, err = s.db.GetCommunityRequestsToJoinWithRevealedAddresses(communityID)
	s.Require().NoError(err, "GetCommunityRequestsToJoinWithRevealedAddresses shouldn't give any error")
	s.Require().Len(rtjResult, 4)
	s.Require().Len(rtjResult[3].RevealedAccounts, 0)
}

func (s *PersistenceSuite) TestCuratedCommunities() {
	communities, err := s.db.GetCuratedCommunities()
	s.Require().NoError(err)
	s.Require().Empty(communities.ContractCommunities)
	s.Require().Empty(communities.ContractFeaturedCommunities)

	setCommunities := &CuratedCommunities{
		ContractCommunities:         []string{"x", "d"},
		ContractFeaturedCommunities: []string{"x"},
	}

	err = s.db.SetCuratedCommunities(setCommunities)
	s.Require().NoError(err)

	communities, err = s.db.GetCuratedCommunities()
	s.Require().NoError(err)
	s.Require().True(reflect.DeepEqual(communities, setCommunities))

	setCommunities = &CuratedCommunities{
		ContractCommunities:         []string{"p", "a", "t", "r", "y", "k"},
		ContractFeaturedCommunities: []string{"p", "k"},
	}

	err = s.db.SetCuratedCommunities(setCommunities)
	s.Require().NoError(err)

	communities, err = s.db.GetCuratedCommunities()
	s.Require().NoError(err)
	s.Require().True(reflect.DeepEqual(communities, setCommunities))
}
