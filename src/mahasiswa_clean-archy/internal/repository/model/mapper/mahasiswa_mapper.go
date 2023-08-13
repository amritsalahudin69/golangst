package mapper

import (
	"encoding/json"
	"mahasiswa_clean-archy/domain/entity"
	"mahasiswa_clean-archy/internal/repository/model"
)

type MahasiswaJSON struct {
	Name      string `json:"nama"`
	NIM       string `json:"nim"`
	BirthDate string `json:"birthDate"`
}

func MapperModelToJson(model *model.Mahasiswa) []byte {
	dataMahasiswa := &MahasiswaJSON{
		Name:      model.Name,
		NIM:       model.NIM,
		BirthDate: "DATE: " + model.BirthDate.String(),
	}

	jsonMahasiwa, _ := json.Marshal(dataMahasiswa)
	return jsonMahasiwa
}

func MapperCollectionModelToCollectionJSON(models []*model.Mahasiswa) []byte {
	var datas []*MahasiswaJSON
	for _, value := range models {
		dataMahasiswa := &MahasiswaJSON{
			Name:      value.Name,
			NIM:       value.NIM,
			BirthDate: "DATE: " + value.BirthDate.String(),
		}

		datas = append(datas, dataMahasiswa)
	}

	jsonMahasiswa, _ := json.Marshal(datas)
	return jsonMahasiswa
}

func MapperListModelToEntityMahasiswa(models []*model.Mahasiswa) []*entity.Mahasiswa {
	var entities []*entity.Mahasiswa
	for _, model := range models {
		entity, err := entity.NewMahasiswa(entity.DTOMahasiswa{
			NameMahasiswa: model.Name,
			Nim:           model.NIM,
			PhoneNumber:   model.Handphone,
			Address:       model.Address,
			Gender:        model.Gender,
		})
		if err != nil {
			return nil
		}
		entities = append(entities, entity)
	}

	return entities
}

func MapperGetModelToEntityMahasiswa(model *model.Mahasiswa) *entity.Mahasiswa {
	var entities *entity.Mahasiswa

	entities, err := entity.NewMahasiswa(entity.DTOMahasiswa{
		NameMahasiswa: model.Name,
		Nim:           model.NIM,
		PhoneNumber:   model.Handphone,
		Address:       model.Address,
		Gender:        model.Gender,
	})
	if err != nil {
		return nil
	}
	return entities
}
