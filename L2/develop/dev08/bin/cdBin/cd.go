package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	r := bufio.NewReader(os.Stdin)
	pathToChange, _ := r.ReadString('\n')
	err := os.Chdir(pathToChange)
	if err != nil {
		log.Fatal("Error of cd:", err)
	}
	currDir, errCurrDir := os.Getwd()
	if err != nil {
		log.Fatal("Error of cd:", errCurrDir)
	}
	fmt.Printf("Succesful. Now you in %s.", currDir)
}
