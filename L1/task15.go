package main

import "fmt"

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
//
//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

var justString string

func someFunc() {
	v := createHugeString(1 << 10) //создаем строку длинной 1024
	justString = v[:100]           // записываем срез в переменную

	ArrCopy := make([]byte, len(justString)) // создаем массив байт для копирования в него среза

	copy(ArrCopy, justString) // копируем срез, чтобы на v больше ни один срез не ссылался. Тем самым v не будет больше хранится в памяти

	fmt.Println(ArrCopy)
}

func main() {
	someFunc()
}
