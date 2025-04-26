package enginectl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcRewritePath(t *testing.T) {
	engine := New()

	cases := []struct {
		from string
		path string
		expect string
	}{
		{
			from: "/aaa/bbb/ccc/ddd/eee",
			path: "/{dir2:}/a.txt",
			expect: "/bbb/ccc/ddd/eee/a.txt",
		},
		{
			from: "/aaa/bbb/ccc/ddd/eee",
			path: "/{:dir2}/a.txt",
			expect: "/aaa/bbb/a.txt",
		},
		{
			from: "/aaa/bbb/ccc/ddd/eee",
			path: "/{dir1}/{dir2}/{dir3}/{dir4}/a.txt",
			expect: "/aaa/bbb/ccc/ddd/a.txt",
		},
	}

	for _, tc := range cases {
		actual := engine.calcRewritePath(tc.from, tc.path)
		assert.Equal(t, tc.expect, actual)
	}
}
