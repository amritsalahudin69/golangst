package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"salt-academy_learn_week2/domain/repository"
	"salt-academy_learn_week2/model"
	"time"
)

type MahasiswaInteractor struct {
	conn *sql.DB
}

func NewMahasiswaRepositoryMysql(connMysql *sql.DB) repository.MahasiswaTemplate {
	return &MahasiswaInteractor{conn: connMysql}
}

func (mhs *MahasiswaInteractor) Get(ctx context.Context) (*model.Mahasiswa, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s LIMIT 1 ORDER BY created_date DESC`, model.Mahasiswa{}.GetTableName())
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: model.Mahasiswa{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	result := dbq.MustQ(ctx, mhs.conn, stmt, opts)

	if result != nil {
		return result.(*model.Mahasiswa), nil
	}

	return nil, errors.New("DATA EMPTY")
}

func (mhs *MahasiswaInteractor) GetList(ctx context.Context) ([]*model.Mahasiswa, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s`, model.Mahasiswa{}.GetTableName())
	opts := &dbq.Options{SingleResult: false, ConcreteStruct: model.Mahasiswa{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	result := dbq.MustQ(ctx, mhs.conn, stmt, opts)

	if result != nil || len(result.([]*model.Mahasiswa)) == 0 {
		return result.([]*model.Mahasiswa), nil
	}

	return nil, errors.New("DATA EMPTY")
}

func (mhs *MahasiswaInteractor) Create(ctx context.Context, Mahasiswa model.Mahasiswa) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	stmt := fmt.Sprintf(`INSERT INTO %s (nama, tanggal_lahir, jenis_kelamin) VALUES (?, ?, ?)`, model.Mahasiswa{}.GetTableName())
	ops := &dbq.Options{SingleResult: true, ConcreteStruct: model.Mahasiswa{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	_, err := dbq.E(
		ctx,
		mhs.conn,
		stmt,
		ops,
		Mahasiswa.Name,
		Mahasiswa.BirthDate,
		Mahasiswa.Gender,
	)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func (mhs *MahasiswaInteractor) Update(ctx context.Context, NIM string, Mahasiswa model.Mahasiswa) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`UPDATE %s SET nama = ? WHERE nim = ?`, model.Mahasiswa{}.GetTableName())
	ops := &dbq.Options{SingleResult: true, ConcreteStruct: model.Mahasiswa{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	_, err := dbq.E(
		ctx,
		mhs.conn,
		stmt,
		ops,
		NIM,
		Mahasiswa.Name,
		Mahasiswa.BirthDate,
		Mahasiswa.Handphone,
	)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (mhs *MahasiswaInteractor) Delete(ctx context.Context, NIM string) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	stmt := fmt.Sprintf(`DELETE FROM %s WHERE nim = ?`, model.Mahasiswa{}.GetTableName())
	ops := &dbq.Options{SingleResult: true, ConcreteStruct: model.Mahasiswa{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	_, err := dbq.E(
		ctx,
		mhs.conn,
		stmt,
		ops,
		NIM,
	)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
