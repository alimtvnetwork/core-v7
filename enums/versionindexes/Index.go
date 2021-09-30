package versionindexes

import "gitlab.com/evatix-go/core/coreinterface"

type Index byte

const (
	Major = iota
	Minor = 1
	Patch = 2
	Build = 3
)

func (it *Index) Name() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it *Index) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.ValueByte())
}

func (it *Index) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(
		isMappedToDefault,
		jsonUnmarshallingValue)
}

func (it *Index) String() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Index) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Index) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it *Index) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.ValueByte()), nil
}

func (it *Index) UnmarshalJSON(data []byte) error {
	rawScriptType, err := it.UnmarshallEnumToValue(
		data)

	if err == nil {
		*it = Index(rawScriptType)
	}

	return err
}

func (it *Index) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return it
}

func (it *Index) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it *Index) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it *Index) ValueByte() byte {
	return byte(*it)
}

func (it *Index) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it Index) IsMajor() bool {
	return it == Major
}

func (it Index) IsMinor() bool {
	return it == Minor
}

func (it Index) IsPatch() bool {
	return it == Patch
}

func (it Index) IsBuild() bool {
	return it == Build
}

func (it *Index) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return it
}
