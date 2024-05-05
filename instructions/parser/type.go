package parser

import (
	"esorun/instructions"
	"fmt"
	"regexp"
)

var dicRegex, _ = regexp.Compile(`^<string,(.*)>$`)

func ParseType(str string) (typeDef *instructions.Type, err error) {
	if str[0] == '[' {
		inner, err := ParseTypeBasic(str[1 : len(str)-1])
		if err != nil {
			return nil, err
		}
		t := instructions.KeyType(instructions.KeyTypeInt)
		return &instructions.Type{
			Variant:   instructions.TypeVariantArray,
			KeyType:   &t,
			ValueType: inner,
		}, nil
	} else if str[0] == '<' {
		matches := dicRegex.FindStringSubmatch(str)
		if len(matches) != 2 {
			return nil, fmt.Errorf("invalid dictionary format")
		}
		inner, err := ParseTypeBasic(matches[1])
		if err != nil {
			return nil, err
		}
		t := instructions.KeyType(instructions.KeyTypeString)
		return &instructions.Type{
			Variant:   instructions.TypeVariantDictionary,
			KeyType:   &t,
			ValueType: inner,
		}, nil
	}
	return ParseTypeBasic(str)
}

func ParseTypeBasic(str string) (typeDef *instructions.Type, err error) {
	switch str {
	case instructions.TypeVariantByte,
		instructions.TypeVariantBool,
		instructions.TypeVariantString,
		instructions.TypeVariantInt,
		instructions.TypeVariantUint,
		instructions.TypeVariantFloat:
		return &instructions.Type{Variant: instructions.TypeVariant(str)}, nil
	default:
		return nil, fmt.Errorf("type %s not supported", str)
	}
}
