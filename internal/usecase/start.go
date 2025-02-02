package usecase

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/enuesaa/cywagon/internal/ctlconf"
	"github.com/enuesaa/cywagon/internal/libserve"
)

func Start(ctx context.Context, confDir string) error {
	config, err := ctlconf.Read(ctx, "testdata/example.lua")
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", config)

	go func() {
		cmd := exec.Command("bash", "-c", config.Entry.Cmd)
		cmd.Dir = config.Entry.Workdir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			fmt.Println("Error:", err)
		}
		cmd.Start()
	}()

	time.Sleep(time.Duration(config.Entry.WaitForHealthy) * time.Second)
	fmt.Println("start serve")

	return libserve.Serve(config.Entry.Host)
}
