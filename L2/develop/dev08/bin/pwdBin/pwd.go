package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	currDir, errCurrDir := os.Getwd()
	if errCurrDir != nil {
		log.Fatal("Error of pwd:", errCurrDir)
	}
	fmt.Printf("Now you in %s.", currDir)
}
