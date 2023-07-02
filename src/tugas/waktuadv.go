package main

import (
	"fmt"
	"time"
)

type PesananCustomer struct {
	NamaPenerbangan  string
	KotaBerangkat    string
	KotaTujuan       string
	TglKeberangkatan time.Time
	TglPemesanan     time.Time
}

func (p *PesananCustomer) AddNamapenerbangan() string {
	return p.NamaPenerbangan
}

func (p *PesananCustomer) AddKotaBerangkat() string {
	return p.KotaBerangkat
}

func (p *PesananCustomer) AddKotaTujuan() string {
	return p.KotaTujuan
}

func (p *PesananCustomer) AddTglKeberangkatan() time.Time {
	return p.TglKeberangkatan
}

func (p *PesananCustomer) AddTglPemesanan(TglPemesanan *time.Time) {
	var TglPesan time.Time
	TglPesan = time.Now()
	XX := TglPesan.Format("2006-01-02 15:04:05")
	TglPesanParsed, err := time.Parse("2006-01-02 15:04:05", XX)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	TahunPesan := TglPesanParsed.Format("2006")
	BulanPesan := TglPesanParsed.Format("01")
	TanggalPesan := TglPesanParsed.Format("02")
	TglPemesanan = TahunPesan + ":" + BulanPesan + ":" + TanggalPesan
	return p.TglPemesanan
}
func main() {
	//var TglPesan time.Time
	//TglPesan = time.Now()

	//XX := TglPesan.Format("2006-01-02 15:04:05")
	//TglBerangkat := TglPesan.AddDate(0, 0, 7)

	//TglPesanParsed, err := time.Parse("2006-01-02 15:04:05", XX)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}

	//TglBerangkatX, err := time.Parse("2006-01-02 15:04:05", TglBerangkat)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//TglPesanString := TglPesanParsed.Format("2006-01-02 15:04:05")

	TiketPesawat := PesananCustomer{"Emprit Airlines", "Jepara", "Yogyakarta", TglBerangkat, TglPesan}

	fmt.Println("TglPesanString:", TiketPesawat)
	fmt.Println("TglPesanString:", TglPesanParsed)

	fmt.Printf("Tiket Tujuan Kota %s menuju Kota %s pada tanggal %s pukul %s \n", TiketPesawat.AddKotaBerangkat(), TiketPesawat.AddKotaTujuan(), TiketPesawat.AddTglKeberangkatan(), JamBerangkatGabung)

	//fmt.Printf("Tiket ini dipesan pada tanggal %s pukul %s", TglGabung, JamPesanGabung)

	//VxTime := time.Now()
	//VxTimeFormat := VxTime.Format("2006-01-02 15:04:05")
	//VxTimeParsed, err := time.Parse("2006-01-02 15:04:05", VxTimeFormat)
	//
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	////Date Time Pemesanan
	//TahunPesan := VxTimeParsed.Format("2006")
	//BulanPesan := VxTimeParsed.Format("01")
	//TanggalPesan := VxTimeParsed.Format("02")
	//TglGabung := TahunPesan + ":" + BulanPesan + ":" + TanggalPesan
	////fmt.Println(TglGabung)
	//
	//JamPesan := fmt.Sprintf("%02d", VxTimeParsed.Hour())
	//MenitPesan := fmt.Sprintf("%02d", VxTimeParsed.Minute())
	//Detik := fmt.Sprintf("%02d", VxTimeParsed.Second())
	//JamPesanGabung := JamPesan + ":" + MenitPesan + ":" + Detik
	//
	////Date Time Pemesanan
	//
	////Date Time Berangkat
	//VdTime := VxTimeParsed.AddDate(0, 0, 7)
	//VdJam := VdTime.Add(10 * time.Hour)
	//
	//JamBerangkat := fmt.Sprintf("%02d", VdJam.Hour())
	//MenitBerangkat := fmt.Sprintf("%02d", VdJam.Minute())
	//DetikBerangkat := fmt.Sprintf("%02d", VdJam.Second())
	//JamBerangkatGabung := JamBerangkat + ":" + MenitBerangkat + ":" + DetikBerangkat

	//Date Time Berangkat

}
