package libhcl

func New() Parser {
	return Parser{}
}

// see https://github.com/hashicorp/hcl/issues/496
// see https://github.com/hashicorp/hcl/issues/298
type Parser struct{}
