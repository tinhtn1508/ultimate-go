package main

import "fmt"

func binPow(a, n int) int {
	if n == 0 {
		return 1
	}
	res := binPow(a, n/2)
	if n%2 == 0 {
		return res * res
	} else {
		return res * res * a
	}
}

func binPow2(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res *= a
		}
		a *= a
		n >>= 1
	}
	return res
}

func main() {
	fmt.Println("binPow: 2 ^ 10 =", binPow(2, 10))
	fmt.Println("binPow: 3 ^ 13 =", binPow(3, 13))
	fmt.Println("binPow2: 2 ^ 10 =", binPow2(2, 10))
	fmt.Println("binPow2: 3 ^ 13 =", binPow2(3, 13))
}
