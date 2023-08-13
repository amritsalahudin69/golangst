package usecase_mahasiswa

import (
	"context"
	"mahasiswa_clean-archy/domain/entity"
)

func (m *mahasiswaInteractor) GetMahasiswaUC(ctx context.Context) (*entity.Mahasiswa, error) {
	res, err := m.mahasiswaSvc.GetMahasiswa(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
