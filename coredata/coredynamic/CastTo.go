package coredynamic

import (
	"reflect"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func CastTo(
	isOutputPointer bool,
	input interface{},
	acceptedTypes ...reflect.Type,
) CastedResult {
	currentRfType := reflect.TypeOf(input)
	rv := reflect.ValueOf(input)
	kind := rv.Kind()
	var sliceErr []string

	isMatchingAcceptedType := IsAnyTypesOf(
		currentRfType,
		acceptedTypes...,
	)

	if !isMatchingAcceptedType {
		// not matching
		sliceErr = append(
			sliceErr,
			errcore.UnsupportedType.Combine(
				"none matches, current type:"+currentRfType.String(),
				getTypeNamesUsingReflectFunc(true, acceptedTypes...),
			),
		)
	}

	isNull := input == nil || reflectinternal.Is.NullRv(
		rv,
	)
	isOutNonPointer := !isOutputPointer
	hasNonPointerIssue := isNull && isOutNonPointer

	if hasNonPointerIssue {
		// has issue
		// cannot non pointer a nil pointer
		// will panic
		sliceErr = append(
			sliceErr,
			errcore.
				InvalidNullPointerType.
				SrcDestination(
					"cannot output non pointer if pointer is null",
					"Value", constants.NilAngelBracket,
					"Type", currentRfType.String(),
				),
		)

		// ending process
		return CastedResult{
			Casted:                 nil,
			SourceReflectType:      currentRfType,
			SourceKind:             kind,
			Error:                  errcore.SliceToError(sliceErr),
			IsNull:                 isNull,
			IsMatchingAcceptedType: isMatchingAcceptedType,
			IsPointer:              isOutNonPointer,
			IsSourcePointer:        kind == reflect.Ptr,
			IsValid:                rv.IsValid(),
		}
	}

	val, _ := PointerOrNonPointerUsingReflectValue(
		isOutputPointer,
		rv,
	)

	return CastedResult{
		Casted:                 val,
		SourceReflectType:      currentRfType,
		SourceKind:             kind,
		Error:                  errcore.SliceToError(sliceErr),
		IsNull:                 isNull,
		IsMatchingAcceptedType: isMatchingAcceptedType,
		IsPointer:              isOutNonPointer,
		IsSourcePointer:        kind == reflect.Ptr,
		IsValid:                rv.IsValid(),
	}
}
