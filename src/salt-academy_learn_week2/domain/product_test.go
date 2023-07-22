package domain_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"salt-academy_learn_week2/domain"
	"testing"
)

func TestProduct(t *testing.T) {

	t.Run("Positive_Test", func(t *testing.T) {
		newVariable, err := domain.NewMahasiswa("Maleakhi", "UKRI0012",
			"081373100105", "Jl Kebon Jeruk 3 No 29", true)

		assert.Nil(t, err)
		assert.NotNil(t, newVariable)
	})

	t.Run("Return Error", func(t *testing.T) {
		newVariable, err := domain.NewMahasiswa("", "UKRI0012",
			"081373100105", "Jl Kebon Jeruk 3 No 29", true)
		assert.NotNil(t, err)
		assert.Nil(t, newVariable)
	})

	t.Run("Return Error Name Must Be Stated", func(t *testing.T) {
		newVariable, err := domain.NewMahasiswa("Maleakhi", "",
			"081373100105", "Jl Kebon Jeruk 3 No 29", true)
		assert.NotNil(t, err)
		assert.Nil(t, newVariable)
		assert.Equal(t, errors.New("NIM MUST BE STATED"), err)
	})

}
