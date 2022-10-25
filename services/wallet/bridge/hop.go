package bridge

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/status-im/status-go/account"
	"github.com/status-im/status-go/contracts"
	"github.com/status-im/status-go/eth-node/types"
	"github.com/status-im/status-go/params"
	"github.com/status-im/status-go/rpc"
	"github.com/status-im/status-go/services/wallet/token"
	"github.com/status-im/status-go/transactions"
)

const HopLpFeeBps = 4
const HopCanonicalTokenIndex = 0
const HophTokenIndex = 1
const HopMinBonderFeeUsd = 0.25

var HopBondTransferGasLimit = map[uint64]int64{
	1:      165000,
	5:      165000,
	10:     100000000,
	42161:  2500000,
	420:    100000000,
	421613: 2500000,
}
var HopSettlementGasLimitPerTx = map[uint64]int64{
	1:      5141,
	5:      5141,
	10:     8545,
	42161:  19843,
	420:    8545,
	421613: 19843,
}
var HopBonderFeeBps = map[string]map[uint64]int64{
	"USDC": {
		1:      14,
		5:      14,
		10:     14,
		42161:  14,
		420:    14,
		421613: 14,
	},
	"USDT": {
		1:      26,
		10:     26,
		421613: 26,
	},
	"DAI": {
		1:     26,
		10:    26,
		42161: 26,
	},
	"ETH": {
		1:      5,
		5:      5,
		10:     5,
		42161:  5,
		420:    5,
		421613: 5,
	},
	"WBTC": {
		1:     23,
		10:    23,
		42161: 23,
	},
}

type HopTxArgs struct {
	transactions.SendTxArgs
	ChainID   uint64         `json:"chainId"`
	Symbol    string         `json:"symbol"`
	Recipient common.Address `json:"recipient"`
	Amount    *hexutil.Big   `json:"amount"`
	BonderFee *hexutil.Big   `json:"bonderFee"`
}

type HopBridge struct {
	contractMaker *contracts.ContractMaker
}

func NewHopBridge(rpcClient *rpc.Client) *HopBridge {
	return &HopBridge{
		contractMaker: &contracts.ContractMaker{RPCClient: rpcClient},
	}
}

func (h *HopBridge) Name() string {
	return "Hop"
}

func (h *HopBridge) Can(from, to *params.Network, token *token.Token, balance *big.Int) (bool, error) {
	if balance.Cmp(big.NewInt(0)) == 0 {
		return false, nil
	}

	if from.ChainID == to.ChainID {
		return false, nil
	}

	fees, ok := HopBonderFeeBps[token.Symbol]
	if !ok {
		return false, nil
	}

	if _, ok := fees[from.ChainID]; !ok {
		return false, nil
	}

	if _, ok := fees[to.ChainID]; !ok {
		return false, nil
	}
	return true, nil
}

func (h *HopBridge) EstimateGas(from, to *params.Network, token *token.Token, amountIn *big.Int) (uint64, error) {
	// TODO: find why this doesn't work
	// ethClient, err := s.contractMaker.RPCClient.EthClient(from.ChainID)
	// if err != nil {
	// 	return 0, err
	// }
	// zero := common.HexToAddress("0x0")
	// zeroInt := big.NewInt(0)
	// var data []byte
	// if from.Layer == 1 {
	// 	bridgeABI, err := abi.JSON(strings.NewReader(hopBridge.HopBridgeABI))
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// 	data, err = bridgeABI.Pack("sendToL2", big.NewInt(int64(to.ChainID)), zero, amountIn, zeroInt, zeroInt, zero, zeroInt)
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// } else {
	// 	wrapperABI, err := abi.JSON(strings.NewReader(hopWrapper.HopWrapperABI))
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// 	data, err = wrapperABI.Pack("swapAndSend", big.NewInt(int64(to.ChainID)), zero, amountIn, zeroInt, zeroInt, zeroInt, zeroInt, zeroInt)
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// }
	// estimate, err := ethClient.EstimateGas(context.Background(), ethereum.CallMsg{
	// 	From:  zero,
	// 	To:    &token.Address,
	// 	Value: big.NewInt(0),
	// 	Data:  data,
	// })
	return 500000 + 1000, nil
}

func (h *HopBridge) Send(sendArgs *TransactionBridge, verifiedAccount *account.SelectedExtKey) (hash types.Hash, err error) {
	networks, err := h.contractMaker.RPCClient.NetworkManager.Get(false)
	if err != nil {
		return hash, err
	}
	var fromNetwork *params.Network
	for _, network := range networks {
		if network.ChainID == sendArgs.ChainID {
			fromNetwork = network
			break
		}
	}

	if fromNetwork.Layer == 1 {
		return h.sendToL2(sendArgs.ChainID, sendArgs.HopTx, verifiedAccount)
	}
	return h.swapAndSend(sendArgs.ChainID, sendArgs.HopTx, verifiedAccount)
}

func getSigner(chainID uint64, from types.Address, verifiedAccount *account.SelectedExtKey) bind.SignerFn {
	return func(addr common.Address, tx *ethTypes.Transaction) (*ethTypes.Transaction, error) {
		s := ethTypes.NewLondonSigner(new(big.Int).SetUint64(chainID))
		return ethTypes.SignTx(tx, s, verifiedAccount.AccountKey.PrivateKey)
	}
}

func (h *HopBridge) sendToL2(chainID uint64, sendArgs *HopTxArgs, verifiedAccount *account.SelectedExtKey) (hash types.Hash, err error) {
	bridge, err := h.contractMaker.NewHopL1Bridge(chainID, sendArgs.Symbol)
	if err != nil {
		return hash, err
	}
	txOpts := sendArgs.ToTransactOpts(getSigner(chainID, sendArgs.From, verifiedAccount))
	txOpts.Value = (*big.Int)(sendArgs.Amount)
	now := time.Now()
	deadline := big.NewInt(now.Unix() + 604800)
	tx, err := bridge.SendToL2(
		txOpts,
		big.NewInt(int64(sendArgs.ChainID)),
		sendArgs.Recipient,
		sendArgs.Amount.ToInt(),
		big.NewInt(0),
		deadline,
		common.HexToAddress("0x0"),
		big.NewInt(0),
	)

	if err != nil {
		return hash, err
	}
	return types.Hash(tx.Hash()), nil
}

func (h *HopBridge) swapAndSend(chainID uint64, sendArgs *HopTxArgs, verifiedAccount *account.SelectedExtKey) (hash types.Hash, err error) {
	ammWrapper, err := h.contractMaker.NewHopL2AmmWrapper(chainID, sendArgs.Symbol)
	if err != nil {
		return hash, err
	}

	txOpts := sendArgs.ToTransactOpts(getSigner(chainID, sendArgs.From, verifiedAccount))
	txOpts.Value = (*big.Int)(sendArgs.Amount)
	now := time.Now()
	deadline := big.NewInt(now.Unix() + 604800)
	tx, err := ammWrapper.SwapAndSend(
		txOpts,
		big.NewInt(int64(sendArgs.ChainID)),
		sendArgs.Recipient,
		sendArgs.Amount.ToInt(),
		sendArgs.BonderFee.ToInt(),
		big.NewInt(0),
		deadline,
		big.NewInt(0),
		deadline,
	)

	if err != nil {
		return hash, err
	}

	return types.Hash(tx.Hash()), nil
}

// CalculateBonderFees logics come from: https://docs.hop.exchange/fee-calculation
func (h *HopBridge) CalculateBonderFees(from, to *params.Network, token *token.Token, amountIn *big.Int, nativeTokenPrice, tokenPrice float64, gasPrice *big.Float) (*big.Int, error) {
	amount := new(big.Float).SetInt(amountIn)
	totalFee := big.NewFloat(0)
	destinationTxFee, err := h.getDestinationTxFee(from, to, nativeTokenPrice, tokenPrice, gasPrice)
	if err != nil {
		return nil, err
	}

	bonderFeeRelative, err := h.getBonderFeeRelative(from, to, amount, token)
	if err != nil {
		return nil, err
	}
	if from.Layer != 1 {
		adjustedBonderFee, err := h.calcFromHTokenAmount(to, bonderFeeRelative, token.Symbol)
		if err != nil {
			return nil, err
		}
		adjustedDestinationTxFee, err := h.calcToHTokenAmount(to, destinationTxFee, token.Symbol)
		if err != nil {
			return nil, err
		}

		bonderFeeAbsolute := h.getBonderFeeAbsolute(tokenPrice)
		if adjustedBonderFee.Cmp(bonderFeeAbsolute) == -1 {
			adjustedBonderFee = bonderFeeAbsolute
		}

		totalFee.Add(adjustedBonderFee, adjustedDestinationTxFee)
	}
	res, _ := new(big.Float).Mul(totalFee, big.NewFloat(math.Pow(10, float64(token.Decimals)))).Int(nil)
	return res, nil
}

func (h *HopBridge) CalculateFees(from, to *params.Network, token *token.Token, amountIn *big.Int, nativeTokenPrice, tokenPrice float64, gasPrice *big.Float) (*big.Int, *big.Int, error) {
	bonderFees, err := h.CalculateBonderFees(from, to, token, amountIn, nativeTokenPrice, tokenPrice, gasPrice)
	if err != nil {
		return nil, nil, err
	}
	amountOut, err := h.amountOut(from, to, new(big.Float).SetInt(amountIn), token.Symbol)
	if err != nil {
		return nil, nil, err
	}
	amountOutInt, _ := amountOut.Int(nil)

	return bonderFees, new(big.Int).Add(
		bonderFees,
		new(big.Int).Sub(amountIn, amountOutInt),
	), nil
}

func (h *HopBridge) calcToHTokenAmount(network *params.Network, amount *big.Float, symbol string) (*big.Float, error) {
	if network.Layer == 1 || amount.Cmp(big.NewFloat(0)) == 0 {
		return amount, nil
	}

	contract, err := h.contractMaker.NewHopL2SaddlSwap(network.ChainID, symbol)
	if err != nil {
		return nil, err
	}
	amountInt, _ := amount.Int(nil)
	res, err := contract.CalculateSwap(&bind.CallOpts{Context: context.Background()}, HopCanonicalTokenIndex, HophTokenIndex, amountInt)
	if err != nil {
		return nil, err
	}

	return new(big.Float).SetInt(res), nil
}

func (h *HopBridge) calcFromHTokenAmount(network *params.Network, amount *big.Float, symbol string) (*big.Float, error) {
	if network.Layer == 1 || amount.Cmp(big.NewFloat(0)) == 0 {
		return amount, nil
	}
	contract, err := h.contractMaker.NewHopL2SaddlSwap(network.ChainID, symbol)
	if err != nil {
		return nil, err
	}
	amountInt, _ := amount.Int(nil)
	res, err := contract.CalculateSwap(&bind.CallOpts{Context: context.Background()}, HophTokenIndex, HopCanonicalTokenIndex, amountInt)
	if err != nil {
		return nil, err
	}

	return new(big.Float).SetInt(res), nil
}

func (h *HopBridge) CalculateAmountOut(from, to *params.Network, amountIn *big.Int, symbol string) (*big.Int, error) {
	amountOut, err := h.amountOut(from, to, new(big.Float).SetInt(amountIn), symbol)
	if err != nil {
		return nil, err
	}
	amountOutInt, _ := amountOut.Int(nil)
	return amountOutInt, nil
}

func (h *HopBridge) amountOut(from, to *params.Network, amountIn *big.Float, symbol string) (*big.Float, error) {
	hTokenAmount, err := h.calcToHTokenAmount(from, amountIn, symbol)
	if err != nil {
		return nil, err
	}
	return h.calcFromHTokenAmount(to, hTokenAmount, symbol)
}

func (h *HopBridge) getBonderFeeRelative(from, to *params.Network, amount *big.Float, token *token.Token) (*big.Float, error) {
	if from.Layer != 1 {
		return big.NewFloat(0), nil
	}

	hTokenAmount, err := h.calcToHTokenAmount(from, amount, token.Symbol)
	if err != nil {
		return nil, err
	}
	feeBps := HopBonderFeeBps[token.Symbol][to.ChainID]

	factor := new(big.Float).Mul(hTokenAmount, big.NewFloat(float64(feeBps)))
	return new(big.Float).Quo(
		factor,
		big.NewFloat(10000),
	), nil
}

func (h *HopBridge) getBonderFeeAbsolute(tokenPrice float64) *big.Float {
	return new(big.Float).Quo(big.NewFloat(HopMinBonderFeeUsd), big.NewFloat(tokenPrice))
}

func (h *HopBridge) getDestinationTxFee(from, to *params.Network, nativeTokenPrice, tokenPrice float64, gasPrice *big.Float) (*big.Float, error) {
	if from.Layer != 1 {
		return big.NewFloat(0), nil
	}

	bondTransferGasLimit := HopBondTransferGasLimit[to.ChainID]
	settlementGasLimit := HopSettlementGasLimitPerTx[to.ChainID]
	totalGasLimit := new(big.Int).Add(big.NewInt(bondTransferGasLimit), big.NewInt(settlementGasLimit))

	rate := new(big.Float).Quo(big.NewFloat(nativeTokenPrice), big.NewFloat(tokenPrice))

	txFeeEth := new(big.Float).Mul(gasPrice, new(big.Float).SetInt(totalGasLimit))
	return new(big.Float).Mul(txFeeEth, rate), nil
}
