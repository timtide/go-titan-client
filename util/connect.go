package util

import (
	"fmt"
	"github.com/ipfs/go-cid"
	"github.com/timtide/titan-client/common"
	http2 "github.com/timtide/titan-client/util/http"
)

// RequestDataFromTitan connect Titan net by http get method
func RequestDataFromTitan(host, token string, cid cid.Cid) ([]byte, error) {
	url := fmt.Sprintf("%s%s%s", host, "?cid=", cid.String())
	return http2.Get(url, token, common.AppName)
}
