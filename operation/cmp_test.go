package operation

import (
	. "esorun/value"
	"github.com/stretchr/testify/assert"
	"testing"
)

func cmpTester(t *testing.T, v1, v2, nv Value, op func(Value, Value) (bool, error)) {
	var isTrue bool
	var err error
	if v2 == nil {
		isTrue, err = op(v1, v2)
		assert.False(t, isTrue)
		assert.Error(t, err)
		return
	}
	isTrue, err = op(v1, v2)
	assert.True(t, isTrue)
	assert.Nil(t, err)

	isTrue, err = op(v1, nv)
	assert.False(t, isTrue)
	assert.Nil(t, err)
}

func TestOpCmpEquals(t *testing.T) {
	cmpTester(t, NewBoolValue(true), NewBoolValue(true), NewBoolValue(false), OpCmpEquals)
	cmpTester(t, NewByteValue(1), NewByteValue(1), NewByteValue(2), OpCmpEquals)
	cmpTester(t, NewFloatValue(1), NewFloatValue(1), NewFloatValue(2), OpCmpEquals)
	cmpTester(t, NewIntValue(1), NewIntValue(1), NewIntValue(2), OpCmpEquals)
	cmpTester(t, NewUintValue(1), NewUintValue(1), NewUintValue(2), OpCmpEquals)
	cmpTester(t, NewStringValue("asd"), NewStringValue("asd"), NewStringValue(""), OpCmpEquals)
}

func TestOpCmpNotEquals(t *testing.T) {
	cmpTester(t, NewBoolValue(true), NewBoolValue(false), NewBoolValue(true), OpCmpNotEquals)
	cmpTester(t, NewByteValue(1), NewByteValue(2), NewByteValue(1), OpCmpNotEquals)
	cmpTester(t, NewFloatValue(1), NewFloatValue(2), NewFloatValue(1), OpCmpNotEquals)
	cmpTester(t, NewIntValue(1), NewIntValue(2), NewIntValue(1), OpCmpNotEquals)
	cmpTester(t, NewUintValue(1), NewUintValue(2), NewUintValue(1), OpCmpNotEquals)
	cmpTester(t, NewStringValue("asd"), NewStringValue(""), NewStringValue("asd"), OpCmpNotEquals)
}

func TestOpCmpGreaterThen(t *testing.T) {
	cmpTester(t, NewBoolValue(true), nil, nil, OpCmpGreaterThen)
	cmpTester(t, NewByteValue(3), NewByteValue(2), NewByteValue(4), OpCmpGreaterThen)
	cmpTester(t, NewFloatValue(3), NewFloatValue(2), NewFloatValue(4), OpCmpGreaterThen)
	cmpTester(t, NewIntValue(3), NewIntValue(2), NewIntValue(4), OpCmpGreaterThen)
	cmpTester(t, NewUintValue(3), NewUintValue(2), NewUintValue(4), OpCmpGreaterThen)
	cmpTester(t, NewStringValue("asd"), nil, nil, OpCmpGreaterThen)
}

func TestOpCmpLessThen(t *testing.T) {
	cmpTester(t, NewBoolValue(true), nil, nil, OpCmpLessThen)
	cmpTester(t, NewByteValue(2), NewByteValue(3), NewByteValue(1), OpCmpLessThen)
	cmpTester(t, NewFloatValue(2), NewFloatValue(3), NewFloatValue(1), OpCmpLessThen)
	cmpTester(t, NewIntValue(2), NewIntValue(3), NewIntValue(1), OpCmpLessThen)
	cmpTester(t, NewUintValue(2), NewUintValue(3), NewUintValue(1), OpCmpLessThen)
	cmpTester(t, NewStringValue("asd"), nil, nil, OpCmpLessThen)
}

func TestOpCmpIsType(t *testing.T) {
	assert.True(t, OpCmpIsType(NewBoolValue(true), BaseType{TypeNameBool}))
	assert.True(t, OpCmpIsType(NewByteValue(2), BaseType{TypeNameByte}))
	assert.True(t, OpCmpIsType(NewFloatValue(2), BaseType{TypeNameFloat}))
	assert.True(t, OpCmpIsType(NewIntValue(2), BaseType{TypeNameInt}))
	assert.True(t, OpCmpIsType(NewUintValue(2), BaseType{TypeNameUint}))
	assert.True(t, OpCmpIsType(NewStringValue("asd"), BaseType{TypeNameString}))
}
