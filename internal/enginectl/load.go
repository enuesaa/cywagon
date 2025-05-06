package enginectl

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/enuesaa/cywagon/internal/service/model"
)

type Site struct {
	Dist   fs.FS
	Config model.Site
}

func (e *Engine) loadSites(config model.Config) {
	e.sitemap = make(map[string]model.Site, 0)
	for _, site := range config.Sites {
		e.sitemap[site.Host] = site
	}
}

func (e *Engine) loadLogics(config model.Config) {
	e.logicmap = make(map[string]model.Logic)
	for _, logic := range config.Logics {
		e.logicmap[logic.Name] = logic
	}
}

func (e *Engine) loadDists(config model.Config, workdir string) error {
	e.distmap = make(map[string]fs.FS)
	for _, site := range config.Sites {
		dist, err := e.loadFS(workdir, site.Dist)
		if err != nil {
			return err
		}
		e.distmap[site.Dist] = dist

		for _, ifb := range site.Ifs {
			if ifb.Respond != nil && ifb.Respond.Dist != nil {
				distpath := *ifb.Respond.Dist
				dist, err := e.loadFS(workdir, distpath)
				if err != nil {
					return err
				}
				e.distmap[distpath] = dist
			}
		}
	}
	return nil
}

func (e *Engine) loadFS(workdir string, distpath string) (fs.FS, error) {
	path := filepath.Join(workdir, distpath)
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	return os.DirFS(path), nil
}
