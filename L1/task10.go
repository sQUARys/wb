package main

import (
	"fmt"
	"math"
	"sort"
)

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов.
//Последовательность в подмножноствах не важна.
//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	sort.Float64s(temp) //[-27 -25.4 -21 13 15.5 19 24.5 32.5]

	result := make(map[float64][]float64)

	for _, t := range temp {
		RoundedUp := math.Trunc(t/10) * 10 // Округляем температуры вниз
		result[RoundedUp] = append(result[RoundedUp], t)
	}

	for i, temperatures := range result {
		fmt.Printf("%v: %v\n", i, temperatures)
	}
}
