package main

import "fmt"

func main() {
	var X int
	X = 5

	for baris := 1; baris <= X; baris++ {
		for kolom := X; kolom >= baris; kolom-- {
			fmt.Print("_")
		}
		for sisi := 1; sisi <= baris; sisi++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}
