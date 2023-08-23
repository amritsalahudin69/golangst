package entity

import "time"

type SemesterDetail struct {
	mahasiswa    *Mahasiswa
	semesterType *SemesterType
	course       []*Course
	startDate    time.Time
	endDate      time.Time
}

type DTOSemesterDetail struct {
	Mahasiswa    *Mahasiswa
	SemesterType *SemesterType
	Course       []*Course
	StartDate    time.Time
	EndDate      time.Time
}
