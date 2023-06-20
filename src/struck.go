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
	isActive    bool
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
	fmt.Println("data 1: ", warungAmri)
}
