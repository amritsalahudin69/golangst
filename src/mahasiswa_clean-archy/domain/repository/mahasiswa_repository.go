package repository

import "mahasiswa_clean-archy/domain/entity"

type MahasiswaRepository interface {
	GetMahasiswaByNim(nim string) (*entity.Mahasiswa, error)
}
