package value

import (
	"fmt"
	"reflect"
)

type DictionaryValue struct {
	itemType Type
	items    map[string]Value
}

func NewDictionaryValue(itemType Type) *DictionaryValue {
	v := DictionaryValue{itemType, map[string]Value{}}
	return &v
}

func (a *DictionaryValue) GetType() Type {
	return DictionaryType{BaseType{TypeNameDictionary}, a.itemType}
}

func (a *DictionaryValue) String() string {
	return fmt.Sprintf("%v", a.items)
}

func (a *DictionaryValue) SetValue(i interface{}) {
	m := reflect.ValueOf(i)
	if m.Kind() != reflect.Map {

		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if m.IsNil() {
		return
	}

	items := map[string]Value{}

	for _, key := range m.MapKeys() {
		items[key.Interface().(string)] = m.MapIndex(key).Interface().(Value)
	}
	*a = DictionaryValue{a.itemType, items}
}

func (a *DictionaryValue) GetValue() interface{} {
	return a.items
}

func (a *DictionaryValue) SetItem(key string, item Value) error {
	if item.GetType() != a.itemType {
		return fmt.Errorf("cannot add item to array of type %s", item.GetType())
	}
	a.items[key] = item
	return nil
}

func (a *DictionaryValue) GetItem(key string) (item Value, err error) {
	v := a.items[key]
	if v == nil {
		return nil, fmt.Errorf("no value found for key %s", key)
	}
	return v, nil
}

func (a *DictionaryValue) GetSize() int { return len(a.items) }

func (a *DictionaryValue) Iterate() map[string]Value {
	return a.items
}
