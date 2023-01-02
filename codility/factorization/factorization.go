package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	x, _ := strconv.Atoi(args[0])
	fmt.Println(Solution(x))
}

func Solution(x int) []int {
	primeFactors := []int{}
	// get smallest prime divisors array upto that number using Sieve of Eratosthenes algorithm
	F := factorizationArr(x)
	fmt.Println("smallest prime divisors array")
	fmt.Println(F)
	for F[x] > 0 {
		primeFactors = append(primeFactors, F[x])
		x = x / F[x]
	}
	primeFactors = append(primeFactors, x)
	return primeFactors
}

func factorizationArr(n int) []int {
	F := make([]int, n+1)
	for i := range F {
		F[i] = 0
	}
	i := 2
	for i*i <= n {
		if F[i] == 0 {
			k := i * i
			for k <= n {
				if F[k] == 0 {
					F[k] = i
				}
				k += i
			}
		}
		i += 1
	}
	return F
}
