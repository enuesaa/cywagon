package cli

import (
	"context"
	"flag"
	"log"

	"github.com/enuesaa/cywagon/internal/conf"
	"github.com/google/subcommands"
	"fmt"
	"reflect"

	"github.com/fatih/structtag"

)

func newPlanCmd() *planCmd {
	return &planCmd{}
}

type planCmd struct {}

func (c *planCmd) Name() string {
	return "plan"
}

func (c *planCmd) Synopsis() string {
	return "plan"
}

func (c *planCmd) Usage() string {
	return "cywagon plan\n"
}

func (c *planCmd) SetFlags(f *flag.FlagSet) {}

func (c *planCmd) Execute(ctx context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	type Value struct {
		Name string `lua:"name"`
		Meta string `lua:"meta"`
		Status int `lua:"status"`
		Success bool `lua:"success"`
		IgnoreThis int
	}
	value := Value{
		Name: "a",
		Meta: "me",
		Status: 300,
		Success: true,
		IgnoreThis: 300,
	}

	target := reflect.TypeOf(value)
	for i := range target.NumField() {
		tag := target.Field(i).Tag

		tags, err := structtag.Parse(string(tag))
		if err != nil {
			log.Printf("Error: %s", err.Error())
			continue
		}
	
		luaTag, err := tags.Get("lua")
		if err != nil {
			log.Printf("Error: %s", err.Error())
			continue
		}
		fmt.Printf("FOUND: %s ===>>> %+v\n", luaTag.Name, target.Field(i))
	}
	
	if err := conf.Parse(ctx); err != nil {
		log.Printf("Error: %s", err.Error())
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
