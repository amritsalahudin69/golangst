package main

import (
	"time"
)

type NamaCust struct {
	Nama    string
	Ktp     string
	Pesanan *Pesanan
}

type Pesanan struct {
	NamaPenerbangan string
	KotaBerangkat   string
	KotaTujuan      string
	TglPemesanan    string
	TglBerangkat    time.Time
}

func main() {
	VxTime := time.Now()
	VTglPesan := VxTime.Format("2006-01-02 15:04:05")
	NamaPemesan := NamaCust{
		Nama: "Amri T Salahudin",
		Ktp:  "33262626262626",
		Pesanan: &Pesanan{
			NamaPenerbangan: "Emprit Airlines",
			KotaBerangkat:   "Jepara",
			KotaTujuan:      "Yogyakarta",
			TglPemesanan:    VTglPesan,
			TglBerangkat:    time.Now(),
		},
	}

	//fmt.Printf("Tiket Tujuan Kota %s menuju Kota %s pada tanggal %s pukul Jam:Menit:Detik \n", NamaPemesan.Pesanan.KotaBerangkat, NamaPemesan.Pesanan.KotaTujuan, NamaPemesan.Pesanan.TglPemesanan)

}
