package service

import (
	"log"

	"github.com/mrbryside/internal/domain/reader"
	"github.com/mrbryside/internal/domain/reader/ethereum"
)

type ITransactionService interface {
	GetTransactions(address []string, blockFrom, blockTo uint64) error
}

type TransactionService struct {
	reader reader.BlockFetcher
}

func NewTransactionService(ec ethereum.EthereumClient) ITransactionService {
	return TransactionService{
		reader: ec,
	}
}

func (ts TransactionService) GetTransactions(address []string, blockFrom, blockTo uint64) error {

	rAgg := reader.NewReader(address, blockFrom, blockTo)
	err := ts.reader.FetchTransaction(rAgg)
	if err != nil {
		return err
	}
	log.Println(rAgg.TransactionsVo())
	return nil
}
