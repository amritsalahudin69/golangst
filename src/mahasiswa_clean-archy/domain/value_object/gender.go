package value_object

type Gender struct {
	gender bool
	detail string
}

func NewGender(datagender bool) (*Gender, error) {
	genderlist := &Gender{gender: datagender}

	if genderlist.gender == false {
		genderlist.detail = "male"
	} else {
		genderlist.detail = "Female"
	}

	return genderlist, nil
}
func (gbl *Gender) GetGenderBool() bool {
	return gbl.gender
}

func (gbl *Gender) SetGenderBool() string {
	return gbl.detail
}
