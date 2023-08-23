package usecase_mahasiswa

import (
	"mahasiswa_clean-archy/domain/repository"
	"mahasiswa_clean-archy/domain/service"
	"mahasiswa_clean-archy/domain/usecase"
)

type mahasiswaInteractor struct {
	mahasiswaSvc service.MahasiswaTemplateService
}

func NewMahasiswa(repo repository.MahasiswaTemplateRepository) usecase.MahasiswaUseCase {
	// return &mahasiswaInteractor{
	// 	mahasiswaSvc: servicemahasiswa.NewMahasiswa(repo),
	// }
}
