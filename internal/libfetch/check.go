package libfetch

func (f *Fetcher) CheckHTTP(url string, matcher string) error {
	status := f.FetchHTTP(url)

	if status == matcher {
		return nil
	}

	return nil
}

func (f *Fetcher) CheckTCP(address string) error {
	return f.ConnectTCP(address)
}
