package main

import "fmt"

func fibonacci(n int, BilanganFib1 int, bilanganfib2 int) int {
	if n > 0 {
		fmt.Println(BilanganFib1)
		return fibonacci(n-1, bilanganfib2, BilanganFib1+bilanganfib2)
	}
	return BilanganFib1
}

func main() {
	hasil := fibonacci(5, 0, 1)
	fmt.Println(hasil)
}
