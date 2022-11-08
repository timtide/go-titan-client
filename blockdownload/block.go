package blockdownload

import (
	"context"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/timtide/titan-client/blockservice"
)

type Blocker interface {
	// GetBlock gets the requested block.
	GetBlock(ctx context.Context, c cid.Cid) (blocks.Block, error)

	// GetBlocks The scheduler queries the corresponding
	// edge node information according to the incoming value. Each value
	// is assigned to the corresponding edge node for global optimization.
	// schedule service mapping cid to edge node.
	GetBlocks(ctx context.Context, ks []cid.Cid) <-chan blocks.Block
}

func NewBlocker() Blocker {
	return &block{}
}

type block struct{}

func (t *block) GetBlock(ctx context.Context, c cid.Cid) (blocks.Block, error) {
	bs, err := blockservice.NewBlockService()
	if err != nil {
		return nil, err
	}
	return bs.GetBlock(ctx, c)
}

func (t *block) GetBlocks(ctx context.Context, ks []cid.Cid) <-chan blocks.Block {
	ch := make(chan blocks.Block)
	defer close(ch)
	bs, err := blockservice.NewBlockService()
	if err != nil {
		return ch
	}
	return bs.GetBlocks(ctx, ks)
}
