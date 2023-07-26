package repository

import (
	"context"
	"salt-academy_learn_week2/model"
)

type MahasiswaTemplate interface {
	Get(ctx context.Context) (*model.Mahasiswa, error)
	Create(ctx context.Context, Mahasiswa model.Mahasiswa) error
	Update(ctx context.Context, NIM string, Mahasiswa model.Mahasiswa) error
	Delete(ctx context.Context, NIM string) error
	GetList(ctx context.Context) ([]*model.Mahasiswa, error)
}
