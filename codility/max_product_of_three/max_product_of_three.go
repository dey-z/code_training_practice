package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var A []int
	m := make(map[string][]int)
	m["1"] = []int{-3, 1, 2, -2, 5, 6}
	m["2"] = []int{-5, 5, -5, 4}
	m["3"] = []int{-615, -879, 376, 832, -342, -369, -335, 216, -378, 826, 706, 451, -984, 823, 718, 901, 696, -811, 179, 369, -234, -899, 800, 629, 856, -965, -97, 655, -529, -58, -46, -971, 466, 737, 244, 820, -620, 743, -249, 228, -317, -359, 923, 165, 306, -382, 740, -711, -113, -667, -481, -227, -558, -865, -563, 619, 914, 922, -598, 766, -641, -414, 153, -23, -395, 290, 644, -352, 478, 758, 711, 898, 949, 652, 495, -741, 74, -592, 232, -848, -84, -110, -755, 233, 990, 773, 917, 331, -86, -597, 288, 418, 231, 924, -409, -517, 390, -633, -779, -311}
	m["4"] = []int{1, 2, 3, 4, 5, 6, 7, 8}
	m["5"] = []int{-10, -2, -4}
	args := os.Args[1:]
	A = m[args[0]]
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	// algorithm
	// sort and get the multiple of the last 3 triplets?
	// you have to take care of -ve * -ve multiplication
	// so it can be the first 3 or the last 3 or in the combination
	N := len(A)
	sort.Ints(A)
	fmt.Println(A)
	var B []int
	// append LHS + RHS
	if N < 7 {
		B = A
	} else {
		B = append(B, A[0:3]...)
		B = append(B, A[N-3:N]...)
	}
	fmt.Println(B)
	max := B[0] * B[1] * B[2]
	i := 0
	for i < len(B) {
		fmt.Println(i)
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
				fmt.Print(j, " ")
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
		fmt.Println(" ")
		i += 1
	}
	return max
}
