package testdata

import "mahasiswa_clean-archy/domain/entity"

var (
	DummyDTOLesson = entity.DTONewLesson{
		Name:     "Dio Agung Nauval",
		Duration: 160,
	}
)

func GetTestDataLesson() *entity.Lesson {
	newLesson, _ := entity.NewLesson(DummyDTOLesson)
	return newLesson
}
