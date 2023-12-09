package getter

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/t-kuni/go-lazy-load-design-pattern/src/loader"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestArrayIndexer_Get(t *testing.T) {
	type Item struct {
		Id   string
		Name string
		Age  int
	}

	t.Run("指定したkeyを持ったレコードを取得できること", func(t *testing.T) {
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
			items, ok, err := testee.Get(30)
			assert.NoError(t, err)
			assert.True(t, ok)
			assert.Len(t, items, 2)
			assert.Equal(t, "2", items[0].Id)
			assert.Equal(t, "Jane", items[0].Name)
			assert.Equal(t, 30, items[0].Age)
			assert.Equal(t, "3", items[1].Id)
			assert.Equal(t, "Joe", items[1].Name)
			assert.Equal(t, 30, items[1].Age)
		}
	})

	t.Run("2回呼び出した場合、sourceのLoaderは1度しか呼び出されないこと", func(t *testing.T) {
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
		}

		{
			items, ok, err := testee.Get(30)
			assert.NoError(t, err)
			assert.True(t, ok)
			assert.Len(t, items, 2)
		}
	})

	t.Run("存在しないkeyを指定しない場合、第二戻り値にfalseを返す事", func(t *testing.T) {
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

	t.Run("クロージャがエラーを返した場合、エラーを返すこと", func(t *testing.T) {
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
			return 0, nil, errors.New("error")
		})

		_, _, err := testee.Get(999)
		assert.Error(t, err)
	})

	t.Run("sourceのLoaderがエラーを返した場合、エラーを返すこと", func(t *testing.T) {
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
