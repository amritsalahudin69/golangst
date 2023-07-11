package model

import "time"

type Kehadiran struct {
	Tanggal time.Time `dbq:"tanggal"`
}

func (Kehadiran) DGetTableNameKehadiran  {
	return `m_kehadiran`
}

func ListRowsInsertKehadiran() []string {
	return []string{
		"id",
		"nama",
		"tanggal_lahir",
		"jenis_kelamin",
	}
}