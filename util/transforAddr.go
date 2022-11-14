package util

import (
	"fmt"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"github.com/timtide/go-titan-client/common"
)

// TransformationMultiAddrStringsToUrl multi format address transfer to url
// eg: "/ip4/127.0.0.1/tcp/3456" => "http://127.0.0.1:3456/rpc/v0"
// current implementation tcp4
func TransformationMultiAddrStringsToUrl(multiAddrString string) (string, error) {
	if multiAddrString == "" {
		return "", fmt.Errorf("multi address is null")
	}

	multiAddr, err := ma.NewMultiaddr(multiAddrString)
	if err != nil {
		return "", err
	}
	pt, host, err := manet.DialArgs(multiAddr)
	if err != nil {
		return "", err
	}
	// todo tcp6 ?
	switch pt {
	case "tcp4":
		return fmt.Sprintf("%s%s%s%s", "http", "://", host, common.RPCProtocol), nil
	default:
		return "", fmt.Errorf("unkown protocol type")
	}
}
