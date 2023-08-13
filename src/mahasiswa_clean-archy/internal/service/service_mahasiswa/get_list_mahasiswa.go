package servicemahasiswa

import (
	"context"
	"mahasiswa_clean-archy/domain/entity"
)

func (mhs *mahasiswaInteractor) GetListMahasiswa(ctx context.Context) ([]*entity.Mahasiswa, error) {
	MhsL, err := mhs.mahasiswaRepo.GetListMahasiswa(ctx)
	if err != nil {
		return nil, err
	}
	return MhsL, nil
}
