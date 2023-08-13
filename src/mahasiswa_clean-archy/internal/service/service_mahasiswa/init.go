package servicemahasiswa

import (
	"mahasiswa_clean-archy/domain/repository"
	"mahasiswa_clean-archy/domain/service"
)

type mahasiswaInteractor struct {
	mahasiswaRepo repository.MahasiswaTemplateRepository
}

func NewMahasiswa(repo repository.MahasiswaTemplateRepository) service.MahasiswaTemplateService {
	return &mahasiswaInteractor{
		mahasiswaRepo: repo,
	}
}
