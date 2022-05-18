package main

import (
	"flag"
)

var (
	timeout       string
	pathToConnect string
)

func flagSet() {
	flag.StringVar(&timeout, "timeout", "default", "A timeout Var")
	flag.Parse()
	args := flag.Args()
	pathToConnect = args[0] + ":" + args[1]
}

func main() {
	flagSet()

}
