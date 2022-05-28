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
	site := Site{} // создаем пустую структуру
	site.setFlag() // считываем с командной строки

	fmt.Print("\nDownloading file...\n")

	segments := strings.Split(site.URL, "/") // делим URL сайта по /
	fileName := segments[len(segments)-1]    // записываем имя файла
	file, err := os.Create(fileName)         // создаем такой файл

	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	defer file.Close()

	response, err := http.Get(site.URL) // отправляем Get-запрос и принимаем ответ

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("Request Status: %s\n\n", response.Status)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body) // считываем все что вернулось с ответа
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body) // приводим массив байт к строке

	filesize := response.ContentLength // размер файла записываем

	n, error := file.WriteString(sb) //записываем в файл данные, которые были переданы
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
