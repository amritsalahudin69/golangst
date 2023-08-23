package service

import (
	"context"
	"mahasiswa_clean-archy/domain/entity"
	"time"
)

type DTOASRVddMahasiswa struct {
	Name       string
	NIM        string
	BirthPlace string
	Handphone  string
	Gender     string
	Address    string
	BirthDate  time.Time
}

type MahasiswaTemplateService interface {
	GetMahasiswa(ctx context.Context, nim string) (*entity.Mahasiswa, error)
	GetListMahasiswa(ctx context.Context) ([]*entity.Mahasiswa, error)
	AddMahasiswa(ctx context.Context, Mahasiswa DTOASRVddMahasiswa) error
}
