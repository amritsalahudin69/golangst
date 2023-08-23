package servicemahasiswa

import (
	"context"
	"mahasiswa_clean-archy/domain/repository"
	"mahasiswa_clean-archy/domain/service"

	"github.com/google/uuid"
)

func (mhs *mahasiswaInteractor) AddMahasiswa(ctx context.Context, Mahasiswa service.DTOASRVddMahasiswa) error {
	Id, _ := uuid.NewUUID()
	newMahasiswa := repository.DTOARepoddMahasiswa{
		Id:         Id.String(),
		Name:       Mahasiswa.Name,
		NIM:        Mahasiswa.NIM,
		BirthPlace: Mahasiswa.BirthPlace,
		Handphone:  Mahasiswa.Handphone,
		Gender:     Mahasiswa.Gender,
		Address:    Mahasiswa.Address,
		BirthDate:  Mahasiswa.BirthDate,
	}
	err := mhs.mahasiswaRepo.AddMahasiswa(ctx, newMahasiswa)
	if err != nil {
		return err
	}

	return nil
}
