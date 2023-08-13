package usecase_mahasiswa

import (
	"context"
)

func (m *mahasiswaInteractor) AddMahasiswaUC(ctx context.Context, Mahasiswa DTOASRVddMahasiswa) error {
	newMahasiswaUc := service.DTOASRVddMahasiswa{
		Name:       Mahasiswa.Name,
		NIM:        Mahasiswa.NIM,
		BirthPlace: Mahasiswa.BirthPlace,
		Handphone:  Mahasiswa.Handphone,
		Gender:     Mahasiswa.Gender,
		Address:    Mahasiswa.Address,
		BirthDate:  Mahasiswa.BirthDate,
	}
	err := m.mahasiswaSvc.AddMahasiswa(ctx, newMahasiswaUc)
	f err != nil {
		return err
	}

	return nil
}
