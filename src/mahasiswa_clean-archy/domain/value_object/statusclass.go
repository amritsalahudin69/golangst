package value_object

import "errors"

type StatusClass struct {
	name string
}

func NewStatusClass(dataStatus string) (*StatusClass, error) {
	status := &StatusClass{name: dataStatus}

	if status.name == "" {
		return nil, errors.New("Nama kelas harus ada!")
	}

	if status.name != "Quiz" || status.name != "Tutor Session" || status.name != "Lecture Session" {
		return nil, errors.New("Nama kelas tidak tersedia")
	}
	return status, nil
}

func (sts *StatusClass) GetStatusClassName() string {
	return sts.name
}
