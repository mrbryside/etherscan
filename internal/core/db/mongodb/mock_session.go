package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockSessionSuccess struct {
	mongo.Session
}

func NewMockSession() MockSessionSuccess {
	return MockSessionSuccess{}
}

func (s MockSessionSuccess) StartTransaction(...*options.TransactionOptions) error {
	return nil
}

func (s MockSessionSuccess) CommitTransaction(context.Context) error {
	return nil
}

func (s MockSessionSuccess) EndSession(context.Context) {}

type MockSessionFailedTransaction struct {
	// Add any necessary fields to the mockSession struct
	mongo.Session
}

func NewMockSessionFailedTransaction() MockSessionFailedTransaction {
	return MockSessionFailedTransaction{}
}

func (s MockSessionFailedTransaction) StartTransaction(...*options.TransactionOptions) error {
	return errors.New("error start transaction")
}

func (s MockSessionFailedTransaction) CommitTransaction(context.Context) error {
	return nil
}

func (s MockSessionFailedTransaction) EndSession(context.Context) {}

type MockSessionFailedCommit struct {
	// Add any necessary fields to the mockSession struct
	mongo.Session
}

func NewMockSessionFailedCommit() MockSessionFailedCommit {
	return MockSessionFailedCommit{}
}

func (s MockSessionFailedCommit) StartTransaction(...*options.TransactionOptions) error {
	return nil
}

func (s MockSessionFailedCommit) CommitTransaction(context.Context) error {
	return errors.New("error commit transaction")
}

func (s MockSessionFailedCommit) EndSession(context.Context) {}
