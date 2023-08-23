package usecase

import (
	"mahasiswa_clean-archy/domain/entity"
	"mahasiswa_clean-archy/domain/repository"
)

type scoreInteractor struct {
	courseRepository    repository.CourseRepository
	mahasiswaRepository repository.MahasiswaRepository
	scoreRepository     repository.ScoreRepository
}

func NewScoreUseCase(courseRepo repository.CourseRepository, mhsRepo repository.MahasiswaRepository, scoreRepo repository.ScoreRepository) usecase.ScoreUsecase {
	return &scoreInteractor{
		courseRepository:    courseRepo,
		mahasiswaRepository: mhsRepo,
		scoreRepository:     scoreRepo,
	}
}

func (score *scoreInteractor) GetReportScore(nim string, bachelor int) ([]*entity.Score, error) {
	mahasiswa, err := score.mahasiswaRepository.GetMahasiswaByNim(nim)
	if err != nil {
		return nil, err
	}

	mahasiswa, errCheckBachelor := score.courseRepository.GetListCourses(mahasiswa, bachelor)

	if errCheckBachelor != nil {
		return nil, errCheckBachelor
	}

	listCourse, errGetLis := score.scoreRepository.GetScore(mahasiswa)
	if errGetLis != nil {
		return nil, errGetLis
	}

	return listCourse, nil
}
