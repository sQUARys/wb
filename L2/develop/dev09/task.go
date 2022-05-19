package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ==

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	str := "https://github.com/sQUARys"

	fmt.Print("\nDownloading file...\n")

	segments := strings.Split(str, "/")
	fileName := segments[len(segments)-1]
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	response, err := http.Get(str)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("Request Status: %s\n\n", response.Status)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	filesize := response.ContentLength

	go func() {

		n, error := file.WriteString(sb)
		if n != int(filesize) {
			fmt.Println("Truncated")
		}
		if error != nil {
			fmt.Printf("Error: %v", err)
		}
	}()

	fmt.Println("Succesful writing in file")
}
