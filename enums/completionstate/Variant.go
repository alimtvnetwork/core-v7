package completionstate

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
)

type Variant byte

const (
	Unknown Variant = iota
	Running
	Success
	SuccessWithWarning
	FailedMiddleWithError // it means exited in middle
	CompleteWithError     // completed but has error
)

func (it Variant) IsSuccess() bool {
	return it == Success
}

func (it Variant) IsCompletedLogically() bool {
	return CompletionMap[it]
}

func (it Variant) IsSuccessLogically() bool {
	return it == Success || it == SuccessWithWarning
}

func (it Variant) IsCompletedWithErrorLogically() bool {
	return it == FailedMiddleWithError || it == CompleteWithError
}

func (it Variant) HasErrorLogically() bool {
	return it == FailedMiddleWithError || it == CompleteWithError
}

func (it Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.ValueByte())
}

func (it Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.ValueByte()), nil
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	dataConv, err := it.UnmarshallEnumToValue(
		data)

	if err == nil {
		*it = Variant(dataConv)
	}

	return err
}

func (it Variant) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it Variant) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(
		true,
		jsonUnmarshallingValue)
}

func (it Variant) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it Variant) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it Variant) ValueByte() byte {
	return byte(it)
}

func (it Variant) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it Variant) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it Variant) String() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it *Variant) JsonParseSelfInject(jsonResult *corejson.Result) error {
	err := jsonResult.Unmarshal(it)

	return err
}

func (it Variant) Json() corejson.Result {
	return corejson.New(it)
}

func (it Variant) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Variant) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Variant) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Variant) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it *Variant) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return it
}

func (it *Variant) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return it
}
