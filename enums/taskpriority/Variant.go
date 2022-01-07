package taskpriority

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
)

type Variant byte

const (
	Default       Variant = iota // Priority : 40%
	DefaultLock                  // Priority : 20%, this should lock the task https://github.com/hibiken/asynq/wiki/Unique-Tasks
	Reminder                     // Priority : 10%
	Notification                 // Priority : 10%
	SystemUpdate                 // Priority : 10%
	LowerPriority                // Priority : 10%
	Invalid
)

func (it Variant) IsDefault() bool {
	return it == Default
}

func (it Variant) IsDefaultLock() bool {
	return it == DefaultLock
}

func (it Variant) IsReminder() bool {
	return it == Reminder
}

func (it Variant) IsNotification() bool {
	return it == Notification
}

func (it Variant) IsSystemUpdate() bool {
	return it == SystemUpdate
}

func (it Variant) IsLowerPriority() bool {
	return it == LowerPriority
}

func (it Variant) IsLockEnforced() bool {
	return lockEnforcedMap[it]
}

func (it Variant) IsInvalid() bool {
	return it == Invalid
}

func (it Variant) IsValid() bool {
	return it != Invalid
}

func (it Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Variant) IsNameEqual(name string) bool {
	return BasicEnumImpl.ToEnumString(it.ValueByte()) == name
}

func (it Variant) IsNameNotEqual(name string) bool {
	return BasicEnumImpl.ToEnumString(it.ValueByte()) != name
}

func (it Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.ValueByte())
}

func (it Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.ValueByte()), nil
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	dataConv, err := it.UnmarshallEnumToValue(
		data)

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

func (it Variant) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(
		true,
		jsonUnmarshallingValue)
}

func (it Variant) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it Variant) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it Variant) ValueByte() byte {
	return byte(it)
}

func (it Variant) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it Variant) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it Variant) String() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it *Variant) JsonParseSelfInject(jsonResult *corejson.Result) error {
	err := jsonResult.Unmarshal(it)

	return err
}

func (it Variant) Json() corejson.Result {
	return corejson.New(it)
}

func (it Variant) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Variant) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Variant) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Variant) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it *Variant) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return it
}

func (it *Variant) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return it
}
