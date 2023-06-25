package main

import "fmt"

func main()  {
	siswa := map[string]interface{}{
		"Nama": "Openg",
		"Umur": 18,
		"Bekerja": false,
	}

	for index, value := range siswa {
		fmt.Println(index, ":", value)
	}
}