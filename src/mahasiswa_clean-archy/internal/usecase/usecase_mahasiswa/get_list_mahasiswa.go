package usecase_mahasiswa

import (
	"context"
	"mahasiswa_clean-archy/domain/entity"
)

func (m *mahasiswaInteractor) GetListMahasiswaUC(ctx context.Context) ([]*entity.Mahasiswa, error) {

	data, err := m.mahasiswaSvc.GetListMahasiswa(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
