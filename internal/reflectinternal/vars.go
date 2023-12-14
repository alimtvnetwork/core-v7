package reflectinternal

import "gitlab.com/auk-go/core/internal/convertinteranl"

var (
	Converter = reflectConverter{}
	Utils     = reflectUtils{}
	Looper    = looper{}
	CodeStack = codeStack{}
	GetFunc   = getFunc{}
	Is        = isChecker{}

	ReflectGetter                  = reflectGetter{}
	ReflectType                    = reflectTypeConverter{}
	ReflectGetterUsingReflectValue = reflectGetUsingReflectValue{}
	SliceConverter                 = sliceConverter{}
	MapConverter                   = mapConverter{}
	Path                           = reflectPath{}

	indexToPositionFunc   = convertinteranl.Util.String.IndexToPosition
	prependWithSpacesFunc = convertinteranl.Util.String.PrependWithSpacesDefault
	repoDir               *string // will be updated, dangerous but this is the way for this now
)
