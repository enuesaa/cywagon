package libhcl

import (
	// "fmt"

	"fmt"

	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/zclconf/go-cty/cty"

	// "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	// "github.com/zclconf/go-cty/cty"
)

func New() Parser {
	return Parser{}
}

type Parser struct {}

// see https://github.com/hashicorp/hcl/issues/496
func (p *Parser) Parse(body []byte, _ any) error {
	var val model.Config

	parser := hclparse.NewParser()

	file, diags := parser.ParseHCL(body, "cywagon.hcl")
	if diags.HasErrors() {
		return diags
	}

	type PartialDefsConfig struct {
		Defs []model.Def  `hcl:"def,block"`
		Remain hcl.Body `hcl:",remain"`
	}
	var partialConfig PartialDefsConfig
	if diags := gohcl.DecodeBody(file.Body, nil, &partialConfig); diags.HasErrors() {
		return diags
	}

	defMap := make(map[string]cty.Value)
	for _, def := range partialConfig.Defs {
		defMap[def.Name] = cty.ObjectVal(def.Props)
	}

	ctx := &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"def": cty.ObjectVal(defMap),
		},
	}
	fmt.Printf("%+v\n", cty.ObjectVal(defMap))

	if diags := gohcl.DecodeBody(file.Body, ctx, &val); diags.HasErrors() {
		return diags
	}
	fmt.Printf("%+v\n", val)

	return nil
}
