package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=./client_wrapper.go -destination=../../generated/mockgen/mongo_client_wrapper/mongo_client_wrapper.go -package=mockMongoClientWrapper
type ClientWrapper interface {
	StartSession(opts ...*options.SessionOptions) (mongo.Session, error)
}

type clientWrapper struct {
	*mongo.Client
}

func NewClientWrapper(c *mongo.Client) ClientWrapper {
	return &clientWrapper{c}
}

func (w *clientWrapper) StartSession(opts ...*options.SessionOptions) (mongo.Session, error) {
	session, err := w.Client.StartSession(opts...)
	if err != nil {
		return nil, err
	}
	return session, nil
}
