package entity

import "errors"

type SemesterType struct {
	name         string
	typesemester string
}

type DTOSemesterType struct {
	Name         string
	Typesemester string
}

func NewSemesterType(dto DTOSemesterType) (*SemesterType, error) {
	if dto.Name == "" {
		return nil, errors.New("Name Semester Tidak Boleh Kosong")
	}
	semesterType := &SemesterType{
		name:         dto.Name,
		typesemester: dto.Typesemester,
	}
	return semesterType, nil
}

func (st *SemesterType) GetName() string {
	return st.name
}

func (st *SemesterType) GetTypeSemester() string {
	return st.typesemester
}
