package defaulterr

import "gitlab.com/evatix-go/core/msgtype"

var (
	Marshalling = msgtype.
			MarshallingFailed.
			ErrorNoRefs("Cannot marshal object to serialize form.")

	UnMarshalling = msgtype.
			UnMarshallingFailed.
			ErrorNoRefs("Cannot unmarshal data to object form.")

	MarshallingFailedDueToNilOrEmpty = msgtype.
						UnMarshallingFailed.
						ErrorNoRefs("Cannot marshal to serialize data because of nil or empty object.")

	UnMarshallingFailedDueToNilOrEmpty = msgtype.
						UnMarshallingFailed.
						ErrorNoRefs("Cannot unmarshal to object because of nil or empty serialized data.")

	CannotProcessNilOrEmpty = msgtype.
				CannotBeNilOrEmptyMessage.
				ErrorNoRefs("Cannot process nil or empty.")

	OutOfRange = msgtype.
			OutOfRange.
			ErrorNoRefs("Cannot process out of range data.")

	NegativeDataCannotProcess = msgtype.
					CannotBeNegativeMessage.
					ErrorNoRefs("Cannot process negative values.")

	NilResult = msgtype.
			NullResultMessage.
			ErrorNoRefs("Cannot process nil result.")

	UnexpectedValue = msgtype.
			UnexpectedValueErrorMessage.
			ErrorNoRefs("Cannot process unexpected value or values.")

	CannotRemoveFromEmptyCollection = msgtype.
					CannotRemoveIndexesFromEmptyCollection.
					ErrorNoRefs("Cannot process request: cannot remove from empty collection.")

	CannotConvertStringToByte = msgtype.
					FailedToConvert.
					ErrorNoRefs("Cannot convert string to byte.")
)
