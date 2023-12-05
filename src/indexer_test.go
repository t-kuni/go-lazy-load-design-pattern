package src

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestIndex(t *testing.T) {
	type Item struct {
		Id   string
		Name string
		Age  int
	}

	t.Run("Load", func(t *testing.T) {
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
		testee := NewIndexer[*Item](mockLoader, func(item *Item) (string, *Item, error) {
			return item.Name, item, nil
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

		item, ok, err := testee.Get("Jane")
		assert.NoError(t, err)
		assert.True(t, ok)
		assert.Equal(t, "2", item.Id)
		assert.Equal(t, "Jane", item.Name)
		assert.Equal(t, 30, item.Age)
	})

	//t.Run("Load", func(t *testing.T) {
	//	mockCtrl := gomock.NewController(t)
	//	defer mockCtrl.Finish()
	//
	//	mockLoader := NewMockILoader[*Item](mockCtrl)
	//	mockLoader.EXPECT().
	//		Load().
	//		Return([]*Item{
	//			{"1", "John", 20},
	//			{"2", "Jane", 30},
	//			{"3", "Joe", 40},
	//		}, nil).
	//		Times(1) // 1回しか呼び出されないこと
	//	testee := NewIndexer[*Item](mockLoader, func(item *Item) (string, *Item, error) {
	//		return item.Name, item, nil
	//	})
	//
	//	// Loadを2回実行
	//	_, err := testee.Load()
	//	items, err := testee.Load()
	//
	//	assert.NoError(t, err)
	//
	//	assert.Len(t, items, 3)
	//	assert.Equal(t, "1", items[0].Id)
	//	assert.Equal(t, "John", items[0].Name)
	//	assert.Equal(t, 20, items[0].Age)
	//	assert.Equal(t, "3", items[2].Id)
	//	assert.Equal(t, "Joe", items[2].Name)
	//	assert.Equal(t, 40, items[2].Age)
	//})

	//t.Run("Load2", func(t *testing.T) {
	//	mockCtrl := gomock.NewController(t)
	//	defer mockCtrl.Finish()
	//
	//	mockLoader := NewMockILoader[*Item](mockCtrl)
	//	mockLoader.EXPECT().
	//		Load().
	//		Return(nil, errors.New("error")).
	//		Times(1)
	//	testee := NewIndexer[*Item](mockLoader)
	//
	//	_, err := testee.Load()
	//	assert.Error(t, err)
	//})
}
