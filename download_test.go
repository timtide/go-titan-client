package titan_client

import (
	"compress/gzip"
	"context"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/timtide/titan-client/common"
	"github.com/timtide/titan-client/util"
	"testing"
)

func TestTitanDownloader_Download(t *testing.T) {
	err := logging.SetLogLevel(common.AppName, "DEBUG")
	if err != nil {
		t.Error(err.Error())
		return
	}
	ctx := context.Background()
	c, err := cid.Decode("QmajjF2D13CsreihRsWsDicraMh2nXFmBLXKoF5MNBRAyL")
	if err != nil {
		t.Error(err)
		return
	}
	err = NewDownloader().Download(ctx, c, false, gzip.NoCompression, "/Users/jason/data/tmp")
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
