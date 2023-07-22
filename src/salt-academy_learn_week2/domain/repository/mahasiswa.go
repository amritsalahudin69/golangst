package repository

import (
	"context"
	"salt-academy_learn_week2/model"
)

type MahasiswaTemplate interface {
	GetMahasiswa(ctx context.Context) (*model.Mahasiswa, error)
	GetListMahasiswa(ctx context.Context) ([]*model.Mahasiswa, error)
	AddMahasiswa(ctx context.Context, Mahasiswa model.Mahasiswa) error
}
