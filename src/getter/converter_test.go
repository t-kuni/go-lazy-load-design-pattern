package getter

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestConverter(t *testing.T) {
	type Item struct {
		Id   string
		Name string
		Age  int
	}

	t.Run("Get", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockGetter := NewMockIGetter[string, *Item](mockCtrl)
		mockGetter.EXPECT().
			Get(gomock.Eq("Jane")).
			Return(&Item{"2", "Jane", 30}, true, nil).
			Times(1)
		testee := NewConverter[*Item, string, *Item](mockGetter, func(key string, v *Item) (*Item, error) {
			v.Name = v.Name + "(Converted)"
			return v, nil
		})

		{
			item, ok, err := testee.Get("Jane")
			assert.NoError(t, err)
			assert.True(t, ok)
			assert.Equal(t, "2", item.Id)
			assert.Equal(t, "Jane(Converted)", item.Name)
			assert.Equal(t, 30, item.Age)
		}
	})

	t.Run("sourceからokがfalseで返ってくる場合", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockGetter := NewMockIGetter[string, *Item](mockCtrl)
		mockGetter.EXPECT().
			Get(gomock.Eq("Jane")).
			Return(nil, false, nil).
			Times(1)
		testee := NewConverter[*Item, string, *Item](mockGetter, func(key string, v *Item) (*Item, error) {
			v.Name = v.Name + "(Converted)"
			return v, nil
		})

		{
			_, ok, err := testee.Get("Jane")
			assert.NoError(t, err)
			assert.False(t, ok)
		}
	})

	t.Run("Get", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockGetter := NewMockIGetter[string, *Item](mockCtrl)
		mockGetter.EXPECT().
			Get(gomock.Eq("Jane")).
			Return(&Item{"2", "Jane", 30}, true, nil).
			Times(1)
		testee := NewConverter[*Item, string, *Item](mockGetter, func(key string, v *Item) (*Item, error) {
			return nil, errors.New("error")
		})

		_, _, err := testee.Get("Jane")
		assert.Error(t, err)
	})
}
