package entity

import (
	"errors"
)

type Score struct {
	scoreRecord int
	mahasiswa   *Mahasiswa
	lecture     *Lecture
}

type DTONewScore struct {
	ScoreRecord int
	Mahasiswa   string
	Lecture     string
}

func NewScore(dto DTONewScore) (*Score, error) {
	if dto.ScoreRecord == 0 {
		return nil, errors.New("Nilai tidak boleh kosong")
	}
	score := &Score{
		scoreRecord: dto.ScoreRecord,
	}
	return score, nil
}

func (sc *Score) GetScoreRecord() int {
	return sc.scoreRecord
}

func (sc *Score) GetMahasiswa() *Mahasiswa {
	return sc.mahasiswa
}
func (sc *Score) GetLecture() *Lecture {
	return sc.lecture
}

func (sc *Score) SetScoreRecord(Score int) {
	sc.scoreRecord = Score
}

func (sc *Score) SetMahasiswa(mahasiswa *Mahasiswa) {
	sc.mahasiswa = mahasiswa
}

func (sc *Score) SetLecture(lecture *Lecture) {
	sc.lecture = lecture
}
