package main

import "fmt"

func kehadiranSiswa(nama string, absensi ...bool,) {
	fmt.Println("Siswa bernama:", nama)
	for index, Hadir := range absensi{
		if Hadir {
			fmt.Println("Hari ke-", index+1, "Hadir")
		} else {
			fmt.Println("Hari ke-", index+1, "Tidak Hadir")
		}
		
	}
}

func main()  {
	kehadiranSiswa("Openg", true, true,false,false,true)
}