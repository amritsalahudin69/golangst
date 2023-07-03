package main

import (
	"fmt"
	"time"
)

type PesananCustomer struct {
	NamaPenerbangan  string
	KotaBerangkat    string
	KotaTujuan       string
	TglKeberangkatan string
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

func (p *PesananCustomer) AddTglKeberangkatan() string {
	return p.TglKeberangkatan
}

func (p *PesananCustomer) AddTglPemesanan() time.Time {
	return p.TglPemesanan
}

func main() {
	var TglPesan time.Time
	TglPesan = time.Now()

	TiketPesawat := PesananCustomer{"Emprit Airlines", "Jepara", "Yogyakarta", "2023-06-20 18:00:00", TglPesan}

	TglKebe := TiketPesawat.AddTglKeberangkatan()

	TglPesanParsed, err := time.Parse("2006-01-02 15:04:05", TglKebe)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	TglBerangkat := TglPesanParsed.Format("2006-01-02")

	JamBerangkat := fmt.Sprintf("%02d", TglPesanParsed.Hour())
	MenitBerangkat := fmt.Sprintf("%02d", TglPesanParsed.Minute())
	DetikBerangkat := fmt.Sprintf("%02d", TglPesanParsed.Second())
	JamKeberangkatan := JamBerangkat + ":" + MenitBerangkat + ":" + DetikBerangkat

	TahunPesan := TglPesanParsed.Format("2006")
	BulanPesan := TglPesanParsed.Format("01")
	TanggalPesan := TglPesanParsed.Format("02")
	Tglpesanan := TahunPesan + ":" + BulanPesan + ":" + TanggalPesan
	JamPesan := fmt.Sprintf("%02d", TglPesanParsed.Hour())
	MenitPesan := fmt.Sprintf("%02d", TglPesanParsed.Minute())
	Detik := fmt.Sprintf("%02d", TglPesanParsed.Second())
	JamPesanan := JamPesan + ":" + MenitPesan + ":" + Detik

	fmt.Printf("Tiket Tujuan Kota %s menuju Kota %s pada tanggal %s pukul %s \n", TiketPesawat.AddKotaBerangkat(), TiketPesawat.AddKotaTujuan(), TglBerangkat, JamKeberangkatan)

	fmt.Printf("Tiket ini dipesan pada tanggal %s pukul %s", Tglpesanan, JamPesanan)

}
