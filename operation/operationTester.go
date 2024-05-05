package operation

import (
	. "esorun/value"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testValue struct {
	before          Value
	argument        Value
	after           Value
	invalidArgument Value
}

func testOperation(
	t *testing.T,
	values []testValue,
	operation func(Value, Value) error,
) {
	for _, v := range values {
		err := operation(v.before, v.argument)
		if v.after != nil {
			assert.Nil(t, err)
			assert.Equal(t, v.after, v.before)
		} else {
			assert.Error(t, err)
		}
		if v.invalidArgument != nil {
			err = operation(v.before, v.invalidArgument)
			assert.Error(t, err)
		}
	}
}
