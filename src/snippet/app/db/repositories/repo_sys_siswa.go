package repositories

import (
	"context"
	"database/sql"
	"dbmysql/app/db/models"
	"errors"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

type SiswaInteractor struct {
	conn *sql.DB
}

func NewSiswaRepoMysql(connMysql *sql.DB) *SiswaInteractor {
	return &SiswaInteractor{conn: connMysql}
}
func (s *SiswaInteractor) GetSiswalist(ctx context.Context) ([]*models.Siswa, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s`, models.Siswa{}.GetTableNameSiswa())
	ops := &dbq.Options{SingleResult: false, ConcreteStruct: models.Siswa{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	result := dbq.MustQ(ctx, s.conn, stmt, ops)

	if result != nil {
		return result.([]*models.Siswa), nil
	}
	return nil, errors.New("DATA EMPTY")
}

func (s *SiswaInteractor) GetSiswa(id int, ctx context.Context) (*models.Siswa, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE id = ? `, models.Siswa{}.GetTableNameSiswa())
	ops := &dbq.Options{SingleResult: true, ConcreteStruct: models.Siswa{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	result := dbq.MustQ(ctx, s.conn, stmt, ops, id)

	if result != nil {
		return result.(*models.Siswa), nil
	}
	return nil, errors.New("DATA EMPTY")
}

func (s *SiswaInteractor) AddSiswa(siswa models.Siswa, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	stmt := fmt.Sprintf(`INSERT INTO %s (nama, tanggal_lahir, jenis_kelamin) VALUES (?, ?, ?)`, models.Siswa{}.GetTableNameSiswa())
	ops := &dbq.Options{SingleResult: true, ConcreteStruct: models.Siswa{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	_, err := dbq.E(
		ctx,
		s.conn,
		stmt,
		ops,
		siswa.Name,
		siswa.DateBirth,
		siswa.Gender,
	)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (s *SiswaInteractor) DeleteSiwa(id int, ctx context.Context) (*models.Siswa, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	stmt := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, models.Siswa{}.GetTableNameSiswa())
	ops := &dbq.Options{SingleResult: true, ConcreteStruct: models.Siswa{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	_, err := dbq.E(ctx, s.conn, stmt, ops, id)
	if err != nil {
		fmt.Println(err)
	}

	return nil, errors.New("DATA EMPTY")
}

func (s *SiswaInteractor) UpdateSiswa(id int, nama string, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`UPDATE %s SET nama = ? WHERE id = ?`, models.Siswa{}.GetTableNameSiswa())
	ops := &dbq.Options{SingleResult: true, ConcreteStruct: models.Siswa{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	_, err := dbq.E(ctx, s.conn, stmt, ops, nama, id)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
