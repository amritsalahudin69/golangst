package mapper

import (
	"encoding/json"
	"salt-academy_learn_week2/model"
)

type CampusJSON struct {
	Name    string `json:"nama"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func MapperModelCampusToJson(model *model.Campus) []byte {
	dataCampus := &CampusJSON{
		Name:    model.Name,
		Email:   model.Email,
		Address: model.Address,
	}
	jsonCampus, _ := json.Marshal(dataCampus)
	return jsonCampus
}

func MapperCollectionModelCampusToCollectionJSON(models []*model.Campus) []byte {
	var datas []*CampusJSON
	for _, value := range models {
		dataCampus := &CampusJSON{
			Name:    value.Name,
			Email:   value.Email,
			Address: value.Address,
		}

		datas = append(datas, dataCampus)
	}

	jsonCampus, _ := json.Marshal(datas)
	return jsonCampus
}
