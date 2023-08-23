package usecase

import (
	"context"
	"mahasiswa_clean-archy/domain/entity"
)

type DTOUCMahasiswa struct {
	Name       string
	NIM        string
	BirthPlace string
	Handphone  string
	Gender     string
	Address    string
	BirthDate  string
}

type MahasiswaUseCase interface {
	GetMahasiswaUC(ctx context.Context) (*entity.Mahasiswa, error)
	GetListMahasiswaUC(ctx context.Context) ([]*entity.Mahasiswa, error)
	AddMahasiswa(ctx context.Context, Mahasiswa DTOUCMahasiswa) error
}
