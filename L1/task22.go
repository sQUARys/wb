package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает,
//делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

const a_example = 24000000000000000000

const as = "24000000000000000000"
const bs string = "53430044342004232000"

func main() {
	a := new(big.Int)
	a.SetString(as, 10)

	b := new(big.Int)
	b.SetString(bs, 10)

	result := new(big.Int)
	result.Add(a, b)
	fmt.Println("Sum = ", result)
	result.Sub(a, b)
	fmt.Println("Sub = ", result)
	result.Mul(a, b)
	fmt.Println("Mul = ", result)
	result.Div(b, a)
	fmt.Println("Div = ", result)

}
