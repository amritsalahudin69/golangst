package entity_test

import (
	"mahasiswa_clean-archy/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	true  = 0 == 0 // Untyped bool.
	false = 0 != 0 // Untyped bool.
)

func TestNewBachelor(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		newVariable, err := entity.NewBachelor(entity.DTOBachelor{
			Code: 1,
			Name: "BE Golang",
		})
		assert.Nil(t, err)
		assert.NotNil(t, newVariable)
	})
}
