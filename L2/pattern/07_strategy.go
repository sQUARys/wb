package main

import (
	"fmt"
	"sort"
)

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type StrategySort interface {
	Sort([]int)
}

type BubbleSort struct{}

func (sc *BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

type SimpleSort struct{}

func (s *SimpleSort) Sort(a []int) {
	sort.Ints(a)
}

type Context struct {
	strategy StrategySort
}

func (ctx *Context) ReplacesStrategies(strategy StrategySort) {
	ctx.strategy = strategy
}

func (ctx *Context) Sort(a []int) {
	ctx.strategy.Sort(a)
}

func main() {
	firstArr := []int{1, 4, 5, 6, -1}
	secondArr := []int{10, 3, 5, 2, 6}

	ctx := new(Context)
	ctx.ReplacesStrategies(&BubbleSort{})
	ctx.Sort(firstArr)

	ctx.ReplacesStrategies(&SimpleSort{})
	ctx.Sort(secondArr)

	for _, val := range firstArr {
		fmt.Print(val, " ")
	}
	fmt.Println()
	for _, val := range secondArr {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
