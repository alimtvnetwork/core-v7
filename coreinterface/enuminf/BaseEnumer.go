package enuminf

import "gitlab.com/auk-go/core/coredata/corejson"

type BaseEnumer interface {
	enumNameStinger
	SimpleEnumer
	NameValuer
	IsNameEqualer
	IsAnyNameOfChecker
	ToNumberStringer
	IsValidInvalidChecker
	BasicEnumValuer
	RangeNamesCsvGetter
	corejson.JsonMarshaller
}
