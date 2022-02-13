package enuminf

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
)

type BasicEnumer interface {
	enumNameStinger
	nameValuer
	IsNameEqualer
	IsAnyNameOfChecker
	ToNumberStringer
	IsValidInvalidChecker
	BasicEnumValuer
	EnumFormatter
	EnumType() EnumTyper
	RangeNamesCsvGetter
	corejson.JsonMarshaller
}
