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

func (receiver Variation) Value() byte {
	return byte(receiver)
}

func (receiver Variation) Name() string {
	return BasicEnumImpl.ToEnumString(receiver.Value())
}

func (receiver Variation) IsUnknown() bool {
	return receiver == Unknown
}

func (receiver Variation) IsUbuntuServer() bool {
	return receiver == UbuntuServer
}

func (receiver Variation) IsUbuntuServer18() bool {
	return receiver == UbuntuServer18
}

func (receiver Variation) IsUbuntuServer19() bool {
	return receiver == UbuntuServer19
}

func (receiver Variation) IsUbuntuServer20() bool {
	return receiver == UbuntuServer20
}

func (receiver Variation) IsUbuntuServer21() bool {
	return receiver == UbuntuServer21
}

func (receiver Variation) IsUbuntuDesktop() bool {
	return receiver == UbuntuDesktop
}

func (receiver Variation) IsCentos() bool {
	return receiver == Centos
}

func (receiver Variation) IsCentos7() bool {
	return receiver == Centos7
}

func (receiver Variation) IsCentos8() bool {
	return receiver == Centos8
}

func (receiver Variation) IsCentos9() bool {
	return receiver == Centos9
}

func (receiver Variation) IsDebianServer() bool {
	return receiver == DebianServer
}

func (receiver Variation) IsDebianDesktop() bool {
	return receiver == DebianDesktop
}

func (receiver Variation) IsDocker() bool {
	return receiver == Docker
}

func (receiver Variation) IsDockerUbuntuServer() bool {
	return receiver == DockerUbuntuServer
}

func (receiver Variation) IsDockerUbuntuServer20() bool {
	return receiver == DockerUbuntuServer20
}

func (receiver Variation) IsDockerUbuntuServer21() bool {
	return receiver == DockerUbuntuServer20
}

func (receiver Variation) IsDockerCentos9() bool {
	return receiver == DockerCentos9
}

func (receiver Variation) IsAndroid() bool {
	return receiver == Android
}

func (receiver Variation) String() string {
	return BasicEnumImpl.ToEnumString(receiver.Value())
}

func (receiver Variation) IsAnyOf(checkingItems ...byte) bool {
	return BasicEnumImpl.IsAnyOf(receiver.Value(), checkingItems...)
}

func (receiver Variation) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(receiver.Value()), nil
}

func (receiver Variation) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(receiver.Value())
}

func (receiver *Variation) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return BasicEnumImpl.
		UnmarshallEnumToValue(jsonUnmarshallingValue)
}

func (receiver *Variation) UnmarshalJSON(data []byte) error {
	dataConv, err := BasicEnumImpl.UnmarshallEnumToValue(data)

	if err != nil {
		return err
	}

	*receiver = Variation(dataConv)

	return nil
}

func (receiver Variation) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return &receiver
}

func (receiver Variation) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (receiver Variation) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (receiver Variation) ValueByte() byte {
	return receiver.Value()
}

func (receiver Variation) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (receiver Variation) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return &receiver
}
