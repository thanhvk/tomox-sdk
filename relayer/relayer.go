package relayer

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Relayer get token
type Relayer struct {
	rpcURL                string
	coinBase              common.Address
	relayerAddress        common.Address
	lendingRelayerAddress common.Address
}

// NewRelayer init relayer
func NewRelayer(rpcURL string,
	coinBase common.Address,
	relayerAddress common.Address,
	lendingRelayerAddress common.Address,
) *Relayer {

	return &Relayer{
		rpcURL:                rpcURL,
		coinBase:              coinBase,
		relayerAddress:        relayerAddress,
		lendingRelayerAddress: lendingRelayerAddress,
	}
}

// GetRelayer get relayer information
func (r *Relayer) GetRelayer() (*RInfo, error) {
	signer := NewSigner()
	client, err := rpc.Dial(r.rpcURL)
	if err != nil {
		fmt.Println(err)
	}
	ethclient := ethclient.NewClient(client)
	bc := NewBlockchain(client, ethclient, signer)
	return bc.GetRelayer(r.coinBase, r.relayerAddress)

}

// GetLending get relayer information
func (r *Relayer) GetLending() (*LendingRInfo, error) {
	signer := NewSigner()
	client, err := rpc.Dial(r.rpcURL)
	if err != nil {
		fmt.Println(err)
	}
	ethclient := ethclient.NewClient(client)
	bc := NewBlockchain(client, ethclient, signer)
	return bc.GetLendingRelayer(r.coinBase, r.lendingRelayerAddress)

}
