package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var A []int
	m := make(map[string][]int)
	m["1"] = []int{}
	m["2"] = []int{0}
	m["3"] = []int{-1}
	m["4"] = []int{1}
	m["5"] = []int{-1, 1}
	m["6"] = []int{1, 1}
	m["7"] = []int{-1, -1}
	args := os.Args[1:]
	A = m[args[0]]
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	return 0
}

func sum(arr []int) int {
	result := 0
	for _, v := range arr {
		result += v
	}
	return result
}

func diff(a, b int) int {
	return int(math.Abs(float64(a) - float64(b)))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
