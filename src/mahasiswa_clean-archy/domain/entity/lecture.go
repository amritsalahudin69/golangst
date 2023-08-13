package entity

import (
	"errors"
)

type Lecture struct {
	nameLecture string
	nikLecture  string
	lesson      *Lesson
}

type DTONewLecture struct {
	NameLecture string
	NikLecture  string
}

func NewLecture(dto DTONewLecture) (*Lecture, error) {
	if dto.NameLecture == "" {
		return nil, errors.New("Nama kelas tidak boleh kosong")
	}

	lecture := &Lecture{
		nameLecture: dto.NameLecture,
		nikLecture:  dto.NikLecture,
	}

	return lecture, nil
}

func (le *Lecture) GetNameLecture() string {
	return le.nameLecture
}

func (le *Lecture) GetNikLecture() string {
	return le.nameLecture
}

func (le *Lecture) GetLesson() *Lesson {
	return le.lesson
}

func (le *Lecture) SetNameLecture(nameLecture string) {
	le.nameLecture = nameLecture
}

func (le *Lecture) SetNikLecture(nikLecture string) {
	le.nikLecture = nikLecture
}

func (le *Lecture) SetLesson(lesson *Lesson) {
	le.lesson = lesson
}
