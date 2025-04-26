package enginectl

import (
	"io/fs"
	"strings"

	"github.com/enuesaa/cywagon/internal/libserve"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) handleIfBlocks(c *libserve.Context, ifs []model.If, distmap map[string]fs.FS, logicmap map[string]model.Logic) *libserve.Response {
	for _, ifb := range ifs {
		if ifb.Logic != nil {
			logic, ok := logicmap[*ifb.Logic]
			if !ok {
				continue
			}
			if res := e.handleIfBlocks(c, logic.Ifs, distmap, logicmap); res != nil {
				return res
			}
		}

		if e.shouldCheckCondStr(ifb.Path, ifb.PathIn, ifb.PathNot, ifb.PathNotIn) {
			if !e.matchCondPath(c.Path, ifb.Path, ifb.PathIn, ifb.PathNot, ifb.PathNotIn) {
				continue
			}
		}
		if e.shouldCheckCondStrMap(ifb.Headers, ifb.HeadersIn, ifb.HeadersNot, ifb.HeadersNotIn) {
			if !e.matchCondStrMap(c.Headers, ifb.Headers, ifb.HeadersIn, ifb.HeadersNot, ifb.HeadersNotIn) {
				continue
			}
		}
		if ifb.Rewrite != nil {
			if ifb.Rewrite.Path != nil {
				c.Path = e.calcRewritePath(c.Path, *ifb.Rewrite.Path)
			}
		}
		if ifb.Respond != nil {
			for key, value := range ifb.Respond.Headers {
				c.ResHeader(key, value)
			}
			if ifb.Respond.Body != nil {
				c.ResBody(c.Path, strings.NewReader(*ifb.Respond.Body))
			}
			if ifb.Respond.Status != nil {
				c.ResStatusPrefer(*ifb.Respond.Status)
			}
			if ifb.Respond.Dist != nil {
				dist := distmap[*ifb.Respond.Dist]
				path := strings.TrimPrefix(c.Path, "/")

				f, err := dist.Open(path)
				if err != nil {
					return c.Resolve(404)
				}
				if err := c.ResBody(path, f); err != nil {
					return c.Resolve(404)
				}
			}
			return c.Resolve(200)
		}
	}
	return nil
}
