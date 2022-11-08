package blockdownload

import (
	"context"
	"errors"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	ipld "github.com/ipfs/go-ipld-format"
	logging "github.com/ipfs/go-log/v2"
	"github.com/linguohua/titan/api"
	"github.com/linguohua/titan/api/client"
	"github.com/timtide/titan-client/util"
	"sync"
)

// multiAddrString domain name or multi address string
const multiAddrString = "/ip4/221.4.187.172/tcp/3456"

var logger = logging.Logger("titan-client/blockdownload")

type BlockGetter interface {
	// GetBlock gets the requested block.
	GetBlock(ctx context.Context, c cid.Cid) (blocks.Block, error)

	// GetBlocks The scheduler queries the corresponding
	// edge node information according to the incoming value. Each value
	// is assigned to the corresponding edge node for global optimization.
	// schedule service mapping cid to edge node.
	GetBlocks(ctx context.Context, ks []cid.Cid) <-chan blocks.Block
}

func NewBlockGetter() (BlockGetter, error) {
	bg := &blockGetter{}
	urls, err := util.TransformationMultiAddrStringsToUrl(multiAddrString)
	if err != nil {
		return nil, err
	}
	bg.schedulerURL = urls
	return bg, nil
}

type blockGetter struct {
	schedulerURL string
}

func (b *blockGetter) GetBlock(ctx context.Context, c cid.Cid) (blocks.Block, error) {
	logger.Debugf("begin get block with cid [%s]", c.String())
	if !c.Defined() {
		return nil, ipld.ErrNotFound{Cid: c}
	}
	apiScheduler, closer, err := client.NewScheduler(ctx, b.schedulerURL, nil)
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
	logger.Debug("the IP address of the edge node is ", downloadInfo.URL)
	data, err := util.RequestDataFromTitan(downloadInfo.URL, downloadInfo.Token, c)
	if err != nil {
		return nil, err
	}
	logger.Debug("block data download success")
	block, err := blocks.NewBlockWithCid(data, c)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (b *blockGetter) GetBlocks(ctx context.Context, ks []cid.Cid) <-chan blocks.Block {
	logger.Debug("start batch download block")
	ch := make(chan blocks.Block)
	go func() {
		defer close(ch)

		if len(ks) == 0 || ks == nil {
			return
		}

		apiScheduler, closer, err := client.NewScheduler(ctx, b.schedulerURL, nil)
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
				logger.Debugf("start download cid [%s] data from edge node [%s]", c.String(), df.URL)
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
					logger.Debugf("the cid [%s] data download success", block.Cid().String())
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
