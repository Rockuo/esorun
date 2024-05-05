package parser

import "fmt"

import . "esorun/instructions"

func ParseConstant(val string) (c Constant, err error) {
	switch val {
	case StdIn,
		StdOut,
		StdErr:
		return Constant{ChanelType, val}, nil
	case Cmd:
		return Constant{GlobalType, val}, nil
	default:
		return Constant{}, fmt.Errorf("Invalid constant value: %s", val)
	}
}
