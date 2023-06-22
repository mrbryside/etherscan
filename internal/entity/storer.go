package entity

import "github.com/google/uuid"

// Storer - Entity
type Storer struct {
	ID        uuid.UUID
	Address   string
	BlockFrom uint64
	BlockTo   uint64
}
