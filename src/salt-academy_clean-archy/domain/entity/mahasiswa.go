package entity

import (
	"salt-academy_clean-archy/domain/value_object"
	"time"
)

type Mahasiswa struct {
	name       string
	nim        string
	birthPlace string
	birthDate  time.Time
	gender     string
	handphone  string
	address    string
	class      *Class
	lessons    []*Lessons
}

type DTOMahasiswa struct {
	NameMahasiswa string
	NIM           string
	BirthPlace    string
	BirthDate     time.Time
	Gender        string
	Handphone     string
	Address       string
}

func NewMahasiswa(dataMahasiswa DTOMahasiswa) (*Mahasiswa, error) {

	genderObject, err := value_object.NewGender(dataMahasiswa.string)
	if err != nil {
		return nil, err
	}

	return &Mahasiswa{
		name:       dataMahasiswa.NameMahasiswa,
		nim:        dataMahasiswa.NIM,
		birthPlace: dataMahasiswa.BirthPlace,
		birthDate:  time.Time{},
		gender:     dataMahasiswa.Gender,
		handphone:  dataMahasiswa.Handphone,
		address:    dataMahasiswa.Address,
	}, nil
}

func (mhs *Mahasiswa) GetName() string {
	return mhs.name
}

func (mhs *Mahasiswa) GetNIM() string {
	return mhs.nim
}

func (mhs *Mahasiswa) GetBirthPlace() string {
	return mhs.birthPlace
}

func (mhs *Mahasiswa) GetBirthDate() time.Time {
	return mhs.birthDate
}

func (mhs *Mahasiswa) SetClass(dataClass DTONewClass) (*Mahasiswa, error) {
	class, err := NewClass(dataClass)
	if err != nil {
		return mhs, err
	}

	mhs.class = mhs.class
	return mhs, nil
}

func NewClass(dataClass DTONewClass) {
	panic("unimplemented")
}

func (mhs *Mahasiswa) SetLessons(dataClass DTOLessons) (*Mahasiswa, error)

//TIBA TIBA KELUAR UNIMPLEMENTED, TERUSIN GUYS!
