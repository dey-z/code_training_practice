package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/dey-z/code_training_practice/codility/missing_integer/testdata"
)

func main() {
	var A []int
	args := os.Args[1:]
	A = testdata.T[args[0]]
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	// algorithm
	// sort the array and make it sequential
	sort.Ints(A)
	N := len(A)
	res := A[0]
	// 1st outlier IF only single integer is present in array which is 1 THEN return 2
	// 2nd outlier IF only single integer is present in array THEN return 1
	if N == 1 && A[0] == 1 {
		return 2
	} else if N == 1 && A[0] != 1 {
		return 1
	}
	var cached []int
	checked := make(map[int]bool)
	i := 0
	for i < N {
		// IF not already checked & the integer is positive THEN put in cached
		if !checked[A[i]] && A[i] > 0 {
			cached = append(cached, A[i])
			checked[A[i]] = true
		}
		i += 1
	}
	// fmt.Println(cached)
	if len(cached) == 0 || cached[0] != 1 {
		return 1
	} else {
		i = 0
		for i < len(cached) {
			if (i+1 >= len(cached)) || (cached[i]+1 != cached[i+1]) {
				res = cached[i] + 1
				break
			}
			i += 1
		}
	}
	return res
}
