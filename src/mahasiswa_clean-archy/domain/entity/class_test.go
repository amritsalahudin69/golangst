package entity_test

import (
	"mahasiswa_clean-archy/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClass(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		newVariable, err := entity.NewClass(entity.DTONewClass{
			NameClass:    "Golang Bacth 3",
			TotalStudent: "15",
			StatusClass:  "Quiz",
		})
		assert.Nil(t, err)
		assert.NotNil(t, newVariable)
	})

	t.Run("Negative Cast", func(t *testing.T) {

		newVariable, err := entity.NewClass(entity.DTONewClass{
			NameClass:    "",
			TotalStudent: "15",
			StatusClass:  "Quiz",
		})
		assert.NotNil(t, err)
		assert.Nil(t, newVariable)
	})
}
