package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/dey-z/code_training_practice/codility/max_product_of_three/testdata"
)

func main() {
	var A []int
	args := os.Args[1:]
	A = testdata.T[args[0]]
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	// algorithm
	// sort and get the multiple of the last 3 triplets?
	// you have to take care of -ve * -ve multiplication
	// so it can be the first 3 or the last 3 or in the combination
	N := len(A)
	sort.Ints(A)
	var B []int
	// append LHS + RHS
	if N < 7 {
		B = A
	} else {
		B = append(B, A[0:3]...)
		B = append(B, A[N-3:N]...)
	}
	max := B[0] * B[1] * B[2]
	i := 0
	for i < len(B) {
		j := 0
		cnt := 1
		mult := B[i]
		for j < len(B) {
			if j == i {
				j += 1
				continue
			} else {
				if cnt < 3 {
					mult = mult * B[j]
					cnt += 1
				}
			}
			if cnt == 3 {
				// check max
				if max < mult {
					max = mult
				}
				// reset mult & cnt & j
				mult = B[i]
				cnt = 1
				j -= 1
			}
			j += 1
		}
		i += 1
	}
	return max
}
