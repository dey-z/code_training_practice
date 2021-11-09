package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var T int
	fmt.Scanf("%d", &T)
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 1000000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	cnt := 0
	var N, K int
	for scanner.Scan() {
		var line string
		line = scanner.Text()
		if cnt%2 != 0 {
			nums := strings.Fields(line)
			index := N - (K % N)
			fmt.Print(strings.Trim(fmt.Sprint(nums[index:N]), "[]"))
			if len(nums[index:N]) > 0 && len(nums[0:index]) > 0 {
				fmt.Print(" ")
			}
			fmt.Print(strings.Trim(fmt.Sprint(nums[0:index]), "[]"))
			fmt.Println("")
		} else {
			N, _ = strconv.Atoi(strings.Split(line, " ")[0])
			K, _ = strconv.Atoi(strings.Split(line, " ")[1])
		}
		cnt += 1
		if cnt >= T*2 {
			break
		}
	}
}
