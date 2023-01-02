package main

import (
	"fmt"
	"os"
)

var A []int
var C []int
var G []int
var T []int

var a = 0
var c = 0
var g = 0
var t = 0

func main() {
	s := make(map[string]string)
	p := make(map[string][]int)
	q := make(map[string][]int)
	// String S
	s["1"] = "CAGCCTA"
	// Array P
	p["1"] = []int{2, 5, 0}
	// Array Q
	q["1"] = []int{4, 5, 6}
	args := os.Args[1:]
	S := s[args[0]]
	P := p[args[0]]
	Q := q[args[0]]
	fmt.Println(Solution(S, P, Q))
}

// efficient solution
func Solution(S string, P, Q []int) []int {
	// algorithm
	// O(N + M)
	K := len(P)
	N := len(S)
	res := []int{}
	i := 0
	// 1st pass
	for i < N {
		if string(S[i]) == "A" {
			a += 1
			A = append(A, a)
			C = append(C, c)
			G = append(G, g)
			T = append(T, t)
		} else if string(S[i]) == "C" {
			c += 1
			A = append(A, a)
			C = append(C, c)
			G = append(G, g)
			T = append(T, t)
		} else if string(S[i]) == "G" {
			g += 1
			A = append(A, a)
			C = append(C, c)
			G = append(G, g)
			T = append(T, t)
		} else if string(S[i]) == "T" {
			t += 1
			A = append(A, a)
			C = append(C, c)
			G = append(G, g)
			T = append(T, t)
		}
		i += 1
	}
	fmt.Println(A)
	fmt.Println(C)
	fmt.Println(G)
	fmt.Println(T)
	fmt.Println("")
	// 2nd pass
	i = 0
	for i < K {
		// edge case
		if P[i] == Q[i] {
			if string(S[P[i]]) == "A" {
				res = append(res, 1)
			} else if string(S[P[i]]) == "C" {
				res = append(res, 2)
			} else if string(S[P[i]]) == "G" {
				res = append(res, 3)
			} else if string(S[P[i]]) == "T" {
				res = append(res, 4)
			}
		} else {
			// change has occurred within boundary or already changed at start
			if A[P[i]] < A[Q[i]] || string(S[P[i]]) == "A" {
				res = append(res, 1)
			} else if C[P[i]] < C[Q[i]] || string(S[P[i]]) == "C" {
				res = append(res, 2)
			} else if G[P[i]] < G[Q[i]] || string(S[P[i]]) == "G" {
				res = append(res, 3)
			} else if T[P[i]] < T[Q[i]] || string(S[P[i]]) == "T" {
				res = append(res, 4)
			}
		}
		i += 1
	}
	return res
}
