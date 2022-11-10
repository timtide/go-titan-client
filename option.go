package titan_client

type Option func(td *titanDownloader)

// WithCustomGatewayUrlOption custom set gateway url
// eg: http://127.0.0.1:5001 or https://ipfs.io/ipfs/
func WithCustomGatewayUrlOption(url string) Option {
	return func(td *titanDownloader) {
		td.customGatewayURL = url
	}
}
