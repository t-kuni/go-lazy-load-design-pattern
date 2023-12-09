package loader

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestConverter_Load(t *testing.T) {
	type Item struct {
		Id   string
		Name string
		Age  int
	}

	t.Run("Should be able to retrieve the value after applying the conversion process", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockLoader := NewMockILoader[*Item](mockCtrl)
		mockLoader.EXPECT().
			Load().
			Return([]*Item{
				{"1", "John", 20},
				{"2", "Jane", 30},
				{"3", "Joe", 40},
			}, nil).
			Times(1)
		testee := NewConverter[*Item](mockLoader, func(items []*Item) ([]*Item, error) {
			for i := range items {
				items[i].Name = items[i].Name + "(Converted)"
			}
			return items, nil
		})

		items, err := testee.Load()

		assert.NoError(t, err)

		assert.Len(t, items, 3)
		assert.Equal(t, "1", items[0].Id)
		assert.Equal(t, "John(Converted)", items[0].Name)
		assert.Equal(t, 20, items[0].Age)
		assert.Equal(t, "3", items[2].Id)
		assert.Equal(t, "Joe(Converted)", items[2].Name)
		assert.Equal(t, 40, items[2].Age)
	})

	t.Run("If the closure returns an error, an error should be returned", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockLoader := NewMockILoader[*Item](mockCtrl)
		mockLoader.EXPECT().
			Load().
			Return([]*Item{
				{"1", "John", 20},
				{"2", "Jane", 30},
				{"3", "Joe", 40},
			}, nil).
			Times(1)
		testee := NewConverter[*Item](mockLoader, func(items []*Item) ([]*Item, error) {
			return nil, errors.New("error")
		})

		_, err := testee.Load()
		assert.Error(t, err)
	})

	t.Run("If the source Loader returns an error, an error should be returned", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockLoader := NewMockILoader[*Item](mockCtrl)
		mockLoader.EXPECT().
			Load().
			Return(nil, errors.New("error")).
			Times(1)
		testee := NewConverter[*Item](mockLoader, func(items []*Item) ([]*Item, error) {
			for i := range items {
				items[i].Name = items[i].Name + "(Converted)"
			}
			return items, nil
		})

		_, err := testee.Load()
		assert.Error(t, err)
	})
}
