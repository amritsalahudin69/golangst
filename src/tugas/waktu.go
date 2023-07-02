package main

import (
	"fmt"
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
	TglBerangkat    string
}

func main() {
	VxTime := time.Now()
	VxTimeFormat := VxTime.Format("2006-01-02 15:04:05")
	VxTimeParsed, err := time.Parse("2006-01-02 15:04:05", VxTimeFormat)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//Date Time Pemesanan
	TahunPesan := VxTimeParsed.Format("2006")
	BulanPesan := VxTimeParsed.Format("01")
	TanggalPesan := VxTimeParsed.Format("02")
	TglGabung := TahunPesan + ":" + BulanPesan + ":" + TanggalPesan

	JamPesan := fmt.Sprintf("%02d", VxTimeParsed.Hour())
	MenitPesan := fmt.Sprintf("%02d", VxTimeParsed.Minute())
	Detik := fmt.Sprintf("%02d", VxTimeParsed.Second())
	JamPesanGabung := JamPesan + ":" + MenitPesan + ":" + Detik
	//Date Time Pemesanan

	//Date Time Berangkat
	VdTime := VxTimeParsed.AddDate(0, 0, 7)
	VdJam := VdTime.Add(10 * time.Hour)

	JamBerangkat := fmt.Sprintf("%02d", VdJam.Hour())
	MenitBerangkat := fmt.Sprintf("%02d", VdJam.Minute())
	DetikBerangkat := fmt.Sprintf("%02d", VdJam.Second())
	JamBerangkatGabung := JamBerangkat + ":" + MenitBerangkat + ":" + DetikBerangkat
	//Date Time Berangkat

	NamaPemesan := NamaCust{
		Nama: "Amri T Salahudin",
		Ktp:  "33262626262626",
		Pesanan: &Pesanan{
			NamaPenerbangan: "Emprit Airlines",
			KotaBerangkat:   "Jepara",
			KotaTujuan:      "Yogyakarta",
			TglPemesanan:    TglGabung,
			TglBerangkat:    VdJam.Format("2006-01-02"),
		},
	}

	fmt.Printf("Tiket Tujuan Kota %s menuju Kota %s pada tanggal %s pukul %s \n", NamaPemesan.Pesanan.KotaBerangkat, NamaPemesan.Pesanan.KotaTujuan, NamaPemesan.Pesanan.TglBerangkat, JamBerangkatGabung)

	fmt.Printf("Tiket ini dipesan pada tanggal %s pukul %s", NamaPemesan.Pesanan.TglPemesanan, JamPesanGabung)

}
