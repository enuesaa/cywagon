package libserve

// import (
// 	"net/http"

// 	"github.com/enuesaa/cywagon/internal/infra"
// )

// type Transport struct {
// 	infra.Container
// 	Sites      Sites
// 	Middleware Middleware
// }

// func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
// 	site := t.Sites.getByHost(req.Host)

// 	res, err := t.Middleware.Handle(site, req)
// 	if err != nil {
// 		t.Log.Error(err)
// 		return res, err
// 	}
// 	t.Log.Info("[http] %d %s %s %s", res.StatusCode, req.Method, site.Host, req.URL.Path)

// 	return res, nil
// }
