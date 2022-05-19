package main

import (
	"flag"
	"fmt"
)

var (
	cd            string
	pathToConnect string
)

//-k — указание колонки для сортировки Done
//-n — сортировать по числовому значению ??
//-r — сортировать в обратном порядке Done
//-u — не выводить повторяющиеся строки Done

type Command struct {
	k int
	n bool
	r bool
	u bool
}

func (c *Command) flagSet() {

	flag.IntVar(&c.k, "k", -1, "Column")
	flag.BoolVar(&c.n, "n", false, "sorting by int value")
	flag.BoolVar(&c.r, "r", false, "reverse sorting")
	flag.BoolVar(&c.u, "u", false, "sorting without repeated")

	flag.Parse()
}

func (c *Command) flagSet1() {
	flag.IntVar(&c.k, "k", -1, "Column")
	flag.BoolVar(&c.n, "n", false, "sorting by int value")
	flag.BoolVar(&c.r, "r", false, "reverse sorting")
	flag.BoolVar(&c.u, "u", false, "sorting without repeated")

	flag.Parse()
}
func main() {
	c := Command{}
	c.flagSet1()
	fmt.Println(c)
}
