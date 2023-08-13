package value_object

import "errors"

type Gender struct {
	name string
}

func NewGender(dataGender string) (*Gender, error) {
	gender := &Gender{name: dataGender}

	if gender.name == "" {
		return nil, errors.New("Gender Must Be State!")
	}

	if gender.name != "Pria" && gender.name != "Wanita" {
		return nil, errors.New("Gender Not State Very Well")
	}

	return gender, nil
}

func (gnd *Gender) GetGenderName() string {
	return gnd.name
}

type StatusClass struct {
	name string
}

func NewStatusClass(dataStatus string) (*StatusClass, error) {
	status := &StatusClass{name: dataStatus}

	if status.name == "" {
		return nil, errors.New("Nama kelas harus ada!")
	}

	if status.name != "Quiz" && status.name != "Tutor Session" && status.name != "Lecture Session" {
		return nil, errors.New("Nama kelas tidak tersedia")
	}
	return status, nil
}

func (sts *StatusClass) GetStatusClassName() string {
	return sts.name
}
