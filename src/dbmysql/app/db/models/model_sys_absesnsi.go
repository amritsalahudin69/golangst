package models

import "time"

type Absensi struct {
	Id       int       `dbq:"id"`
	DateTime time.Time `dbq:"tanggal"`
	Present  bool      `dbq:"kehadiran"`
	SiswaId  int       `bdq:"id_siswa"`
}

func (Absensi) GetTableNameAbsensi() string {
	return `sys_absensi`
}

func ListRowsInsertAbsensi() []string {
	return []string{
		"id",
		"tanggal",
		"kehadiran",
		"id_siswa",
	}
}
