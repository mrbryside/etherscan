package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

// Define the Ethereum addresses to monitor
var ethereumAddresses = []string{
	"0x28c6c06298d514db089934071355e5743bf21d60",
}

// Ethereum JSON-RPC endpoint
const ethereumRPCURL = "https://mainnet.infura.io/v3/2df69398f7aa46e9bdf095f460e731fa"

type MyTX struct {
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

type rpcBlock struct {
	Hash         common.Hash   `json:"hash"`
	Transactions []MyTX        `json:"transactions"`
	UncleHashes  []common.Hash `json:"uncles"`
}

func main() {
	client, err := rpc.Dial(ethereumRPCURL)
	if err != nil {
		log.Fatal(err)
	}

	for _, address := range ethereumAddresses {
		_, err := getTransactions(client, address, 17065470, 17065471)
		if err != nil {
			log.Printf("Failed to fetch transactions for address %s: %v", address, err)
			continue
		}
	}
}

// getTransactions retrieves transactions for the specified Ethereum address within the given block number range.
func getTransactions(client *rpc.Client, address string, startBlock, endBlock uint64) ([]MyTX, error) {
	// var transactions []Transaction

	var block rpcBlock
	for blockNumber := startBlock; blockNumber <= endBlock; blockNumber++ {
		blockNumberHex := fmt.Sprintf("0x%x", blockNumber)

		err := client.Call(&block, "eth_getBlockByNumber", blockNumberHex, true)
		if err != nil {
			return nil, err
		}

		for _, v := range block.Transactions {
			if v.From == address || v.To == address {
				log.Println(v.From)
				blockNumberInt, _ := strconv.ParseInt(strings.TrimPrefix(v.BlockNumber, "0x"), 16, 64)
				valueWei, success := new(big.Int).SetString(v.Value[2:], 16)
				if !success {
					log.Printf("Failed to convert value to decimal for transaction: %s", v.TransactionHash)
					continue
				}
				valueEth := new(big.Float).Quo(new(big.Float).SetInt(valueWei), big.NewFloat(1e18))

				log.Println(valueEth, blockNumberInt)
			}
		}

	}
	return nil, nil
}
