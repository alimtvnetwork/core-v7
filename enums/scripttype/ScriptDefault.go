package scripttype

import (
	"fmt"

	"gitlab.com/evatix-go/core/converters"
)

type ScriptDefault struct {
	ScriptType       Variant
	IsImplemented    bool
	ProcessName      string
	DefaultArguments []string
}

func (it *ScriptDefault) String() string {
	return fmt.Sprint(
		it.ScriptType.String(),
		converters.AnyToValueString(*it))
}
