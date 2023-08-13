package servicemahasiswa

import (
	"context"
	"mahasiswa_clean-archy/domain/service"
)

func (mhs *mahasiswaInteractor) AddMahasiswa(ctx context.Context, Mahasiswa service.DTOARepoddMahasiswa) error {
	newMahasiswa := service.DTOARepoddMahasiswa{
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
