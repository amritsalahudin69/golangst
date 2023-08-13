package entity

import (
	"errors"
	"mahasiswa_clean-archy/domain/value_object"
)

type Class struct {
	nameClass    string
	totalStudent string
	statusClass  *value_object.StatusClass
}

type DTONewClass struct {
	NameClass    string
	TotalStudent string
	StatusClass  string
}

func NewClass(dto DTONewClass) (*Class, error) {
	statusObject, err := value_object.NewStatusClass(dto.StatusClass)
	if err != nil {
		return nil, err
	}

	if dto.NameClass == "" {
		return nil, errors.New("Nama kelas tidak boleh kosong")
	}

	class := &Class{
		nameClass:    dto.NameClass,
		totalStudent: dto.TotalStudent,
		statusClass:  statusObject,
	}

	return class, nil
}

func (c *Class) GetNameClass() string {
	return c.nameClass
}

func (c *Class) GetTotalStudent() string {
	return c.totalStudent
}

func (c *Class) GetStatusClass() *value_object.StatusClass {
	return c.statusClass
}

func (c *Class) SetNameClass(nameClass string) {
	c.nameClass = nameClass
}

func (c *Class) SetTotalStudent(totalStudent string) {
	c.totalStudent = totalStudent
}

func (c *Class) SetStatusClass(statusClass *value_object.StatusClass) {
	c.statusClass = statusClass
}
