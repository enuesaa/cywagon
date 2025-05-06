package enginectl

import "github.com/gobwas/glob"

func (e *Engine) matchCondPath(val string, eq *string, in []string, nq *string, notin []string) bool {
	isEq := func(pattern string, v string) bool {
		g, err := glob.Compile(pattern)
		if err != nil {
			return false
		}
		return g.Match(v)
	}

	if nq != nil && !isEq(*nq, val) {
		return true
	}
	if len(notin) > 0 {
		for _, nq := range notin {
			if isEq(nq, val) {
				return false
			}
		}
		return true
	}
	if eq != nil && isEq(*eq, val) {
		return true
	}
	if len(in) > 0 {
		for _, eq := range in {
			if isEq(eq, val) {
				return true
			}
		}
	}
	return false
}
