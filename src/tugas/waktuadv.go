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

func (p *PesananCustomer) AddTglPemesanan() time.Time {
	return p.TglPemesanan
}

func main() {
	var TglPesan time.Time
	TglPesan = time.Now()

	XX := TglPesan.Format("2006-01-02 15:04:05")
	TglB := TglPesan.AddDate(0, 0, 7)
	TglBerangkat := TglB.Format("2006-01-02")
	JamB := TglB.Add(10 * time.Hour)

	JamBerangkat := fmt.Sprintf("%02d", JamB.Hour())
	MenitBerangkat := fmt.Sprintf("%02d", JamB.Minute())
	DetikBerangkat := fmt.Sprintf("%02d", JamB.Second())
	JamKeberangkatan := JamBerangkat + ":" + MenitBerangkat + ":" + DetikBerangkat

	TglPesanParsed, err := time.Parse("2006-01-02 15:04:05", XX)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	TahunPesan := TglPesanParsed.Format("2006")
	BulanPesan := TglPesanParsed.Format("01")
	TanggalPesan := TglPesanParsed.Format("02")
	Tglpesanan := TahunPesan + ":" + BulanPesan + ":" + TanggalPesan
	JamPesan := fmt.Sprintf("%02d", TglPesanParsed.Hour())
	MenitPesan := fmt.Sprintf("%02d", TglPesanParsed.Minute())
	Detik := fmt.Sprintf("%02d", TglPesanParsed.Second())
	JamPesanan := JamPesan + ":" + MenitPesan + ":" + Detik

	TiketPesawat := PesananCustomer{"Emprit Airlines", "Jepara", "Yogyakarta", TglB, TglPesan}

	fmt.Printf("Tiket Tujuan Kota %s menuju Kota %s pada tanggal %s pukul %s \n", TiketPesawat.AddKotaBerangkat(), TiketPesawat.AddKotaTujuan(), TglBerangkat, JamKeberangkatan)

	fmt.Printf("Tiket ini dipesan pada tanggal %s pukul %s", Tglpesanan, JamPesanan)

}
