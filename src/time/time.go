package main

import (
	"fmt"
	"time"
)

type Absensi struct {
	Hadir   bool
	Tanggal time.Time
}

type Siswa struct {
	Nama          string
	Usia          int
	NaikKelas     bool
	TanggalLahir  time.Time
	ListKehadiran []Absensi
}

func (s *Siswa) IsNaikKelas() {
	s.NaikKelas = true
}

func main() {
	var siswa1 *Siswa
	var tglLahir time.Time

	tglLahir, err := time.Parse("2006-01-02", "2003-06-10")
	if err != nil {
		fmt.Println(err)
		return
	}

	siswa1 = &Siswa{
		Nama: "budi",
		Usia: 20,
		ListKehadiran: []Absensi{
			{
				Hadir:   false,
				Tanggal: time.Now(),
			},
		},
		TanggalLahir: tglLahir,
	}

	fmt.Println("Kehadiran siswa 1:")
	for _, val := range siswa1.ListKehadiran {
		fmt.Println("Kehadiran: ", val.Hadir)
		fmt.Println("Tanggal: ", val.Tanggal.Format("2006 01 + 02 15:04:05"))
	}

	fmt.Println("Siswa yang bernama ", siswa1.Nama, " lahir pada tanggal ", siswa1.TanggalLahir)
}
