package gencode

import(
	"errors"
)

type Source int
const(
	SourceTemplate Source = iota
	SourceInterface
)
const(
	SourceFlagName = "Source"
)
func (sr *Source) String() string {
	switch *sr {
	case SourceTemplate:
		return "template"
	case SourceInterface:
		return "interface"
	}

	return "unknow"
}
func (sr *Source) Set(s string) error {
	switch s {
	case "template", "t":
		*sr = SourceTemplate
	case "interface", "i":
		*sr = SourceInterface
	default:
		return errors.New("error Source type")
	}

	return nil
}
func (sr *Source) Type() string {
	return SourceFlagName
}