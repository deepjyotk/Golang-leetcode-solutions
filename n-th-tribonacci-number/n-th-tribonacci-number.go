package main

import "fmt"

func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}
	a, b, c := 0, 1, 1
	curr := 0
	for i := 3; i <= n; i++ {
		curr = a + b + c

		a = b
		b = c
		c = curr
	}

	return curr
}

func main() {
	t1 := tribonacci(1)
	t2 := tribonacci(2)
	t3 := tribonacci(3)
	t4 := tribonacci(10)

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
	fmt.Println(t4)

}
