package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel
//из переменной типа interface{}.

func TypeOf(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func main() {
	v := "hello world"
	fmt.Println(TypeOf(v))
	a := 10
	fmt.Println(TypeOf(a))
	channel := make(chan string)
	fmt.Println(TypeOf(channel))
	close(channel)

	boolean := true
	fmt.Println(TypeOf(boolean))
}
