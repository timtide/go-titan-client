package titan_client

type Option func(td *titanDownloader)

// WithCustomGatewayAddressOption custom set gateway url
// eg: http://127.0.0.1:5001 or https://ipfs.io/ipfs/
func WithCustomGatewayAddressOption(addr string) Option {
	return func(td *titanDownloader) {
		td.customGatewayAddr = addr
	}
}

func WithLocatorAddressOption(locatorAddr string) Option {
	return func(td *titanDownloader) {
		td.locatorAddr = locatorAddr
	}
}
