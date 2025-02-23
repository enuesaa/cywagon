package handle

import (
	"github.com/enuesaa/cywagon/internal/service"
	"github.com/enuesaa/cywagon/internal/service/model"
)

func (h *Handler) listConfs(paths []string) ([]model.Conf, error) {
	var confpaths []string

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
			confpaths = append(confpaths, files...)
		} else {
			confpaths = append(confpaths, path)
		}
	}

	var list []model.Conf
	confsrv := service.NewConfService()

	for _, confpath := range confpaths {
		conf, err := confsrv.Read(confpath)
		if err != nil {
			return nil, err
		}
		list = append(list, conf)
	}

	return list, nil
}
