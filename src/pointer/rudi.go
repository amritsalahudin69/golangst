package main

import "fmt"

type Siswa struct {
	nama      string
	Usia      int
	NaikKelas bool
}

func (s *Siswa) isNaikKelas() {
	s.NaikKelas = true
}

func main() {
	siswa1 := Siswa{
		nama: "budi",
		Usia: 17,
	}

	siswa1.isNaikKelas()
	fmt.Println(siswa1)
}
