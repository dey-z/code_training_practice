package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for T != 0 {
		var N, K int
		fmt.Scanf("%d %d", &N, &K)
		arr := make([]int, N)
		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')
		strs := strings.Split(line[0:len(line)-1], " ")
		var wg sync.WaitGroup
		wg.Add(N)
		for i := 0; i < N; i++ {
			go func(i int) {
				defer wg.Done()
				arr[i], _ = strconv.Atoi(strs[i])
			}(i)
		}
		wg.Wait()
		index := N - (K % N)
		fmt.Print(strings.Trim(fmt.Sprint(arr[index:N]), "[]"))
		if len(arr[index:N]) > 0 {
			fmt.Print(" ")
		}
		fmt.Print(strings.Trim(fmt.Sprint(arr[0:index]), "[]"))
		fmt.Println("")
		T -= 1
	}
}
