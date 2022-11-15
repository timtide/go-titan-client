package blockdownload

type Option func(bg *blockGetter)

// WithLocatorUrlOption set your favorite locator url, address or domain name
func WithLocatorUrlOption(locatorAddr string) Option {
	return func(dg *blockGetter) {
		dg.locatorAddr = locatorAddr
	}
}
