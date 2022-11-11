package util

import (
	"context"
	"errors"
	"fmt"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/linguohua/titan/api"
	"github.com/linguohua/titan/api/client"
	"github.com/timtide/titan-client/common"
	http2 "github.com/timtide/titan-client/util/http"
	"strings"
	"sync"
)

var logger = logging.Logger("titan-client/util")

const defaultScheduleAddress = "http://39.108.143.56:5000/rpc/v0"

// DataGetter from titan or common gateway or local gateway to get data
type DataGetter interface {
	GetDataFromTitanByCid(ctx context.Context, c cid.Cid) ([]byte, error)
	GetDataFromTitanOrGatewayByCid(ctx context.Context, customGatewayURL string, c cid.Cid) ([]byte, error)
	GetBlockFromTitanOrGatewayByCids(ctx context.Context, customGatewayURL string, ks []cid.Cid) <-chan blocks.Block
	GetBlockFromTitanByCids(ctx context.Context, ks []cid.Cid) <-chan blocks.Block
}

type dataGetter struct {
	schedulerURL string
}

func NewDataGetter() DataGetter {
	return &dataGetter{
		schedulerURL: defaultScheduleAddress,
	}
}

func (d *dataGetter) GetDataFromTitanByCid(ctx context.Context, c cid.Cid) ([]byte, error) {
	apiScheduler, closer, err := client.NewScheduler(ctx, d.schedulerURL, nil)
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
	data, err := d.getDataFromEdgeNode(downloadInfo.URL, downloadInfo.Token, c)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *dataGetter) GetDataFromTitanOrGatewayByCid(ctx context.Context, customGatewayURL string, c cid.Cid) ([]byte, error) {
	data, err := d.GetDataFromTitanByCid(ctx, c)
	if err != nil && !strings.Contains(strings.ToLower(err.Error()), "not found") {
		logger.Warn(err.Error())
		return nil, err
	}
	if data == nil {
		data, err = d.getDataFromCommonGateway(customGatewayURL, c)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
	}
	return data, nil
}

func (d *dataGetter) GetBlockFromTitanOrGatewayByCids(ctx context.Context, customGatewayURL string, ks []cid.Cid) <-chan blocks.Block {
	ch := make(chan blocks.Block)
	go func() {
		defer close(ch)

		if len(ks) == 0 || ks == nil {
			return
		}

		apiScheduler, closer, err := client.NewScheduler(ctx, d.schedulerURL, nil)
		if err != nil {
			return
		}
		defer closer()

		cs := make([]string, 0, len(ks))
		for _, v := range ks {
			cs = append(cs, v.String())
		}

		cidToEdges, err := apiScheduler.GetDownloadInfosWithBlocks(ctx, cs)
		if err != nil {
			logger.Error(err.Error())
			return
		}
		mp := UniformMapping(cidToEdges)
		for _, v := range ks {
			if _, ok := mp[v.String()]; !ok {
				mp[v.String()] = api.DownloadInfo{}
			}
		}

		var wg sync.WaitGroup
		for k, v := range mp {
			kk := k
			key, err := cid.Decode(kk)
			if err != nil {
				continue
			}
			value := v

			wg.Add(1)
			go func(cc context.Context, c cid.Cid, df api.DownloadInfo) {
				defer wg.Done()
				data, err := d.getDataFromEdgeNode(df.URL, df.Token, c)
				if err != nil && !strings.Contains(strings.ToLower(err.Error()), "not found") {
					logger.Error(err.Error())
					return
				}
				if data == nil {
					data, err = d.getDataFromCommonGateway(customGatewayURL, c)
					if err != nil {
						logger.Error(err.Error())
						return
					}
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

func (d *dataGetter) GetBlockFromTitanByCids(ctx context.Context, ks []cid.Cid) <-chan blocks.Block {
	ch := make(chan blocks.Block)
	go func() {
		defer close(ch)

		if len(ks) == 0 || ks == nil {
			return
		}

		apiScheduler, closer, err := client.NewScheduler(ctx, d.schedulerURL, nil)
		if err != nil {
			logger.Error(err.Error())
			return
		}
		defer closer()

		cs := make([]string, 0, len(ks))
		for _, v := range ks {
			cs = append(cs, v.String())
		}

		cidToEdges, err := apiScheduler.GetDownloadInfosWithBlocks(ctx, cs)
		if err != nil {
			logger.Error(err.Error())
			return
		}
		mp := UniformMapping(cidToEdges)
		for _, v := range ks {
			if _, ok := mp[v.String()]; !ok {
				mp[v.String()] = api.DownloadInfo{}
			}
		}
		var wg sync.WaitGroup
		for k, v := range mp {
			kk := k
			key, err := cid.Decode(kk)
			if err != nil {
				continue
			}
			value := v
			wg.Add(1)
			go func(cc context.Context, c cid.Cid, df api.DownloadInfo) {
				defer wg.Done()
				data, err := d.getDataFromEdgeNode(df.URL, df.Token, c)
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

// getDataFromEdgeNode connect Titan net by http get method
func (d *dataGetter) getDataFromEdgeNode(host, token string, cid cid.Cid) ([]byte, error) {
	if host == "" || token == "" {
		return nil, fmt.Errorf("not found target host")
	}
	logger.Debugf("get data from titan edge node [%s]", host)
	url := fmt.Sprintf("%s%s%s", host, "?cid=", cid.String())
	return http2.Get(url, token, common.AppName)
}

func (d *dataGetter) getDataFromCommonGateway(customGatewayURL string, c cid.Cid) ([]byte, error) {
	if customGatewayURL == "" {
		return nil, fmt.Errorf("not found target host")
	}
	logger.Debugf("get data from common gateway with cid [%s]", c.String())
	url := fmt.Sprintf("%s%s", customGatewayURL, c.String())
	return http2.PostFromGateway(url)
}
