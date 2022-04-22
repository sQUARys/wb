package main

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.
//
//var justString string

//
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}

//func main() {
//	justString := "HiHowatryou"
//	//someFunc()
//	justString[:5]
//}

import "fmt"

var c string

func main() {
	a := "12345"
	//c := copyString(a[:2])
	c := a[:3]
	fmt.Println(c, a)
}

func copyString(b string) string {
	c := make([]byte, len(b))

	copy(c, b)

	return string(c)
}

// ??????
