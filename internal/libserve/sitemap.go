package libserve

import "net/url"

func NewSites() Sites {
	return map[string]Site{}
}

type Sites map[string]Site

func (m *Sites) Push(site Site) error {
	parsed, err := url.Parse(site.OriginUrl)
	if err != nil {
		return err
	}
	site.parsedOriginUrl = parsed

	if len(*m) == 0 {
		(*m)["default"] = site
	}
	(*m)[site.Host] = site

	return nil
}

func (m *Sites) getByHost(host string) Site {
	site, ok := (*m)[host]
	if ok {
		return site
	}
	site, ok = (*m)["default"]
	if ok {
		return site
	}

	return Site{}
}
