package collectibles

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/status-im/status-go/contracts/collectibles"
	"github.com/status-im/status-go/rpc"
	"github.com/status-im/status-go/services/wallet/thirdparty"
	"github.com/status-im/status-go/services/wallet/thirdparty/opensea"
)

const requestTimeout = 5 * time.Second

type Manager struct {
	rpcClient        *rpc.Client
	metadataProvider thirdparty.NFTMetadataProvider
	openseaAPIKey    string
	nftCache         map[uint64]map[string]opensea.Asset
	nftCacheLock     sync.RWMutex
}

func NewManager(rpcClient *rpc.Client, metadataProvider thirdparty.NFTMetadataProvider, openseaAPIKey string) *Manager {
	return &Manager{
		rpcClient:        rpcClient,
		metadataProvider: metadataProvider,
		openseaAPIKey:    openseaAPIKey,
		nftCache:         make(map[uint64]map[string]opensea.Asset),
	}
}

func (o *Manager) FetchAllCollectionsByOwner(chainID uint64, owner common.Address) ([]opensea.OwnedCollection, error) {
	client, err := opensea.NewOpenseaClient(chainID, o.openseaAPIKey)
	if err != nil {
		return nil, err
	}
	return client.FetchAllCollectionsByOwner(owner)
}

func (o *Manager) FetchAllAssetsByOwnerAndCollection(chainID uint64, owner common.Address, collectionSlug string, cursor string, limit int) (*opensea.AssetContainer, error) {
	client, err := opensea.NewOpenseaClient(chainID, o.openseaAPIKey)
	if err != nil {
		return nil, err
	}

	assetContainer, err := client.FetchAllAssetsByOwnerAndCollection(owner, collectionSlug, cursor, limit)
	if err != nil {
		return nil, err
	}

	err = o.processAssets(chainID, assetContainer.Assets)
	if err != nil {
		return nil, err
	}

	return assetContainer, nil
}

func (o *Manager) FetchAllAssetsByOwnerAndContractAddress(chainID uint64, owner common.Address, contractAddresses []common.Address, cursor string, limit int) (*opensea.AssetContainer, error) {
	client, err := opensea.NewOpenseaClient(chainID, o.openseaAPIKey)
	if err != nil {
		return nil, err
	}

	assetContainer, err := client.FetchAllAssetsByOwnerAndContractAddress(owner, contractAddresses, cursor, limit)
	if err != nil {
		return nil, err
	}

	err = o.processAssets(chainID, assetContainer.Assets)
	if err != nil {
		return nil, err
	}

	return assetContainer, nil
}

func (o *Manager) FetchAllAssetsByOwner(chainID uint64, owner common.Address, cursor string, limit int) (*opensea.AssetContainer, error) {
	client, err := opensea.NewOpenseaClient(chainID, o.openseaAPIKey)
	if err != nil {
		return nil, err
	}

	assetContainer, err := client.FetchAllAssetsByOwner(owner, cursor, limit)
	if err != nil {
		return nil, err
	}

	err = o.processAssets(chainID, assetContainer.Assets)
	if err != nil {
		return nil, err
	}

	return assetContainer, nil
}

func (o *Manager) FetchAssetsByNFTUniqueID(chainID uint64, uniqueIDs []thirdparty.NFTUniqueID, limit int) (*opensea.AssetContainer, error) {
	assetContainer := new(opensea.AssetContainer)

	idsToFetch := o.getIDsNotInCache(chainID, uniqueIDs)
	if len(idsToFetch) > 0 {
		client, err := opensea.NewOpenseaClient(chainID, o.openseaAPIKey)
		if err != nil {
			return nil, err
		}

		fetchedAssetContainer, err := client.FetchAssetsByNFTUniqueID(idsToFetch, limit)
		if err != nil {
			return nil, err
		}

		err = o.processAssets(chainID, fetchedAssetContainer.Assets)
		if err != nil {
			return nil, err
		}

		assetContainer.NextCursor = fetchedAssetContainer.NextCursor
		assetContainer.PreviousCursor = fetchedAssetContainer.PreviousCursor
	}

	assetContainer.Assets = o.getCachedAssets(chainID, uniqueIDs)

	return assetContainer, nil
}

func isMetadataEmpty(asset opensea.Asset) bool {
	return asset.Name == "" &&
		asset.Description == "" &&
		asset.ImageURL == "" &&
		asset.TokenURI == ""
}

func (o *Manager) fetchTokenURI(chainID uint64, id thirdparty.NFTUniqueID) (string, error) {
	backend, err := o.rpcClient.EthClient(chainID)
	if err != nil {
		return "", err
	}

	caller, err := collectibles.NewCollectiblesCaller(id.ContractAddress, backend)
	if err != nil {
		return "", err
	}

	timeoutContext, timeoutCancel := context.WithTimeout(context.Background(), requestTimeout)
	defer timeoutCancel()

	tokenURI, err := caller.TokenURI(&bind.CallOpts{
		Context: timeoutContext,
	}, id.TokenID.Int)

	if err != nil {
		if strings.HasPrefix(err.Error(), "execution reverted") {
			// Contract doesn't support "TokenURI" method
			return "", nil
		}
		return "", err
	}

	return tokenURI, err
}

func (o *Manager) processAssets(chainID uint64, assets []opensea.Asset) error {
	o.nftCacheLock.Lock()
	defer o.nftCacheLock.Unlock()

	if _, ok := o.nftCache[chainID]; !ok {
		o.nftCache[chainID] = make(map[string]opensea.Asset)
	}

	for idx, asset := range assets {
		id := thirdparty.NFTUniqueID{
			ContractAddress: common.HexToAddress(asset.Contract.Address),
			TokenID:         asset.TokenID,
		}

		if isMetadataEmpty(asset) {
			tokenURI, err := o.fetchTokenURI(chainID, id)

			if err != nil {
				return err
			}

			assets[idx].TokenURI = tokenURI

			canProvide, err := o.metadataProvider.CanProvideNFTMetadata(chainID, id, tokenURI)

			if err != nil {
				return err
			}

			if canProvide {
				metadata, err := o.metadataProvider.FetchNFTMetadata(chainID, id, tokenURI)
				if err != nil {
					return err
				}

				if metadata != nil {
					assets[idx].Name = metadata.Name
					assets[idx].Description = metadata.Description
					assets[idx].Collection.ImageURL = metadata.CollectionImageURL
					assets[idx].ImageURL = metadata.ImageURL
				}
			}
		}

		o.nftCache[chainID][id.HashKey()] = assets[idx]
	}

	return nil
}

func (o *Manager) getIDsNotInCache(chainID uint64, uniqueIDs []thirdparty.NFTUniqueID) []thirdparty.NFTUniqueID {
	o.nftCacheLock.RLock()
	defer o.nftCacheLock.RUnlock()

	idsToFetch := make([]thirdparty.NFTUniqueID, 0, len(uniqueIDs))
	if _, ok := o.nftCache[chainID]; !ok {
		idsToFetch = uniqueIDs
	} else {
		for _, id := range uniqueIDs {
			if _, ok := o.nftCache[chainID][id.HashKey()]; !ok {
				idsToFetch = append(idsToFetch, id)
			}
		}
	}
	return idsToFetch
}

func (o *Manager) getCachedAssets(chainID uint64, uniqueIDs []thirdparty.NFTUniqueID) []opensea.Asset {
	o.nftCacheLock.RLock()
	defer o.nftCacheLock.RUnlock()

	assets := make([]opensea.Asset, 0, len(uniqueIDs))

	if _, ok := o.nftCache[chainID]; ok {
		for _, id := range uniqueIDs {

			if asset, ok := o.nftCache[chainID][id.HashKey()]; ok {
				assets = append(assets, asset)
			}
		}
	}

	return assets
}
