package transfer

import (
	"context"
	"database/sql"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"

	"github.com/status-im/status-go/rpc/chain"
	"github.com/status-im/status-go/services/wallet/async"
	"github.com/status-im/status-go/services/wallet/balance"
	w_common "github.com/status-im/status-go/services/wallet/common"
	"github.com/status-im/status-go/services/wallet/token"
	"github.com/status-im/status-go/services/wallet/walletevent"
	"github.com/status-im/status-go/transactions"
)

const (
	// EventNewTransfers emitted when new block was added to the same canonical chan.
	EventNewTransfers walletevent.EventType = "new-transfers"
	// EventFetchingRecentHistory emitted when fetching of lastest tx history is started
	EventFetchingRecentHistory walletevent.EventType = "recent-history-fetching"
	// EventRecentHistoryReady emitted when fetching of lastest tx history is started
	EventRecentHistoryReady walletevent.EventType = "recent-history-ready"
	// EventFetchingHistoryError emitted when fetching of tx history failed
	EventFetchingHistoryError walletevent.EventType = "fetching-history-error"
	// EventNonArchivalNodeDetected emitted when a connection to a non archival node is detected
	EventNonArchivalNodeDetected walletevent.EventType = "non-archival-node-detected"

	// EventInternalERC721TransferDetected emitted when ERC721 transfer is detected
	EventInternalERC721TransferDetected walletevent.EventType = walletevent.InternalEventTypePrefix + "erc721-transfer-detected"

	numberOfBlocksCheckedPerIteration = 40
	noBlockLimit                      = 0
)

var (
	// This will work only for binance testnet as mainnet doesn't support
	// archival request.
	binanceChainErc20BatchSize   = big.NewInt(5000)
	goerliErc20BatchSize         = big.NewInt(100000)
	goerliErc20ArbitrumBatchSize = big.NewInt(10000)
	goerliErc20OptimismBatchSize = big.NewInt(10000)
	erc20BatchSize               = big.NewInt(500000)
	binancChainID                = uint64(56)
	goerliChainID                = uint64(5)
	goerliArbitrumChainID        = uint64(421613)
	goerliOptimismChainID        = uint64(420)
	binanceTestChainID           = uint64(97)
)

type ethHistoricalCommand struct {
	address       common.Address
	chainClient   chain.ClientInterface
	balanceCacher balance.Cacher
	feed          *event.Feed
	foundHeaders  []*DBHeader
	error         error
	noLimit       bool

	from                          *Block
	to, resultingFrom, startBlock *big.Int
	threadLimit                   uint32
}

type Transaction []*Transfer

func (c *ethHistoricalCommand) Command() async.Command {
	return async.FiniteCommand{
		Interval: 5 * time.Second,
		Runable:  c.Run,
	}.Run
}

func (c *ethHistoricalCommand) Run(ctx context.Context) (err error) {
	log.Debug("eth historical downloader start", "chainID", c.chainClient.NetworkID(), "address", c.address,
		"from", c.from.Number, "to", c.to, "noLimit", c.noLimit)

	start := time.Now()
	if c.from.Number != nil && c.from.Balance != nil {
		c.balanceCacher.Cache().AddBalance(c.address, c.chainClient.NetworkID(), c.from.Number, c.from.Balance)
	}
	if c.from.Number != nil && c.from.Nonce != nil {
		c.balanceCacher.Cache().AddNonce(c.address, c.chainClient.NetworkID(), c.from.Number, c.from.Nonce)
	}
	from, headers, startBlock, err := findBlocksWithEthTransfers(ctx, c.chainClient,
		c.balanceCacher, c.address, c.from.Number, c.to, c.noLimit, c.threadLimit)

	if err != nil {
		c.error = err
		log.Error("failed to find blocks with transfers", "error", err, "chainID", c.chainClient.NetworkID(),
			"address", c.address, "from", c.from.Number, "to", c.to)
		return nil
	}

	c.foundHeaders = headers
	c.resultingFrom = from
	c.startBlock = startBlock

	log.Debug("eth historical downloader finished successfully", "chain", c.chainClient.NetworkID(),
		"address", c.address, "from", from, "to", c.to, "total blocks", len(headers), "time", time.Since(start))

	return nil
}

type erc20HistoricalCommand struct {
	erc20       BatchDownloader
	address     common.Address
	chainClient chain.ClientInterface
	feed        *event.Feed

	iterator     *IterativeDownloader
	to           *big.Int
	from         *big.Int
	foundHeaders []*DBHeader
}

func (c *erc20HistoricalCommand) Command() async.Command {
	return async.FiniteCommand{
		Interval: 5 * time.Second,
		Runable:  c.Run,
	}.Run
}

func getErc20BatchSize(chainID uint64) *big.Int {
	if isBinanceChain(chainID) {
		return binanceChainErc20BatchSize
	}

	if chainID == goerliChainID {
		return goerliErc20BatchSize
	}

	if chainID == goerliOptimismChainID {
		return goerliErc20OptimismBatchSize
	}

	if chainID == goerliArbitrumChainID {
		return goerliErc20ArbitrumBatchSize
	}

	return erc20BatchSize
}

func (c *erc20HistoricalCommand) Run(ctx context.Context) (err error) {
	log.Debug("wallet historical downloader for erc20 transfers start", "chainID", c.chainClient.NetworkID(), "address", c.address,
		"from", c.from, "to", c.to)

	start := time.Now()
	if c.iterator == nil {
		c.iterator, err = SetupIterativeDownloader(
			c.chainClient, c.address,
			c.erc20, getErc20BatchSize(c.chainClient.NetworkID()), c.to, c.from)
		if err != nil {
			log.Error("failed to setup historical downloader for erc20")
			return err
		}
	}
	for !c.iterator.Finished() {
		headers, _, _, err := c.iterator.Next(ctx)
		if err != nil {
			log.Error("failed to get next batch", "error", err, "chainID", c.chainClient.NetworkID()) // TODO: stop inifinite command in case of an error that we can't fix like missing trie node
			return err
		}
		c.foundHeaders = append(c.foundHeaders, headers...)
	}
	log.Debug("wallet historical downloader for erc20 transfers finished", "chainID", c.chainClient.NetworkID(), "address", c.address,
		"from", c.from, "to", c.to, "time", time.Since(start), "headers", len(c.foundHeaders))
	return nil
}

type transfersCommand struct {
	db                 *Database
	blockDAO           *BlockDAO
	eth                *ETHDownloader
	blockNums          []*big.Int
	address            common.Address
	chainClient        chain.ClientInterface
	blocksLimit        int
	transactionManager *TransactionManager
	pendingTxManager   *transactions.PendingTxTracker
	tokenManager       *token.Manager
	feed               *event.Feed

	// result
	fetchedTransfers []Transfer
}

func (c *transfersCommand) Command() async.Command {
	return async.FiniteCommand{
		Interval: 5 * time.Second,
		Runable:  c.Run,
	}.Run
}

func (c *transfersCommand) Run(ctx context.Context) (err error) {
	// Take blocks from cache if available and disrespect the limit
	// If no blocks are available in cache, take blocks from DB respecting the limit
	// If no limit is set, take all blocks from DB
	log.Debug("start transfersCommand", "chain", c.chainClient.NetworkID(), "address", c.address, "blockNums", c.blockNums)
	startTs := time.Now()

	for {
		blocks := c.blockNums
		if blocks == nil {
			blocks, _ = c.blockDAO.GetBlocksToLoadByAddress(c.chainClient.NetworkID(), c.address, numberOfBlocksCheckedPerIteration)
		}

		for _, blockNum := range blocks {
			log.Debug("transfersCommand block start", "chain", c.chainClient.NetworkID(), "address", c.address, "block", blockNum)

			allTransfers, err := c.eth.GetTransfersByNumber(ctx, blockNum)
			if err != nil {
				log.Error("getTransfersByBlocks error", "error", err)
				return err
			}

			c.processUnknownErc20CommunityTransactions(ctx, allTransfers)

			err = c.processMultiTransactions(ctx, allTransfers)
			if err != nil {
				log.Error("processMultiTransactions error", "error", err)
				return err
			}

			if len(allTransfers) > 0 {
				err := c.saveAndConfirmPending(allTransfers, blockNum)
				if err != nil {
					log.Error("saveAndConfirmPending error", "error", err)
					return err
				}
			} else {
				// If no transfers found, that is suspecting, because downloader returned this block as containing transfers
				log.Error("no transfers found in block", "chain", c.chainClient.NetworkID(), "address", c.address, "block", blockNum)

				err = markBlocksAsLoaded(c.chainClient.NetworkID(), c.db.client, c.address, []*big.Int{blockNum})
				if err != nil {
					log.Error("Mark blocks loaded error", "error", err)
					return err
				}
			}

			c.fetchedTransfers = append(c.fetchedTransfers, allTransfers...)

			c.notifyOfNewTransfers(blockNum, allTransfers)
			c.notifyOfNewERC721Transfers(allTransfers)

			log.Debug("transfersCommand block end", "chain", c.chainClient.NetworkID(), "address", c.address,
				"block", blockNum, "tranfers.len", len(allTransfers), "fetchedTransfers.len", len(c.fetchedTransfers))
		}

		if c.blockNums != nil || len(blocks) == 0 ||
			(c.blocksLimit > noBlockLimit && len(blocks) >= c.blocksLimit) {
			log.Debug("loadTransfers breaking loop on block limits reached or 0 blocks", "chain", c.chainClient.NetworkID(),
				"address", c.address, "limit", c.blocksLimit, "blocks", len(blocks))
			break
		}
	}

	log.Debug("end transfersCommand", "chain", c.chainClient.NetworkID(), "address", c.address,
		"blocks.len", len(c.blockNums), "transfers.len", len(c.fetchedTransfers), "in", time.Since(startTs))

	return nil
}

// saveAndConfirmPending ensures only the transaction that has owner (Address) as a sender is matched to the
// corresponding multi-transaction (by multi-transaction ID). This way we ensure that if receiver is in the list
// of accounts filter will discard the proper one
func (c *transfersCommand) saveAndConfirmPending(allTransfers []Transfer, blockNum *big.Int) error {
	tx, resErr := c.db.client.Begin()
	if resErr != nil {
		return resErr
	}
	notifyFunctions := make([]func(), 0)
	defer func() {
		if resErr == nil {
			commitErr := tx.Commit()
			if commitErr != nil {
				log.Error("failed to commit", "error", commitErr)
			}
			for _, notify := range notifyFunctions {
				notify()
			}
		} else {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Error("failed to rollback", "error", rollbackErr)
			}
		}
	}()

	// Confirm all pending transactions that are included in this block
	for i, tr := range allTransfers {
		chainID := w_common.ChainID(tr.NetworkID)
		txHash := tr.Receipt.TxHash
		txType, mTID, err := transactions.GetOwnedPendingStatus(tx, chainID, txHash, tr.Address)
		if err == sql.ErrNoRows {
			if tr.MultiTransactionID > 0 {
				continue
			} else {
				// Outside transaction, already confirmed by another duplicate or not yet downloaded
				existingMTID, err := GetOwnedMultiTransactionID(tx, chainID, tr.ID, tr.Address)
				if err == sql.ErrNoRows || existingMTID == 0 {
					// Outside transaction, ignore it
					continue
				} else if err != nil {
					log.Warn("GetOwnedMultiTransactionID", "error", err)
					continue
				}
				mTID = w_common.NewAndSet(existingMTID)

			}
		} else if err != nil {
			log.Warn("GetOwnedPendingStatus", "error", err)
			continue
		}

		if mTID != nil {
			allTransfers[i].MultiTransactionID = MultiTransactionIDType(*mTID)
		}
		if txType != nil && *txType == transactions.WalletTransfer {
			notify, err := c.pendingTxManager.DeleteBySQLTx(tx, chainID, txHash)
			if err != nil && err != transactions.ErrStillPending {
				log.Error("DeleteBySqlTx error", "error", err)
			}
			notifyFunctions = append(notifyFunctions, notify)
		}
	}

	resErr = saveTransfersMarkBlocksLoaded(tx, c.chainClient.NetworkID(), c.address, allTransfers, []*big.Int{blockNum})
	if resErr != nil {
		log.Error("SaveTransfers error", "error", resErr)
	}

	return resErr
}

// Mark all subTxs of a given Tx with the same multiTxID
func setMultiTxID(tx Transaction, multiTxID MultiTransactionIDType) {
	for _, subTx := range tx {
		subTx.MultiTransactionID = multiTxID
	}
}

func (c *transfersCommand) checkAndProcessSwapMultiTx(ctx context.Context, tx Transaction) (bool, error) {
	for _, subTx := range tx {
		switch subTx.Type {
		// If the Tx contains any uniswapV2Swap/uniswapV3Swap subTx, generate a Swap multiTx
		case w_common.UniswapV2Swap, w_common.UniswapV3Swap:
			multiTransaction, err := buildUniswapSwapMultitransaction(ctx, c.chainClient, c.tokenManager, subTx)
			if err != nil {
				return false, err
			}

			if multiTransaction != nil {
				id, err := c.transactionManager.InsertMultiTransaction(multiTransaction)
				if err != nil {
					return false, err
				}
				setMultiTxID(tx, id)
				return true, nil
			}
		}
	}

	return false, nil
}

func (c *transfersCommand) checkAndProcessBridgeMultiTx(ctx context.Context, tx Transaction) (bool, error) {
	for _, subTx := range tx {
		switch subTx.Type {
		// If the Tx contains any hopBridge subTx, create/update Bridge multiTx
		case w_common.HopBridgeFrom, w_common.HopBridgeTo:
			multiTransaction, err := buildHopBridgeMultitransaction(ctx, c.chainClient, c.transactionManager, c.tokenManager, subTx)
			if err != nil {
				return false, err
			}

			if multiTransaction != nil {
				setMultiTxID(tx, MultiTransactionIDType(multiTransaction.ID))
				return true, nil
			}
		}
	}

	return false, nil
}

func (c *transfersCommand) processUnknownErc20CommunityTransactions(ctx context.Context, allTransfers []Transfer) {
	for _, tx := range allTransfers {
		if tx.Type == w_common.Erc20Transfer {
			// Find token in db or if this is a community token, find its metadata
			_ = c.tokenManager.FindOrCreateTokenByAddress(ctx, tx.NetworkID, *tx.Transaction.To())
		}
	}
}

func (c *transfersCommand) processMultiTransactions(ctx context.Context, allTransfers []Transfer) error {
	txByTxHash := subTransactionListToTransactionsByTxHash(allTransfers)

	// Detect / Generate multitransactions
	// Iterate over all detected transactions
	for _, tx := range txByTxHash {
		// Then check for a Swap transaction
		txProcessed, err := c.checkAndProcessSwapMultiTx(ctx, tx)
		if err != nil {
			return err
		}
		if txProcessed {
			continue
		}

		// Then check for a Bridge transaction
		_, err = c.checkAndProcessBridgeMultiTx(ctx, tx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *transfersCommand) notifyOfNewTransfers(blockNum *big.Int, transfers []Transfer) {
	if c.feed != nil {
		if len(transfers) > 0 {
			c.feed.Send(walletevent.Event{
				Type:        EventNewTransfers,
				Accounts:    []common.Address{c.address},
				ChainID:     c.chainClient.NetworkID(),
				BlockNumber: blockNum,
			})
		}
	}
}

func (c *transfersCommand) notifyOfNewERC721Transfers(transfers []Transfer) {
	if c.feed != nil {
		// Internal event for ERC721 transfers
		latestERC721TransferTimestamp := uint64(0)
		for _, transfer := range transfers {
			if transfer.Type == w_common.Erc721Transfer {
				if transfer.Timestamp > latestERC721TransferTimestamp {
					latestERC721TransferTimestamp = transfer.Timestamp
				}
			}
		}
		if latestERC721TransferTimestamp > 0 {
			c.feed.Send(walletevent.Event{
				Type:     EventInternalERC721TransferDetected,
				Accounts: []common.Address{c.address},
				ChainID:  c.chainClient.NetworkID(),
				At:       int64(latestERC721TransferTimestamp),
			})
		}
	}
}

type loadTransfersCommand struct {
	accounts           []common.Address
	db                 *Database
	blockDAO           *BlockDAO
	chainClient        chain.ClientInterface
	blocksByAddress    map[common.Address][]*big.Int
	transactionManager *TransactionManager
	pendingTxManager   *transactions.PendingTxTracker
	blocksLimit        int
	tokenManager       *token.Manager
	feed               *event.Feed
}

func (c *loadTransfersCommand) Command() async.Command {
	return async.FiniteCommand{
		Interval: 5 * time.Second,
		Runable:  c.Run,
	}.Run
}

func (c *loadTransfersCommand) LoadTransfers(ctx context.Context, limit int, blocksByAddress map[common.Address][]*big.Int) error {
	return loadTransfers(ctx, c.accounts, c.blockDAO, c.db, c.chainClient, limit, blocksByAddress,
		c.transactionManager, c.pendingTxManager, c.tokenManager, c.feed)
}

func (c *loadTransfersCommand) Run(parent context.Context) (err error) {
	err = c.LoadTransfers(parent, c.blocksLimit, c.blocksByAddress)
	return
}

func loadTransfers(ctx context.Context, accounts []common.Address, blockDAO *BlockDAO, db *Database,
	chainClient chain.ClientInterface, blocksLimitPerAccount int, blocksByAddress map[common.Address][]*big.Int,
	transactionManager *TransactionManager, pendingTxManager *transactions.PendingTxTracker,
	tokenManager *token.Manager, feed *event.Feed) error {

	log.Debug("loadTransfers start", "accounts", accounts, "chain", chainClient.NetworkID(), "limit", blocksLimitPerAccount)

	start := time.Now()
	group := async.NewGroup(ctx)

	for _, address := range accounts {
		transfers := &transfersCommand{
			db:          db,
			blockDAO:    blockDAO,
			chainClient: chainClient,
			address:     address,
			eth: &ETHDownloader{
				chainClient: chainClient,
				accounts:    []common.Address{address},
				signer:      types.LatestSignerForChainID(chainClient.ToBigInt()),
				db:          db,
			},
			blockNums:          blocksByAddress[address],
			transactionManager: transactionManager,
			pendingTxManager:   pendingTxManager,
			tokenManager:       tokenManager,
			feed:               feed,
		}
		group.Add(transfers.Command())
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-group.WaitAsync():
		log.Debug("loadTransfers finished for account", "in", time.Since(start), "chain", chainClient.NetworkID())
		return nil
	}
}

func isBinanceChain(chainID uint64) bool {
	return chainID == binancChainID || chainID == binanceTestChainID
}

// Ensure 1 DBHeader per Block Hash
func uniqueHeaderPerBlockHash(allHeaders []*DBHeader) []*DBHeader {
	uniqHeadersByHash := map[common.Hash]*DBHeader{}
	for _, header := range allHeaders {
		uniqHeader, ok := uniqHeadersByHash[header.Hash]
		if ok {
			if len(header.PreloadedTransactions) > 0 {
				uniqHeader.PreloadedTransactions = append(uniqHeader.PreloadedTransactions, header.PreloadedTransactions...)
			}
			uniqHeadersByHash[header.Hash] = uniqHeader
		} else {
			uniqHeadersByHash[header.Hash] = header
		}
	}

	uniqHeaders := []*DBHeader{}
	for _, header := range uniqHeadersByHash {
		uniqHeaders = append(uniqHeaders, header)
	}

	return uniqHeaders
}

// Organize subTransactions by Transaction Hash
func subTransactionListToTransactionsByTxHash(subTransactions []Transfer) map[common.Hash]Transaction {
	rst := map[common.Hash]Transaction{}

	for index := range subTransactions {
		subTx := &subTransactions[index]
		txHash := subTx.Transaction.Hash()

		if _, ok := rst[txHash]; !ok {
			rst[txHash] = make([]*Transfer, 0)
		}
		rst[txHash] = append(rst[txHash], subTx)
	}

	return rst
}
