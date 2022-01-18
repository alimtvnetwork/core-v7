package defaulterr

import "gitlab.com/evatix-go/core/errcore"

var (
	Marshalling = errcore.
			MarshallingFailedType.
			ErrorNoRefs("Cannot marshal object to serialize form.")

	UnMarshalling = errcore.
			UnMarshallingFailedType.
			ErrorNoRefs("Cannot unmarshal data to object form.")

	UnMarshallingPlusCannotFindingEnumMap = errcore.
						UnMarshallingFailedType.
						ErrorNoRefs(
			"Cannot find in the enum map. " +
				"Reference data given as : ")

	MarshallingFailedDueToNilOrEmpty = errcore.
						UnMarshallingFailedType.
						ErrorNoRefs("Cannot marshal to serialize data because of nil or empty object.")

	UnMarshallingFailedDueToNilOrEmpty = errcore.
						UnMarshallingFailedType.
						ErrorNoRefs("Cannot unmarshal to object because of nil or empty serialized data.")

	CannotProcessNilOrEmpty = errcore.
				CannotBeNilOrEmptyType.
				ErrorNoRefs("Cannot process nil or empty.")

	OutOfRange = errcore.
			OutOfRangeType.
			ErrorNoRefs("Cannot process out of range data.")

	NegativeDataCannotProcess = errcore.
					CannotBeNegativeType.
					ErrorNoRefs("Cannot process negative values.")

	NilResult = errcore.
			NullResultType.
			ErrorNoRefs("Cannot process nil result.")

	UnexpectedValue = errcore.
			UnexpectedValueType.
			ErrorNoRefs("Cannot process unexpected value or values.")

	CannotRemoveFromEmptyCollection = errcore.
					CannotRemoveIndexesFromEmptyCollectionType.
					ErrorNoRefs("Cannot process request: cannot remove from empty collection.")

	CannotConvertStringToByte = errcore.
					FailedToConvertType.
					ErrorNoRefs("Cannot convert string to byte.")

	AttributeNull = errcore.
			NullResultType.
			ErrorNoRefs("attribute is nil!")
)
