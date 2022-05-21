package main

import (
	"reflect"
	"testing"
)

type testPair struct {
	firstWord  string
	secondWord string
	result     bool
}

var testsAnagramm = []testPair{
	{"пятка", "тяпка", true},
	{"", "", true},
	{"пятк", "тапк", false},
}

func TestAnagramm(t *testing.T) {
	for _, test := range testsAnagramm {
		ret := isAnagramm(test.firstWord, test.secondWord)

		if ret != test.result {
			t.Error(
				"For", test.firstWord+" "+test.secondWord,
				"expected", test.result,
				"got", ret,
			)
		}
	}
}

func Test_uploadMap(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want map[string][]string
	}{
		{"First test...", []string{"пятак", "актяп", "тяпка", "тяпка"}, map[string][]string{"пятак": {"пятак", "актяп", "тяпка"}}},
		{"Second test...", []string{"листок", "слиток", "столик", "hi", "ih", "h"}, map[string][]string{"листок": {"листок", "слиток", "столик"}, "hi": {"hi", "ih"}}},
		{"Third test...", []string{"😀😃", "😃😀", "", "листок", ""}, map[string][]string{"😀😃": {"😀😃", "😃😀"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uploadMap(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uploadMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
