package corejson

import (
	"encoding/json"
	"errors"

	"gitlab.com/evatix-go/core/coredata"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

type newResultCreator struct{}

// UnmarshalUsingBytes
//
//  Aka. alias for DeserializeUsingBytes
//
//  Should be used when Result itself is Serialized
//  and save to somewhere and then unmarshal or deserialize
func (it newResultCreator) UnmarshalUsingBytes(
	deserializingBytes []byte,
) *Result {
	return it.DeserializeUsingBytes(deserializingBytes)
}

// DeserializeUsingBytes
//
//  Should be used when Result itself is Serialized
//  and save to somewhere and then unmarshal or deserialize
func (it newResultCreator) DeserializeUsingBytes(
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

func (it newResultCreator) DeserializeUsingResult(
	jsonResult *Result,
) *Result {
	if jsonResult.HasIssuesOrEmpty() {
		return it.ErrorPtr(jsonResult.MeaningfulError())
	}

	empty := it.TypeName(resultTypeName)

	err := Deserialize.
		UsingBytes(jsonResult.SafeBytes(), empty)

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

func (it newResultCreator) UsingBytesType(
	typeName string,
	jsonBytes []byte,
) Result {
	return Result{
		Bytes:    jsonBytes,
		TypeName: typeName,
	}
}

func (it newResultCreator) UsingBytesTypePtr(
	typeName string,
	jsonBytes []byte,
) *Result {
	return &Result{
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

func (it newResultCreator) UsingBytesPtrErrPtr(
	jsonBytes *[]byte, err error, typeName string,
) *Result {
	if jsonBytes == nil || *jsonBytes == nil {
		return &Result{
			Error:    err,
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    *jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) UsingBytesErrPtr(
	jsonBytes []byte, err error, typeName string,
) *Result {
	if len(jsonBytes) == 0 {
		return &Result{
			Bytes:    []byte{},
			Error:    err,
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) PtrUsingStringPtr(
	jsonStringPtr *string,
	typeName string,
) *Result {
	if jsonStringPtr == nil {
		return it.PtrUsingBytesPtr(
			nil,
			errors.New("json string ptr is nil cannot process further"),
			typeName)
	}

	return &Result{
		Bytes:    []byte(*jsonStringPtr),
		TypeName: typeName,
	}
}

func (it newResultCreator) UsingErrorStringPtr(
	err error,
	jsonStringPtr *string,
	typeName string,
) *Result {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	if jsonStringPtr == nil {
		return it.PtrUsingBytesPtr(
			nil,
			errors.New("json string ptr is nil cannot process further"+errMsg),
			typeName)
	}

	return &Result{
		Bytes:    []byte(*jsonStringPtr),
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

func (it newResultCreator) UsingSerializer(
	serializer bytesSerializer,
) *Result {
	if serializer == nil {
		return nil
	}

	allBytes, err := serializer.Serialize()

	return &Result{
		Bytes: allBytes,
		Error: err,
		TypeName: reflectinternal.TypeName(
			serializer),
	}
}

func (it newResultCreator) UsingSerializerFunc(
	serializerFunc func() ([]byte, error),
) *Result {
	if serializerFunc == nil {
		return nil
	}

	allBytes, err := serializerFunc()

	return &Result{
		Bytes:    allBytes,
		Error:    err,
		TypeName: reflectinternal.TypeName(serializerFunc),
	}
}

func (it newResultCreator) UsingJsoner(
	jsoner Jsoner,
) *Result {
	if jsoner == nil {
		return nil
	}

	return jsoner.JsonPtr()
}
