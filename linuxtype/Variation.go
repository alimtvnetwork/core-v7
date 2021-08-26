package linuxtype

import (
	"gitlab.com/evatix-go/core/coreinterface"
)

type Variation byte

const (
	Unknown Variation = iota
	UbuntuServer
	UbuntuServer18
	UbuntuServer19
	UbuntuServer20
	UbuntuServer21
	UbuntuDesktop
	Centos
	Centos7
	Centos8
	Centos9
	DebianServer
	DebianDesktop
	Docker
	DockerUbuntuServer
	DockerUbuntuServer20
	DockerUbuntuServer21
	DockerCentos9
	Android
)

func (it Variation) Value() byte {
	return byte(it)
}

func (it Variation) Name() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it Variation) IsUnknown() bool {
	return it == Unknown
}

func (it Variation) IsUbuntuServer() bool {
	return it == UbuntuServer
}

func (it Variation) IsUbuntuServer18() bool {
	return it == UbuntuServer18
}

func (it Variation) IsUbuntuServer19() bool {
	return it == UbuntuServer19
}

func (it Variation) IsUbuntuServer20() bool {
	return it == UbuntuServer20
}

func (it Variation) IsUbuntuServer21() bool {
	return it == UbuntuServer21
}

func (it Variation) IsUbuntuDesktop() bool {
	return it == UbuntuDesktop
}

func (it Variation) IsCentos() bool {
	return it == Centos
}

func (it Variation) IsCentos7() bool {
	return it == Centos7
}

func (it Variation) IsCentos8() bool {
	return it == Centos8
}

func (it Variation) IsCentos9() bool {
	return it == Centos9
}

func (it Variation) IsDebianServer() bool {
	return it == DebianServer
}

func (it Variation) IsDebianDesktop() bool {
	return it == DebianDesktop
}

func (it Variation) IsDocker() bool {
	return it == Docker
}

func (it Variation) IsDockerUbuntuServer() bool {
	return it == DockerUbuntuServer
}

func (it Variation) IsDockerUbuntuServer20() bool {
	return it == DockerUbuntuServer20
}

func (it Variation) IsDockerUbuntuServer21() bool {
	return it == DockerUbuntuServer20
}

func (it Variation) IsDockerCentos9() bool {
	return it == DockerCentos9
}

func (it Variation) IsAndroid() bool {
	return it == Android
}

func (it Variation) String() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it Variation) IsAnyOf(checkingItems ...byte) bool {
	return BasicEnumImpl.IsAnyOf(it.Value(), checkingItems...)
}

func (it Variation) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.Value()), nil
}

func (it Variation) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.Value())
}

func (it *Variation) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return BasicEnumImpl.
		UnmarshallToValue(true, jsonUnmarshallingValue)
}

func (it Variation) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variation) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it *Variation) UnmarshalJSON(data []byte) error {
	dataConv, err := it.UnmarshallEnumToValue(data)

	if err != nil {
		return err
	}

	*it = Variation(dataConv)

	return nil
}

func (it Variation) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return &it
}

func (it Variation) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it Variation) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it Variation) ValueByte() byte {
	return it.Value()
}

func (it Variation) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it Variation) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return &it
}
