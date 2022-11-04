package titan_client

import (
	"context"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"titan-client/util"
)

type TitanClient interface {
	// GetBlock gets the requested block.
	GetBlock(ctx context.Context, c cid.Cid) (blocks.Block, error)

	// GetBlocks The scheduler queries the corresponding
	// edge node information according to the incoming value. Each value
	// is assigned to the corresponding edge node for global optimization.
	// schedule service mapping cid to edge node.
	GetBlocks(ctx context.Context, ks []cid.Cid) <-chan blocks.Block
}

// NewTitanClient create interface TitanClient
func NewTitanClient() (TitanClient, error) {
	bs := &blockService{}
	urls, err := util.TransformationMultiAddrStringsToUrl(MultiAddrString)
	if err != nil {
		return nil, err
	}
	bs.schedulerURL = urls
	return bs, nil
}
