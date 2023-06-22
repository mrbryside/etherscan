package ethereum

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mrbryside/internal/domain/reader"
	"github.com/mrbryside/internal/entity"
)

type internalTransaction struct {
	BlockNumber       string `json:"blockNumber"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	TransactionHash   string `json:"transactionHash"`
	TransactionIndex  string `json:"transactionIndex"`
	GasPrice          string `json:"gasPrice"`
	Gas               string `json:"gas"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
}

type internalRpcBlock struct {
	Hash         common.Hash           `json:"hash"`
	Transactions []internalTransaction `json:"transactions"`
	UncleHashes  []common.Hash         `json:"uncles"`
}

type EthereumClient struct {
	url string
}

func NewEthereumClient() reader.BlockFetcher {
	return EthereumClient{url: "https://mainnet.infura.io/v3/2df69398f7aa46e9bdf095f460e731fa"}
}

func (ec EthereumClient) FetchTransaction(r *reader.Reader) error {

	client, err := rpc.Dial(ec.url)
	if err != nil {
		log.Fatal(err)
	}

	for _, address := range r.BlockEntity().Address {
		err := addTransactions(client, r, address)
		if err != nil {
			log.Printf("Failed to fetch transactions for address %s: %v", address, err)
			continue
		}
	}
	return nil
}

// addTransactions retrieves transactions for the specified Ethereum address within the given block number range.
func addTransactions(client *rpc.Client, r *reader.Reader, address string) error {
	var block internalRpcBlock
	blockR := r.BlockEntity()
	for blockNumber := blockR.From; blockNumber <= blockR.To; blockNumber++ {
		blockNumberHex := fmt.Sprintf("0x%x", blockNumber)

		err := client.Call(&block, "eth_getBlockByNumber", blockNumberHex, true)
		if err != nil {
			return err
		}

		for _, t := range block.Transactions {
			if t.From == address || t.To == address {
				// prepare all data from hex to beatiful type
				blockNumberInt, err := strconv.ParseUint(strings.TrimPrefix(t.BlockNumber, "0x"), 16, 64)
				if err != nil {
					log.Printf("Failed to convert block number to int for transaction: %s", t.TransactionHash)
					continue
				}
				valueWei, success := new(big.Int).SetString(t.Value[2:], 16)
				if !success {
					log.Printf("Failed to convert value to decimal for transaction: %s", t.TransactionHash)
					continue
				}
				valueEth := new(big.Float).Quo(new(big.Float).SetInt(valueWei), big.NewFloat(1e18))

				method := "in"
				if t.From == address {
					method = "out"
				}

				tran := entity.Transaction{
					Hash:    t.TransactionHash,
					Address: address,
					Method:  method,
					Block:   blockNumberInt,
					From:    t.From,
					To:      t.To,
					Value:   valueEth.String(),
				}
				r.AddTransaction(tran)
			}
		}

	}
	return nil
}
