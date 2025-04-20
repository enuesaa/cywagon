package enginectl

import "path"

func (e *Engine) matchCondPath(val string, eq *string, in []string, nq *string, notin []string) bool {
	isEq := func(pattern string, v string) bool {
		ok, err := path.Match(pattern, v)
		if err != nil {
			return false
		}
		return ok
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