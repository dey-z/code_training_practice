package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	var N int
	for scanner.Scan() {
		var line string
		line = scanner.Text()
		if cnt%2 != 0 {
			// input string array
			numsStr := strings.Fields(line)
			var nums []int
			// conv to int array
			for _, n := range numsStr {
				val, _ := strconv.Atoi(n)
				nums = append(nums, val)
			}
			// sort int num array
			sort.Ints(nums)
			// minXor init
			minXor := math.MaxInt32
			// xor for consecutive pair
			i := 0
			for i < N-1 {
				val := nums[i] ^ nums[i+1]
				if val < minXor {
					minXor = val
				}
				i += 1
			}
			fmt.Print(minXor)
			fmt.Println("")
		} else {
			split := strings.Split(line, " ")
			N, _ = strconv.Atoi(split[0])
		}
		cnt += 1
		if cnt >= T*2 {
			break
		}
	}
}
