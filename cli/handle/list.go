package handle

import (
	"fmt"

	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (h *Handler) listConfs(search []string) ([]model.Conf, error) {
	confsrv := service.NewConfService(h.Container)

	confpaths, err := h.listConfPaths(search)
	if err != nil {
		return nil, err
	}
	var list []model.Conf

	for _, confpath := range confpaths {
		conf, err := confsrv.Read(confpath)
		if err != nil {
			return nil, err
		}
		list = append(list, conf)
	}
	return list, nil
}

func (h *Handler) listConfPaths(search []string) ([]string, error) {
	confsrv := service.NewConfService(h.Container)

	var list []string

	for _, path := range search {
		if !h.Fs.IsExist(path) {
			return nil, fmt.Errorf("path not found: %s", path)
		}
		if h.Fs.IsFile(path) {
			list = append(list, path)
			continue
		}
		files, err := h.Fs.ListFiles(path)
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			if confsrv.IsConfPath(file) {
				list = append(list, file)
			}
		}
	}
	return list, nil
}
