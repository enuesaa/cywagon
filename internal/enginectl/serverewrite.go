package enginectl

import (
	"fmt"
	"regexp"
	"strings"
)

func (e *Engine) extractRewritePathVars(path string, pattern string) []string {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil
	}
	match := re.FindAllStringSubmatch(path, -1)
	if len(match[0]) <= 1 {
		return nil
	}
	return match[0][1:]
}

func (e *Engine) injectRewritePathVars(path string, vars []string) string {
	if len(vars) == 0 {
		return path
	}
	for i, val := range vars {
		from := fmt.Sprintf("%%%d", i + 1)
		path = strings.ReplaceAll(path, from, val)
	}
	fmt.Println(path)
	return path
}
