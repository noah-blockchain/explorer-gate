package core

import (
	"github.com/noah-blockchain/explorer-gate/env"
	"github.com/noah-blockchain/explorer-gate/errors"
	"github.com/noah-blockchain/noah-node-go-api"
	"github.com/sirupsen/logrus"
	"github.com/tendermint/tendermint/libs/pubsub"
	"strings"
)

type NoahGate struct {
	api     *noah_node_go_api.NoahNodeApi
	emitter *pubsub.Server
	Logger  *logrus.Entry
}

type CoinEstimate struct {
	Value      string
	Commission string
}

//New instance of Noah Gate
func New(e *pubsub.Server, logger *logrus.Entry) *NoahGate {
	return &NoahGate{
		emitter: e,
		api:     noah_node_go_api.New(env.GetEnv(env.NoahApiNodeEnv, "")),
		Logger:  logger,
	}
}

//Send transaction to blockchain
//Return transaction hash
func (mg NoahGate) TxPush(transaction string) (*string, error) {
	response, err := mg.api.PushTransaction(transaction)
	if err != nil {
		mg.Logger.WithFields(logrus.Fields{
			"transaction": transaction,
		}).Warn(err)
		return nil, err
	}
	if response.Error != nil || response.Result.Code != 0 {
		err = errors.GetNodeErrorFromResponse(response)
		mg.Logger.WithFields(logrus.Fields{
			"transaction": transaction,
		}).Warn(err)
		return nil, err
	}
	hash := `Nt` + strings.ToLower(response.Result.Hash)
	return &hash, nil
}

//Return estimate of transaction
func (mg *NoahGate) EstimateTxCommission(transaction string) (*string, error) {
	response, err := mg.api.GetEstimateTx(transaction)
	if err != nil {
		mg.Logger.WithFields(logrus.Fields{
			"transaction": transaction,
		}).Warn(err)
		return nil, err
	}
	if response.Error != nil {
		err = errors.NewNodeError(response.Error.Message, response.Error.Code)
		mg.Logger.WithFields(logrus.Fields{
			"transaction": transaction,
		}).Warn(err)
		return nil, err
	}
	return &response.Result.Commission, nil
}

//Return estimate of buy coin
func (mg *NoahGate) EstimateCoinBuy(coinToSell string, coinToBuy string, value string) (*CoinEstimate, error) {
	response, err := mg.api.GetEstimateCoinBuy(coinToSell, coinToBuy, value)
	if err != nil {
		mg.Logger.WithFields(logrus.Fields{
			"coinToSell": coinToSell,
			"coinToBuy":  coinToBuy,
			"value":      value,
		}).Warn(err)
		return nil, err
	}

	if response.Error != nil {
		err = errors.NewNodeError(response.Error.Message, response.Error.Code)
		mg.Logger.WithFields(logrus.Fields{
			"coinToSell": coinToSell,
			"coinToBuy":  coinToBuy,
			"value":      value,
		}).Warn(err)
		return nil, err
	}

	return &CoinEstimate{response.Result.WillPay, response.Result.Commission}, nil
}

//Return estimate of sell coin
func (mg *NoahGate) EstimateCoinSell(coinToSell string, coinToBuy string, value string) (*CoinEstimate, error) {
	response, err := mg.api.GetEstimateCoinSell(coinToSell, coinToBuy, value, 0) // fixed here
	if err != nil {
		mg.Logger.WithFields(logrus.Fields{
			"coinToSell": coinToSell,
			"coinToBuy":  coinToBuy,
			"value":      value,
		}).Warn(err)
		return nil, err
	}
	if response.Error != nil {
		err = errors.NewNodeError(response.Error.Message, response.Error.Code)
		mg.Logger.WithFields(logrus.Fields{
			"coinToSell": coinToSell,
			"coinToBuy":  coinToBuy,
			"value":      value,
		}).Warn(err)
		return nil, err
	}
	return &CoinEstimate{response.Result.WillGet, response.Result.Commission}, nil
}

//Return estimate of sell coin
func (mg *NoahGate) EstimateCoinSellAll(coinToSell string, coinToBuy string, value string, gasPrice string) (*CoinEstimate, error) {
	response, err := mg.api.GetEstimateCoinSellAll(coinToSell, coinToBuy, value, gasPrice)
	if err != nil {
		mg.Logger.WithFields(logrus.Fields{
			"coinToSell": coinToSell,
			"coinToBuy":  coinToBuy,
			"value":      value,
			"gasPrice":   gasPrice,
		}).Warn(err)
		return nil, err
	}
	if response.Error != nil {
		err = errors.NewNodeError(response.Error.Message, response.Error.Code)
		mg.Logger.WithFields(logrus.Fields{
			"coinToSell": coinToSell,
			"coinToBuy":  coinToBuy,
			"value":      value,
			"gasPrice":   gasPrice,
		}).Warn(err)
		return nil, err
	}
	return &CoinEstimate{response.Result.WillGet, ""}, nil
}

//Return nonce for address
func (mg *NoahGate) GetNonce(address string) (*string, error) {
	response, err := mg.api.GetAddress(address)
	if err != nil {
		mg.Logger.WithFields(logrus.Fields{
			"address": address,
		}).Warn(err)
		return nil, err
	}
	if response.Error != nil {
		err = errors.NewNodeError(response.Error.Message, response.Error.Code)
		mg.Logger.WithFields(logrus.Fields{
			"address": address,
		}).Warn(err)
		return nil, err
	}
	return &response.Result.TransactionCount, nil
}

//Return nonce for address
func (mg *NoahGate) GetMinGas() (*string, error) {
	response, err := mg.api.GetMinGasPrice()
	if err != nil {
		mg.Logger.Error(err)
		return nil, err
	}
	if response.Error != nil {
		err = errors.NewNodeError(response.Error.Message, response.Error.Code)
		mg.Logger.Error(err)
		return nil, err
	}
	return &response.Result, nil
}
