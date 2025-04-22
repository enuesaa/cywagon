package service

import (
	"path/filepath"
	"strings"

	"github.com/enuesaa/cywagon/internal/infra"
	"github.com/enuesaa/cywagon/internal/libhcl"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func NewConfSrv() ConfSrvInterface {
	return &ConfSrv{
		Container: infra.Default,
		Hcl: libhcl.New(),
	}
}

type ConfSrvInterface interface {
	ReadInWorkdir(workdir string) (model.Config, error)
	Read(workdir string) (model.Config, error)
}

type ConfSrv struct {
	infra.Container

	Hcl libhcl.Parser
}

func (c *ConfSrv) ReadInWorkdir(workdir string) (model.Config, error) {
	var config model.Config

	path := filepath.Join(workdir, "server.hcl")
	fbytes, err := c.Fs.Read(path)
	if err != nil {
		return config, err
	}
	if err := c.Hcl.Parse(fbytes, &config); err != nil {
		return config, err
	}
	return config, nil
}

func (c *ConfSrv) ListHCLFiles(workdir string) (map[string][]byte, error) {
	files := make(map[string][]byte, 0)
	fpaths, err := c.Fs.ListFiles(workdir)
	if err != nil {
		return nil, err
	}
	for _, fpath := range fpaths {
		if !strings.HasSuffix(fpath, ".hcl") {
			continue
		}
		fbytes, err := c.Fs.Read(fpath)
		if err != nil {
			return nil, err
		}
		files[fpath] = fbytes
	}
	return files, nil
}

func (c *ConfSrv) Read(workdir string) (model.Config, error) {
	var config model.Config

	files, err := c.ListHCLFiles(workdir)
	if err != nil {
		return config, err
	}
	hclbody, err := c.Hcl.MergeHCLFiles(files)
	if err != nil {
		return config, err
	}

	if err := c.Hcl.Decode(hclbody, &config); err != nil {
		return config, err
	}

	return config, nil
}
