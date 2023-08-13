package repository_mahasiswa

import (
	"context"
	"fmt"
	"mahasiswa_clean-archy/domain/repository"
	"mahasiswa_clean-archy/internal/repository/model"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

func (mhs *mahasiswaInteractor) AddMahasiswa(ctx context.Context, Mahasiswa repository.DTOAddMahasiswa) error {
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

// func isZero(v interface{}) bool {
// 	if v == nil {
// 		return true
// 	}
// 	return reflect.ValueOf(v).IsZero()
// }
