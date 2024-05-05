package parser

import (
	. "esorun/instructions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseType(t *testing.T) {
	var intKey = KeyType(KeyTypeInt)
	var strKey = KeyType(KeyTypeString)

	var valid = map[string]Type{
		"byte":         {Variant: TypeVariant(TypeVariantByte)},
		"string":       {Variant: TypeVariant(TypeVariantString)},
		"int":          {Variant: TypeVariant(TypeVariantInt)},
		"uint":         {Variant: TypeVariant(TypeVariantUint)},
		"float":        {Variant: TypeVariant(TypeVariantFloat)},
		"bool":         {Variant: TypeVariant(TypeVariantBool)},
		"[byte]":       {Variant: TypeVariant(TypeVariantArray), KeyType: &intKey, ValueType: &Type{Variant: TypeVariant(TypeVariantByte)}},
		"<string,int>": {Variant: TypeVariant(TypeVariantDictionary), KeyType: &strKey, ValueType: &Type{Variant: TypeVariant(TypeVariantInt)}},
	}
	for key, value := range valid {
		tp, e := ParseType(key)
		assert.Nil(t, e)
		assert.Equal(t, &value, tp)
	}

	tp, e := ParseType("asdf")
	assert.Nil(t, tp)
	assert.Error(t, e)

	tp, e = ParseType("[asdf]")
	assert.Nil(t, tp)
	assert.Error(t, e)

	tp, e = ParseType("<string, asdf>")
	assert.Nil(t, tp)
	assert.Error(t, e)

	tp, e = ParseType("<int, byte>")
	assert.Nil(t, tp)
	assert.Error(t, e)

}
