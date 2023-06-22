package entity

// Transaction - value object
type Transaction struct {
	Hash    string
	Address string
	Method  string
	Block   uint64
	From    string
	To      string
	Value   string
}
