package main

import (
	"strconv"
	"strings"
)

func main() {
	var arr [2022]int
	arr[0] = 1
	res := strings.Join(arr, "")

	val := strconv.Atoi(res)

}
