package libhcl

import (
	"fmt"

	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/zclconf/go-cty/cty"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	// "github.com/zclconf/go-cty/cty"
)

func New() Parser {
	return Parser{}
}

type Parser struct{}

// see https://github.com/hashicorp/hcl/issues/496
func (p *Parser) Parse(body []byte, val any) error {
	parser := hclparse.NewParser()

	file, diags := parser.ParseHCL(body, "cywagon.hcl")
	if diags.HasErrors() {
		return diags
	}

	type PartialConstsConfig struct {
		Consts []model.Const `hcl:"const,block"`
		Remain hcl.Body      `hcl:",remain"`
	}
	var partialConfig PartialConstsConfig
	if diags := gohcl.DecodeBody(file.Body, nil, &partialConfig); diags.HasErrors() {
		return diags
	}

	constsmap := make(map[string]cty.Value)
	for _, co := range partialConfig.Consts {
		for key, attr := range co.Attrs {
			constsmap[key] = attr
		}
	}

	ctx := &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"const": cty.ObjectVal(constsmap),
		},
	}
	fmt.Printf("%+v\n", cty.ObjectVal(constsmap))

	if diags := gohcl.DecodeBody(file.Body, ctx, val); diags.HasErrors() {
		return diags
	}
	return nil
}
