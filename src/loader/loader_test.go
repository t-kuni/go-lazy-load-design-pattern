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

	t.Run("If caching is disabled, calling Load twice will call the closure twice", func(t *testing.T) { // Translated
		calledCount := 0

		testee := NewLoader(func() ([]*Item, error) {
			calledCount++
			return []*Item{
				{"1", "John", 20},
				{"2", "Jane", 30},
				{"3", "Joe", 40},
			}, nil
		}, &LoaderOption{Cache: false})

		{
			items, err := testee.Load()
			assert.NoError(t, err)
			assert.Len(t, items, 3)
		}

		{
			items, err := testee.Load()
			assert.NoError(t, err)
			assert.Len(t, items, 3)
		}

		assert.Equal(t, 2, calledCount)
	})

	t.Run("If caching is enabled, calling Load twice will only call the closure once", func(t *testing.T) { // Translated
		calledCount := 0

		testee := NewLoader(func() ([]*Item, error) {
			calledCount++
			return []*Item{
				{"1", "John", 20},
				{"2", "Jane", 30},
				{"3", "Joe", 40},
			}, nil
		}, &LoaderOption{Cache: true})

		{
			items, err := testee.Load()
			assert.NoError(t, err)
			assert.Len(t, items, 3)
		}

		{
			items, err := testee.Load()
			assert.NoError(t, err)
			assert.Len(t, items, 3)
		}

		assert.Equal(t, 1, calledCount)
	})

	t.Run("If the option specification is omitted, caching should be enabled by default", func(t *testing.T) { // Translated
		calledCount := 0

		testee := NewLoader(func() ([]*Item, error) {
			calledCount++
			return []*Item{
				{"1", "John", 20},
				{"2", "Jane", 30},
				{"3", "Joe", 40},
			}, nil
		}, nil)

		{
			items, err := testee.Load()
			assert.NoError(t, err)
			assert.Len(t, items, 3)
		}

		{
			items, err := testee.Load()
			assert.NoError(t, err)
			assert.Len(t, items, 3)
		}

		assert.Equal(t, 1, calledCount)
	})

	t.Run("If the closure returns an error, an error should be returned", func(t *testing.T) {
		testee := NewLoader(func() ([]*Item, error) {
			return nil, errors.New("error")
		}, &LoaderOption{Cache: false})

		_, err := testee.Load()
		assert.Error(t, err)
	})
}
