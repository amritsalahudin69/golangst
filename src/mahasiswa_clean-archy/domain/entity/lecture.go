package entity

import (
	"errors"
	"mahasiswa_clean-archy/domain/value_object"
	"time"
)

type Lecture struct {
	nikLecture  string
	nameLecture string
	phonenumber string
	joinDate    time.Time
	status      *value_object.Status
}
type DTONewLecture struct {
	NikLecture  string
	NameLecture string
	PhoneNumber string
	JoinDate    time.Time
	Status      *value_object.Status
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

func (le *Lecture) GetNikLecture() string {
	return le.nameLecture
}

func (le *Lecture) GetNameLecture() string {
	return le.nameLecture
}

func (le *Lecture) GetPhoneNumber() string {
	return le.phonenumber
}

func (le *Lecture) GetJoinDate() time.Time {
	return le.joinDate
}

func (le *Lecture) GetStatus() *value_object.Status {
	return le.status
}

func (le *Lecture) SetNikLecture(nikLecture string) {
	le.nikLecture = nikLecture
}

func (le *Lecture) SetNameLecture(nameLecture string) {
	le.nameLecture = nameLecture
}

func (le *Lecture) SetPhoneNumebr(phonenumber string) {
	le.phonenumber = phonenumber
}

func (le *Lecture) SetJoinDate(joindate time.Time) {
	le.joinDate = joindate
}

func (le *Lecture) SetStatusLec(status bool) (*Lecture, error) {
	datastatus, err := value_object.NewStatus(status)
	if err != nil {
		return nil, err
	}
	le.status = datastatus
	return le, nil
}
