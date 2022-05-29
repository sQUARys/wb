package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	str, error := r.ReadString('\n')
	if error != nil && error != io.EOF {
		fmt.Print("Error of input: ")
		fmt.Println(error)
	}
	fmt.Println("Echo:", str)
}
