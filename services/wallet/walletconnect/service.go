package walletconnect

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"

	"github.com/status-im/status-go/account"
	"github.com/status-im/status-go/eth-node/crypto"
	"github.com/status-im/status-go/eth-node/types"
	"github.com/status-im/status-go/multiaccounts/accounts"
	"github.com/status-im/status-go/params"
	"github.com/status-im/status-go/rpc/network"
	"github.com/status-im/status-go/transactions"
)

type txSigningDetails struct {
	chainID       uint64
	from          common.Address
	txBeingSigned *ethTypes.Transaction
}

type Service struct {
	db             *sql.DB
	networkManager *network.Manager
	accountsDB     *accounts.Database
	eventFeed      *event.Feed

	transactor  *transactions.Transactor
	gethManager *account.GethManager

	config        *params.NodeConfig
	txSignDetails *txSigningDetails
}

func NewService(db *sql.DB, networkManager *network.Manager, accountsDB *accounts.Database, transactor *transactions.Transactor, gethManager *account.GethManager, eventFeed *event.Feed, config *params.NodeConfig) *Service {
	return &Service{
		db:             db,
		networkManager: networkManager,
		accountsDB:     accountsDB,
		eventFeed:      eventFeed,
		transactor:     transactor,
		gethManager:    gethManager,
		config:         config,
	}
}

func (s *Service) SignMessage(message types.HexBytes, address common.Address, password string) (string, error) {
	selectedAccount, err := s.gethManager.VerifyAccountPassword(s.config.KeyStoreDir, address.Hex(), password)
	if err != nil {
		return "", err
	}

	signature, err := crypto.Sign(message[:], selectedAccount.PrivateKey)

	return types.EncodeHex(signature), err
}

func (s *Service) SendTransaction(signature string) (response *SessionRequestResponse, err error) {
	return s.sendTransaction(signature)
}

func (s *Service) PairSessionProposal(proposal SessionProposal) (*PairSessionResponse, error) {
	namespace := Namespace{
		Methods: []string{params.SendTransactionMethodName, params.PersonalSignMethodName},
		Events:  []string{"accountsChanged", "chainChanged"},
	}

	proposedChains := proposal.Params.RequiredNamespaces.Eip155.Chains
	chains, eipChains := sessionProposalToSupportedChain(proposedChains, func(chainID uint64) bool {
		return s.networkManager.Find(chainID) != nil
	})
	if len(chains) != len(proposedChains) {
		log.Warn("Some chains are not supported; wanted: ", proposedChains, "; supported: ", chains)
		return nil, ErrorChainsNotSupported
	}
	namespace.Chains = eipChains

	activeAccounts, err := s.accountsDB.GetActiveAccounts()
	if err != nil {
		return nil, fmt.Errorf("failed to get active accounts: %w", err)
	}

	// Filter out non-own accounts
	usableAccounts := make([]*accounts.Account, 0, 1)
	for _, acc := range activeAccounts {
		if !acc.IsWalletAccountReadyForTransaction() {
			continue
		}
		usableAccounts = append(usableAccounts, acc)
	}

	addresses := activeToOwnedAccounts(usableAccounts)
	namespace.Accounts = caip10Accounts(addresses, chains)

	// TODO #12434: respond async
	return &PairSessionResponse{
		SessionProposal: proposal,
		SupportedNamespaces: Namespaces{
			Eip155: namespace,
		},
	}, nil
}

func (s *Service) RecordSuccessfulPairing(proposal SessionProposal) error {
	var icon string
	if len(proposal.Params.Proposer.Metadata.Icons) > 0 {
		icon = proposal.Params.Proposer.Metadata.Icons[0]
	}
	return InsertPairing(s.db, Pairing{
		Topic:       proposal.Params.PairingTopic,
		Expiry:      proposal.Params.Expiry,
		Active:      true,
		AppName:     proposal.Params.Proposer.Metadata.Name,
		URL:         proposal.Params.Proposer.Metadata.URL,
		Description: proposal.Params.Proposer.Metadata.Description,
		Icon:        icon,
		Verified:    proposal.Params.Verify.Verified,
	})
}

func (s *Service) HasActivePairings() (bool, error) {
	return HasActivePairings(s.db, time.Now().Unix())
}

func (s *Service) SessionRequest(request SessionRequest) (response *SessionRequestResponse, err error) {
	// TODO #12434: should we check topic for validity? It might make sense if we
	// want to cache the paired sessions

	if request.Params.Request.Method == params.SendTransactionMethodName {
		return s.buildTransaction(request)
	} else if request.Params.Request.Method == params.PersonalSignMethodName {
		return s.buildPersonalSingMessage(request)
	}

	// TODO #12434: respond async
	return nil, ErrorMethodNotSupported
}
