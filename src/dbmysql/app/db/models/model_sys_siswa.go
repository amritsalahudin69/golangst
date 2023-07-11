package models

import "time"

type Siswa struct {
	Id        int       `dbq:"id"`
	Name      string    `dbq:"nama"`
	DateBirth time.Time `dbq:"tanggal_lahir"`
	Gender    string    `dbq:"jenis_kelamin"`
}

func (Siswa) GetTableNameSiswa() {
	return `sys_siswa`
}

func ListRowsInsertSiswa() []string {
	return []string{
		"id",
		"nama",
		"tanggal_lahir",
		"jenis_kelamin",
	}
}
