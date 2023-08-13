package testdata

import "mahasiswa_clean-archy/domain/entity"

var (
	DummyDTOLecture = entity.DTONewLecture{
		NameLecture: "Abdul Aziz",
		NikLecture:  "120120001",
	}
)

func GetTestDataLeture() *entity.Lecture {
	newLecture, _ := entity.NewLecture(DummyDTOLecture)
	newLecture.SetLesson(GetTestDataLesson())
	return newLecture
}

// func GetTestDataLeture() *entity.Lecture {
// 	newLecture, _ := entity.NewLecture(entity.DTONewLecture{
// 		NameLecture: "Abdul Aziz",
// 		NikLecture:  "120120001",
// 	})
// 	newlesson, _ := entity.NewLesson(entity.DTONewLesson{
// 		Name:     "Kalkulus",
// 		Duration: 60,
// 	})
// 	newLecture.SetLesson(newlesson)
// 	return newLecture
// }
