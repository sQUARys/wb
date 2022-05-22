package main

import (
	"flag"
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
	site := Site{}
	site.setFlag()

	fmt.Print("\nDownloading file...\n")

	segments := strings.Split(site.URL, "/")
	fileName := segments[len(segments)-1]
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
	defer file.Close()

	response, err := http.Get(site.URL)

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

	n, error := file.WriteString(sb)
	if n != int(filesize) {
		fmt.Println("Truncated")
	}
	if error != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Println("Succesful writing in file")
}

type Site struct {
	URL string
}

func (s *Site) setFlag() {
	flag.StringVar(&s.URL, "url", "https://github.com/sQUARys", "url of site we need to download")
	flag.Parse()
}
