package usecase

import "mahasiswa_clean-archy/domain/entity"

type ScoreUseCase interface {
	GetReportScore(nim string, bachelor int) (*entity.Mahasiswa, error)
}
