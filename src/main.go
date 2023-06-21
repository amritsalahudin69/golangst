package main

import "fmt"

//func pertambahan(bill int, bill2 int) int {
//	return bill + bill2
//}
//
//func main() {
//	hasil := pertambahan(3, 5)
//	fmt.Println(hasil)
//}

func kehadiran(nama string, absensi ...bool) {
	fmt.Println("siswa bernama : ", nama)
	for index, hadir := range absensi {
		if hadir {
			fmt.Println("hari ke-1", index+1, ": Hadir")
		} else {
			fmt.Println("hari ke-1", index+1, ": tidak Hadir")
		}
	}
}

func main() {
	kehadiran("arif", true, true, false, true, false)
}
