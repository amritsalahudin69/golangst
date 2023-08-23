package entity

import (
	"mahasiswa_clean-archy/domain/value_object"
	"time"
)

type Score struct {
	gradingvalue *value_object.Grade
	mahasiswa    *Mahasiswa
	course       *Course
	lecture      *Lecture
	date         time.Time
	submit       time.Time
}

type DTOScore struct {
	GradingValue *value_object.Grade
	Mahasiswa    *Mahasiswa
	Course       *Course
	Lecture      *Lecture
	Date         time.Time
	Submit       time.Time
}

func NewScore(dto DTOScore) (*Score, error) {
	score := &Score{
		gradingvalue: dto.GradingValue,
		mahasiswa:    dto.Mahasiswa,
		course:       dto.Course,
		lecture:      dto.Lecture,
		date:         dto.Date,
		submit:       dto.Submit,
	}

	return score, nil
}

func (scr *Score) GetValue() *value_object.Grade {
	return scr.gradingvalue
}

func (scr *Score) GetMahasiswa() *Mahasiswa {
	return scr.mahasiswa
}

func (scr *Score) GetCourse() *Course {
	return scr.course
}

func (scr *Score) GetLecture() *Lecture {
	return scr.lecture
}

func (scr *Score) GetDate() time.Time {
	return scr.date
}

func (scr *Score) GetSubmit() time.Time {
	return scr.submit
}

func (scr *Score) SetMahasiswa(mahasiswa *Mahasiswa) {
	scr.mahasiswa = mahasiswa
}

func (scr *Score) SetCourse(course *Course) {
	scr.course = course
}

func (scr *Score) SetLecture(lecture *Lecture) {
	scr.lecture = lecture
}
