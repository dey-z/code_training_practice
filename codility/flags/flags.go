package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var A []int
	m := make(map[string][]int)
	m["1"] = []int{}
	m["2"] = []int{0}
	m["3"] = []int{1}
	m["4"] = []int{1, 1}
	m["5"] = []int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}
	m["6"] = []int{1, 2, 7}
	m["7"] = []int{1, 5, 2, 9}
	m["8"] = []int{3, 2, 1}
	m["9"] = []int{0, 0, 0, 0, 0, 1, 0, 1, 0, 1}
	args := os.Args[1:]
	A = m[args[0]]
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	// algorithm
	// find peaks
	// find prefix_sums of peaks
	// max flags = len(peaks)
	// calculate how many flags are practically possible
	N := len(A)
	peaks := []int{}
	i := 1
	for i < N-1 {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
		i += 1
	}
	// edge case
	if len(peaks) == 0 {
		// also cannot set flags if length of peaks is 0
		return 0
	} else if len(peaks) == 1 {
		// 1 peak then only 1 flag
		return 1
	}
	max_flags := 0
	k := min(len(peaks), int(math.Sqrt(float64(N)))+1)
	for k >= 1 {
		i := 1
		lastPos := i - 1
		cnt := 1
		// iterate through all peaks
		for i < len(peaks) {
			if peaks[i]-peaks[lastPos] >= k {
				cnt += 1
				lastPos = i
			}
			if cnt == k {
				if max_flags < cnt {
					max_flags = cnt
				} else {
					return max_flags
				}
				break
			}
			i += 1
		}
		k -= 1
	}
	return max_flags
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
