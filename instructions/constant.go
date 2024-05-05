package instructions

type ConstType int

type Constant struct {
	Variant ConstType
	Value   string
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
