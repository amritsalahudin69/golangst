package repository

import (
	"context"
	"mahasiswa_clean-archy/domain/entity"
	"time"
)

type DTOARepoddMahasiswa struct {
	Name       string
	NIM        string
	BirthPlace string
	Handphone  string
	Gender     string
	Address    string
	BirthDate  time.Time
}

type MahasiswaTemplateRepository interface {
	GetMahasiswa(ctx context.Context) (*entity.Mahasiswa, error)
	GetListMahasiswa(ctx context.Context) ([]*entity.Mahasiswa, error)
	AddMahasiswa(ctx context.Context, Mahasiswa DTOARepoddMahasiswa) error
}
