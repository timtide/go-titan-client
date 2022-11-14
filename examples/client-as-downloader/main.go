package main

import (
	"compress/gzip"
	"context"
	"fmt"
	"github.com/ipfs/go-cid"
	"github.com/timtide/go-titan-client"
)

func main() {
	ctx := context.Background()
	c, err := cid.Decode("bafybeiglv5lkp2uwrhpwtfixn2gtu7w62yorckmfxac3jphys2q267plwa")
	if err != nil {
		panic(err.Error())
	}
	d := titan_client.NewDownloader(titan_client.WithCustomGatewayUrlOption("http://127.0.0.1:5001"))
	err = d.Download(ctx, c, false, gzip.NoCompression, "/Users/jason/data/tmp")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("download success")
}
