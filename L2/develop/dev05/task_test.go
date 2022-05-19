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
	count     int
}

var testForAfter = []MapInfo{
	{[]string{"Основные", "объявления", "флагов"}, "флагов", []string{"Основные", "объявления", "флагов", "справа1", "справа2"}, -1},
	{[]string{""}, "флагов", []string{""}, -1},
	{[]string{"флагов"}, "флагов", []string{"флагов", "справа1", "справа2"}, -1},
}

var testForBefore = []MapInfo{
	{[]string{"Основные", "объявления", "флагов"}, "флагов", []string{"Основные", "объявления", "слева1", "слева2", "флагов"}, -1},
	{[]string{""}, "флагов", []string{""}, -1},
}

var testForContext = []MapInfo{
	{[]string{"Основные", "объявления", "флагов"}, "флагов", []string{"Основные", "объявления", "слева1", "слева2", "флагов", "справа1", "справа2"}, -1},
	{[]string{""}, "флагов", []string{""}, -1},
}

type ForLines struct {
	arr    []byte
	substr string
	result int
}

var testForCountLines = []ForLines{
	{[]byte("Основные\n объявления"), "", 1},
	{[]byte("Основные объявления"), "", 0},
	{[]byte(""), "", 0},
	{[]byte("\n\n\n\n"), "", 4},
}
var testForGetLineNumber = []ForLines{
	{[]byte("Основные\n объявления"), "объявления", 2},
	{[]byte("Основные объявления"), "объявления", 1},
	{[]byte("Основные объявления"), "флагов", -1},
	{[]byte(""), "", -1},
}

var testForIgnoreCase = []MapInfo{
	{[]string{"ОснОвные", "ОбъявлЕния", "фЛагОв"}, "флагов", []string{""}, 1},
	{[]string{"фЛагОв", "фЛагОв", "фЛагОв"}, "флагов", []string{""}, 3},
	{[]string{"ОснОвные", "ОснОвные", "ОснОвные"}, "флагов", []string{""}, 0},
}

var testForInvert = []MapInfo{
	{[]string{"Основные", "объявления", "флагов"}, "флагов", []string{"Основные", "объявления"}, -1},
	{[]string{""}, "флагов", []string{}, -1},
	{[]string{"флагов"}, "флагов", []string{}, -1},
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

func TestBefore(t *testing.T) {
	for _, test := range testForBefore {
		ret := Before(test.input, test.subString, subArrLeft)
		if !isSimilar(ret, test.result) {
			t.Error(
				"For", test.input,
				"expected", test.result,
				"got", ret,
			)
		}
	}
}

func TestContext(t *testing.T) {
	for _, test := range testForContext {
		ret := Context(test.input, test.subString, subArrLeft, subArrRight)
		if !isSimilar(ret, test.result) {
			t.Error(
				"For", test.input,
				"expected", test.result,
				"got", ret,
			)
		}
	}
}

func TestCountLines(t *testing.T) {
	for _, test := range testForCountLines {
		ret := CountLines(test.arr)
		if ret == test.result {
			t.Error(
				"For", test.arr,
				"expected", test.result,
				"got", ret,
			)
		}
	}
}

func TestIgnoreCase(t *testing.T) {
	for _, test := range testForIgnoreCase {
		ret := IgnoreCase(test.input, test.subString)
		if ret != test.count {
			t.Error(
				"For", test.input,
				"expected", test.count,
				"got", ret,
			)
		}
	}
}

func TestInvert(t *testing.T) {
	for _, test := range testForInvert {
		_, ret := Invert(test.input, test.subString)
		if !isSimilar(ret, test.result) {
			t.Error(
				"For", test.input,
				"expected", test.result,
				"got", ret,
			)
		}
	}
}

func TestGetLineNumber(t *testing.T) {
	for _, test := range testForGetLineNumber {
		ret := GetLineNumber(test.arr, test.substr)
		if ret != test.result {
			t.Error(
				"For", test.arr,
				"expected", test.result,
				"got", ret,
			)
		}
	}
}

//-n - "line num", печатать номер строки Done
