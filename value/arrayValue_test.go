package value

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayValue(t *testing.T) {
	s := NewArrayValue(BaseType{TypeNameInt})

	assert.Equal(t, ArrayType{
		BaseType{TypeNameArray},
		BaseType{TypeNameInt},
	}, s.GetType())

	s.SetValue([]*IntValue{NewIntValue(1), NewIntValue(2)})

	assert.Equal(t, s.Iterate()[0], NewIntValue(1))
	assert.Equal(t, s.Iterate()[1], NewIntValue(2))

	_ = s.AddItem(NewIntValue(3))
	assert.Equal(t, 3, s.GetSize())

	_ = s.SetItemToIndex(1, NewIntValue(4))
	item, err := s.GetItemOnIndex(1)
	assert.Equal(t, NewIntValue(4), item)
	assert.Nil(t, err)

	err = s.AddItem(NewBoolValue(true))
	assert.Error(t, err)
	err = s.SetItemToIndex(0, NewBoolValue(true))
	assert.Error(t, err)
	err = s.SetItemToIndex(7, NewIntValue(1))
	assert.Error(t, err)
	_, err = s.GetItemOnIndex(7)
	assert.Error(t, err)
}
