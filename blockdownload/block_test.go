package blockdownload

import (
	"context"
	"github.com/ipfs/go-cid"
	"testing"
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
	"QmajjF2D13CsreihRsWsDicraMh2nXFmBLXKoF5MNBRAyL",
}

func TestBlockGetter_GetBlock(t *testing.T) {
	c, err := cid.Decode("QmTp2hEo8eXRp6wg7jXv1BLCMh5a4F3B7buAUZNZUu772j")
	if err != nil {
		t.Error(err)
		return
	}
	block, err := NewBlockGetter().GetBlock(context.Background(), c)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(block.Cid())
}

func TestBlock_GetBlocks(t *testing.T) {
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
	ch := NewBlockGetter().GetBlocks(context.Background(), ks)
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
