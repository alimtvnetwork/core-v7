package chmodclasstype

import "gitlab.com/evatix-go/core/coreinterface"

type Variant byte

const (
	UnInitialized Variant = iota
	All
	Owner
	Group
	Other
	OwnerGroup
	GroupOther
	OwnerOther
)

func (receiver Variant) IsUnInitialized() bool {
	return receiver == UnInitialized
}

func (receiver Variant) IsAll() bool {
	return receiver == All
}

func (receiver Variant) IsOwner() bool {
	return receiver == Owner
}

func (receiver Variant) IsGroup() bool {
	return receiver == Group
}

func (receiver Variant) IsOther() bool {
	return receiver == Other
}

func (receiver Variant) IsOwnerGroup() bool {
	return receiver == OwnerGroup
}

func (receiver Variant) IsGroupOther() bool {
	return receiver == GroupOther
}

func (receiver Variant) IsOwnerOther() bool {
	return receiver == OwnerOther
}

func (receiver *Variant) Name() string {
	return BasicEnumImpl.ToEnumString(receiver.ValueByte())
}

func (receiver *Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(receiver.ValueByte())
}

func (receiver *Variant) String() string {
	return BasicEnumImpl.ToEnumString(receiver.ValueByte())
}

func (receiver *Variant) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
	return BasicEnumImpl.UnmarshallEnumToValue(jsonUnmarshallingValue)
}

func (receiver *Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(receiver.Value()), nil
}

func (receiver *Variant) UnmarshalJSON(data []byte) error {
	dataConv, err := BasicEnumImpl.UnmarshallEnumToValue(data)

	if err == nil {
		*receiver = Variant(dataConv)
	}

	return err
}

func (receiver *Variant) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return receiver
}

func (receiver *Variant) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (receiver *Variant) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (receiver Variant) ValueByte() byte {
	return byte(receiver)
}

func (receiver Variant) Value() byte {
	return byte(receiver)
}

func (receiver *Variant) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (receiver *Variant) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return receiver
}
