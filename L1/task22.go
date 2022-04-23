package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает,
//делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

const a_example = 24000000000000000000 //если писать const без определения типа, то в данном случае будет тип = переполненный int

const as = "24000000000000000000"        // строка, которая будет приведена в число
const bs string = "53430044342004232000" // строка, которая будет приведена в число

func main() {
	a := new(big.Int)   // создаем переменную типа big
	a.SetString(as, 10) //переводим значение строки в значение bid.Int

	b := new(big.Int)   // создаем переменную типа big
	b.SetString(bs, 10) //переводим значение строки в значение bid.Int

	result := new(big.Int) // создаем переменную типа big для запоминания результата после того или иного изменения
	result.Add(a, b)       // функция сложения двух чисел big.Int
	fmt.Println("Sum = ", result)
	result.Sub(a, b) // функция разности двух чисел big.Int
	fmt.Println("Sub = ", result)
	result.Mul(a, b) // функция произведения двух чисел big.Int
	fmt.Println("Mul = ", result)
	result.Div(b, a) // функция деления двух чисел big.Int
	fmt.Println("Div = ", result)

}
