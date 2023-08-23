package entity

import (
	"errors"
	"mahasiswa_clean-archy/domain/value_object"
)

type Class struct {
	nameClass     string
	datamahasiswa []*Mahasiswa
	statusClass   *value_object.StatusClass
}

type DTONewClass struct {
	NameClass     string
	DataMahasiswa []*Mahasiswa
	StatusClass   string
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
		nameClass:     dto.NameClass,
		datamahasiswa: dto.DataMahasiswa,
		statusClass:   statusObject,
	}

	return class, nil
}

func (c *Class) GetNameClass() string {
	return c.nameClass
}

func (c *Class) GetDataMahasiswa() []*Mahasiswa { // Ubah nama fungsi dan tipe pengembalian
	return c.datamahasiswa
}

func (c *Class) GetStatusClass() *value_object.StatusClass {
	return c.statusClass
}

func (c *Class) SetNameClass(nameClass string) {
	c.nameClass = nameClass
}

func (c *Class) SetDataMahasiswa(datamahasiswa []*Mahasiswa) { // Ubah nama fungsi dan parameter
	c.datamahasiswa = datamahasiswa
}

func (c *Class) SetStatusClass(statusClass *value_object.StatusClass) {
	c.statusClass = statusClass
}
