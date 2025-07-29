package gencode

import "errors"

type Mode int
const(
	ModeOverwrite Mode = iota
	ModeAppend
	ModeGenIfNotExists
)
const(
	ModeFlagName = "source"
)
func (m *Mode) String() string {
	switch *m {
	case ModeGenIfNotExists:
		return "generate if not exists"
	case ModeAppend:
		return "append"
	case ModeOverwrite:
		return "overwrite"
	}

	return "unknow"
}

func (m *Mode) Set(s string) error {
	switch s {
	case "append", "a":
		*m = ModeAppend
	case "overwrite", "o":
		*m = ModeOverwrite
	case "n":
		*m = ModeGenIfNotExists
	default:
		return errors.New("error Mode type")
	}

	return nil
}

func (m *Mode) Type() string {
	return ModeFlagName
}