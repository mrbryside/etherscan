package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=./collection_wrapper.go -destination=../../generated/mockgen/mongo_collection_wrapper/mongo_collection_wrapper.go -package=mockMongoCollectionWrapper
type CollectionWrapper interface {
	UpdateOne(ctx context.Context, filter interface{}, update interface{},
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	InsertMany(ctx context.Context, documents []interface{},
		opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
}

type Collection struct {
	*mongo.Collection
}

func (c Collection) InsertMany(ctx context.Context, documents []interface{},
	opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return c.Collection.InsertMany(ctx, documents, opts...)
}

func (c Collection) UpdateOne(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return c.Collection.UpdateOne(ctx, filter, update, opts...)
}
