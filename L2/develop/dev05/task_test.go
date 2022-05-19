package main

import (
	"fmt"
	"testing"
)

var subArrLeft = []string{"слева1", "слева2"}
var subArrRight = []string{"справа1", "справа2"}

type MapInfo struct {
	input     []string
	subString string
	result    []string
}

var testForAfter = []MapInfo{
	{[]string{"Основные", "объявления", "флагов"}, "флагов", []string{"Основные", "объявления", "флагов", "справа1", "справа2"}},
}

func isSimilar(sl1 []string, sl2 []string) bool {
	isSimilarBool := true
	if len(sl1) != len(sl2) {
		isSimilarBool = false
	}
	for i := range sl1 {
		fmt.Println(sl1[i], sl2[i], sl1[i] == sl2[i])
		if sl1[i] != sl2[i] {
			isSimilarBool = false
			break
		}
	}
	return isSimilarBool
}

func TestAfter(t *testing.T) {
	for _, test := range testForAfter {
		ret := After(test.input, test.subString, subArrRight)
		fmt.Println(ret, test.result, isSimilar(ret, test.result))
		if !isSimilar(ret, test.result) {
			t.Error(
				"For", test.input,
				"expected", test.result,
				"got", ret,
			)
		}
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
