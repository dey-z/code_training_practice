package main

import (
	"fmt"
	"math"
	"os"

	"github.com/dey-z/code_training_practice/codility/flags/testdata"
)

func main() {
	var A []int
	args := os.Args[1:]
	A = testdata.T[args[0]]
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
