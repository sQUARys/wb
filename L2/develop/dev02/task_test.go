package main

import (
	"testing"
)

type testPair struct {
	values  []rune
	results string
}

var tests = []testPair{
	{[]rune("a4bc2d5e"), "aaaabccddddde"},
	{[]rune("abcd"), "abcd"},
	{[]rune("45"), ""},
	{[]rune(""), ""},
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
