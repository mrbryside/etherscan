package mongo

import (
	"testing"

	"github.com/mrbryside/internal/core/db/mongodb"
	"github.com/mrbryside/internal/domain/storer"
	"github.com/mrbryside/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestInsertTransactionSuccess(t *testing.T) {
	client := mongodb.NewMongoClient("mongodb+srv://connextor:ConnextorPassword1@connextor.lfncyii.mongodb.net/?retryWrites=true&w=majority")
	clientWrap := mongodb.NewClientWrapper(client)
	db := client.Database("ethereum")
	coll := db.Collection("transactions")
	collWrap := mongodb.Collection{Collection: coll}

	repo := NewMongoRepository(collWrap, clientWrap)
	s := storer.NewStorer("1234", 1, 2)
	s.AddTransaction(entity.Transaction{Hash: "12345"})
	s.AddTransaction(entity.Transaction{Hash: "1234584"})
	err := repo.Save(s)

	assert.Nil(t, err)
}
