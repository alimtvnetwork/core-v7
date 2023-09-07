package entityinf

import "gitlab.com/auk-go/core/internal/internalinterface"

type RecordEntityDefiner interface {
	BaseRecordEntityDefiner
	internalinterface.DefaultsInjector
	internalinterface.RawPayloadsGetter
}
