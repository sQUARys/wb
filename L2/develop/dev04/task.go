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
	in := []string{"Пятак", "актяп", "тяпка", "тяпка"}
	in = toLowRegister(in)
	fmt.Println(uploadMap(in))
}

type DefMap struct { // структура хранящая в себе мапу
	m map[string][]string
}

func toLowRegister(arr []string) []string {
	for i := range arr {
		arr[i] = strings.ToLower(arr[i])
	}
	return arr
}

func uploadMap(arr []string) map[string][]string { // построение карты анаграмм
	mapForAnagram = *New() // создадим пустую структуру DefMap

	currentArr := arr // перезапишем массив

	mapForAnagram.m[currentArr[0]] = append(mapForAnagram.m[currentArr[0]], currentArr[0]) // добавим первый элемент как ключ и значение в карту

	isUnique := true

	for i := range currentArr { // проходим по всему массиву слов
		for key, _ := range mapForAnagram.m { // проходим по каждому ключу карты
			isUnique = true
			if isAnagramm(key, currentArr[i]) { // если найдена анаграмма ключа
				for j := range mapForAnagram.m[key] { // проверяем уникальная ли новая анаграмма
					if mapForAnagram.m[key][j] == currentArr[i] {
						isUnique = false
						break
					}
				}
				if isUnique { // если анаграмма уникальная
					mapForAnagram.m[key] = append(mapForAnagram.m[key], currentArr[i]) // добавляем по ключу в словарь анаграмм
					break
				}
			} else {
				if isFirstMeet(currentArr[i]) { // если ключ впервые встретился
					mapForAnagram.m[currentArr[i]] = append(mapForAnagram.m[currentArr[i]], currentArr[i]) // создаем новый ключ
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

	keyWords := strings.Split(key, "") // разделяем ключ и значение на буквы
	strWords := strings.Split(str, "")

	sort.Slice(keyWords, func(i, j int) bool {
		return keyWords[i] < keyWords[j] // сортируем каждую букву ключа
	})

	sort.Slice(strWords, func(i, j int) bool {
		return strWords[i] < strWords[j] // сортируем каждую букву значения
	})

	sortedKeyWords := strings.Join(keyWords, "") // соединяем в строку
	sortedStrWords := strings.Join(strWords, "") // соединяем в строку

	if sortedKeyWords == sortedStrWords { // если получили одинаковые слова - анаграммы
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

func findSingleArr() { // поиск повторов в карте
	for key, val := range mapForAnagram.m {
		if len(val) == 1 {
			delete(mapForAnagram.m, key) // если найден повтор, удаляем
		}
	}
}
