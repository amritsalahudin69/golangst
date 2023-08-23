package repository

import "mahasiswa_clean-archy/domain/entity"

type ScoreRepository interface {
	GetScore(mhs *entity.Mahasiswa) ([]*entity.Score, error)
}
