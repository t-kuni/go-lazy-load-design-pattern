package loader

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoader_Load(t *testing.T) {
	type Item struct {
		Id   string
		Name string
		Age  int
	}

	t.Run("Should return the array returned by the closure", func(t *testing.T) { // Translated
		testee := NewLoader(func() ([]*Item, error) {
			return []*Item{
				{"1", "John", 20},
				{"2", "Jane", 30},
				{"3", "Joe", 40},
			}, nil
		})

		items, err := testee.Load()
		assert.NoError(t, err)

		assert.Len(t, items, 3)
		assert.Equal(t, "1", items[0].Id)
		assert.Equal(t, "John", items[0].Name)
		assert.Equal(t, 20, items[0].Age)
		assert.Equal(t, "3", items[2].Id)
		assert.Equal(t, "Joe", items[2].Name)
		assert.Equal(t, 40, items[2].Age)
	})

	t.Run("If the closure returns an error, an error should be returned", func(t *testing.T) { // Translated
		testee := NewLoader(func() ([]*Item, error) {
			return nil, errors.New("error")
		})

		_, err := testee.Load()
		assert.Error(t, err)
	})
}
