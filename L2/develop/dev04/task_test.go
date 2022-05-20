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
		ret := IsAnagramm(test.firstWord, test.secondWord)

		if ret != test.result {
			t.Error(
				"For", test.firstWord+" "+test.secondWord,
				"expected", test.result,
				"got", ret,
			)
		}
	}
}

type uploadPair struct {
	values    []string
	resultMap map[string][]string
}

var testForUpload = []uploadPair{
	{[]string{"пятак", "актяп", "тяпка", "тяпка"}, map[string][]string{"пятак": {"пятак", "актяп", "тяпка"}}},
	{[]string{"листок", "слиток", "столик", "hi", "ih", "h"}, map[string][]string{"листок": {"листок", "слиток", "столик"}, "hi": {"hi", "ih"}}},
	{[]string{"😀😃", "😃😀", "", "листок", ""}, map[string][]string{"😀😃": {"😀😃", "😃😀"}}},
}

func TestUploadMap(t *testing.T) {
	for _, test := range testForUpload {
		mapa := New()

		ret := UploadMap(test.values, *mapa)
		if !reflect.DeepEqual(ret, test.resultMap) {
			t.Error(
				"For", test.values,
				"expected", test.resultMap,
				"got", ret,
			)
		}
	}
}
