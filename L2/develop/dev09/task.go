package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

	fmt.Println("\nDownloading file...\n")

	fileUrl, err := url.Parse(*cliUrl)

	if err != nil {
		panic(err)
	}

	filePath := fileUrl.Path
	segments := strings.Split(filePath, "/")
	fileName := segments[len(segments)-1]
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	checkStatus := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	response, err := checkStatus.Get(*cliUrl)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	fmt.Printf("Request Status: %s\n\n", response.Status)

	filesize := response.ContentLength

	go func() {
		n, err := io.Copy(file, response.Body)
		if n != filesize {
			fmt.Println("Truncated")
		}
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
	}()

}
