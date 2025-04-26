package handle

import (
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

func (h *Handler) Format(workdir string) error {
	fpaths, err := h.Fs.ListFiles(workdir)
	if err != nil {
		return err
	}
	for _, fpath := range fpaths {
		if !strings.HasSuffix(fpath, ".hcl") {
			continue
		}
		fbytes, err := h.Fs.Read(fpath)
		if err != nil {
			return err
		}
		formatted := hclwrite.Format(fbytes)

		if err := h.Fs.Create(fpath, formatted); err != nil {
			return err
		}
		h.Log.Info(fpath)
	}
	return nil
}
