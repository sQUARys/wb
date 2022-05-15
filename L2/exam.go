package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Find(a []string, sub string) bool {
	isFind := false
	for i := range a {
		if a[i] == sub {
			isFind = true
			break
		}
	}
	return isFind
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := strings.Split(scanner.Text(), " ")
	N, _ := strconv.Atoi(text[0])
	if N == 0 {
		fmt.Println("Input error.")
		return
	}
	Q, _ := strconv.Atoi(text[1])
	if Q == 0 {
		fmt.Println("No")
		return
	}
	var facts []int
	for i := 0; i < Q; i++ {
		scanner.Scan()
		arrIn := strings.Split(scanner.Text(), " ")
		for j := 0; j < 2; j++ {
			in, _ := strconv.Atoi(arrIn[j])
			facts = append(facts, in)
		}
	}

	var knownEl []string

	for i := 1; i < len(facts); i++ {
		toPut1 := strconv.Itoa(facts[i-1])
		if !Find(knownEl, toPut1) {
			knownEl = append(knownEl, toPut1)
		}
		toPut2 := strconv.Itoa(facts[i])
		if !Find(knownEl, toPut2) {
			knownEl = append(knownEl, toPut2)
		}
	}

	if len(knownEl) >= N {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
