package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/dey-z/code_training_practice/codility/perm_missing_elem/testdata"
)

func main() {
	var A []int
	args := os.Args[1:]
	A = testdata.T[args[0]]
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	// sort the array
	sort.Ints(A)
	N := len(A)
	// 1st outlier empty array
	// 2nd outlier 1 is not present
	if N == 0 || A[0] != 1 {
		return 1
	}
	// loop to check neigbours if its not +1 then break
	res := 0
	i := 0
	for i < N {
		if (i+1 >= N) || (A[i]+1 != A[i+1]) {
			res = A[i] + 1
			break
		}
		i += 1
	}
	return res
}
