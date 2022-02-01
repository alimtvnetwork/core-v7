package ostype

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
)

type Variation byte

// https://stackoverflow.com/a/50117892 | https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63
// go tool dist list
const (
	Any Variation = iota
	Unknown
	Windows
	Linux
	DarwinOrMacOs
	JavaScript
	FreeBsd
	NetBsd
	OpenBsd
	DragonFly
	Android
	Plan9
	Solaris
	Nacl
	Illumos
	IOs
	Aix
)

func (it Variation) IsValid() bool {
	return it.Value() != 0
}

func (it Variation) IsInvalid() bool {
	return it.Value() == 0
}

func (it Variation) IsByte(another byte) bool {
	return it == Variation(another)
}

func (it Variation) IsAnyOperatingSystem() bool {
	return Any == it
}

func (it Variation) Is(other Variation) bool {
	return other == it
}

func (it Variation) IsAnyMatch(others ...Variation) bool {
	for _, other := range others {
		if other == it {
			return true
		}
	}

	return false
}

func (it Variation) IsStringsMatchAny(others ...string) bool {
	for _, other := range others {
		otherVariant := GetVariant(other)

		if otherVariant == it {
			return true
		}
	}

	return false
}

// IsPossibleUnixGroup variation != Windows
func (it Variation) IsPossibleUnixGroup() bool {
	return it != Windows
}

func (it Variation) IsLinuxOrMac() bool {
	return it == Linux || it == DarwinOrMacOs
}

func (it Variation) Group() Group {
	if it == Windows {
		return WindowsGroup
	}

	if it == Android {
		return AndroidGroup
	}

	return UnixGroup
}

func (it Variation) IsActualGroupUnix() bool {
	return it.Group().IsUnix()
}

func (it Variation) IsWindows() bool {
	return it == Windows
}

func (it Variation) IsLinux() bool {
	return it == Linux
}

func (it Variation) IsDarwinOrMacOs() bool {
	return it == DarwinOrMacOs
}

func (it Variation) IsJavaScript() bool {
	return it == JavaScript
}

func (it Variation) IsFreeBsd() bool {
	return it == FreeBsd
}

func (it Variation) IsNetBsd() bool {
	return it == NetBsd
}

func (it Variation) IsOpenBsd() bool {
	return it == NetBsd
}

func (it Variation) IsDragonFly() bool {
	return it == DragonFly
}

func (it Variation) MarshalJSON() ([]byte, error) {
	return basicEnumImplOsType.ToEnumJsonBytes(it.Value()), nil
}

func (it *Variation) UnmarshalJSON(data []byte) error {
	valueByte, err := basicEnumImplOsType.UnmarshallToValue(
		true,
		data)

	if err == nil {
		*it = Variation(valueByte)
	}

	return err
}

func (it Variation) Name() string {
	return basicEnumImplOsType.ToEnumString(it.Value())
}

func (it Variation) GoosName() string {
	return basicEnumImplOsType.ToEnumString(it.Value())
}

func (it Variation) NameValue() string {
	return basicEnumImplOsType.NameWithValue(it.Value())
}

func (it Variation) ToNumberString() string {
	return basicEnumImplOsType.ToNumberString(it.Value())
}

func (it Variation) RangeNamesCsv() string {
	return basicEnumImplOsType.RangeNamesCsv()
}

func (it Variation) TypeName() string {
	return basicEnumImplOsType.TypeName()
}

func (it Variation) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
	return basicEnumImplOsType.UnmarshallToValue(true, jsonUnmarshallingValue)
}

func (it Variation) MaxByte() byte {
	return basicEnumImplOsType.Max()
}

func (it Variation) MinByte() byte {
	return basicEnumImplOsType.Min()
}

func (it Variation) ValueByte() byte {
	return byte(it)
}

func (it Variation) RangesByte() []byte {
	return basicEnumImplOsType.Ranges()
}

func (it Variation) Value() byte {
	return byte(it)
}

func (it Variation) String() string {
	return basicEnumImplOsType.ToEnumString(it.Value())
}

func (it *Variation) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return it
}

func (it *Variation) AsJsonContractsBinder() corejson.JsonMarshaller {
	return it
}

func (it Variation) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return &it
}
