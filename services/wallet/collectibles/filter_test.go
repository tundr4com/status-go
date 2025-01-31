package collectibles

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/status-im/status-go/protocol/communities/token"
	w_common "github.com/status-im/status-go/services/wallet/common"
	"github.com/status-im/status-go/services/wallet/thirdparty"
	"github.com/status-im/status-go/t/helpers"
	"github.com/status-im/status-go/walletdatabase"

	"github.com/stretchr/testify/require"
)

func setupTestFilterDB(t *testing.T) (db *sql.DB, close func()) {
	db, err := helpers.SetupTestMemorySQLDB(walletdatabase.DbInitializer{})
	require.NoError(t, err)

	return db, func() {
		require.NoError(t, db.Close())
	}
}

func TestFilterOwnedCollectibles(t *testing.T) {
	db, close := setupTestFilterDB(t)
	defer close()

	oDB := NewOwnershipDB(db)
	cDB := NewCollectibleDataDB(db)

	const nData = 50
	data := generateTestCollectiblesData(nData)
	communityData := generateTestCommunityData(nData)

	ownerAddresses := []common.Address{
		common.HexToAddress("0x1234"),
		common.HexToAddress("0x5678"),
		common.HexToAddress("0xABCD"),
	}
	randomAddress := common.HexToAddress("0xFFFF")

	dataPerID := make(map[string]thirdparty.CollectibleData)
	communityDataPerID := make(map[string]thirdparty.CollectibleCommunityInfo)
	idsPerChainIDAndOwner := make(map[w_common.ChainID]map[common.Address][]thirdparty.CollectibleUniqueID)

	var err error

	for i := 0; i < nData; i++ {
		dataPerID[data[i].ID.HashKey()] = data[i]
		communityDataPerID[data[i].ID.HashKey()] = communityData[i]

		chainID := data[i].ID.ContractID.ChainID
		ownerAddress := ownerAddresses[i%len(ownerAddresses)]

		if _, ok := idsPerChainIDAndOwner[chainID]; !ok {
			idsPerChainIDAndOwner[chainID] = make(map[common.Address][]thirdparty.CollectibleUniqueID)
		}
		if _, ok := idsPerChainIDAndOwner[chainID][ownerAddress]; !ok {
			idsPerChainIDAndOwner[chainID][ownerAddress] = make([]thirdparty.CollectibleUniqueID, 0, len(data))
		}

		idsPerChainIDAndOwner[chainID][ownerAddress] = append(idsPerChainIDAndOwner[chainID][ownerAddress], data[i].ID)
	}

	timestamp := int64(1234567890)

	for chainID, idsPerOwner := range idsPerChainIDAndOwner {
		for ownerAddress, ids := range idsPerOwner {
			err = oDB.Update(chainID, ownerAddress, ids, timestamp)
			require.NoError(t, err)
		}
	}

	err = cDB.SetData(data)
	require.NoError(t, err)
	for i := 0; i < nData; i++ {
		err = cDB.SetCommunityInfo(data[i].ID, communityData[i])
		require.NoError(t, err)
	}

	var filter Filter
	var filterIDs []thirdparty.CollectibleUniqueID
	var expectedIDs []thirdparty.CollectibleUniqueID
	var tmpIDs []thirdparty.CollectibleUniqueID

	ctx := context.Background()

	filterChains := []w_common.ChainID{w_common.ChainID(1), w_common.ChainID(2)}
	filterAddresses := []common.Address{ownerAddresses[0], ownerAddresses[1], ownerAddresses[2], randomAddress}

	// Test common case
	filter = allFilter()

	tmpIDs, err = oDB.GetOwnedCollectibles(filterChains, filterAddresses, 0, nData)
	require.NoError(t, err)

	expectedIDs = tmpIDs

	filterIDs, err = filterOwnedCollectibles(ctx, db, filterChains, filterAddresses, filter, 0, nData)
	require.NoError(t, err)
	require.Equal(t, expectedIDs, filterIDs)

	// Test only non-community
	filter = allFilter()
	filter.FilterCommunity = OnlyNonCommunity

	tmpIDs, err = oDB.GetOwnedCollectibles(filterChains, filterAddresses, 0, nData)
	require.NoError(t, err)

	expectedIDs = nil
	for _, id := range tmpIDs {
		if dataPerID[id.HashKey()].CommunityID == "" {
			expectedIDs = append(expectedIDs, id)
		}
	}

	filterIDs, err = filterOwnedCollectibles(ctx, db, filterChains, filterAddresses, filter, 0, nData)
	require.NoError(t, err)
	require.Equal(t, expectedIDs, filterIDs)

	// Test only community
	filter = allFilter()
	filter.FilterCommunity = OnlyCommunity

	tmpIDs, err = oDB.GetOwnedCollectibles(filterChains, filterAddresses, 0, nData)
	require.NoError(t, err)

	expectedIDs = nil
	for _, id := range tmpIDs {
		if dataPerID[id.HashKey()].CommunityID != "" {
			expectedIDs = append(expectedIDs, id)
		}
	}

	filterIDs, err = filterOwnedCollectibles(ctx, db, filterChains, filterAddresses, filter, 0, nData)
	require.NoError(t, err)
	require.Equal(t, expectedIDs, filterIDs)

	// Test specific community
	communityIDa := data[0].CommunityID
	communityIDb := data[1].CommunityID
	communityIDs := []string{communityIDa, communityIDb}

	filter = allFilter()
	filter.CommunityIDs = communityIDs

	tmpIDs, err = oDB.GetOwnedCollectibles(filterChains, filterAddresses, 0, nData)
	require.NoError(t, err)

	expectedIDs = nil
	for _, id := range tmpIDs {
		if dataPerID[id.HashKey()].CommunityID == communityIDa || dataPerID[id.HashKey()].CommunityID == communityIDb {
			expectedIDs = append(expectedIDs, id)
		}
	}

	filterIDs, err = filterOwnedCollectibles(ctx, db, filterChains, filterAddresses, filter, 0, nData)
	require.NoError(t, err)
	require.Equal(t, expectedIDs, filterIDs)

	// Test specific privileges level
	privilegeLevel := token.PrivilegesLevel(2)

	filter = allFilter()
	filter.CommunityPrivilegesLevels = []token.PrivilegesLevel{privilegeLevel}

	tmpIDs, err = oDB.GetOwnedCollectibles(filterChains, filterAddresses, 0, nData)
	require.NoError(t, err)

	expectedIDs = nil
	for _, id := range tmpIDs {
		if communityDataPerID[id.HashKey()].PrivilegesLevel == privilegeLevel {
			expectedIDs = append(expectedIDs, id)
		}
	}

	filterIDs, err = filterOwnedCollectibles(ctx, db, filterChains, filterAddresses, filter, 0, nData)
	require.NoError(t, err)
	require.Equal(t, expectedIDs, filterIDs)
}
