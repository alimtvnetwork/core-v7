package scripttype

import (
	"gitlab.com/evatix-go/core/coreinterface"
)

type Variant byte

const (
	Uninitialized Variant = iota
	Default
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

func (receiver Variant) IsUninitialized() bool {
	return receiver == Uninitialized
}

func (receiver Variant) IsDefault() bool {
	return receiver == Default
}

func (receiver Variant) IsShell() bool {
	return receiver == Shell
}

func (receiver Variant) IsBash() bool {
	return receiver == Bash
}

func (receiver Variant) IsPerl() bool {
	return receiver == Perl
}

func (receiver Variant) IsPython() bool {
	return receiver == Python
}

func (receiver Variant) IsPython2() bool {
	return receiver == Python2
}

func (receiver Variant) IsPython3() bool {
	return receiver == Python3
}

func (receiver Variant) IsCLang() bool {
	return receiver == CLang
}

func (receiver Variant) IsMakeScript() bool {
	return receiver == MakeScript
}

func (receiver Variant) IsPowershell() bool {
	return receiver == Powershell
}

func (receiver Variant) IsCmd() bool {
	return receiver == Cmd
}

func (receiver Variant) IsCmdOrPowerShell() bool {
	return receiver.IsCmd() ||
		receiver.IsPowershell()
}

func (receiver Variant) IsAnyPython() bool {
	return receiver.IsPython() ||
		receiver.IsPython2() ||
		receiver.IsPython3()
}

func (receiver *Variant) RangesVariants() []Variant {
	return scriptTypeRanges[:]
}

func (receiver *Variant) ScriptDefault() *ScriptDefault {
	if receiver.IsDefault() {
		return DefaultOsScript()
	}

	return RangesMap[*receiver]
}

func (receiver *Variant) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return receiver
}
