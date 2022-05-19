package main

import "testing"

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
	{{"пятак", "актяп", "тяпка", "тяпка"}, map[string][]string{ "пятак":["пятак" "актяп" "тяпка"] }},
	{"пятак", "актяп", "тяпка", "тяпка", "hi", "ih", "h"},
	{},
}
//"hi":["hi" "ih"] "листок":["листок" "слиток" "столик"]
//😀😃:[😀😃 😃😀]

func TestUploadMap(t *testing.T) {

}
