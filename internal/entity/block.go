package entity

import "github.com/google/uuid"

// Block - Entity
type Block struct {
	ID      uuid.UUID
	Address []string
	From    uint64
	To      uint64
}
