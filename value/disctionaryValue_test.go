package value

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionaryValue(t *testing.T) {
	s := NewDictionaryValue(BaseType{TypeNameInt})

	assert.Equal(t, DictionaryType{
		BaseType{TypeNameDictionary},
		BaseType{TypeNameInt},
	}, s.GetType())

	s.SetValue(map[string]*IntValue{
		"key1": NewIntValue(1),
		"key2": NewIntValue(2),
	})

	assert.Equal(t, s.Iterate()["key1"], NewIntValue(1))
	assert.Equal(t, s.Iterate()["key2"], NewIntValue(2))

	assert.Equal(t, 2, s.GetSize())

	_ = s.SetItem("key1", NewIntValue(4))
	item, err := s.GetItem("key1")
	assert.Equal(t, NewIntValue(4), item)
	assert.Nil(t, err)

	err = s.SetItem("key1", NewBoolValue(true))
	assert.Error(t, err)
	_, err = s.GetItem("not exist")
	assert.Error(t, err)
}
