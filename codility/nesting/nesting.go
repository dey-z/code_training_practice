package main

import (
	"fmt"
	"os"
)

var stack []string

func main() {
	var S string
	m := make(map[string]string)
	m["1"] = "(())"
	m["2"] = "(()(())())"
	m["3"] = "())"
	m["4"] = ""
	m["5"] = ")("
	args := os.Args[1:]
	S = m[args[0]]
	fmt.Println(Solution(S))
}

func Solution(S string) int {
	// algorithm
	// 1 stack one for open/close bracket
	// get ( then open bracket(push open)
	// get ) then close bracket(pop top open)
	// the order for check should be ( -> )
	// ultimately all brackets should be closed or stack empty
	N := len(S)
	// edge case
	if N == 0 {
		return 1
	}
	// all cases
	i := 0
	for i < N {
		// open
		if string(S[i]) == "(" {
			// push
			stack = push(stack)
		} else if string(S[i]) == ")" {
			// IF nothing there to pop THEN not nesting return 0
			if len(stack) == 0 {
				return 0
			}
			// pop top open
			stack = pop(stack)
		}
		i += 1
	}
	// IF stack empty THEN nesting
	if len(stack) == 0 {
		return 1
	}
	// not nesting
	return 0
}

func pop(stack []string) []string {
	n := len(stack) - 1
	return stack[:n]
}

func push(stack []string) []string {
	stack = append(stack, "open")
	return stack
}
