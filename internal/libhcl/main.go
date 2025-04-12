package libhcl

import (
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
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
	if diags := gohcl.DecodeBody(file.Body, nil, val); diags.HasErrors() {
		return diags
	}
	return nil
}
