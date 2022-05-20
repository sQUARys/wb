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

type DefMap struct { // структура хранящая в себе мапу и мьютекс для ее блокировки и разблокировки
	m map[string][]string
}

func New() *DefMap { // функция создания новой структуры DefMap с пустой мапой
	return &DefMap{
		m: make(map[string][]string),
	}
}

func IsAnagramm(key string, str string) bool {
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

func IsFirstMeet(word string, mapa DefMap) bool {
	isFirst := true

	for key, _ := range mapa.m {
		if IsAnagramm(key, word) {
			isFirst = false
		}
	}
	return isFirst
}

func FindSingleArr(mapa DefMap) {
	for key, val := range mapa.m {
		if len(val) == 1 {
			delete(mapa.m, key)
		}
	}
}

func UploadMap(arr []string, mapa DefMap) map[string][]string {
	currentArr := arr

	mapa.m[currentArr[0]] = append(mapa.m[currentArr[0]], currentArr[0])

	isUnique := true

	for i := range currentArr {
		for key, _ := range mapa.m {
			isUnique = true
			if IsAnagramm(key, currentArr[i]) {
				for j := range mapa.m[key] {
					if mapa.m[key][j] == currentArr[i] {
						isUnique = false
						break
					}
				}
				if isUnique {
					mapa.m[key] = append(mapa.m[key], currentArr[i])
					break
				}
			} else {
				if IsFirstMeet(currentArr[i], mapa) {
					mapa.m[currentArr[i]] = append(mapa.m[currentArr[i]], currentArr[i])
				}
			}
		}
	}

	FindSingleArr(mapa)
	return mapa.m
}

func main() {
	in := []string{"пятак", "актяп", "тяпка", "тяпка"}

	mapa := New()
	fmt.Println(UploadMap(in, *mapa))
}
