package repository

import (
	"mahasiswa_clean-archy/domain/entity"
	"time"
)

type DTOARepoddMahasiswa struct {
	Id         string
	Name       string
	NIM        string
	BirthPlace string
	Handphone  string
	Gender     string
	Address    string
	BirthDate  time.Time
}

type MahasiswaTemplateRepository interface {
	GetMahasiswaBynim(nim string) (*entity.Mahasiswa, error)
	// GetListMahasiswa(ctx context.Context) ([]*entity.Mahasiswa, error)
	// AddMahasiswa(ctx context.Context, Mahasiswa DTOARepoddMahasiswa) error
}
