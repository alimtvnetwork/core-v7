package corejson

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

type deserializerLogic struct{}

func (it deserializerLogic) ToStrings(
	jsonResult *Result,
) (lines []string, err error) {
	var slice []string
	err = it.Apply(jsonResult, &slice)

	return slice, err
}

func (it deserializerLogic) ToString(
	jsonResult *Result,
) (line string, err error) {
	err = it.Apply(jsonResult, &line)

	return line, err
}

func (it deserializerLogic) ToBool(
	jsonResult *Result,
) (isResult bool, err error) {
	err = it.Apply(jsonResult, &isResult)

	return isResult, err
}

func (it deserializerLogic) ToByte(
	jsonResult *Result,
) (byteVal byte, err error) {
	err = it.Apply(jsonResult, &byteVal)

	return byteVal, err
}

func (it deserializerLogic) ToByteMust(
	jsonResult *Result,
) byte {
	result, err := it.ToByte(jsonResult)

	if err != nil {
		panic(err)
	}

	return result
}

func (it deserializerLogic) ToBoolMust(
	jsonResult *Result,
) bool {
	result, err := it.ToBool(jsonResult)

	if err != nil {
		panic(err)
	}

	return result
}

func (it deserializerLogic) ToStringMust(
	jsonResult *Result,
) string {
	result, err := it.ToString(jsonResult)

	if err != nil {
		panic(err)
	}

	return result
}

func (it deserializerLogic) ToStringsMust(
	jsonResult *Result,
) []string {
	results, err := it.ToStrings(jsonResult)

	if err != nil {
		panic(err)
	}

	return results
}

func (it deserializerLogic) Apply(
	jsonResult *Result,
	unmarshalToPointer interface{},
) error {
	return jsonResult.Unmarshal(
		unmarshalToPointer)
}

func (it deserializerLogic) ApplyMust(
	jsonResult *Result,
	unmarshalToPointer interface{},
) {
	err := jsonResult.Unmarshal(
		unmarshalToPointer)

	if err != nil {
		panic(err)
	}
}

func (it deserializerLogic) UsingBytes(
	rawBytes []byte,
	unmarshalToPointer interface{},
) error {
	err := json.Unmarshal(rawBytes, unmarshalToPointer)

	if err != nil {
		reference := errcore.Var2NoType(
			"JsonResult Error", err.Error(),
			"To Reference Type", reflectinternal.TypeName(unmarshalToPointer))

		return errcore.
			UnMarshallingFailedType.
			Error(
				"failed to unmarshal bytes.",
				reference)
	}

	return nil
}

func (it deserializerLogic) UsingBytesMust(
	rawBytes []byte,
	unmarshalToPointer interface{},
) {
	err := it.UsingBytes(rawBytes, unmarshalToPointer)

	if err != nil {
		panic(err)
	}
}

func (it deserializerLogic) UsingSafeBytesMust(
	rawBytes []byte,
	unmarshalToPointer interface{},
) {
	if len(rawBytes) == 0 {
		return
	}

	err := it.UsingBytes(rawBytes, unmarshalToPointer)

	if err != nil {
		panic(err)
	}
}
