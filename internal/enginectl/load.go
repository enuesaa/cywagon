package enginectl

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/enuesaa/cywagon/internal/service/model"
)

func (e *Engine) Load(config model.Config, workdir string) error {
	e.config = config

	if err := e.loadConfig(); err != nil {
		return err
	}
	if err := e.loadDists(workdir); err != nil {
		return err
	}
	e.listen(workdir)

	return nil
}

func (e *Engine) loadConfig() error {
	e.sitemap = make(map[string]model.Site, 0)
	for _, site := range e.config.Sites {
		e.sitemap[site.Host] = site
	}

	e.logicmap = make(map[string]model.Logic)
	for _, logic := range e.config.Logics {
		e.logicmap[logic.Name] = logic
	}

	if err := e.Log.SetLogFile(*e.config.Server.LogFile); err != nil {
		return err
	}
	e.Log.IncludeDebugLog(*e.config.Server.LogDebug)

	for _, site := range e.config.Sites {
		if site.TLSCert != nil && site.TLSKey == nil {
			return fmt.Errorf("tlskey should be specified when the site use tls")
		}
	}
	return nil
}

func (e *Engine) loadDists(workdir string) error {
	e.distmap = make(map[string]fs.FS)

	for _, site := range e.config.Sites {
		path := filepath.Join(workdir, site.Dist)
		dist, err := e.Fs.DirFS(path)
		if err != nil {
			return err
		}
		e.distmap[site.Dist] = dist

		for _, ifb := range site.Ifs {
			if ifb.Respond != nil && ifb.Respond.Dist != nil {
				path := filepath.Join(workdir, *ifb.Respond.Dist)
				dist, err := e.Fs.DirFS(path)
				if err != nil {
					return err
				}
				e.distmap[*ifb.Respond.Dist] = dist
			}
		}
	}
	return nil
}

func (e *Engine) listen(workdir string) {
	for _, site := range e.sitemap {
		e.logf("The server loaded: %s", site.Host)
		if site.TLSCert == nil {
			e.Server.Listen(site.Port)
		} else {
			tlscertpath := filepath.Join(workdir, *site.TLSCert)
			tlskeypath := filepath.Join(workdir, *site.TLSKey)
			e.Server.ListenTLS(site.Port, tlscertpath, tlskeypath)
		}
	}
}
