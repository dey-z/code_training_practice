package main

import (
	"fmt"
	"math"
	"os"

	"github.com/dey-z/code_training_practice/codility/tape_equilibrium/testdata"
)

func main() {
	var A []int
	args := os.Args[1:]
	A = testdata.T[args[0]]
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	N := len(A)
	// algorithm
	// TSUM = LS + RS
	// LS - RS = LS - (TSUM - LS) = 2LS - TSUM
	ls := A[0]
	tsum := sum(A[0:N])
	res := math.MaxInt32
	i := 1
	for i < N {
		diff := diff(2*ls, tsum)
		// fmt.Println(diff)
		if res > diff {
			res = diff
		}
		ls += A[i]
		i += 1
	}

	// fmt.Println("")
	// fmt.Println("RESULT")
	return res
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
