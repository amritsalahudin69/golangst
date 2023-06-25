package main

import "fmt"

type siswa struct {
	nama string
	listhadir []kehadiran
}

type kehadiran struct {
	hadir bool
	tanggal int
}

func main()  {
	siswa1 := siswa{
		nama: "Openg",
		listhadir:[]kehadiran{
			{
				hadir: true,
				tanggal: 1,
			},
			{
				hadir: false,
				tanggal: 2,
			},
			{
				hadir: true,
				tanggal: 3,
			},
			{
				hadir: true,
				tanggal: 4,
			},
		},
	}

	fmt.Println("siswa bernama:", siswa1. nama)

	for _, absensi := range siswa1.listhadir {
		if absensi.hadir{
			fmt.Println("hadir pada tanggal", absensi.tanggal)
		} else {
			fmt.Println("tidak hadir pada tanggal:", absensi.tanggal)
		}
	}
}