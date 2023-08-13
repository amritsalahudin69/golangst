package entity

import "errors"

type Lesson struct {
	name     string
	duration int
}

type DTONewLesson struct {
	Name     string
	Duration int
}

func NewLesson(dto DTONewLesson) (*Lesson, error) {
	if dto.Name == "" {
		return nil, errors.New("Nama Kelas Tidak Boleh Kosong!")
	}
	lesson := &Lesson{
		name:     dto.Name,
		duration: dto.Duration,
	}
	return lesson, nil
}

func (l *Lesson) GetNameLesson() string {
	return l.name
}

func (l *Lesson) GetDurationLesson() int {
	return l.duration
}

func (l *Lesson) SetNameLesson(name string) {
	l.name = name
}

func (l *Lesson) SetDurationLesson(duration int) {
	l.duration = duration
}
