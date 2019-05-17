package fetcher

import "github.com/haihongs/go-ethereum/core/types"

type FetcherInsertBlockEvent struct {
	Peer  string
	Block *types.Block
}
