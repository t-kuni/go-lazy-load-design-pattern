package getter

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/t-kuni/go-lazy-load-design-pattern/src/loader"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestArrayIndexer(t *testing.T) {
	type Item struct {
		Id   string
		Name string
		Age  int
	}

	t.Run("Get", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockLoader := loader.NewMockILoader[*Item](mockCtrl)
		mockLoader.EXPECT().
			Load().
			Return([]*Item{
				{"1", "John", 20},
				{"2", "Jane", 30},
				{"3", "Joe", 30},
			}, nil).
			Times(1)
		testee := NewArrayIndexer[*Item](mockLoader, func(item *Item) (int, *Item, error) {
			return item.Age, item, nil
		})

		{
			items, ok, err := testee.Get(20)
			assert.NoError(t, err)
			assert.True(t, ok)
			assert.Len(t, items, 1)
			assert.Equal(t, "1", items[0].Id)
			assert.Equal(t, "John", items[0].Name)
			assert.Equal(t, 20, items[0].Age)
		}

		{
			items, ok, err := testee.Get(30)
			assert.NoError(t, err)
			assert.True(t, ok)
			assert.Len(t, items, 2)
			assert.Equal(t, "2", items[0].Id)
			assert.Equal(t, "3", items[1].Id)
		}
	})

	t.Run("Get", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockLoader := loader.NewMockILoader[*Item](mockCtrl)
		mockLoader.EXPECT().
			Load().
			Return([]*Item{
				{"1", "John", 20},
				{"2", "Jane", 30},
				{"3", "Joe", 40},
			}, nil).
			Times(1)
		testee := NewArrayIndexer[*Item](mockLoader, func(item *Item) (int, *Item, error) {
			return item.Age, item, nil
		})

		_, ok, err := testee.Get(999)
		assert.NoError(t, err)
		assert.False(t, ok)
	})

	t.Run("Get", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockLoader := loader.NewMockILoader[*Item](mockCtrl)
		mockLoader.EXPECT().
			Load().
			Return(nil, errors.New("error")).
			Times(1)
		testee := NewArrayIndexer[*Item](mockLoader, func(item *Item) (int, *Item, error) {
			return item.Age, item, nil
		})

		_, _, err := testee.Get(999)
		assert.Error(t, err)
	})
}
