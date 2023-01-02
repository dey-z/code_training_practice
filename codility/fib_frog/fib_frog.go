package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var A []int
	m := make(map[string][]int)
	m["1"] = []int{0, 1, 1, 1, 1, 0}
	m["2"] = []int{0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m["3"] = []int{0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0}
	m["4"] = []int{1}
	m["5"] = []int{0, 1, 0, 1, 0}
	args := os.Args[1:]
	A = m[args[0]]
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	// add last leaf so as to assign target bank
	A = append(A, 1)
	N := len(A)
	// find Fibonacci list
	var F []int
	fibCached := make(map[int]bool)
	// F[0] = 0
	F = append(F, 0)
	// F[1] = 1
	F = append(F, 1)
	// first find fibonacci upto F[2:K] < N
	i := 2
	for i < math.MaxInt32 {
		fib := F[i-1] + F[i-2]
		if fib > N {
			break
		}
		F = append(F, fib)
		fibCached[fib] = true
		i += 1
	}
	// not considering 1st two fibs which are 0 & 1
	F = F[2:]
	fmt.Println(F)
	fmt.Println(fibCached)

	// algorithm
	// 1. get starting leaves that can be reached in one step
	// 2. then find intermediate leaves
	// 3. find min of combinations of 2.

	// initialize reachSteps with 0
	reachSteps := make([]int, N)
	for i := range reachSteps {
		reachSteps[i] = 0
	}

	// leafs that can be reached in one fibonacci step
	i = 0
	for i < N {
		if A[i] == 1 && fibCached[i+1] {
			reachSteps[i] = 1
		}
		i += 1
	}
	fmt.Println(reachSteps)

	// search leafs with more than one step
	i = 0
	for i < N {
		if A[i] == 0 || reachSteps[i] > 0 {
			i += 1
			continue
		}
		minI := -1
		minV := 100000
		for _, fib := range F {
			previousI := i - fib
			fmt.Println(i, fib, previousI)
			if previousI < 0 {
				break
			}
			if reachSteps[previousI] > 0 && minV > reachSteps[previousI] {
				minV = reachSteps[previousI]
				minI = previousI
			}
			if minI != -1 {
				reachSteps[i] = minV + 1
			}
		}
		i += 1
	}

	fmt.Println("")
	fmt.Println(A)
	fmt.Println(reachSteps)
	fmt.Println("")
	if reachSteps[N-1] > 0 {
		return reachSteps[N-1]
	} else {
		return -1
	}
}
