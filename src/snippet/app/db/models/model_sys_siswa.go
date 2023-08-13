package models

import "time"

type Siswa struct {
	Id        int       `dbq:"id"`
	Name      string    `dbq:"nama"`
	DateBirth time.Time `dbq:"tanggal_lahir"`
	Gender    string    `dbq:"jenis_kelamin"`
}

func (Siswa) GetTableNameSiswa() string {
	return `sys_siswa`
}

//type Addsiswa struct {
//	Name      string `json:"name"`
//	DateBirth string `json:"dateBirth"`
//	Gander    string `json:"gander"`
//}

func InsertSiswa() []string {
	return []string{
		"nama",
		"tanggal_lahir",
		"jenis_kelamin",
	}
}
