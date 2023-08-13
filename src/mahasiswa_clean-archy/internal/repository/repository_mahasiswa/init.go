package repository_mahasiswa

import (
	"database/sql"
)

type mahasiswaInteractor struct {
	conn *sql.DB
}

// func NewMahasiswaRepositoryMysql(connMysql *sql.DB) repository.MahasiswaTemplateRepository {
// 	return &mahasiswaInteractor{conn: connMysql}
// }
