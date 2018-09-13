package main

import (
	"fmt"
)

// subFactorial counts subFactorial
var subFactorial int

// generate finds possible permutations using Heap algorithm.
func generate(n int, A []int) {
	if n == 1 {
		// fmt.Println(A)
		countsubFactorial(A)
	} else {
		for i := 0; i < n-1; i++ {
			generate(n-1, A)
			if n%2 == 0 {
				// swap
				t := A[i]
				A[i] = A[n-1]
				A[n-1] = t
			} else {
				// swap
				t := A[0]
				A[0] = A[n-1]
				A[n-1] = t
			}
		}
		generate(n-1, A)
	}
}

// countsubFactorial counts subFactorial
func countsubFactorial(A []int) {
	same := 0
	for index, elem := range A {
		// check if elem and its index + 1 is same
		// break if same since that won't count for subFactorial
		if elem == index+1 {
			same++
			break
		}
	}
	// if no same index + 1 and elem, count subFactorial
	if same == 0 {
		subFactorial++
	}
}

func main() {
	var n int
	fmt.Printf("Enter n: ")
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	// initialize array
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	generate(n, arr)
	fmt.Printf("!%d -> %d\n", n, subFactorial)
}
