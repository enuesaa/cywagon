package handle

import (
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (h *Handler) listConfs(search []string) ([]model.Conf, error) {
	confsrv := service.NewConfService()

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

func (h *Handler) listConfPaths(paths []string) ([]string, error) {
	confsrv := service.NewConfService()

	var list []string

	for _, path := range paths {
		isDir, err := h.Fs.IsDir(path)
		if err != nil {
			return nil, err
		}
		if isDir {
			files, err := h.Fs.ListFiles(path)
			if err != nil {
				return nil, err
			}
			for _, file := range files {
				if confsrv.IsConfPath(file) {
					list = append(list, file)
				}
			}
		} else {
			list = append(list, path)
		}
	}
	return list, nil
}
