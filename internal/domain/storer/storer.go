package storer

import (
	"github.com/google/uuid"
	"github.com/mrbryside/internal/entity"
)

type Storer struct {
	storer       entity.Storer
	transactions []entity.Transaction
}

func NewStorer(address string, blockFrom, blockTo uint64) *Storer {
	return &Storer{
		storer: entity.Storer{
			ID:        uuid.New(),
			Address:   address,
			BlockFrom: blockFrom,
			BlockTo:   blockTo,
		},
		transactions: make([]entity.Transaction, 0),
	}
}

func (s *Storer) StorerEntity() entity.Storer {
	return s.storer
}

func (s *Storer) TransactionsVo() []entity.Transaction {
	return s.transactions
}

func (s *Storer) AddTransaction(t entity.Transaction) {
	s.transactions = append(s.transactions, t)
}
