package getter

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/t-kuni/go-lazy-load-design-pattern/src/loader"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestIndexer(t *testing.T) {
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
				{"3", "Joe", 40},
			}, nil).
			Times(1)
		testee := NewIndexer[*Item](mockLoader, func(item *Item) (string, *Item, error) {
			return item.Name, item, nil
		})

		{
			item, ok, err := testee.Get("Jane")
			assert.NoError(t, err)
			assert.True(t, ok)
			assert.Equal(t, "2", item.Id)
			assert.Equal(t, "Jane", item.Name)
			assert.Equal(t, 30, item.Age)
		}

		{
			item, ok, err := testee.Get("Joe")
			assert.NoError(t, err)
			assert.True(t, ok)
			assert.Equal(t, "3", item.Id)
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
		testee := NewIndexer[*Item](mockLoader, func(item *Item) (string, *Item, error) {
			return item.Name, item, nil
		})

		_, ok, err := testee.Get("XXX")
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
		testee := NewIndexer[*Item](mockLoader, func(item *Item) (string, *Item, error) {
			return item.Name, item, nil
		})

		_, _, err := testee.Get("XXX")
		assert.Error(t, err)
	})
}
