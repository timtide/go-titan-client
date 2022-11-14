package titan_client

import (
	"compress/gzip"
	"context"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/timtide/go-titan-client/common"
	"github.com/timtide/go-titan-client/util"
	"testing"
)

func TestNewDownloader(t *testing.T) {
	downloader := NewDownloader(WithCustomGatewayUrlOption("http://127.0.0.1:5001"))
	t.Log(downloader)
}

func TestTitanDownloader_Download(t *testing.T) {
	err := logging.SetLogLevel(common.AppName, "DEBUG")
	if err != nil {
		t.Error(err.Error())
		return
	}
	err = logging.SetLogLevel("titan-client/util", "DEBUG")
	if err != nil {
		t.Error(err.Error())
		return
	}
	ctx := context.Background()
	c, err := cid.Decode("bafybeiglv5lkp2uwrhpwtfixn2gtu7w62yorckmfxac3jphys2q267plwa")
	if err != nil {
		t.Error(err)
		return
	}
	err = NewDownloader(WithCustomGatewayUrlOption("http://127.0.0.1:5001")).Download(ctx, c, false, gzip.NoCompression, "/Users/jason/data/tmp")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("download success")
}

func TestTitanDownloader_GetReader(t *testing.T) {
	// set log level
	err := logging.SetLogLevel(common.AppName, "DEBUG")
	if err != nil {
		t.Error(err.Error())
		return
	}
	c, err := cid.Decode("QmajjF2D13CsreihRsWsDicraMh2nXFmBLXKoF5MNBRAyL")
	if err != nil {
		t.Error(err)
		return
	}
	reader, err := NewDownloader().GetReader(context.Background(), c, false, gzip.NoCompression)
	if err != nil {
		t.Error(err)
		return
	}
	defer reader.Close()

	ow := util.Writer{
		Archive:     false,
		Compression: gzip.NoCompression,
	}

	err = ow.Write(reader, "/Users/jason/data/tmp")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("download success")
}
