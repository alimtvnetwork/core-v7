package scripttype

import (
	"fmt"

	"gitlab.com/evatix-go/core/converters"
)

type ScriptDefault struct {
	ScriptType       Variant
	ProcessName      string
	DefaultArguments []string
	IsImplemented    bool
}

func (receiver *ScriptDefault) String() string {
	return fmt.Sprint(receiver.ScriptType.String(), converters.AnyToString(*receiver))
}
