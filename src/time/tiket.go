package main

import "time"

type NamaCust struct {
	Nama    string
	Ktp     string
	Pesanan Order
}

type Order struct {
	NamaPenerbangan string
	KotaBerangkat   string
	KotaTujuan      string
	TglBerangkat    time.Timer
	TglPemesanan    time.Timer
}

func NewNamaCust(NewNamaCust string, NewKtp string) *NamaCust {
	return &NamaCust{
		Nama: NewNamaCust,
		Ktp:  NewKtp,
	}
}

func (n *NamaCust) getNama() string {
	return n.Nama
}

func (n *NamaCust) getKtp() string {
	return n.Ktp
}
