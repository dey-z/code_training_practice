package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var args []int
	for _, n := range os.Args[1:] {
		val, _ := strconv.Atoi(n)
		args = append(args, val)
	}
	fmt.Println(Solution(args[0], args[1], args[2]))
}

func Solution(X int, Y int, D int) int {
	var diff float64
	diff = float64(Y - X)
	jumps := diff / float64(D)
	return int(math.Ceil(jumps))
}
