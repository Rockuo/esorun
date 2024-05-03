package value

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func valueTest(
	t *testing.T,
	typeType Type,
	defaultValueGenerator func() Value,
	stringRepresentation string,
	setValue interface{},
	operations map[string]map[string]Value,
	invalidArgument Value,
	notEqual Value,
	bigger Value,
	smaller Value,
) {
	assert.Equal(t, typeType, defaultValueGenerator().getType())
	assert.Equal(t, stringRepresentation, defaultValueGenerator().toString())

	var i = defaultValueGenerator()
	i.setValue(setValue)
	assert.Equal(t, "ab", i.getValue())

	valueOperationTest(
		t,
		func(v Value, v2 Value) error { return v.opAdd(v2) },
		operations["add"],
		invalidArgument,
	)
	valueOperationTest(
		t,
		func(v Value, v2 Value) error { return v.opSub(v2) },
		operations["sub"],
		invalidArgument,
	)
	valueOperationTest(
		t,
		func(v Value, v2 Value) error { return v.opMul(v2) },
		operations["mul"],
		invalidArgument,
	)
	valueOperationTest(
		t,
		func(v Value, v2 Value) error { return v.opDiv(v2) },
		operations["div"],
		invalidArgument,
	)

	valueConditionalTest(
		t,
		func(v Value, v2 Value) (isTrue bool, err error) { return v.cmpEquals(v2) },
		defaultValueGenerator(),
		defaultValueGenerator(),
		notEqual,
		invalidArgument,
	)

	valueConditionalTest(
		t,
		func(v Value, v2 Value) (isTrue bool, err error) { return v.cmpNotEquals(v2) },
		defaultValueGenerator(),
		notEqual,
		defaultValueGenerator(),
		invalidArgument,
	)

	valueConditionalTest(
		t,
		func(v Value, v2 Value) (isTrue bool, err error) { return v.cmpGreaterThan(v2) },
		defaultValueGenerator(),
		smaller,
		bigger,
		invalidArgument,
	)

	valueConditionalTest(
		t,
		func(v Value, v2 Value) (isTrue bool, err error) { return v.cmpLessTan(v2) },
		defaultValueGenerator(),
		bigger,
		smaller,
		invalidArgument,
	)

	assert.True(t, i.cmpIsType(TypeNameString))
	assert.False(t, i.cmpIsType(TypeNameByte))

}

func valueOperationTest(
	t *testing.T,
	execute func(Value, Value) error,
	data map[string]Value,
	invalidArgument Value) {
	err := execute(data["before"], data["argument"])
	if data["after"] != nil {
		assert.Nil(t, err)
		assert.Equal(t, data["after"], data["before"])
	} else {
		assert.Error(t, err)
	}

	err = execute(data["before"], invalidArgument)
	assert.Error(t, err)
}
func valueConditionalTest(
	t *testing.T,
	execute func(Value, Value) (isTrue bool, err error),
	v1 Value,
	v2 Value,
	v3 Value,
	invalidArgument Value,
) {
	var isTrue bool
	var err error

	isTrue, err = execute(v1, v2)
	if v2 != nil {
		assert.True(t, isTrue)
		assert.Nil(t, err)
	} else {
		assert.False(t, isTrue)
		assert.Error(t, err)
	}
	isTrue, err = execute(v1, v3)
	if v3 != nil {
		assert.False(t, isTrue)
		assert.Nil(t, err)
	} else {
		assert.False(t, isTrue)
		assert.Error(t, err)
	}

	isTrue, err = execute(v1, invalidArgument)
	assert.False(t, isTrue)
	assert.Error(t, err)
}
