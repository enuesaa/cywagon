package libhcl

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
)

func New() Parser {
	return Parser{}
}

type Parser struct {}

func (p *Parser) Parse(body []byte, val any) error {
	parser := hclparse.NewParser()

	file, diags := parser.ParseHCL(body, "cywagon.hcl")
	if diags.HasErrors() {
		return diags
	}
	ctx := &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"aaa": cty.StringVal("hello"),
		},
	}
	if diags := gohcl.DecodeBody(file.Body, ctx, val); diags.HasErrors() {
		return diags
	}
	return nil
}
