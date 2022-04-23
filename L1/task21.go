package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

type IProcess interface { // интерфейс для работы с адаптером
	InProcess()
}
type Student struct { // структура
}

func (s *Student) PrintInfo(str string) { // функция для структуры Student
	fmt.Printf("Hi, my name is %s\n", str)
}

type School struct { //Структура аdapter
	student Student
}

func (school School) InProcess() { // для структуры School создадим функцию
	fmt.Println("Welcome to our school.")
	school.student.PrintInfo("Oleg") // в которой будем использован метод другой структуры Student
}

func main() {
	var processor IProcess = School{} // создаем адаптер
	processor.InProcess()             // вызываем функцию адаптера
}
