package value

import (
	"fmt"
	"reflect"
)

type ArrayValue struct {
	itemType Type
	items    []Value
}

func NewArrayValue(itemType Type) *ArrayValue {
	v := ArrayValue{itemType, make([]Value, 0)}
	return &v
}

func (a *ArrayValue) GetType() Type {
	return ArrayType{BaseType{TypeNameArray}, a.itemType}
}

func (a *ArrayValue) String() string {
	return fmt.Sprintf("%v", a.items)
}

func (a *ArrayValue) SetValue(arr interface{}) {
	s := reflect.ValueOf(arr)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return
	}

	items := make([]Value, s.Len())

	for i := 0; i < s.Len(); i++ {
		items[i] = s.Index(i).Interface().(Value)
	}

	*a = ArrayValue{a.itemType, items}
}

func (a *ArrayValue) GetValue() interface{} {
	return a.items
}

func (a *ArrayValue) AddItem(item Value) error {
	if item.GetType() != a.itemType {
		return fmt.Errorf("cannot add item to array of type %s", item.GetType())
	}
	a.items = append(a.items, item)
	return nil
}

func (a *ArrayValue) SetItemToIndex(index int, item Value) error {
	if item.GetType() != a.itemType {
		return fmt.Errorf("cannot add item to array of type %s", item.GetType())
	}
	if index >= len(a.items) {
		return fmt.Errorf("index out of range")
	}
	a.items[index] = item
	return nil
}

func (a *ArrayValue) GetItemOnIndex(index int) (item Value, err error) {
	if index >= len(a.items) {
		return nil, fmt.Errorf("index out of range")
	}
	return a.items[index], nil
}

func (a *ArrayValue) GetSize() int { return len(a.items) }

func (a *ArrayValue) Iterate() []Value {
	return a.items
}
