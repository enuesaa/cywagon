package enginectl

import "slices"

func (e *Engine) shouldCheckCondStr(eq *string, in []string, nq *string, notin []string) bool {
	return eq != nil || len(in) > 0 || nq != nil || len(notin) > 0
}

func (e *Engine) matchCondStr(val string, eq *string, in []string, nq *string, notin []string) bool {
	if nq != nil && *nq != val {
		return true
	}
	if len(notin) > 0 && !slices.Contains(notin, val) {
		return true
	}
	if eq != nil && *eq == val {
		return true
	}
	if len(in) > 0 && slices.Contains(in, val) {
		return true
	}
	return false
}

func (e *Engine) shouldCheckCondStrMap(eq map[string]string, in []map[string]string, nq map[string]string, notin []map[string]string) bool {
	return len(eq) > 0 || len(in) > 0 || len(nq) > 0 || len(notin) > 0
}

func (e *Engine) matchCondStrMap(val map[string]string, eq map[string]string, in []map[string]string, nq map[string]string, notin []map[string]string) bool {
	isEq := func(target map[string]string) bool {
		for eqk, eqv := range target {
			valv, ok := val[eqk]
			if !ok {
				return false
			}
			if valv != eqv {
				return false
			}
		}
		return true
	}

	if len(nq) > 0 && !isEq(nq) {
		return true
	}
	if len(notin) > 0 {
		for _, nq := range notin {
			if len(nq) > 0 && isEq(nq) {
				return false
			}
		}
		return true
	}
	if eq != nil && isEq(eq) {
		return true
	}
	if len(in) > 0 {
		for _, eq := range in {
			if len(eq) > 0 && isEq(eq) {
				return true
			}
		}
	}
	return false
}
