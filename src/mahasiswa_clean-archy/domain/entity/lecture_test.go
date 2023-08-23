package entity_test

import (
	"mahasiswa_clean-archy/domain/entity"
	"mahasiswa_clean-archy/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	DummyTestLecture = entity.DTONewLecture{
		NameLecture: "Abdul Aziz",
		NikLecture:  "120120001",
	}
)

func TestNewLecture(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		newVariable, err := entity.NewLecture(DummyTestLecture)
		assert.Nil(t, err)
		assert.NotNil(t, newVariable)
	})

	t.Run("Positive Case ke 2", func(t *testing.T) {
		newVariable, err := entity.NewLecture(DummyTestLecture)
		// newlesson, errLesson := entity.NewLesson(entity.DTONewLesson{
		// Name:     "Kalkulus",
		// Duration: 60,
		// })
		// newVariable.SetLesson(newlesson)
		assert.Nil(t, err)
		// assert.Nil(t, errLesson)
		assert.NotNil(t, newVariable)
		assert.Equal(t, testdata.GetTestDataLeture(), newVariable)
	})

	// t.Run("Negative Case", func(t *testing.T) {
	// 	newVariable, err := entity.NewLecture(entity.DTONewLecture{
	// 		NameLecture: "",
	// 		NikLecture:  "120120001",
	// 	})
	// 	assert.NotNil(t, err)
	// 	assert.Nil(t, newVariable)
	// })
}
