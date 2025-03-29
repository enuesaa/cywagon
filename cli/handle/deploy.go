package handle

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"time"
)

func (h *Handler) Deploy(sitename string, path string) error {
	h.Log.Info("start a new deployment..")

	current := os.DirFS(".")

	readme, err := current.Open("README.md")
	if err != nil {
		return err
	}
	time.Sleep(10 * time.Second)

	readmeBytes, err := io.ReadAll(readme)
	if err != nil {
		return err
	}
	fmt.Println(string(readmeBytes))

	fs.WalkDir(current, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
		return nil
	})

	return nil
}
