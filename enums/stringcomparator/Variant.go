package stringcomparator

import "gitlab.com/evatix-go/core/coreinterface"

type Variant byte

const (
	Equal Variant = iota
	StartsWith
	EndsWith
	AnyWhere
	NotEqual
	Regex
)

func (receiver *Variant) Name() string {
	return basicEnumImpl.ToEnumString(receiver.ValueByte())
}

func (receiver *Variant) ToNumberString() string {
	return basicEnumImpl.ToNumberString(receiver.ValueByte())
}

func (receiver *Variant) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
	return basicEnumImpl.UnmarshallEnumToValue(jsonUnmarshallingValue)
}

func (receiver *Variant) String() string {
	return basicEnumImpl.ToEnumString(receiver.ValueByte())
}

func (receiver *Variant) MarshalJSON() ([]byte, error) {
	return basicEnumImpl.ToEnumJsonBytes(receiver.ValueByte()), nil
}

func (receiver *Variant) UnmarshalJSON(data []byte) error {
	rawScriptType, err := basicEnumImpl.UnmarshallEnumToValue(data)

	if err != nil {
		*receiver = Variant(rawScriptType)
	}

	return err
}

func (receiver *Variant) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return receiver
}

func (receiver *Variant) MaxByte() byte {
	return basicEnumImpl.Max()
}

func (receiver *Variant) MinByte() byte {
	return basicEnumImpl.Min()
}

func (receiver *Variant) ValueByte() byte {
	return byte(*receiver)
}

func (receiver *Variant) RangesByte() []byte {
	return basicEnumImpl.Ranges()
}

func (receiver *Variant) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return receiver
}
