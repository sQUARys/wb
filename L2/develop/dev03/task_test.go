package main

import (
	"testing"
)

//func Reverse(sl []string) []string {
//	sort.Strings(sl)
//	sort.Sort(sort.Reverse(sort.StringSlice(sl)))
//	return sl
//}
//
//func SortWithoutRepeat(sl []string) []string {
//	var mem []string
//	sort.Strings(sl)
//	for i := range sl {
//		if len(mem) == sort.SearchStrings(mem, sl[i]) { // если не найдено такой строки
//			mem = append(mem, sl[i])
//		}
//	}
//	return mem
//}
//
//func SortBySpecialColumn(sl []string, k int) []string {
//	sort.Slice(sl, func(i, j int) bool {
//		if k >= len(sl[i]) || k >= len(sl[j]) {
//			return false
//		}
//		return sl[i][k] < sl[j][k]
//	})
//	return sl
//}
//
//func SortByNumber(sl []string) {
//	//????? что значит сортировать по числовому значению
//}

type testPair struct {
	values  []string
	k       int
	results []string
}

var testsSortByK = []testPair{
	{[]string{"abc", "aas", "bfd"}, 1, []string{"aas", "abc", "bfd"}},
	{[]string{"", "a", "f", "b"}, 0, []string{"a", "b", "f", ""}},
	{[]string{"aaa", "bab", "a", "bb", "bkf"}, 1, []string{"aaa", "bab", "bkf", "a", "bb"}},
	//{[]string{""}, -1,""},
}

var testsRevers = []testPair{
	{[]string{"abc", "aas", "bfd"}, -1, []string{"bfd", "aas", "abc"}},
	{[]string{"", "a", "f", "b"}, -1, []string{"b", "f", "a", ""}},
	{[]string{"aaa", "bab", "a", "bb", "bkf"}, -1, []string{"bkf", "bb", "a", "bab", "aaa"}},
	{[]string{""}, -1, []string{""}},
}

func isSimilar(sl1 []string, sl2 []string) bool {
	var isSimilarBool bool
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
		ret := Reverse(test.values)

		if isSimilar(ret, test.results) {
			t.Error(
				"For", test.values,
				"expected", test.results,
				"got", ret,
			)
		}
	}
}
