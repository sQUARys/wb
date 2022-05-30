package main

import "testing"

type testPair struct {
	values  []string
	k       int
	results []string
}

var testsSortByK = []testPair{
	{[]string{"abc", "aas", "bfd"}, 1, []string{"aas", "abc", "bfd"}},
	{[]string{"", "a", "f", "b"}, 0, []string{"", "a", "b", "f"}},
	{[]string{"aaa", "bab", "a", "bb", "bkf"}, 1, []string{"aaa", "bab", "a", "bb", "bkf"}},
	{[]string{""}, 0, []string{""}},
}

var testsSortWithoutRepeated = []testPair{
	{[]string{"aaa", "aaa", "aaa"}, 1, []string{"aaa"}},
	{[]string{"c", "a", "b", "b"}, 0, []string{"a", "b", "c"}},
	{[]string{"", "a", "", "a", "c"}, 1, []string{"", "a", "c"}},
	{[]string{""}, 0, []string{""}},
}

var testsRevers = []testPair{
	{[]string{"abc", "aas", "bfd"}, -1, []string{"bfd", "abc", "aas"}},
	{[]string{"", "a", "f", "b"}, -1, []string{"f", "b", "a", ""}},
	{[]string{"aaa", "bab", "a", "bb", "bkf"}, -1, []string{"bkf", "bb", "bab", "aaa", "a"}},
	{[]string{""}, -1, []string{""}},
}

var testSortByNumbers = []testPair{
	{[]string{"1", "2", "3"}, -1, []string{"1", "2", "3"}},
	{[]string{"", "10", "-1", "2"}, -1, []string{"-1", "2", "10", ""}},
	{[]string{"0", "0", "0", "5", "-5"}, -1, []string{"-5", "0", "0", "0", "5"}},
	{[]string{""}, -1, []string{""}},
}

func isSimilar(sl1 []string, sl2 []string) bool {
	isSimilarBool := true
	if len(sl1) != len(sl2) {
		isSimilarBool = false
	}
	for i := range sl1 {
		if sl1[i] != sl2[i] {
			isSimilarBool = false
			break
		}
	}
	return isSimilarBool
}

func TestReverse(t *testing.T) {
	for _, test := range testsRevers {
		ret := reverse(test.values)
		if !isSimilar(ret, test.results) {
			t.Error(
				"For", test.values,
				"expected", test.results,
				"got", ret,
			)
		}
	}
}

func TestSortingByNumber(t *testing.T) {
	for _, test := range testSortByNumbers {
		ret := sortByNumber(test.values)

		if !isSimilar(ret, test.results) {
			t.Error(
				"For", test.values,
				"expected", test.results,
				"got", ret,
			)
		}
	}
}

func TestSortingWithoutRepeated(t *testing.T) {
	for _, test := range testsSortWithoutRepeated {
		ret := sortWithoutRepeat(test.values)

		if !isSimilar(ret, test.results) {
			t.Error(
				"For", test.values,
				"expected", test.results,
				"got", ret,
			)
		}
	}
}

func TestSortingByK(t *testing.T) {
	for _, test := range testsSortByK {
		ret := sortBySpecialColumn(test.values, test.k)

		if !isSimilar(ret, test.results) {
			t.Error(
				"For", test.values,
				"expected", test.results,
				"got", ret,
			)
		}
	}
}
