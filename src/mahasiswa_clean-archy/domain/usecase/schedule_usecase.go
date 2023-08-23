package usecase

import "mahasiswa_clean-archy/domain/entity"

type ScheduleUseCase interface {
	GetSchedule(nim string, periode int) (*entity.Mahasiswa, error)
}
