package repository_mahasiswa

// func (mhs *mahasiswaInteractor) GetMahasiswa(ctx context.Context) (*entity.Mahasiswa, error) {
// ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
// defer cancel()

// stmt := fmt.Sprintf(`SELECT * FROM %s LIMIT 1 ORDER BY created_at DESC`, model.Mahasiswa{}.GetTableName())
// opts := &dbq.Options{SingleResult: true, ConcreteStruct: model.Mahasiswa{}, DecoderConfig: dbq.StdTimeConversionConfig()}
// result := dbq.MustQ(ctx, mhs.conn, stmt, opts)

// if result != nil {
// 	return mapper.MapperGetModelToEntityMahasiswa(result.(*model.Mahasiswa)), nil
// }

// return nil, errors.New("DATA EMPTY")
// }
