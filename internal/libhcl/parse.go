package libhcl

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
)

func (p *Parser) MergeHCLFiles(files map[string][]byte) (hcl.Body, error) {
	var list []*hcl.File
	parser := hclparse.NewParser()

	for fpath, fbody := range files {
		f, diags := parser.ParseHCL(fbody, fpath)
		if diags.HasErrors() {
			return nil, NewErrParseFailed(diags[0])
		}
		list = append(list, f)
	}
	body := hcl.MergeFiles(list)

	return body, nil
}

func (p *Parser) UseVar(name string, val cty.Value) {
	p.vars[name] = val
}

func (p *Parser) Decode(body hcl.Body, val any) error {
	tctx := &hcl.EvalContext{
		Variables: p.vars,
	}
	if diags := gohcl.DecodeBody(body, tctx, val); diags.HasErrors() {
		return NewErrParseFailed(diags[0])
	}
	return nil
}
