package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel
//из переменной типа interface{}.

func TypeOf(v interface{}) string { // функция для определения типа введенего в него
	return reflect.TypeOf(v).String() // вовзращение типа
}

func main() {
	v := "hello world"     // строка
	fmt.Println(TypeOf(v)) // выводит тип переменной v
	a := 10                // число
	fmt.Println(TypeOf(a)) // выводит тип переменной a
	channel := make(chan string)
	fmt.Println(TypeOf(channel)) //выводит тип переменной channel
	close(channel)

	boolean := true
	fmt.Println(TypeOf(boolean)) //выводит тип переменной boolean
}
