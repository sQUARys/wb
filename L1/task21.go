package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

type IProcess interface {
	InProcess()
}
type Student struct {
}

func (s *Student) PrintInfo(str string) {
	fmt.Printf("Hi, my name is %s\n", str)
}

//Adapter
type School struct {
	student Student
}

func (school School) InProcess() {
	fmt.Println("Welcome to our school.")
	school.student.PrintInfo("Oleg")
}

func main() {
	var processor IProcess = School{}
	processor.InProcess()
}
