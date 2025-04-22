package libhcl

import "github.com/zclconf/go-cty/cty"

func New() Parser {
	return Parser{}
}

// see https://github.com/hashicorp/hcl/issues/496
// see https://github.com/hashicorp/hcl/issues/298
type Parser struct{
	vars map[string]cty.Value
}
