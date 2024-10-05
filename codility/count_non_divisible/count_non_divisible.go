package main

import (
	"fmt"
	"math"
	"os"

	"github.com/dey-z/code_training_practice/codility/count_non_divisible/testdata"
)

func main() {
	var A []int
	args := os.Args[1:]
	A = testdata.T[args[0]]
	fmt.Println(Solution(A))
}

func Solution(A []int) []int {
	// algorithm
	N := len(A)
	nonDivisibles := make([]int, N)
	presence := make(map[int]int)
	cached := make(map[int]int)
	i := 0
	for i < N {
		// get map of number of times a number is present in the array
		presence[A[i]] += 1
		// inititate non divisible to -1 as 0 is also possible
		cached[A[i]] = -1
		i += 1
	}
	i = 0
	// iterate thru array
	for i < N {
		if cached[A[i]] < 0 {
			cnt := N
			j := 1
			n := A[i]
			for j <= int(math.Sqrt(float64(n))) {
				if n%j == 0 {
					if n/j == j {
						cnt = cnt - presence[j]
					} else {
						cnt = cnt - presence[j]
						cnt = cnt - presence[n/j]
					}
				}
				j += 1
			}
			nonDivisibles[i] = cnt
			cached[A[i]] = cnt
		} else {
			nonDivisibles[i] = cached[A[i]]
		}
		i += 1
	}
	return nonDivisibles
}
