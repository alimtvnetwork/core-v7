package entityinf

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/internal/internalinterface"
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
