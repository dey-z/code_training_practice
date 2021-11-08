package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for T != 0 {
		var N, K int
		fmt.Scanf("%d %d", &N, &K)
		arr := make([]int, N)
		buf := make([]interface{}, N)
		for i := 0; i < N; i++ {
			buf[i] = &arr[i]
		}
		fmt.Scanln(buf...)
		index := N - (K % N)
		for i := index; i < N; i++ {
			fmt.Printf("%d ", arr[i])
		}
		for i := 0; i < index; i++ {
			fmt.Printf("%d ", arr[i])
		}
		fmt.Println("")
		T -= 1
	}
}
