package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var A []int
	m := make(map[string][]int)
	m["1"] = []int{2, 3, 5, 7, 1, 6}
	m["2"] = []int{1, 3, 4, 6, 2, 5}
	m["3"] = []int{1, 2, 3, 4, 5, 6}
	m["4"] = []int{}
	m["5"] = []int{2, 3, 4, 5, 6, 7}
	args := os.Args[1:]
	A = m[args[0]]
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
