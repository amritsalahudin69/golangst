package value_object

import "errors"

type Gender struct {
	name string
}

func NewGender(dataGender Gender) (*Gender, error) {
	gender := &Gender{name: dataGender}

	if gender.name == "" {
		return nil, errors.New("Gender Must State!")
	}

	if gender.name != "Pria" && gender.name != "Wanita" {
		return nil, errors.New("Gender Not State Very Well")
	}

	return gender, nil
}

func (gnd *Gender) GetGenderName() string {
	return gnd.name
}
