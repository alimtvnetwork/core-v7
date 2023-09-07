package entityinf

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/internal/internalinterface"
)

type TaskEntityDefiner interface {
	internalinterface.UsernameGetter
	internalinterface.AnyValueGetter
	internalinterface.ErrorGetter

	coreinterface.ReflectSetter
	Deserialize(
		anyPointer interface{},
	) error

	corejson.Jsoner
}
