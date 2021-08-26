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

func (it Variant) IsUnInitialized() bool {
	return it == UnInitialized
}

func (it Variant) IsAll() bool {
	return it == All
}

func (it Variant) IsOwner() bool {
	return it == Owner
}

func (it Variant) IsGroup() bool {
	return it == Group
}

func (it Variant) IsOther() bool {
	return it == Other
}

func (it Variant) IsOwnerGroup() bool {
	return it == OwnerGroup
}

func (it Variant) IsGroupOther() bool {
	return it == GroupOther
}

func (it Variant) IsOwnerOther() bool {
	return it == OwnerOther
}

func (it *Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it *Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.ValueByte())
}

func (it *Variant) String() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it *Variant) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(true, jsonUnmarshallingValue)
}

func (it *Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.Value()), nil
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	dataConv, err := it.UnmarshallEnumToValue(data)

	if err == nil {
		*it = Variant(dataConv)
	}

	return err
}

func (it Variant) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
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

func (it Variant) ValueByte() byte {
	return byte(it)
}

func (it Variant) Value() byte {
	return byte(it)
}

func (it *Variant) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it *Variant) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return it
}
