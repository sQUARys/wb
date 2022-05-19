package main

import (
	"fmt"
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

type uploadPair struct {
	values    []string
	resultMap map[string][]string
}

var testForUpload = []uploadPair{
	{[]string{"пятак", "актяп", "тяпка", "тяпка"}, map[string][]string{"пятак": {"пятак", "актяп", "тяпка"}}},
}

//{[]string{"пятак", "актяп", "тяпка", "тяпка", "hi", "ih", "h"}  , map[string][]string{ "пятак":["пятак" "актяп" "тяпка"] "hi":["hi" "ih"] }} ,
//{},
//"hi":["hi" "ih"] "листок":["листок" "слиток" "столик"]
//😀😃:[😀😃 😃😀]

func TestUploadMap(t *testing.T) {
	for _, test := range testForUpload {
		ret := uploadMap(&test.values)
		fmt.Println(ret, test.resultMap)
		if reflect.DeepEqual(ret, test.resultMap) {
			t.Error(
				"For", test.values,
				"expected", test.resultMap,
				"got", ret,
			)
		}
	}
}
