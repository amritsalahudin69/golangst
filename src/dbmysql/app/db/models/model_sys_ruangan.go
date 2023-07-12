package models

type Ruangan struct {
	Id       int    `dbq:"id"`
	RoomName string `dbq:"nama_ruangan"`
	RoomSize int    `dbq:"ukuran_ruangan"`
}

func (Ruangan) GetTableNameRuangan() string {
	return `sys_ruangan`
}

func ListRowsInsertRuangan() []string {
	return []string{
		"id",
		"nama_ruangan",
		"ukuran_ruangan",
	}
}
