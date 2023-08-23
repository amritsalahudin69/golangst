package usecase_mahasiswa

import (
	"context"
	"mahasiswa_clean-archy/domain/service"
	"mahasiswa_clean-archy/domain/usecase"
	"time"
)

func (m *mahasiswaInteractor) AddMahasiswaUC(ctx context.Context, Mahasiswa usecase.DTOUCMahasiswa) error {
	NewBirthDay, err := time.Parse("2006-01-02", Mahasiswa.BirthDate)
	if err != nil {
		return err
	}
	newMahasiswaUc := service.DTOASRVddMahasiswa{
		Name:       Mahasiswa.Name,
		NIM:        Mahasiswa.NIM,
		BirthPlace: Mahasiswa.BirthPlace,
		Handphone:  Mahasiswa.Handphone,
		Gender:     Mahasiswa.Gender,
		Address:    Mahasiswa.Address,
		BirthDate:  NewBirthDay,
	}
	err = m.mahasiswaSvc.AddMahasiswa(ctx, newMahasiswaUc)
	if err != nil {
		return err
	}

	return nil
}
