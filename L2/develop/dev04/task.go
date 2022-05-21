package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var mapForAnagram DefMap

func main() {
	in := []string{"пятак", "актяп", "тяпка", "тяпка"}
	fmt.Println(uploadMap(in))
}

type DefMap struct { // структура хранящая в себе мапу
	m map[string][]string
}

func uploadMap(arr []string) map[string][]string {
	mapForAnagram = *New()

	currentArr := arr

	mapForAnagram.m[currentArr[0]] = append(mapForAnagram.m[currentArr[0]], currentArr[0])

	isUnique := true

	for i := range currentArr {
		for key, _ := range mapForAnagram.m {
			isUnique = true
			if isAnagramm(key, currentArr[i]) {
				for j := range mapForAnagram.m[key] {
					if mapForAnagram.m[key][j] == currentArr[i] {
						isUnique = false
						break
					}
				}
				if isUnique {
					mapForAnagram.m[key] = append(mapForAnagram.m[key], currentArr[i])
					break
				}
			} else {
				if isFirstMeet(currentArr[i]) {
					mapForAnagram.m[currentArr[i]] = append(mapForAnagram.m[currentArr[i]], currentArr[i])
				}
			}
		}
	}

	findSingleArr()
	return mapForAnagram.m
}

func New() *DefMap { // функция создания новой структуры DefMap с пустой мапой
	return &DefMap{
		m: make(map[string][]string),
	}
}

func isAnagramm(key string, str string) bool {
	isAnWord := false

	keyWords := strings.Split(key, "")
	strWords := strings.Split(str, "")

	sort.Slice(keyWords, func(i, j int) bool {
		return keyWords[i] < keyWords[j]
	})

	sort.Slice(strWords, func(i, j int) bool {
		return strWords[i] < strWords[j]
	})

	sortedKeyWords := strings.Join(keyWords, "")
	sortedStrWords := strings.Join(strWords, "")

	if sortedKeyWords == sortedStrWords {
		isAnWord = true
	}

	return isAnWord
}

func isFirstMeet(word string) bool {
	isFirst := true

	for key, _ := range mapForAnagram.m {
		if isAnagramm(key, word) {
			isFirst = false
		}
	}
	return isFirst
}

func findSingleArr() {
	for key, val := range mapForAnagram.m {
		if len(val) == 1 {
			delete(mapForAnagram.m, key)
		}
	}
}
