package libserve

import (
	"fmt"
	"net/url"
)

type Site struct {
	Host            string // Example: `example.com`
	OriginUrl       string // Example: `https://example.com`
	Handler         Handler
	parsedOriginUrl *url.URL
	Cache           bool
}
type Handler func(*HandlerResponse, Next, HandlerRequest) error
type Next func(HandlerRequest) HandlerResponse

type HandlerRequest struct {
	Path string
}
type HandlerResponse struct {
	Status int
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
	return (*m)["default"]
}

var ErrSitesNeedAtLeast1SiteDef = fmt.Errorf("sites need at least 1 def")

func (m *Sites) Validate() error {
	if _, ok := (*m)["default"]; !ok {
		return ErrSitesNeedAtLeast1SiteDef
	}
	return nil
}
