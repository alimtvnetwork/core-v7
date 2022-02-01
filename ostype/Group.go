package ostype

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
)

type Group byte

const (
	WindowsGroup Group = iota
	UnixGroup
	AndroidGroup
	InvalidGroup
)

func (it Group) Is(another Group) bool {
	return it == another
}

func (it Group) IsWindows() bool {
	return it == WindowsGroup
}

func (it Group) IsUnix() bool {
	return it == UnixGroup
}

func (it Group) IsAndroid() bool {
	return it == AndroidGroup
}

func (it Group) IsInvalidGroup() bool {
	return it == InvalidGroup
}

func (it Group) Byte() byte {
	return byte(it)
}

func (it Group) MarshalJSON() ([]byte, error) {
	return basicEnumImplOsGroup.ToEnumJsonBytes(it.Value()), nil
}

func (it *Group) UnmarshalJSON(data []byte) error {
	valueByte, err := basicEnumImplOsGroup.UnmarshallToValue(
		true,
		data)

	if err == nil {
		*it = Group(valueByte)
	}

	return err
}

func (it Group) Name() string {
	return basicEnumImplOsGroup.ToEnumString(it.Value())
}

func (it Group) NameValue() string {
	return basicEnumImplOsGroup.NameWithValue(it.Value())
}

func (it Group) ToNumberString() string {
	return basicEnumImplOsGroup.ToNumberString(it.Value())
}

func (it Group) RangeNamesCsv() string {
	return basicEnumImplOsGroup.RangeNamesCsv()
}

func (it Group) TypeName() string {
	return basicEnumImplOsGroup.TypeName()
}

func (it Group) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
	return basicEnumImplOsGroup.UnmarshallToValue(true, jsonUnmarshallingValue)
}

func (it Group) MaxByte() byte {
	return basicEnumImplOsGroup.Max()
}

func (it Group) MinByte() byte {
	return basicEnumImplOsGroup.Min()
}

func (it Group) ValueByte() byte {
	return byte(it)
}

func (it Group) RangesByte() []byte {
	return basicEnumImplOsGroup.Ranges()
}

func (it Group) Value() byte {
	return byte(it)
}

func (it Group) String() string {
	return basicEnumImplOsGroup.ToEnumString(it.Value())
}

func (it Group) IsValid() bool {
	return it != InvalidGroup
}

func (it Group) IsInvalid() bool {
	return it == InvalidGroup
}

func (it *Group) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return it
}

func (it *Group) AsJsonContractsBinder() corejson.JsonMarshaller {
	return it
}

func (it Group) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return &it
}
