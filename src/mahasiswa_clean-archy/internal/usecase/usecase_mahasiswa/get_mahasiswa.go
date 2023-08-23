package usecase_mahasiswa

import (
	"context"
	"mahasiswa_clean-archy/domain/entity"
)

func (m *mahasiswaInteractor) GetMahasiswaUC(ctx context.Context, nim string) (*entity.Mahasiswa, error) {
	res, err := m.mahasiswaSvc.GetMahasiswa(ctx, nim)
	if err != nil {
		return nil, err
	}
	return res, nil
}
