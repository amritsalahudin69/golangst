package repository_mahasiswa

// func (mhs *mahasiswaInteractor) GetListMahasiswa(ctx context.Context) ([]*entity.Mahasiswa, error) {
// ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
// defer cancel()

// stmt := fmt.Sprintf(`SELECT * FROM %s`, model.Mahasiswa{}.GetTableName())
// opts := &dbq.Options{SingleResult: false, ConcreteStruct: model.Mahasiswa{}, DecoderConfig: dbq.StdTimeConversionConfig()}
// result := dbq.MustQ(ctx, mhs.conn, stmt, opts)

// if result != nil || len(result.([]*model.Mahasiswa)) == 0 {

// 	return mapper.MapperListModelToEntityMahasiswa(result.([]*model.Mahasiswa)), nil

// }

// return nil, errors.New("DATA EMPTY")
// }
