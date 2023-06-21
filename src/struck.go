package main

import (
	"fmt"
	"time"
)

type WarungWaralabaSuryaAbadi struct {
	namaCounter  string
	kodeWarung   string
	alamatWarung string
	nomorTelfon  string
	karyawan     HrEmp
	menu         []MenuWarung
}

type MenuWarung struct {
	namaMasakan string
	harga       float32
	tersedia    bool
}

type HrEmp struct {
	name           string
	jabatan        string
	gaji           float32
	detailKaryawan detail
	listabsen      []Absen
}

type detail struct {
	joindate       time.Weekday
	alamatKaryawan string
	isActive       bool
}

type Absen struct {
	dateAbsen time.Month
	timein    time.Time
	timeout   time.Time
}

func main() {
	warungAmri := WarungWaralabaSuryaAbadi{
		namaCounter:  "Warung Amri 99",
		kodeWarung:   "6977",
		alamatWarung: "Jl. Maju Terus",
		nomorTelfon:  "0812242741724",
		karyawan: HrEmp{
			name:    "Kinanthi",
			jabatan: "Kepala Warung",
			gaji:    10000000,
			detailKaryawan: detail{
				joindate:       time.Monday,
				alamatKaryawan: "Jepara",
				isActive:       true,
			},
		},
	}
	listabsen1 := make([]Absen, 0)
	listabsen1 = append(listabsen1, Absen{
		dateAbsen: time.December,
		timein:    time.Now(),
		timeout:   time.Now().Add(time.Hour * 6),
	})
	warungAmri.karyawan.listabsen = listabsen1

	MenuWarungs := make([]MenuWarung, 0)
	MenuWarungs = append(MenuWarungs, MenuWarung{
		namaMasakan: "Rendang",
		harga:       10000,
		tersedia:    true,
	})
	MenuWarungs = append(MenuWarungs, MenuWarung{
		namaMasakan: "Ayam KiEfci",
		harga:       5000,
		tersedia:    true,
	})
	warungAmri.menu = MenuWarungs
	fmt.Println("data 1: ", warungAmri)
	fmt.Println("data 1: ", warungAmri.karyawan.detailKaryawan)
	fmt.Println("data 1: ", warungAmri.karyawan.listabsen)
	fmt.Println("data 1: ", warungAmri.menu)
}
