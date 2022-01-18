package corejson

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/coredata"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

type newResultCreator struct{}

// UsingUnmarshalBytes
//
//  Aka. alias for UsingDeserializeBytes
//
//  Should be used when Result itself is Serialized
//  and save to somewhere and then unmarshal or deserialize
func (it newResultCreator) UsingUnmarshalBytes(
	deserializingBytes []byte,
) *Result {
	return it.UsingDeserializeBytes(deserializingBytes)
}

// UsingDeserializeBytes
//
//  Should be used when Result itself is Serialized
//  and save to somewhere and then unmarshal or deserialize
func (it newResultCreator) UsingDeserializeBytes(
	deserializingBytes []byte,
) *Result {
	empty := it.TypeName(resultTypeName)

	err := Deserialize.
		UsingBytes(deserializingBytes, empty)

	if err == nil {
		return empty
	}

	empty.Error = err

	return empty
}

func (it newResultCreator) UsingBytes(
	jsonBytes []byte,
) Result {
	return Result{
		Bytes: jsonBytes,
	}
}

func (it newResultCreator) UsingErrType(
	typeName string,
	jsonBytes []byte,
) Result {
	return Result{
		Bytes:    jsonBytes,
		TypeName: typeName,
	}
}

func (it newResultCreator) UsingBytesPtr(
	jsonBytes *[]byte,
) *Result {
	if jsonBytes == nil || *jsonBytes == nil {
		return &Result{}
	}

	return &Result{
		Bytes: *jsonBytes,
	}
}

func (it newResultCreator) UsingBytesErrPtr(
	jsonBytes *[]byte, err error, typeName string,
) *Result {
	if jsonBytes == nil || *jsonBytes == nil {
		return &Result{}
	}

	return &Result{
		Bytes:    *jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) Ptr(
	jsonBytes []byte,
	err error,
	typeName string,
) *Result {
	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) CreatePtr(
	jsonBytes []byte,
	err error,
	typeName string,
) *Result {
	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) NonPtr(
	jsonBytes []byte,
	err error,
	typeName string,
) Result {
	return Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) Create(
	jsonBytes []byte,
	err error,
	typeName string,
) Result {
	return Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) PtrUsingBytesPtr(
	jsonBytes *[]byte,
	err error,
	typeName string,
) *Result {
	if err != nil {
		return &Result{
			Bytes:    []byte{},
			Error:    err,
			TypeName: typeName,
		}
	}

	if jsonBytes == nil {
		return &Result{
			Bytes:    []byte{},
			Error:    nil,
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    *jsonBytes,
		Error:    nil,
		TypeName: typeName,
	}
}

func (it newResultCreator) Any(
	anyItem interface{},
) Result {
	jsonBytes, err := json.Marshal(anyItem)
	typeName := reflectinternal.TypeName(anyItem)

	if err != nil {
		return Result{
			Bytes: jsonBytes,
			Error: errcore.MarshallingFailedType.Error(
				err.Error(),
				typeName),
			TypeName: typeName,
		}
	}

	return Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) AnyPtr(
	anyItem interface{},
) *Result {
	jsonBytes, err := json.Marshal(anyItem)
	typeName := reflectinternal.TypeName(anyItem)

	if err != nil {
		return &Result{
			Bytes: jsonBytes,
			Error: errcore.MarshallingFailedType.Error(
				err.Error(),
				typeName),
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

// UsingBytesError Get created with nil.
func (it newResultCreator) UsingBytesError(
	bytesError *coredata.BytesError,
) Result {
	if bytesError == nil {
		return Result{}
	}

	return Result{
		Bytes: bytesError.Bytes,
		Error: bytesError.Error,
	}
}

func (it newResultCreator) Error(err error) Result {
	return Result{
		Bytes: nil,
		Error: err,
	}
}

func (it newResultCreator) ErrorPtr(err error) *Result {
	return &Result{
		Bytes: nil,
		Error: err,
	}
}

func (it newResultCreator) Empty() Result {
	return Result{}
}

func (it newResultCreator) EmptyPtr() *Result {
	return &Result{}
}

func (it newResultCreator) TypeName(typeName string) *Result {
	return &Result{
		TypeName: typeName,
	}
}

func (it newResultCreator) TypeNameBytes(typeName string) *Result {
	return &Result{
		TypeName: typeName,
	}
}

func (it newResultCreator) Many(
	anyItems ...interface{},
) *Result {
	return it.AnyPtr(anyItems)
}

func (it newResultCreator) Serialize(
	anyItem interface{},
) *Result {
	jsonBytes, err := json.Marshal(anyItem)
	typeName := reflectinternal.TypeName(anyItem)

	if err != nil {
		return &Result{
			Bytes: jsonBytes,
			Error: errcore.MarshallingFailedType.Error(
				err.Error(),
				typeName),
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) Marshal(
	anyItem interface{},
) *Result {
	jsonBytes, err := json.Marshal(anyItem)
	typeName := reflectinternal.TypeName(anyItem)

	if err != nil {
		return &Result{
			Bytes: jsonBytes,
			Error: errcore.MarshallingFailedType.Error(
				err.Error(),
				typeName),
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}
