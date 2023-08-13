package usecase

import (
	"context"
	"mahasiswa_clean-archy/domain/entity"
)

type MahasiswaUseCase interface {
	GetMahasiswaUC(ctx context.Context) (*entity.Mahasiswa, error)
	GetListMahasiswaUC(ctx context.Context) ([]*entity.Mahasiswa, error)
}
