package titan_client

import (
	"compress/gzip"
	"context"
	"github.com/ipfs/go-cid"
	"testing"
	"titan-client/util"
)

func TestDownload(t *testing.T) {
	ctx := context.Background()
	c, err := cid.Decode("QmajjF2D13CsreihRsWsDicraMh2nXFmBLXKoF5MNBRAyL")
	if err != nil {
		t.Error(err)
		return
	}
	err = Download(ctx, c, false, gzip.NoCompression, "/Users/jason/data/tmp")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("download success")
}

func TestGetReader(t *testing.T) {
	c, err := cid.Decode("QmajjF2D13CsreihRsWsDicraMh2nXFmBLXKoF5MNBRAyL")
	if err != nil {
		t.Error(err)
		return
	}
	reader, err := GetReader(context.Background(), c, false, gzip.NoCompression)
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
