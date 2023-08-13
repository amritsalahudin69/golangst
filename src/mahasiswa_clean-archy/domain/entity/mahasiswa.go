package entity

import (
	"errors"
	"mahasiswa_clean-archy/domain/value_object"
)

type Mahasiswa struct {
	name        string
	nim         string
	phoneNumber string
	address     string
	gender      *value_object.Gender
}

type DTOMahasiswa struct {
	NameMahasiswa string
	Nim           string
	PhoneNumber   string
	Address       string
	Gender        string
}

func NewMahasiswa(dto DTOMahasiswa) (*Mahasiswa, error) {
	statusObject, err := value_object.NewGender(dto.Gender)
	if err != nil {
		return nil, err
	}
	mahasiwa := &Mahasiswa{
		name:        dto.NameMahasiswa,
		nim:         dto.Nim,
		phoneNumber: dto.PhoneNumber,
		address:     dto.Address,
		gender:      statusObject,
	}

	if mahasiwa.name == "" {
		return nil, errors.New("NAME MUST BE STATED")
	}

	if mahasiwa.nim == "" {
		return nil, errors.New("NIM MUST BE STATED")
	}

	return mahasiwa, nil
}

func (mhs *Mahasiswa) GetName() string {
	return mhs.name
}

func (mhs *Mahasiswa) GetNIM() string {
	return mhs.nim
}

func (mhs *Mahasiswa) GetPhonenumber() string {
	return mhs.phoneNumber
}

func (mhs *Mahasiswa) GetAddress() string {
	return mhs.address
}

func (mhs *Mahasiswa) GetGender() *value_object.Gender {
	return mhs.gender
}

func (mhs *Mahasiswa) SetName(name string) {
	mhs.name = name
}

func (mhs *Mahasiswa) SetNim(nim string) {
	mhs.nim = nim
}

func (mhs *Mahasiswa) SetPhone(phoneNumber string) {
	mhs.phoneNumber = phoneNumber
}

func (mhs *Mahasiswa) SetAddress(address string) {
	mhs.address = address
}

func (mhs *Mahasiswa) SetGender(gender *value_object.Gender) {
	mhs.gender = gender
}
