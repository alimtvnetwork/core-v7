package reqtype

import (
	"errors"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/msgtype"
)

type Request byte

// https://www.restapitutorial.com/lessons/httpmethods.html
const (
	Uninitialized Request = iota
	Create
	Read
	Update
	Delete
	Drop
	CreateOrUpdate
	ExistCheck
	SkipOnExist
	CreateOrSkipOnExist
	UpdateOrSkipOnNonExist
	DeleteOrSkipOnNonExist
	DropOrSkipOnNonExist
	UpdateOnExist
	DropOnExist
	DropCreate
	Append
	AppendByCompare
	AppendByCompareWhereCommentFound
	AppendLinesByCompare
	AppendLines
	CreateOrAppend
	Prepend
	CreateOrPrepend
	PrependLines
	Rename
	Change
	Merge
	MergeLines
	GetHttp
	PutHttp
	PostHttp
	DeleteHttp
	PatchHttp
	Touch
	Start
	Stop
	Restart
	Reload
	StopSleepStart
	Suspend
	Pause
	Resumed
	TryRestart3Times
	TryRestart5Times
	TryStart3Times
	TryStart5Times
	TryStop3Times
	TryStop5Times
	InheritOnly
	InheritPlusOverride
	DynamicAction
)

func (receiver Request) IsUninitialized() bool {
	return receiver == Uninitialized
}

func (receiver Request) IsCreate() bool {
	return receiver == Create
}

func (receiver Request) IsRead() bool {
	return receiver == Read
}

func (receiver Request) IsUpdate() bool {
	return receiver == Update
}

func (receiver Request) IsDelete() bool {
	return receiver == Delete
}

func (receiver Request) IsDrop() bool {
	return receiver == Drop
}

func (receiver Request) IsCreateOrUpdate() bool {
	return receiver == CreateOrUpdate
}

func (receiver Request) IsExistCheck() bool {
	return receiver == ExistCheck
}

func (receiver Request) IsSkipOnExist() bool {
	return receiver == SkipOnExist
}

func (receiver Request) IsCreateOrSkipOnExist() bool {
	return receiver == CreateOrSkipOnExist
}

func (receiver Request) IsUpdateOrSkipOnNonExist() bool {
	return receiver == UpdateOrSkipOnNonExist
}

func (receiver Request) IsDeleteOrSkipOnNonExist() bool {
	return receiver == DeleteOrSkipOnNonExist
}

func (receiver Request) IsDropOrSkipOnNonExist() bool {
	return receiver == DropOrSkipOnNonExist
}

func (receiver Request) IsUpdateOnExist() bool {
	return receiver == UpdateOnExist
}

func (receiver Request) IsDropOnExist() bool {
	return receiver == DropOnExist
}

func (receiver Request) IsDropCreate() bool {
	return receiver == DropCreate
}

func (receiver Request) IsAppend() bool {
	return receiver == Append
}

func (receiver Request) IsAppendByCompare() bool {
	return receiver == AppendByCompare
}

func (receiver Request) IsAppendByCompareWhereCommentFound() bool {
	return receiver == AppendByCompareWhereCommentFound
}

func (receiver Request) IsAppendLinesByCompare() bool {
	return receiver == AppendLinesByCompare
}

func (receiver Request) IsAppendLines() bool {
	return receiver == AppendLines
}

func (receiver Request) IsCreateOrAppend() bool {
	return receiver == CreateOrAppend
}

func (receiver Request) IsPrepend() bool {
	return receiver == Prepend
}

func (receiver Request) IsCreateOrPrepend() bool {
	return receiver == CreateOrPrepend
}

func (receiver Request) IsPrependLines() bool {
	return receiver == PrependLines
}

func (receiver Request) IsRename() bool {
	return receiver == Rename
}

func (receiver Request) IsChange() bool {
	return receiver == Change
}

func (receiver Request) IsMerge() bool {
	return receiver == Merge
}

func (receiver Request) IsMergeLines() bool {
	return receiver == MergeLines
}

func (receiver Request) IsGetHttp() bool {
	return receiver == GetHttp
}

func (receiver Request) IsPutHttp() bool {
	return receiver == PutHttp
}

func (receiver Request) IsPostHttp() bool {
	return receiver == PostHttp
}

func (receiver Request) IsDeleteHttp() bool {
	return receiver == DeleteHttp
}

func (receiver Request) IsPatchHttp() bool {
	return receiver == PatchHttp
}

func (receiver Request) IsTouch() bool {
	return receiver == Touch
}

func (receiver Request) IsStart() bool {
	return receiver == Start
}

func (receiver Request) IsStop() bool {
	return receiver == Stop
}

func (receiver Request) IsRestart() bool {
	return receiver == Restart
}

func (receiver Request) IsReload() bool {
	return receiver == Reload
}

func (receiver Request) IsStopSleepStart() bool {
	return receiver == StopSleepStart
}

func (receiver Request) IsSuspend() bool {
	return receiver == Suspend
}

func (receiver Request) IsPause() bool {
	return receiver == Pause
}

func (receiver Request) IsResumed() bool {
	return receiver == Resumed
}

func (receiver Request) IsTryRestart3Times() bool {
	return receiver == TryRestart3Times
}

func (receiver Request) IsTryRestart5Times() bool {
	return receiver == TryRestart5Times
}

func (receiver Request) IsTryStart3Times() bool {
	return receiver == TryStart3Times
}

func (receiver Request) IsTryStart5Times() bool {
	return receiver == TryStart5Times
}

func (receiver Request) IsTryStop3Times() bool {
	return receiver == TryStop3Times
}

func (receiver Request) IsTryStop5Times() bool {
	return receiver == TryStop5Times
}

func (receiver Request) IsInheritOnly() bool {
	return receiver == InheritOnly
}

func (receiver Request) IsInheritPlusOverride() bool {
	return receiver == InheritPlusOverride
}

// IsRestartOrReload  receiver. IsRestart() || receiver. IsReload()
func (receiver Request) IsRestartOrReload() bool {
	return receiver.IsRestart() || receiver.IsReload()
}

// IsCrud returns true if Read, Update, Create, Delete, IsCreateOrUpdate
func (receiver Request) IsCrud() bool {
	return receiver.IsRead() ||
		receiver.IsCreate() ||
		receiver.IsCreateOrUpdate() ||
		receiver.IsUpdate() ||
		receiver.IsDelete()
}

// IsCrudSkip
//
// returns true if
// IsCreateOrSkipOnExist, IsUpdateOrSkipOnNonExist, IsDeleteOrSkipOnNonExist,
// IsDropOnExist, IsDropOrSkipOnNonExist,
func (receiver Request) IsCrudSkip() bool {
	return receiver.IsCreateOrSkipOnExist() ||
		receiver.IsUpdateOrSkipOnNonExist() ||
		receiver.IsDeleteOrSkipOnNonExist() ||
		receiver.IsDropOnExist() ||
		receiver.IsDropOrSkipOnNonExist()
}

// IsCrudOrSkip
//
// returns true if
// IsCrud || IsCrudSkip
func (receiver Request) IsCrudOrSkip() bool {
	return receiver.IsCrud() ||
		receiver.IsCrudSkip()
}

// IsAnyDrop
//
// returns true if
// IsDrop, IsDelete, IsDeleteOrSkipOnNonExist,
// IsDropOnExist, IsDropCreate, IsDropOrSkipOnNonExist
func (receiver Request) IsAnyDrop() bool {
	return receiver.IsDrop() ||
		receiver.IsDelete() ||
		receiver.IsDeleteOrSkipOnNonExist() ||
		receiver.IsDropOnExist() ||
		receiver.IsDropCreate() ||
		receiver.IsDropOrSkipOnNonExist()
}

// IsDropSafe
//
// returns true if
// IsDeleteOrSkipOnNonExist, IsDropOnExist, IsDropCreate,
// IsDropOrSkipOnNonExist
func (receiver Request) IsDropSafe() bool {
	return receiver.IsDeleteOrSkipOnNonExist() ||
		receiver.IsDropOnExist() ||
		receiver.IsDropCreate() ||
		receiver.IsDropOrSkipOnNonExist()
}

// IsAnyCreate
//
// returns true if
// IsCreate, IsCreateOrUpdate, IsCreateOrAppend,
// IsCreateOrPrepend, IsCreateOrSkipOnExist, IsDropCreate
func (receiver Request) IsAnyCreate() bool {
	return receiver.IsCreate() ||
		receiver.IsCreateOrUpdate() ||
		receiver.IsCreateOrAppend() ||
		receiver.IsCreateOrPrepend() ||
		receiver.IsCreateOrSkipOnExist() ||
		receiver.IsDropCreate()
}

// IsHttp
//
// returns true if
// IsGetHttp, IsPostHttp, IsPutHttp,
// IsDeleteHttp, IsPatchHttp
func (receiver Request) IsHttp() bool {
	return receiver.IsGetHttp() ||
		receiver.IsPostHttp() ||
		receiver.IsPutHttp() ||
		receiver.IsDeleteHttp() ||
		receiver.IsPatchHttp()
}

func (receiver Request) Name() string {
	return BasicEnumImpl.ToEnumString(receiver.Value())
}

func (receiver Request) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(receiver.Value())
}

func (receiver Request) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return BasicEnumImpl.UnmarshallEnumToValue(jsonUnmarshallingValue)
}

func (receiver Request) IsValidRange() bool {
	return BasicEnumImpl.IsValidRange(receiver.Value())
}

// IsInBetween edge case including the start, end
func (receiver Request) IsInBetween(
	start, end Request,
) bool {
	val := receiver.Value()

	return val >= start.Value() && val <= end.Value()
}

func (receiver Request) CurrentNotImpl(
	reference interface{},
	messages ...string,
) error {
	compiledMessage := strings.Join(messages, constants.Space)
	fullCompiled := receiver.String() +
		" : is not implemented. " +
		compiledMessage

	if reference == nil {
		return msgtype.NotImplemented.ErrorNoRefs(fullCompiled)
	}

	return msgtype.NotImplemented.Error(fullCompiled, reference)
}

// IsAnyOfReqs returns true if current one is matching with any of it
func (receiver Request) IsAnyOfReqs(reqs ...Request) bool {
	if len(reqs) == 0 {
		return true
	}

	for _, req := range reqs {
		if req == receiver {
			return true
		}
	}

	return false
}

// GetStatusAnyOf returns status success true if current one is any of the given values.
func (receiver Request) GetStatusAnyOf(reqs ...Request) *ResultStatus {
	if len(reqs) == 0 {
		return &ResultStatus{
			IsSuccess:  true,
			IndexMatch: constants.InvalidNotFoundCase,
			Ranges:     reqs,
			Error:      nil,
		}
	}

	for i, req := range reqs {
		if req == receiver {
			return &ResultStatus{
				IsSuccess:  true,
				IndexMatch: i,
				Ranges:     reqs,
				Error:      nil,
			}
		}
	}

	errMsg := msgtype.RangeNotMeet(
		"Failed GetStatusAnyOf",
		start(&reqs),
		end(&reqs),
		reqs)

	return &ResultStatus{
		IsSuccess:  true,
		IndexMatch: constants.InvalidNotFoundCase,
		Ranges:     reqs,
		Error:      errors.New(errMsg),
	}
}

// GetInBetweenStatus edge case including the start, end
func (receiver Request) GetInBetweenStatus(start, end Request) *ResultStatus {
	isInBetween := receiver.IsInBetween(start, end)
	ranges := RangesInBetween(start, end)

	if isInBetween {
		return &ResultStatus{
			IsSuccess:  isInBetween,
			IndexMatch: receiver.ValueInt(),
			Ranges:     ranges,
			Error:      nil,
		}
	}

	errMsg := msgtype.RangeNotMeet(
		"Failed GetInBetweenStatus",
		start,
		end,
		ranges)

	return &ResultStatus{
		IsSuccess:  false,
		IndexMatch: constants.InvalidNotFoundCase,
		Ranges:     ranges,
		Error:      errors.New(errMsg),
	}
}

func (receiver Request) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (receiver Request) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (receiver Request) ValueByte() byte {
	return receiver.Value()
}

func (receiver Request) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (receiver Request) Value() byte {
	return byte(receiver)
}

func (receiver Request) ValueInt() int {
	return int(receiver)
}

func (receiver Request) IsAnyOf(checkingItems ...byte) bool {
	return BasicEnumImpl.IsAnyOf(receiver.Value(), checkingItems...)
}

func (receiver Request) String() string {
	return BasicEnumImpl.ToEnumString(receiver.Value())
}

func (receiver *Request) UnmarshalJSON(data []byte) error {
	dataConv, err := BasicEnumImpl.UnmarshallEnumToValue(data)

	if err == nil {
		*receiver = Request(dataConv)
	}

	return err
}

func (receiver Request) ToPtr() *Request {
	return &receiver
}

func (receiver *Request) ToSimple() Request {
	if receiver == nil {
		return Uninitialized
	}

	return *receiver
}

func (receiver Request) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(receiver.Value()), nil
}

func (receiver Request) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return &receiver
}

func (receiver *Request) AsJsonMarshaller() corejson.JsonMarshaller {
	return receiver
}

func (receiver *Request) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return receiver
}
