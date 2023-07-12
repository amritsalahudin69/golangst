package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"salt-academy_learn_week2/model"
	"time"
)

type kehadiranInteractor struct {
	conn *sql.DB
}

func NewKehadiranRepositoryMsql(connMysql *sql.DB) *kehadiranInteractor {
	return &kehadiranInteractor{conn: connMysql}
}

func (hdr *kehadiranInteractor) GetKehadiranByNim(ctx context.Context, nim string) ([]*model.DetailKehadiran, error) {
	ctx, cansel := context.WithTimeout(ctx, 60*time.Second)
	defer cansel()

	kehadiran := model.Kehadiran{}.GetTableName()
	mahasiswa := model.Mahasiswa{}.GetTableName()
	khdr := fmt.Sprintf(`SELECT * FROM %s RIGHT JOIN %s ON %s.nim = %s.nim AND %s.nim = ? ORDER BY tanggal`, kehadiran, mahasiswa, kehadiran, mahasiswa, mahasiswa)
	opts := &dbq.Options{SingleResult: false, ConcreteStruct: model.DetailKehadiran{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	result := dbq.MustQ(ctx, hdr.conn, khdr, opts)

	if result != nil {
		return result.([]*model.DetailKehadiran), nil
	}

	return nil, errors.New("DATA EMPTY")
}
