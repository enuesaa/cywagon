package enginectl

import (
	"fmt"
	"strings"
)

func (e *Engine) calcRewritePath(from string, path string) string {
	dirs := strings.Split(from, "/")

	for i, val := range dirs {
		if val == "" {
			continue
		}
		keyr := fmt.Sprintf("{dir%d}", i)
		path = strings.ReplaceAll(path, keyr, val)

		path = strings.ReplaceAll(path, "/:}", "/"+val+"/:}")
		keyr = fmt.Sprintf("{dir%d:}", i)
		path = strings.ReplaceAll(path, keyr, val+"/:}")

		keyr = fmt.Sprintf("{:dir%d}", i)
		path = strings.ReplaceAll(path, keyr, val)
		path = strings.ReplaceAll(path, "{:", val+"/{:")
	}
	path = strings.ReplaceAll(path, "/{:", "")
	path = strings.ReplaceAll(path, "/:}", "")

	path = strings.ReplaceAll(path, "{path}", from)

	if len(dirs) > 0 {
		last := dirs[len(dirs)-1]
		path = strings.ReplaceAll(path, "{last}", last)
	}
	return path
}
