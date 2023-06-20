package main

import "fmt"

type siswasalt struct {
	nama        string
	usia        int
	basicIT     bool
	detail      detailSiswa
	listAbsensi []absensi
}

type detailSiswa struct {
	alamat       string
	jenisKelamin string
}
type absensi struct {
	hadir   bool
	hari    string
	tanggal int
	bulan   string
	tahun   int
}

func main() {
	siswa1 := siswasalt{
		nama:    "amri",
		usia:    20,
		basicIT: true,
		detail: detailSiswa{
			alamat:       "welahan",
			jenisKelamin: "lakilaki",
		},
	}
	absensiSiswa1 := make([]absensi, 0)
	absensiSiswa1 = append(absensiSiswa1, absensi{
		hadir:   true,
		hari:    "senin",
		tanggal: 1,
		bulan:   "Juni",
		tahun:   2023,
	})
	absensiSiswa1 = append(absensiSiswa1, absensi{
		hadir:   false,
		hari:    "selasa",
		tanggal: 2,
		bulan:   "Juni",
		tahun:   2023,
	})

	siswa2 := siswasalt{
		nama:    "gendis",
		usia:    13,
		basicIT: true,
		detail: detailSiswa{
			alamat:       "bugo",
			jenisKelamin: "perempuan",
		},
	}
	absensiSiswa2 := make([]absensi, 0)
	absensiSiswa2 = append(absensiSiswa2, absensi{
		hadir:   true,
		hari:    "senin",
		tanggal: 1,
		bulan:   "Juni",
		tahun:   2023,
	})
	absensiSiswa2 = append(absensiSiswa2, absensi{
		hadir:   true,
		hari:    "selasa",
		tanggal: 2,
		bulan:   "Juni",
		tahun:   2023,
	})

	fmt.Println("Nama Siswa 1: ", siswa1.nama)
	fmt.Println("Usia Siswa 1: ", siswa1.usia)
	fmt.Println("Basic IT Siswa 1: ", siswa1.basicIT)
	fmt.Println("Alamat Siswa 1: ", siswa1.detail.alamat)
	fmt.Println("Jenis Kelamin Siswa 1: ", siswa1.detail.jenisKelamin)
	fmt.Println("Absensi Siswa 1: ", siswa1.listAbsensi)
	fmt.Println()
	fmt.Println("Nama Siswa 2: ", siswa2.nama)
	fmt.Println("Usia Siswa 2: ", siswa2.usia)
	fmt.Println("Basic IT Siswa 2: ", siswa2.basicIT)
	fmt.Println("Alamat Siswa 2: ", siswa2.detail.alamat)
	fmt.Println("Jenis Kelamin Siswa 2: ", siswa2.detail.jenisKelamin)
	fmt.Println("Absensi Siswa 2: ", siswa2.listAbsensi)

}
