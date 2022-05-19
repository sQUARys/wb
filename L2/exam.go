package main

import (
	"flag"
	"fmt"
)

var (
	cd            string
	pathToConnect string
)

func flagSet() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
}

func main() {
	flagSet()

}
