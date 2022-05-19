package main

import (
	"fmt"
	"testing"
)

//
//type testPair struct {
//	values  []string
//	k       int
//	results []string
//}
//
//var testsForAfter = []testPair{
//	{[]string{"abc", "aas", "bfd"}, 1, []string{"aas", "abc", "bfd"}},
//	{[]string{"", "a", "f", "b"}, 0, []string{"a", "b", "f", ""}},
//	{[]string{"aaa", "bab", "a", "bb", "bkf"}, 1, []string{"aaa", "bab", "bkf", "a", "bb"}},
//	{[]string{""}, 0, []string{""}},
//}

//in := []string{"строка1", "строка2"}

var subArr1 = []string{"слева1", "слева2"}
var subArr2 = []string{"справа1", "справа2"}

type MapInfo struct {
	input     []string
	subString string
	result    []string
}

var mapa = map[string][]MapInfo{
	"after": {{[]string{"Основные", "объявления", "флагов"}, "флагов", []string{"Основные", "объявления", "флагов", "справа1", "справа2"}}},
}

func TestAfter(t *testing.T) {
	for _, test := range mapa["after"] {
		fmt.Println(test)
		//ret := After(&test.values)
		//fmt.Println(ret, test.resultMap)
		//if reflect.DeepEqual(ret, test.resultMap) {
		//	t.Error(
		//		"For", test.values,
		//		"expected", test.resultMap,
		//		"got", ret,
		//	)
		//}
	}
}

//-A - "after" печатать +N строк после совпадения Done
//-B - "before" печатать +N строк до совпадения Done
//-C - "context" (A+B) печатать ±N строк вокруг совпадения Done
//-c - "count" (количество строк) Done
//-i - "ignore-case" (игнорировать регистр) Done
//-v - "invert" (вместо совпадения, исключать) Done
//-F - "fixed", точное совпадение со строкой, не паттерн Done
//-n - "line num", печатать номер строки Done
