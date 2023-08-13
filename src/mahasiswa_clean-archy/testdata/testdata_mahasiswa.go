package testdata

import "mahasiswa_clean-archy/domain/entity"

var (
	DummyDTOMahasiswa = entity.DTOMahasiswa{
		NameMahasiswa: "Dio",
		Nim:           "123456",
		PhoneNumber:   "0824564",
		Address:       "Jalan Jomblo",
		Gender:        "Pria",
	}
)

func GetTestDataMahasiswa() *entity.Mahasiswa {
	newMahasiswa, _ := entity.NewMahasiswa(DummyDTOMahasiswa)
	return newMahasiswa
}
