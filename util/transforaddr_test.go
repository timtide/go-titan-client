package util

import "testing"

func TestTransformationMultiAddrStringsToUrl(t *testing.T) {
	eg := "/ip4/127.0.0.1/tcp/3456"
	url, err := TransformationMultiAddrStringsToUrl(eg)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(url)
	return
}
