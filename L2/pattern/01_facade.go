package pattern

import (
	"fmt"
	"strings"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

//Структура School которая также имплементирует фасад в виде двух студентов
type School struct {
	student1 *Student1
	student2 *Student2
}

//создание новой структуры School
func NewSchool() *School {
	return &School{
		student1: &Student1{},
		student2: &Student2{},
	}
}

//Общий вывод программы
func (school *School) GetInfo() {
	result := []string{
		school.student1.SayNameFirst(),
		school.student2.SayNameSecond(),
	}
	fmt.Println(strings.Join(result, "\n"))
}

//Подсистема для первого студента

type Student1 struct{}

func (s1 *Student1) SayNameFirst() string {
	return ("Hi, my name is Dmitrii")
}

//Подсистема для второго студента

type Student2 struct{}

func (s2 *Student2) SayNameSecond() string {
	return ("Hi, my name is Elena")
}
