package envtype

import (
	"gitlab.com/evatix-go/core/coreinterface"
)

type Variant byte

const (
	Uninitialized Variant = iota
	Development
	Development1
	Development2
	Test
	Test1
	Test2
	Production
	Production1
	Production2
)

func (it *Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it *Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.ValueByte())
}

func (it Variant) IsUninitialized() bool {
	return it == Uninitialized
}

func (it Variant) IsInitialized() bool {
	return it != Uninitialized
}

func (it Variant) IsDevelopment() bool {
	return it == Development
}

func (it Variant) IsDevelopment1() bool {
	return it == Development1
}

func (it Variant) IsDevelopment2() bool {
	return it == Development2
}

func (it Variant) IsAnyDevelopment() bool {
	return it == Development ||
		it == Development1 ||
		it == Development2
}

func (it Variant) IsTest() bool {
	return it == Test
}

func (it Variant) IsTest1() bool {
	return it == Test1
}

func (it Variant) IsTest2() bool {
	return it == Test2
}

func (it Variant) IsAnyTest() bool {
	return it == Test || it == Test1 || it == Test2
}

func (it Variant) IsProduction() bool {
	return it == Production
}

func (it Variant) IsProduction1() bool {
	return it == Production1
}

func (it Variant) IsProduction2() bool {
	return it == Production2
}

func (it Variant) IsAnyProduction() bool {
	return it == Production ||
		it == Production1 ||
		it == Production2
}

func (it *Variant) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(
		isMappedToDefault,
		jsonUnmarshallingValue)
}

func (it *Variant) String() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Variant) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it *Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.ValueByte()), nil
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
	return BasicEnumImpl.Max()
}

func (it *Variant) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it *Variant) ValueByte() byte {
	return byte(*it)
}

func (it *Variant) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it *Variant) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return it
}
