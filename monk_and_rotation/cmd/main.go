package main

import (
	"fmt"
	"strings"
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
		fmt.Print(strings.Trim(fmt.Sprint(arr[index:N]), "[]"))
		if len(arr[index:N]) > 0 && len(arr[0:index]) > 0 {
			fmt.Print(" ")
		}
		fmt.Print(strings.Trim(fmt.Sprint(arr[0:index]), "[]"))
		fmt.Println("")
		T -= 1
	}
}
