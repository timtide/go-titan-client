package titan_client

import (
	"compress/gzip"
	"context"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/timtide/go-titan-client/util"
	"testing"
)

func TestNewDownloader(t *testing.T) {
	downloader := NewDownloader(WithCustomGatewayAddressOption("http://127.0.0.1:5001"), WithLocatorAddressOption(""))
	t.Log(downloader)
}

func TestTitanDownloader_Download(t *testing.T) {
	err := logging.SetLogLevel("titan-client/util", "DEBUG")
	if err != nil {
		t.Error(err.Error())
		return
	}
	ctx := context.Background()
	c, err := cid.Decode("QmUbaDBz6YKn3dVzoKrLDyupMmyWk5am2QSdgfKsU1RN3N")
	if err != nil {
		t.Error(err)
		return
	}
	downloader := NewDownloader(WithCustomGatewayAddressOption("http://127.0.0.1:5001"))
	err = downloader.Download(ctx, c, false, gzip.NoCompression, "./titan.mp4")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("download success")
}

func TestTitanDownloader_GetReader(t *testing.T) {
	c, err := cid.Decode("QmPgaP4SiadmrtFzEVY5aGTCRou5vbMDJCgEaJwuN9Lk4H")
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

	err = ow.Write(reader, "./")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("download success")
}
