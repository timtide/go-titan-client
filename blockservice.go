package titan_client

import (
	"context"
	"fmt"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	exchange "github.com/ipfs/go-ipfs-exchange-interface"
	"github.com/timtide/titan-client/blockdownload"
)

type blockService struct{}

// NewBlockService creates a BlockService with given datastore instance.
func NewBlockService() *blockService {
	return &blockService{}
}

// Blockstore returns the blockstore behind this blockservice.
func (s *blockService) Blockstore() blockstore.Blockstore {
	logger.Error("not implemented")
	return nil
}

// Exchange returns the exchange behind this blockservice.
func (s *blockService) Exchange() exchange.Interface {
	logger.Error("not implemented")
	return nil
}

// AddBlock adds a particular block to the service, Putting it into the datastore.
func (s *blockService) AddBlock(ctx context.Context, o blocks.Block) error {
	return fmt.Errorf("%s", "not implemented")
}

func (s *blockService) AddBlocks(ctx context.Context, bs []blocks.Block) error {
	return fmt.Errorf("%s", "not implemented")
}

// GetBlock retrieves a particular block from the service,
// Getting it from the datastore using the key (hash).
func (s *blockService) GetBlock(ctx context.Context, c cid.Cid) (blocks.Block, error) {
	bg, err := blockdownload.NewBlockGetter()
	if err != nil {
		return nil, err
	}
	return bg.GetBlock(ctx, c)
}

// GetBlocks gets a list of blocks asynchronously and returns through
// the returned channel.
// NB: No guarantees are made about order.
func (s *blockService) GetBlocks(ctx context.Context, ks []cid.Cid) <-chan blocks.Block {
	ch := make(chan blocks.Block)
	defer close(ch)
	bg, err := blockdownload.NewBlockGetter()
	if err != nil {
		logger.Error(err.Error())
		return ch
	}
	return bg.GetBlocks(ctx, ks)
}

// DeleteBlock deletes a block in the blockservice from the datastore
func (s *blockService) DeleteBlock(ctx context.Context, c cid.Cid) error {
	return fmt.Errorf("%s", "not implemented")
}

func (s *blockService) Close() error {
	return fmt.Errorf("%s", "not implemented")
}
