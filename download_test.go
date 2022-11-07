package titan_client

import (
	"compress/gzip"
	"context"
	"github.com/ipfs/go-cid"
	"testing"
	"titan-client/util"
)

var keys = []string{
	"QmRdJtFocxnDYNYogxbaqHRgMPWytqpWK57D27fCfQ7Z3h",
	"QmWmpfDgZtGAnj3GHcH4uvmHpsz4SyA3KmUW9eyDmjjAwZ",
	"QmPvMJoaJVAio28aFNTZeTE2iwSVeFje81yM27yck4LrFX",
	"QmUudXZmoTb8vSbSUEjZzvnFXcy5skX3K4FBcVAVY9SeRo",
	"QmPaT31QjDyqZFRh2mQHXDp5dnFEUFsHLXy6NSKdeqdN7r",
	"Qma5xs8qDyMDUYyDGG9j9xry7fbd2FhDK6o8HghbfNCitr",
	"QmcynPBgaYCfGCeRYiuSfZAua1wEvFzpp2uuxohgCbHyk5",
	"QmcEa7N1QzgYKxBAVQUnHNLLN6vYvn5DPGSG3F6Dy8YZbv",
	"QmViyqSGezfNC7Nbz1BKmUgYyQGgiBspnFqbZg284VooGP",
	"QmPAEAY7imdVHSCBTYAGqb62EWNxHrju8b9v9JAo2Rhg3D",
	"QmckJc9ES1X9f6W6K8CrUZksAWUZXYvpyyJ49WTWBWxUPH",
	"QmbUHuK8CScdp6SuDbnkNdPKc2SQ6BQA4AYgjXFhh9FfQz",
	"QmXqZ2UCEGCkRmpuw37yrxBPMG33yhVDATinGEkBNuXqUx",
	"QmfKL2ivaeftVFK3oD3wjizFdLxTj3dweCN69gkGZiL6ac",
	"QmRSxVQAEL8Gv1v7c5tktv65wrRPEUPkd8ow4RwTJzV2Aq",
	"QmeWPjg9dMRpWUeJesFhGAxGDzUh7koxds6KdMHqKzywsH",
	"QmRzyWwz4dhZR4d51jeKMRgnTL3q2GQcymZhAMERfXrbrV",
	"QmY4huuAeTEFgFqyWFoEKpHLvpcUTFEPzyjLbV9vg9Sak6",
	"QmXHUgmFRhdGWMBmCs5RUaweGx45ts5akSbqk2pRLzx2U9",
	"QmXBQLm519HnsjEhD1dugCvJS7pdoDMEquuy2nTL1ctHrc",
	"QmYPAjzSMGtbR26swWbe7t7kfUFnC2pkcFtZofDUBpvBir",
	"QmbUM7QfjZa9g58VduQ6PcN8cTYBYDtCaCSZ5Db4DsiN3P",
}

func TestTitanDownloader_GetBlock(t *testing.T) {
	// set log level
	err := NewLog().SetLevel("DEBUG")
	if err != nil {
		t.Error(err.Error())
		return
	}
	c, err := cid.Decode("QmRdJtFocxnDYNYogxbaqHRgMPWytqpWK57D27fCfQ7Z3h")
	if err != nil {
		t.Error(err)
		return
	}
	block, err := NewDownloader().GetBlock(context.Background(), c)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(block.Cid())
}

func TestTitanDownloader_GetBlocks(t *testing.T) {
	// set log level
	err := NewLog().SetLevel("DEBUG")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("keys length : ", len(keys))
	ks := make([]cid.Cid, 0, len(keys))
	for _, v := range keys {
		c, err := cid.Decode(v)
		if err != nil {
			t.Error(err)
			continue
		}
		ks = append(ks, c)
	}
	t.Log("ks length : ", len(ks))
	ch := NewDownloader().GetBlocks(context.Background(), ks)
	var count int
	for {
		select {
		case b, ok := <-ch:
			if !ok {
				t.Log("channel is not ok, and download block is : ", count)
				return
			}
			count++
			t.Log(b.Cid())
		}
	}
}

func TestTitanDownloader_Download(t *testing.T) {
	// set log level
	err := NewLog().SetLevel("DEBUG")
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
	err := NewLog().SetLevel("DEBUG")
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
