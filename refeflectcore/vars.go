package refeflectcore

import (
	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

var (
	Converter                      = reflectinternal.Converter
	Utils                          = reflectinternal.Utils
	Looper                         = reflectinternal.Looper
	CodeStack                      = reflectinternal.CodeStack
	GetFunc                        = reflectinternal.GetFunc
	Is                             = reflectinternal.Is
	TypeName                       = reflectinternal.TypeName
	TypeNames                      = reflectinternal.TypeNames
	TypeNamesString                = reflectinternal.TypeNamesString
	TypeNamesReferenceString       = reflectinternal.TypeNamesReferenceString
	ReflectType                    = reflectinternal.ReflectType
	ReflectGetter                  = reflectinternal.ReflectGetter
	ReflectGetterUsingReflectValue = reflectinternal.ReflectGetterUsingReflectValue
	SliceConverter                 = reflectinternal.SliceConverter
	MapConverter                   = reflectinternal.MapConverter

	indexToPositionFunc   = convertinteranl.Util.String.IndexToPosition
	prependWithSpacesFunc = convertinteranl.Util.String.PrependWithSpacesDefault
)
