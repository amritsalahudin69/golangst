package entity

import (
	"errors"
	"mahasiswa_clean-archy/domain/value_object"
)

type Bachelor struct {
	code   int
	name   string
	status *value_object.Status
}

type DTOBachelor struct {
	Code   int
	Name   string
	Status *value_object.Status
}

func NewBachelor(dto DTOBachelor) (*Bachelor, error) {
	bachelor := &Bachelor{
		code:   dto.Code,
		name:   dto.Name,
		status: dto.Status,
	}

	if bachelor.code == 0 {
		return nil, errors.New("CODE MUST BE STATED")
	}

	if bachelor.name == "" {
		return nil, errors.New("NAME MUST BE STATED")
	}

	return bachelor, nil
}

func (bcl *Bachelor) GetCode() int {
	return bcl.code
}

func (bcl *Bachelor) GetName() string {
	return bcl.name
}

func (bcl *Bachelor) GetStatus() *value_object.Status {
	return bcl.status
}

func (bcl *Bachelor) SetCode(code int) {
	bcl.code = code
}

func (bcl *Bachelor) SetName(name string) {
	bcl.name = name
}

func (bcl *Bachelor) SetStatus(status bool) (*Bachelor, error) {
	datastatus, err := value_object.NewStatus(status)
	if err != nil {
		return nil, err
	}
	bcl.status = datastatus
	return bcl, nil
}
