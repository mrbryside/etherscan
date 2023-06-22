package reader

type BlockFetcher interface {
	FetchTransaction(r *Reader) error
}
