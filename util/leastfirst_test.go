package util

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/linguohua/titan/api"
	"testing"
)

var di1 = api.DownloadInfoResult{
	URL: "url1",
}
var di2 = api.DownloadInfoResult{
	URL: "url2",
}
var di3 = api.DownloadInfoResult{
	URL: "url3",
}
var di4 = api.DownloadInfoResult{
	URL: "url4",
}
var di5 = api.DownloadInfoResult{
	URL: "url5",
}
var di6 = api.DownloadInfoResult{
	URL: "url6",
}
var di7 = api.DownloadInfoResult{
	URL: "url7",
}
var di8 = api.DownloadInfoResult{
	URL: "url8",
}
var di9 = api.DownloadInfoResult{
	URL: "url9",
}
var di10 = api.DownloadInfoResult{
	URL: "url10",
}

var data1 = map[string][]api.DownloadInfoResult{
	"cid1":  {di1, di2},
	"cid2":  {di1, di3, di4},
	"cid3":  {di1, di2},
	"cid4":  {di1},
	"cid5":  {di1, di2, di3, di4},
	"cid6":  {di1, di2, di3, di5, di6, di7, di8, di9, di10},
	"cid7":  {di1, di2, di3, di7, di8},
	"cid8":  {di1, di2, di3, di9, di10},
	"cid9":  {di1, di2, di3},
	"cid10": {di1, di2, di5, di6},
}

var data2 = map[string][]api.DownloadInfoResult{
	"cid1":  {di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid2":  {di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid3":  {di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid4":  {di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid5":  {di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid6":  {di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid7":  {di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid8":  {di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid9":  {di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
	"cid10": {di1, di2, di3, di4, di5, di6, di7, di8, di9, di10},
}

func TestTransfer(t *testing.T) {
	err := logging.SetLogLevel("titan-client/util", "DEBUG")
	if err != nil {
		t.Error(err.Error())
		return
	}
	testData := []map[string][]api.DownloadInfoResult{
		data1,
		data2,
	}
	for _, v := range testData {
		t.Log(v)
		res := UniformMapping(v)
		t.Log(res)
	}
}
