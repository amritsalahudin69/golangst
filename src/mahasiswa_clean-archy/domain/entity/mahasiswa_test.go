package entity_test

import (
	"mahasiswa_clean-archy/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMahasiswa(t *testing.T) {

	t.Run("Positive_Test", func(t *testing.T) {
		newVariable, err := entity.NewMahasiswa(entity.DTOMahasiswa{
			NameMahasiswa: "Noval",
			Nim:           "123456789",
			PhoneNumber:   "082445644",
			Address:       "jalan Sendirin mblo",
			Gender:        "Pria",
		})

		assert.Nil(t, err)
		assert.NotNil(t, newVariable)
	})

	// t.Run("Return Error", func(t *testing.T) {
	// 	newVariable, err := domain.NewMahasiswa("", "UKRI0012",
	// 		"081373100105", "Jl Kebon Jeruk 3 No 29", "true")
	// 	assert.NotNil(t, err)
	// 	assert.Nil(t, newVariable)
	// })

	// t.Run("Return Error Name Must Be Stated", func(t *testing.T) {
	// 	newVariable, err := domain.NewMahasiswa("Maleakhi", "",
	// 		"081373100105", "Jl Kebon Jeruk 3 No 29", "true")
	// 	assert.NotNil(t, err)
	// 	assert.Nil(t, newVariable)
	// 	assert.Equal(t, errors.New("NIM MUST BE STATED"), err)
	// })

}

// func TestProductGet(t *testing.T) {
// 	newVariable, err := domain.NewMahasiswa("Maleakhi", "UKRI0012",
// 		"081373100105", "Jl Kebon Jeruk 3 No 29", "")
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Maleakhi", newVariable.GetName())
// 	assert.Equal(t, "UKRI0012", newVariable.GetNIM())
// }
