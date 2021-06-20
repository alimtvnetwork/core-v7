package scripttype

import (
	"gitlab.com/evatix-go/core/coreinterface"
)

type Variant byte

const (
	Uninitialized Variant = iota
	Shell
	Bash
	Perl
	Python
	Python2
	Python3
	CLang
	MakeScript
	Powershell
	Cmd
)

func (receiver *Variant) Name() string {
	return scriptTypeBasicEnumImpl.ToEnumString(receiver.ValueByte())
}

func (receiver *Variant) ToNumberString() string {
	return scriptTypeBasicEnumImpl.ToNumberString(receiver.ValueByte())
}

func (receiver *Variant) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
	return scriptTypeBasicEnumImpl.UnmarshallEnumToValue(jsonUnmarshallingValue)
}

func (receiver *Variant) String() string {
	return scriptTypeBasicEnumImpl.ToEnumString(receiver.ValueByte())
}

func (receiver *Variant) MarshalJSON() ([]byte, error) {
	return scriptTypeBasicEnumImpl.ToEnumJsonBytes(receiver.ValueByte()), nil
}

func (receiver *Variant) UnmarshalJSON(data []byte) error {
	rawScriptType, err := scriptTypeBasicEnumImpl.UnmarshallEnumToValue(
		data)

	if err == nil {
		*receiver = Variant(rawScriptType)
	}

	return err
}

func (receiver *Variant) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return receiver
}

func (receiver *Variant) MaxByte() byte {
	return scriptTypeBasicEnumImpl.Max()
}

func (receiver *Variant) MinByte() byte {
	return scriptTypeBasicEnumImpl.Min()
}

func (receiver *Variant) ValueByte() byte {
	return byte(*receiver)
}

func (receiver *Variant) RangesByte() []byte {
	return scriptTypeBasicEnumImpl.Ranges()
}

func (receiver *Variant) RangesVariants() []Variant {
	return scriptTypeRanges[:]
}

func (receiver *Variant) ScriptDefault() *ScriptDefault {
	return RangesMap[*receiver]
}

func (receiver *Variant) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return receiver
}
