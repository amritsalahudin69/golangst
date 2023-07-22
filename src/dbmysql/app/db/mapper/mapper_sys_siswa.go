package mapper

import (
	"dbmysql/app/db/models"
	"time"
)

type Addsiswa struct {
	Name      string    `json:"name"`
	DateBirth time.Time `json:"dateBirth"`
	Gander    string    `json:"gander"`
}

func MapperJsonModeltoModel(siswa *models.Siswa) *Addsiswa {
	return &Addsiswa{
		Name:      models.Siswa{}.Name,
		DateBirth: models.Siswa{}.DateBirth,
		Gander:    models.Siswa{}.Gender,
	}
}
