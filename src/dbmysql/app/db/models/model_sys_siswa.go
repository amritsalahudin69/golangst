package models

import "time"

type siswa struct {
	Id            int       `dbq:"id"`
	Nama          string    `dbq:"nama"`
	TglLahir      time.Time `dbq:"tanggal_lahir"`
	Jenis_kelamin string    `dbq:"jenis_kelamin"`
}
