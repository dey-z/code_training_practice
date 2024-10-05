package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/dey-z/code_training_practice/codility/triangle/testdata"
)

func main() {
	var A []int
	args := os.Args[1:]
	A = testdata.T[args[0]]
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	N := len(A)
	// edge case
	if N < 3 {
		return 0
	}
	// sort the array
	sort.Ints(A)
	B := A
	// remove -ve nums since a > 0 & b > 0 & c > 0
	i := 0
	// if first element is -ve then there are -ve nums
	if A[i] < 0 {
		for i < N {
			if A[i] < 0 {
				B = A[i+1 : N]
			}
			if A[i] > 0 {
				// after 1st +ve found no need to continue
				break
			}
			i += 1
		}
	}

	// fmt.Println(B)
	// edge case
	if len(B) < 3 {
		return 0
	}
	i = 0
	for i < len(B)-2 {
		P := B[i]
		Q := B[i+1]
		R := B[i+2]
		if check(P, Q, R) {
			return 1
		}
		i += 1
	}

	return 0
}

func check(P, Q, R int) bool {
	if P+Q > R && P+R > Q && Q+R > P {
		return true
	}
	return false
}
