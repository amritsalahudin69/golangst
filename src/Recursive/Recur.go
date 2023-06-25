package main

import "fmt"

func hitungMundur(angka int) int {
	if angka > 0 {
		fmt.Println(angka)
		return hitungMundur(angka -1)
	}
	return 0
}

func main() {
	fmt.Println(hitungMundur(10))
}