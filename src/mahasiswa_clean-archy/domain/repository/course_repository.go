package repository

import "mahasiswa_clean-archy/domain/entity"

type CourseRepository interface {
	GetListCourses(mhs *entity.Mahasiswa, bachelorcode int) (*entity.Mahasiswa, error)
}
