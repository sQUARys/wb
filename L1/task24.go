package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct { // структура
	X float64
	Y float64
}

func New(x float64, y float64) Point { // создание структуры с координатами x , y
	return Point{
		X: x,
		Y: y,
	}
}

func GetDistance(p1 *Point, p2 *Point) float64 { //функция получения расстояния между двумя точками
	distance := math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2)) // считаем расстояние
	return distance                                                        //возвращаем его расстояние
}

func main() {
	FirstPoint := New(10, 20)                           // создаем точку с координатвми (10 , 20)
	SecondPoint := New(-10, 40)                         // создаем точку с координатвми (-10 , 40)
	fmt.Println(GetDistance(&FirstPoint, &SecondPoint)) // высчитываем расстояние между двумя точками
}
