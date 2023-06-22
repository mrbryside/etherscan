package reader

import (
	"github.com/google/uuid"
	"github.com/mrbryside/internal/entity"
)

// Reader - aggregate for comunicate with fetcher and service
type Reader struct {
	block        entity.Block         // root entity
	transactions []entity.Transaction // value object
}

func NewReader(address []string, from, to uint64) *Reader {
	block := entity.Block{
		ID:      uuid.New(),
		Address: address,
		From:    from,
		To:      to,
	}

	return &Reader{
		block:        block,
		transactions: make([]entity.Transaction, 0),
	}
}

func (r *Reader) BlockEntity() entity.Block {
	return r.block
}

func (r *Reader) TransactionsVo() []entity.Transaction {
	return r.transactions
}

func (r *Reader) AddTransaction(t entity.Transaction) {
	r.transactions = append(r.transactions, t)
}

func (r *Reader) AddAddress(address string) {
	r.block.Address = append(r.block.Address, address)
}
