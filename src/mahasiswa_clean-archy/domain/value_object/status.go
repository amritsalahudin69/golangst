package value_object

type Status struct {
	value  bool
	detail string
}

func NewStatus(datastatus bool) (*Status, error) {
	status := &Status{value: datastatus}

	if status.value == false {
		status.detail = "InActive"
	} else {
		status.detail = "Active"
	}
	return status, nil
}

func (sts *Status) GetValue() bool {
	return sts.value
}

func (sts *Status) GetDetailValue() string {
	// if sts.detail
	return sts.detail
}
