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

func (it *Variant) Name() string {
	return scriptTypeBasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Variant) NameValue() string {
	return scriptTypeBasicEnumImpl.NameWithValue(it)
}

func (it *Variant) ToNumberString() string {
	return scriptTypeBasicEnumImpl.ToNumberString(it.ValueByte())
}

func (it *Variant) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return scriptTypeBasicEnumImpl.UnmarshallToValue(
		isMappedToDefault,
		jsonUnmarshallingValue)
}

func (it Variant) String() string {
	return scriptTypeBasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Variant) RangeNamesCsv() string {
	return scriptTypeBasicEnumImpl.RangeNamesCsv()
}

func (it Variant) TypeName() string {
	return scriptTypeBasicEnumImpl.TypeName()
}

func (it *Variant) MarshalJSON() ([]byte, error) {
	return scriptTypeBasicEnumImpl.ToEnumJsonBytes(it.ValueByte()), nil
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	rawScriptType, err := it.UnmarshallEnumToValue(
		data)

	if err == nil {
		*it = Variant(rawScriptType)
	}

	return err
}

func (it *Variant) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return it
}

func (it *Variant) MaxByte() byte {
	return scriptTypeBasicEnumImpl.Max()
}

func (it *Variant) MinByte() byte {
	return scriptTypeBasicEnumImpl.Min()
}

func (it *Variant) ValueByte() byte {
	return byte(*it)
}

func (it *Variant) RangesByte() []byte {
	return scriptTypeBasicEnumImpl.Ranges()
}

func (it Variant) IsUninitialized() bool {
	return it == Uninitialized
}

func (it Variant) IsDefault() bool {
	return it == Default
}

func (it Variant) IsShell() bool {
	return it == Shell
}

func (it Variant) IsBash() bool {
	return it == Bash
}

func (it Variant) IsPerl() bool {
	return it == Perl
}

func (it Variant) IsPython() bool {
	return it == Python
}

func (it Variant) IsPython2() bool {
	return it == Python2
}

func (it Variant) IsPython3() bool {
	return it == Python3
}

func (it Variant) IsCLang() bool {
	return it == CLang
}

func (it Variant) IsMakeScript() bool {
	return it == MakeScript
}

func (it Variant) IsPowershell() bool {
	return it == Powershell
}

func (it Variant) IsCmd() bool {
	return it == Cmd
}

func (it Variant) IsCmdOrPowerShell() bool {
	return it.IsCmd() ||
		it.IsPowershell()
}

func (it Variant) IsAnyPython() bool {
	return it.IsPython() ||
		it.IsPython2() ||
		it.IsPython3()
}

func (it *Variant) RangesVariants() []Variant {
	return scriptTypeRanges[:]
}

func (it *Variant) ScriptDefault() *ScriptDefault {
	if it.IsDefault() {
		return DefaultOsScript()
	}

	return RangesMap[*it]
}

func (it *Variant) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return it
}
