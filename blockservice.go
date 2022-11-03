package titan_client

import (
	"context"
	"errors"
	"fmt"
	blocks "github.com/ipfs/go-block-format"
	bserv "github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	exchange "github.com/ipfs/go-ipfs-exchange-interface"
	ipld "github.com/ipfs/go-ipld-format"
	"github.com/linguohua/titan/api"
	"github.com/linguohua/titan/api/client"
	"sync"
	"titan-client/util"
)

// MultiAddrString domain name or multi address string
const MultiAddrString = "/ip4/221.4.187.172/tcp/3456"

type blockService struct {
	schedulerURL string
}

// NewBlockService creates a BlockService with given datastore instance.
func NewBlockService() (bserv.BlockService, error) {
	bs := &blockService{}
	urls, err := util.TransformationMultiAddrStringsToUrl(MultiAddrString)
	if err != nil {
		return nil, err
	}
	bs.schedulerURL = urls
	return bs, nil
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
	if !c.Defined() {
		return nil, ipld.ErrNotFound{Cid: c}
	}
	apiScheduler, closer, err := client.NewScheduler(ctx, s.schedulerURL, nil)
	if err != nil {
		return nil, err
	}
	defer closer()
	downloadInfo, err := apiScheduler.GetDownloadInfoWithBlock(ctx, c.String())
	if err != nil {
		return nil, err
	}
	if downloadInfo.URL == "" || downloadInfo.Token == "" {
		return nil, errors.New("data not fount")
	}
	data, err := util.RequestDataFromTitan(downloadInfo.URL, downloadInfo.Token, c)
	if err != nil {
		return nil, err
	}
	block, err := blocks.NewBlockWithCid(data, c)
	if err != nil {
		return nil, err
	}
	return block, nil
}

// GetBlocks gets a list of blocks asynchronously and returns through
// the returned channel.
// NB: No guarantees are made about order.
func (s *blockService) GetBlocks(ctx context.Context, ks []cid.Cid) <-chan blocks.Block {
	ch := make(chan blocks.Block)
	go func() {
		defer close(ch)

		if len(ks) == 0 || ks == nil {
			return
		}

		apiScheduler, closer, err := client.NewScheduler(ctx, s.schedulerURL, nil)
		if err != nil {
			logger.Error(err.Error())
			return
		}
		defer closer()

		cs := make([]string, 0, len(ks))
		for _, v := range ks {
			cs = append(cs, v.String())
		}

		mp, err := apiScheduler.GetDownloadInfoWithBlocks(ctx, cs)
		if err != nil {
			logger.Error(err.Error())
			return
		}

		var wg sync.WaitGroup
		for k, v := range mp {
			kk := k
			key, err := cid.Decode(kk)
			if err != nil {
				logger.Warn(err.Error())
				continue
			}
			value := v
			if value.URL == "" || value.Token == "" {
				continue
			}

			wg.Add(1)
			go func(cc context.Context, c cid.Cid, df api.DownloadInfo) {
				defer wg.Done()
				data, err := util.RequestDataFromTitan(df.URL, df.Token, c)
				if err != nil {
					logger.Error(err.Error())
					return
				}
				block, err := blocks.NewBlockWithCid(data, c)
				if err != nil {
					logger.Error(err.Error())
					return
				}
				select {
				case ch <- block:
					return
				case <-cc.Done():
					return
				}
			}(ctx, key, value)
		}
		wg.Wait()
	}()

	return ch
}

// DeleteBlock deletes a block in the blockservice from the datastore
func (s *blockService) DeleteBlock(ctx context.Context, c cid.Cid) error {
	return fmt.Errorf("%s", "not implemented")
}

func (s *blockService) Close() error {
	return fmt.Errorf("%s", "not implemented")
}
