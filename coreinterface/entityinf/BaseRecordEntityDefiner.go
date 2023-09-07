package entityinf

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/internal/internalinterface"
	"gitlab.com/auk-go/core/internal/internalinterface/internalserializer"
)

type BaseRecordEntityDefiner interface {
	internalinterface.IdentifierWithEqualer
	internalinterface.TypeNameWithEqualer
	internalinterface.EntityTypeNameWithEqualer
	internalinterface.CategoryNameWithEqualer
	internalinterface.TableNamer
	corejson.Jsoner
	internalserializer.FieldBytesToPointerDeserializer
}
