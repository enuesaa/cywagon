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
		{
			expect: false,
			val: "a",
		},
	}

	for _, tc := range cases {
		actual := engine.matchCondStr(tc.val, tc.eq, tc.in, tc.nq, tc.notin)
		assert.Equal(t, tc.expect, actual)
	}
}

func TestMatchCondStrMap(t *testing.T) {
	engine := New()

	cases := []struct{
		expect bool
		val map[string]string
		eq map[string]string
		in []map[string]string
		nq map[string]string
		notin []map[string]string
	}{
		{
			expect: true,
			val: map[string]string{
				"a": "aaa",
			},
			eq: map[string]string{
				"a": "aaa",
			},
		},
		{
			expect: false,
			val: map[string]string{
				"a": "aaa",
			},
			eq: map[string]string{
				"b": "bbb",
			},
		},
		{
			expect: false,
			val: map[string]string{
				"a": "aaa",
			},
			eq: map[string]string{
				"a": "aaa",
				"b": "bbb",
			},
		},
		{
			expect: true,
			val: map[string]string{
				"a": "aaa",
			},
			in: []map[string]string{
				{"a": "aaa"},
				{"b": "bbb"},
			},
		},
		{
			expect: false,
			val: map[string]string{
				"a": "aaa",
			},
			in: []map[string]string{
				{"b": "bbb"},
				{"c": "ccc"},
			},
		},
		{
			expect: true,
			val: map[string]string{
				"a": "aaa",
			},
			nq: map[string]string{
				"b": "bbb",
			},
		},
		{
			expect: true,
			val: map[string]string{
				"a": "aaa",
			},
			nq: map[string]string{
				"a": "aaa",
				"b": "bbb",
			},
		},
		{
			expect: true,
			val: map[string]string{
				"a": "aaa",
			},
			nq: map[string]string{
				"b": "bbb",
			},
		},
		{
			expect: true,
			val: map[string]string{
				"a": "aaa",
			},
			notin: []map[string]string{
				{"b": "bbb"},
				{"c": "ccc"},
			},
		},
		{
			expect: false,
			val: map[string]string{
				"a": "aaa",
			},
			notin: []map[string]string{
				{"a": "aaa"},
				{"b": "bbb"},
			},
		},
		{
			expect: false,
			val: map[string]string{
				"a": "aaa",
			},
		},
	}

	for _, tc := range cases {
		actual := engine.matchCondStrMap(tc.val, tc.eq, tc.in, tc.nq, tc.notin)
		assert.Equal(t, tc.expect, actual)
	}
}
