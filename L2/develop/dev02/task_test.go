package main

import (
	"testing"
)

type testPair struct {
	values  string
	results string
}

var tests = []testPair{
	{"a4bc2d5e", "aaaabccddddde"},
	{"abcd", "abcd"},
	{"45", ""},
	{"", ""},
}

func TestUnPack(t *testing.T) {
	for _, test := range tests {
		ret := unPack(test.values)
		if ret != test.results {
			t.Error(
				"For", test.values,
				"expected", test.results,
				"got", ret,
			)
		}
	}
}
