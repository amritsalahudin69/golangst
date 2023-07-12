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
