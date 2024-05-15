package assert

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type gormObj struct {
	gorm.Model
	Name string
}

type nongormObj struct {
	Name string
}

func TestEqual(t *testing.T) {
	expected := &gormObj{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "test",
	}
	os := []*gormObj{
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name: "test",
		},
		{
			Name: "test",
		},
	}
	for _, o := range os {
		Equal(t, expected, o)
	}
}

func TestElementMatch(t *testing.T) {
	expected := []*gormObj{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "test1",
		},
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "test2",
		},
	}
	actual := []*gormObj{
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name: "test1",
		},
		{
			Name: "test2",
		},
	}
	ElementsMatch(t, expected, actual)
}

func TestClearModelFromSlice(t *testing.T) {
	os0 := []*gormObj{
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "test",
		},
		{
			Model: gorm.Model{
				ID:        2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "test",
		},
	}
	assert.NoError(t, clearModelFromSlice(os0))
	for _, o := range os0 {
		assert.Equal(t, gorm.Model{}, o.Model)
	}

	os1 := []nongormObj{
		{Name: "test"},
	}
	assert.Error(t, clearModelFromSlice(os1))
}

func TestClearModel(t *testing.T) {
	o0 := &gormObj{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: "test",
	}
	assert.NoError(t, clearModel(o0))
	assert.Equal(t, gorm.Model{}, o0.Model)

	o1 := &nongormObj{Name: "test"}
	assert.Error(t, clearModel(o1))
}
