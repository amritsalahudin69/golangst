package usecase

import (
	"mahasiswa_clean-archy/domain/entity"
	"mahasiswa_clean-archy/domain/repository"
	"mahasiswa_clean-archy/domain/usecase"
)

type scheduleInteractor struct {
	courseRepository    repository.CourseRepository
	mahasiswaRepository repository.MahasiswaRepository
}

// GetSchedule implements usecase.ScheduleUseCase.
func (*scheduleInteractor) GetSchedule(nim string, periode int) (*entity.Mahasiswa, error) {
	panic("unimplemented")
}

func NewScheduleUseCase(courserepo repository.CourseRepository, mahasiswarepo repository.MahasiswaRepository) usecase.ScheduleUseCase {
	return &scheduleInteractor{
		courseRepository:    courserepo,
		mahasiswaRepository: mahasiswarepo,
	}
}
