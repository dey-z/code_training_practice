package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var A []int
	m := make(map[string][]int)
	m["1"] = []int{1}
	m["2"] = []int{1, 1}
	m["3"] = []int{3, 1, 2, 3, 6}
	m["4"] = []int{1, 2, 5, 2, 3}
	m["5"] = []int{3, 1, 2, 3, 6, 12, 4}
	args := os.Args[1:]
	A = m[args[0]]
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

// func allDivisors(n int) []int {
// 	D := []int{}
// 	i := 1
// 	for i <= int(math.Sqrt(float64(n))) {
// 		if n%i == 0 {
// 			if n/i == i {
// 				D = append(D, i)
// 			} else {
// 				D = append(D, i)
// 				D = append(D, n/i)
// 			}
// 		}
// 		i += 1
// 	}
// 	return D
// }
