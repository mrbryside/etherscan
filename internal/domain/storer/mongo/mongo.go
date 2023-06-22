package mongo

import (
	"context"
	"log"

	"github.com/mrbryside/internal/core/db/mongodb"
	"github.com/mrbryside/internal/domain/storer"
)

type MongoRepository struct {
	pocketCollection mongodb.CollectionWrapper
	dbClient         mongodb.ClientWrapper
}

func NewMongoRepository(pocketColl mongodb.CollectionWrapper, c mongodb.ClientWrapper) storer.StorerRepository {
	return MongoRepository{pocketCollection: pocketColl, dbClient: c}
}

type TransactionModel struct {
	EthereumAddress string `bson:"ethereum_address"`
	TransactionHash string `bson:"transaction_hash"`
	BlockNumber     uint64 `bson:"block_number"`
	From            string `bson:"from"`
	To              string `bson:"to"`
	Value           string `bson:"value"`
}

func toInternalTransactionsModel(s *storer.Storer) []interface{} {
	tms := make([]interface{}, 0)
	for _, t := range s.TransactionsVo() {
		tm := TransactionModel{
			EthereumAddress: s.StorerEntity().Address,
			TransactionHash: t.Hash,
			BlockNumber:     t.Block,
			From:            t.From,
			To:              t.To,
			Value:           t.Value,
		}
		tms = append(tms, tm)
	}
	return tms
}

func (mr MongoRepository) Save(s *storer.Storer) error {
	trans := toInternalTransactionsModel(s)
	_, err := mr.pocketCollection.InsertMany(context.Background(), trans)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
