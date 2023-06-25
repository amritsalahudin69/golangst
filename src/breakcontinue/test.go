package main

import "fmt"

func main()  {
	for i := 0; i < 10; i++ {
		if i/3 == 2 {
			break
		}

		fmt.Println("bilangan:", i)
	}
}