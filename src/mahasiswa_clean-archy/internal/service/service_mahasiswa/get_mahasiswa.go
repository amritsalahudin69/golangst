package servicemahasiswa

import (
	"context"
	"mahasiswa_clean-archy/domain/entity"
)

func (mhs *mahasiswaInteractor) GetMahasiswa(ctx context.Context) (*entity.Mahasiswa, error) {
	Mahsswa, err := mhs.mahasiswaRepo.GetMahasiswa(ctx)
	if err != nil {
		return nil, err
	}
	return Mahsswa, nil
}
