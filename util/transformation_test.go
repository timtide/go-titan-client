package util

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/linguohua/titan/api"
	"testing"
)

var di1 = api.DownloadInfo{
	URL: "url1",
}
var di2 = api.DownloadInfo{
	URL: "url2",
}
var di3 = api.DownloadInfo{
	URL: "url3",
}
var di4 = api.DownloadInfo{
	URL: "url4",
}
var di5 = api.DownloadInfo{
	URL: "url5",
}
var di6 = api.DownloadInfo{
	URL: "url6",
}
var di7 = api.DownloadInfo{
	URL: "url7",
}
var di8 = api.DownloadInfo{
	URL: "url8",
}
var di9 = api.DownloadInfo{
	URL: "url9",
}
var di10 = api.DownloadInfo{
	URL: "url10",
}

var data1 = map[string][]api.DownloadInfo{
	"cid1":  []api.DownloadInfo{di1, di2},
	"cid2":  []api.DownloadInfo{di1, di3, di4},
	"cid3":  []api.DownloadInfo{di1, di2},
	"cid4":  []api.DownloadInfo{di1},
	"cid5":  []api.DownloadInfo{di1, di2, di3, di4},
	"cid6":  []api.DownloadInfo{di1, di2, di3, di5, di6, di7, di8, di9, di10},
	"cid7":  []api.DownloadInfo{di1, di2, di3, di7, di8},
	"cid8":  []api.DownloadInfo{di1, di2, di3, di9, di10},
	"cid9":  []api.DownloadInfo{di1, di2, di3},
	"cid10": []api.DownloadInfo{di1, di2, di5, di6},
}

var data2 = map[string][]api.DownloadInfo{
	"cid1":  []api.DownloadInfo{di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid2":  []api.DownloadInfo{di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid3":  []api.DownloadInfo{di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid4":  []api.DownloadInfo{di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid5":  []api.DownloadInfo{di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid6":  []api.DownloadInfo{di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid7":  []api.DownloadInfo{di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid8":  []api.DownloadInfo{di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid9":  []api.DownloadInfo{di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid10": []api.DownloadInfo{di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
}

func TestTransfer(t *testing.T) {
	err := logging.SetLogLevel("titan-client/util", "DEBUG")
	if err != nil {
		t.Error(err.Error())
		return
	}
	res := UniformMapping(data2)
	t.Log(res)
}
