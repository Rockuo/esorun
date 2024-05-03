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
) {
	assert.Equal(t, typeType, defaultValueGenerator().GetType())
	assert.Equal(t, stringRepresentation, defaultValueGenerator().String())

	var i = defaultValueGenerator()
	i.SetValue(setValue)
	assert.Equal(t, setValue, i.GetValue())
}
