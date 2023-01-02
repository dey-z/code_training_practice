package main

import (
	"fmt"
	"os"
)

var stack []int

func main() {
	var A []int
	m := make(map[string][]int)
	m["1"] = []int{4, 3, 4, 4, 4, 2}
	m["2"] = []int{6, 8, 4, 6, 8, 6, 6}
	m["3"] = []int{5, 5, 5}
	m["4"] = []int{1, 2}
	args := os.Args[1:]
	A = m[args[0]]
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	// algorithm
	// 1st find dominator/leader
	// 2nd occurences of dominator(n)
	// 3rd count for occurences by left(c)/right(n-c)
	// 3rd magjik IF [(i+1<2*c) && SizeA - i - 1 < 2*(n-c)] THEN equiLdrCnt++
	N := len(A)
	i := 0
	value := 0
	d := 0
	n := 0
	c := 0
	equiLdrCnt := 0
	for i < N {
		if len(stack) == 0 {
			value = A[i]
			stack = push(stack, value)
			fmt.Println(stack)
		} else {
			if value != A[i] {
				stack = pop(stack)
			} else {
				stack = push(stack, value)
			}
			fmt.Println(stack)
		}
		i += 1
	}
	// edge case or there is no dominator
	if len(stack) == 0 {
		return 0
	}
	// dominator/leader
	d = stack[0]
	// occurences of dominator
	i = 0
	for i < N {
		if A[i] == d {
			n += 1
		}
		i += 1
	}
	// find equiLdrCnt
	i = 0
	for i < N {
		if A[i] == d {
			c += 1
		}
		if i+1 < 2*c && N-i-1 < 2*(n-c) {
			equiLdrCnt += 1
		}
		i += 1
	}
	return equiLdrCnt
}

func pop(stack []int) []int {
	n := len(stack) - 1
	return stack[:n]
}

func push(stack []int, num int) []int {
	stack = append(stack, num)
	return stack
}
