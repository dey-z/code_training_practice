package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	T := 1
	N := 5
	K := 2
	// fmt.Println("Enter T")
	fmt.Scanf("%d", &T)
	// fmt.Println("Enter N K")
	fmt.Scanf("%d %d", &N, &K)
	// fmt.Println("Enter Input Array")
	arr := make([]int, N)
	_, err := fmt.Fscan(os.Stdin, packAddrs(arr)...)
	if err != nil {
		os.Exit(0)
	}
	// fmt.Fprintf(os.Stdout, "Input is T=%d N=%d K=%d arr=%v\n", T, N, K, arr)
	s, _ := json.Marshal(businessLogic(T, N, K, arr))
	str := strings.Trim(string(s), "[]")
	tkns := strings.Split(str, ",")
	fmt.Printf("%v\n", strings.Join(tkns, " "))
}

func packAddrs(n []int) []interface{} {
	p := make([]interface{}, len(n))
	for i := range n {
		p[i] = &n[i]
	}
	return p
}

func businessLogic(Tests, Num, Ksteps int, arr []int) (newArr []int) {
	newArr = make([]int, Num)
	n := 1
	for n <= Tests {
		steps := 1
		for steps <= Ksteps {
			newArr = make([]int, Num)
			for i := 0; i < len(arr); i++ {
				rightPos := i + 1
				if rightPos == len(arr) {
					rightPos = 0
				}
				newArr[rightPos] = arr[i]
			}
			arr = newArr
			steps += 1
		}
		n += 1
	}
	return
}
