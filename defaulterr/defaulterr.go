package defaulterr

import "gitlab.com/evatix-go/core/errcore"

var (
	Marshalling = errcore.
			MarshallingFailed.
			ErrorNoRefs("Cannot marshal object to serialize form.")

	UnMarshalling = errcore.
			UnMarshallingFailed.
			ErrorNoRefs("Cannot unmarshal data to object form.")

	UnMarshallingPlusCannotFindingEnumMap = errcore.
						UnMarshallingFailed.
						ErrorNoRefs(
			"Cannot find in the enum map. " +
				"Reference data given as : ")

	MarshallingFailedDueToNilOrEmpty = errcore.
						UnMarshallingFailed.
						ErrorNoRefs("Cannot marshal to serialize data because of nil or empty object.")

	UnMarshallingFailedDueToNilOrEmpty = errcore.
						UnMarshallingFailed.
						ErrorNoRefs("Cannot unmarshal to object because of nil or empty serialized data.")

	CannotProcessNilOrEmpty = errcore.
				CannotBeNilOrEmptyMessage.
				ErrorNoRefs("Cannot process nil or empty.")

	OutOfRange = errcore.
			OutOfRange.
			ErrorNoRefs("Cannot process out of range data.")

	NegativeDataCannotProcess = errcore.
					CannotBeNegativeMessage.
					ErrorNoRefs("Cannot process negative values.")

	NilResult = errcore.
			NullResultMessage.
			ErrorNoRefs("Cannot process nil result.")

	UnexpectedValue = errcore.
			UnexpectedValueErrorMessage.
			ErrorNoRefs("Cannot process unexpected value or values.")

	CannotRemoveFromEmptyCollection = errcore.
					CannotRemoveIndexesFromEmptyCollection.
					ErrorNoRefs("Cannot process request: cannot remove from empty collection.")

	CannotConvertStringToByte = errcore.
					FailedToConvert.
					ErrorNoRefs("Cannot convert string to byte.")
)
