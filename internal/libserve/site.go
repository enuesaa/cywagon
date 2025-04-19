package libserve

import "io/fs"

type Site struct {
	Host    string // Example: `example.com`
	Dist    fs.FS
}
