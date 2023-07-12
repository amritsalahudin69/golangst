package model

import "time"

type Kehadiran struct {
	NIM     string    `dbq:"nim"`
	Tanggal time.Time `dbq:"tanggal"`
	Hadir   bool      `dbq:"hadir"`
	Kelas   string    `dbq:"kelas"`
}

type DetailKehadiran struct {
	NIM       string    `dbq:"nim"`
	Tanggal   time.Time `dbq:"tanggal"`
	Hadir     bool      `dbq:"hadir"`
	Kelas     string    `dbq:"kelas"`
	Name      string    `dbq:"nama"`
	Handphone string    `dbq:"no_hp"`
}

func (Kehadiran) GetTableName() string {
	return `m_kehadiran`
}

func listRowInsertKehadiran() []string {
	return []string{
		"NIM",
		"tanggal",
		"hadir",
		"kelas",
	}
}
