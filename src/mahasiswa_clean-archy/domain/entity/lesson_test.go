package entity_test

import (
	"mahasiswa_clean-archy/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLesson(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		newVariable, err := entity.NewLesson(entity.DTONewLesson{
			Name:     "Kalkulus",
			Duration: 30,
		})

		assert.Nil(t, err)
		assert.NotNil(t, newVariable)
	})

	t.Run("Negative Case", func(t *testing.T) {
		newVariable, err := entity.NewLesson(entity.DTONewLesson{
			Name:     "",
			Duration: 30,
		})

		assert.NotNil(t, err)
		assert.Nil(t, newVariable)
	})

}
