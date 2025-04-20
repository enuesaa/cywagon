package enginectl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchCondStr(t *testing.T) {
	engine := New()

	str := func(val string) *string {
		return &val
	}

	cases := []struct{
		expect bool
		val string
		eq *string
		in []string
		nq *string
		notin []string
	}{
		{
			expect: true,
			val: "a",
			eq: str("a"),
			in: []string{},
			nq: nil,
			notin: []string{},
		},
		{
			expect: false,
			val: "a",
			eq: str("b"),
			in: []string{},
			nq: nil,
			notin: []string{},
		},
		{
			expect: true,
			val: "a",
			eq: nil,
			in: []string{"a", "b"},
			nq: nil,
			notin: []string{},
		},
		{
			expect: true,
			val: "a",
			eq: nil,
			in: []string{"a", "b"},
			nq: nil,
			notin: []string{},
		},
		{
			expect: true,
			val: "a",
			eq: nil,
			in: []string{},
			nq: str("b"),
			notin: []string{},
		},
		{
			expect: false,
			val: "a",
			eq: nil,
			in: []string{},
			nq: str("a"),
			notin: []string{},
		},
		{
			expect: true,
			val: "a",
			eq: nil,
			in: []string{},
			nq: nil,
			notin: []string{"b", "c"},
		},
		{
			expect: false,
			val: "a",
			eq: nil,
			in: []string{},
			nq: nil,
			notin: []string{"a", "b"},
		},
	}

	for _, tc := range cases {
		actual := engine.matchCondStr(tc.val, tc.eq, tc.in, tc.nq, tc.notin)
		assert.Equal(t, tc.expect, actual)
	}
}
