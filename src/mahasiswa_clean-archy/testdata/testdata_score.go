package testdata

import "mahasiswa_clean-archy/domain/entity"

var (
	DummyDTOScore = entity.DTONewScore{
		ScoreRecord: 3,
		Mahasiswa:   "LordDIO",
		Lecture:     "LordTabah",
	}
)

func GetTestDataScore() *entity.Score {
	NewSc, _ := entity.NewScore(DummyDTOScore)
	NewSc.SetMahasiswa(GetTestDataMahasiswa())
	NewSc.SetLecture(GetTestDataLeture())
	return NewSc
}
