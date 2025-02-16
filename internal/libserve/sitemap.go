package libserve

type SiteMap struct {
	items map[string]Site
}

func (m *SiteMap) getByHost(host string) Site {
	site, ok := m.items[host]
	if ok {
		return site
	}
	site, ok = m.items["default"]
	if ok {
		return site
	}

	return Site{}
}
