package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	X float64
	Y float64
}

func New(x float64, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func GetDistance(p1 *Point, p2 *Point) float64 {
	distance := math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
	return distance
}

func main() {

	FirstPoint := New(10, 20)
	SecondPoint := New(-10, 40)
	fmt.Println(GetDistance(&FirstPoint, &SecondPoint))
}
