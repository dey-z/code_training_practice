package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	N, _ := strconv.Atoi(args[0])
	M, _ := strconv.Atoi(args[1])
	fmt.Println(Solution(N, M))
}

// efficient solution
func Solution(N, M int) int {
	// algorithm
	// lcm(N,M) / M = N / gcd(N,M)
	return N / gcd(N, M)
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
}
