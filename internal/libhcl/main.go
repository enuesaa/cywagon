package libhcl

import (
	"maps"

	"github.com/enuesaa/cywagon/internal/service/model"
	"github.com/zclconf/go-cty/cty"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

func New() Parser {
	return Parser{}
}

type Parser struct{}

func (p *Parser) Parse(body []byte, val any) error {
	parser := hclparse.NewParser()

	file, diags := parser.ParseHCL(body, "server.hcl")
	if diags.HasErrors() {
		return NewErrParseFailed(diags[0])
	}

	// see https://github.com/hashicorp/hcl/issues/496
	type PartialConstsConfig struct {
		Consts []model.Const `hcl:"const,block"`
		Remain hcl.Body      `hcl:",remain"`
	}
	var partialConfig PartialConstsConfig

	if diags := gohcl.DecodeBody(file.Body, nil, &partialConfig); diags.HasErrors() {
		return NewErrParseFailed(diags[0])
	}

	constsmap := make(map[string]cty.Value)
	for _, co := range partialConfig.Consts {
		maps.Copy(constsmap, co.Attrs)
	}

	tctx := &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"const": cty.ObjectVal(constsmap),
		},
	}
	if diags := gohcl.DecodeBody(file.Body, tctx, val); diags.HasErrors() {
		return NewErrParseFailed(diags[0])
	}
	return nil
}
