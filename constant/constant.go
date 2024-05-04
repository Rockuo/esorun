package constant

import "fmt"

type ConstType int

type Constant struct {
	variant ConstType
	value   string
}

const (
	// types
	InvalidType = iota
	ChanelType  = iota
	GlobalType

	// chanels
	StdIn  = "STD_IN"
	StdOut = "STD_OUT"
	StdErr = "STD_ERR"

	// globals
	Cmd = "CMD"
)

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
