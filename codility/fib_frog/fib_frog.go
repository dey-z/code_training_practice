package main

import (
	"fmt"
	"os"

	"github.com/dey-z/code_training_practice/codility/fib_frog/testdata"
)

func main() {
	var A []int
	args := os.Args[1:]
	A = testdata.T[args[0]]
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	// add last leaf so as to assign target bank
	A = append(A, 1)
	N := len(A)

	// find Fibonacci list
	F := []int{0, 1}
	for i := 2; F[i-1]+F[i-2] <= N; i++ {
		F = append(F, F[i-1]+F[i-2])
	}

	// initialize reachSteps with -1
	reachSteps := make([]int, N)
	for i := range reachSteps {
		reachSteps[i] = -1
	}

	// find the minimum number of jumps to each position
	for i := 0; i < len(A); i++ {
		if A[i] == 1 || i == len(A)-1 {
			minStep := -1
			for _, fib := range F {
				prevPos := i - fib
				if prevPos < -1 {
					break
				}
				if prevPos == -1 {
					minStep = 1
					break
				}
				if reachSteps[prevPos] > 0 {
					if minStep == -1 || reachSteps[prevPos]+1 < minStep {
						minStep = reachSteps[prevPos] + 1
					}
				}
			}
			reachSteps[i] = minStep
		}
	}

	return reachSteps[N-1]
}
