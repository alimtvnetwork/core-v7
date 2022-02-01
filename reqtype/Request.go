package reqtype

import (
	"errors"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/errcore"
)

type Request byte

// https://www.restapitutorial.com/lessons/httpmethods.html
const (
	Invalid Request = iota
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

func (it Request) IsValid() bool {
	return it != Invalid
}

func (it Request) IsInvalid() bool {
	return it == Invalid
}

func (it Request) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it Request) IsUninitialized() bool {
	return it == Invalid
}

func (it Request) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Request) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it Request) IsCreate() bool {
	return it == Create
}

func (it Request) IsRead() bool {
	return it == Read
}

func (it Request) IsUpdate() bool {
	return it == Update
}

func (it Request) IsDelete() bool {
	return it == Delete
}

func (it Request) IsDrop() bool {
	return it == Drop
}

func (it Request) IsCreateOrUpdate() bool {
	return it == CreateOrUpdate
}

func (it Request) IsExistCheck() bool {
	return it == ExistCheck
}

func (it Request) IsSkipOnExist() bool {
	return it == SkipOnExist
}

func (it Request) IsCreateOrSkipOnExist() bool {
	return it == CreateOrSkipOnExist
}

func (it Request) IsUpdateOrSkipOnNonExist() bool {
	return it == UpdateOrSkipOnNonExist
}

func (it Request) IsDeleteOrSkipOnNonExist() bool {
	return it == DeleteOrSkipOnNonExist
}

func (it Request) IsDropOrSkipOnNonExist() bool {
	return it == DropOrSkipOnNonExist
}

func (it Request) IsUpdateOnExist() bool {
	return it == UpdateOnExist
}

func (it Request) IsDropOnExist() bool {
	return it == DropOnExist
}

func (it Request) IsDropCreate() bool {
	return it == DropCreate
}

func (it Request) IsAppend() bool {
	return it == Append
}

func (it Request) IsAppendByCompare() bool {
	return it == AppendByCompare
}

func (it Request) IsAppendByCompareWhereCommentFound() bool {
	return it == AppendByCompareWhereCommentFound
}

func (it Request) IsAppendLinesByCompare() bool {
	return it == AppendLinesByCompare
}

func (it Request) IsAppendLines() bool {
	return it == AppendLines
}

func (it Request) IsCreateOrAppend() bool {
	return it == CreateOrAppend
}

func (it Request) IsPrepend() bool {
	return it == Prepend
}

func (it Request) IsCreateOrPrepend() bool {
	return it == CreateOrPrepend
}

func (it Request) IsPrependLines() bool {
	return it == PrependLines
}

func (it Request) IsRename() bool {
	return it == Rename
}

func (it Request) IsChange() bool {
	return it == Change
}

func (it Request) IsMerge() bool {
	return it == Merge
}

func (it Request) IsMergeLines() bool {
	return it == MergeLines
}

func (it Request) IsGetHttp() bool {
	return it == GetHttp
}

func (it Request) IsPutHttp() bool {
	return it == PutHttp
}

func (it Request) IsPostHttp() bool {
	return it == PostHttp
}

func (it Request) IsDeleteHttp() bool {
	return it == DeleteHttp
}

func (it Request) IsPatchHttp() bool {
	return it == PatchHttp
}

func (it Request) IsTouch() bool {
	return it == Touch
}

func (it Request) IsStart() bool {
	return it == Start
}

func (it Request) IsStop() bool {
	return it == Stop
}

func (it Request) IsRestart() bool {
	return it == Restart
}

func (it Request) IsReload() bool {
	return it == Reload
}

func (it Request) IsStopSleepStart() bool {
	return it == StopSleepStart
}

func (it Request) IsSuspend() bool {
	return it == Suspend
}

func (it Request) IsPause() bool {
	return it == Pause
}

func (it Request) IsResumed() bool {
	return it == Resumed
}

func (it Request) IsTryRestart3Times() bool {
	return it == TryRestart3Times
}

func (it Request) IsTryRestart5Times() bool {
	return it == TryRestart5Times
}

func (it Request) IsTryStart3Times() bool {
	return it == TryStart3Times
}

func (it Request) IsTryStart5Times() bool {
	return it == TryStart5Times
}

func (it Request) IsTryStop3Times() bool {
	return it == TryStop3Times
}

func (it Request) IsTryStop5Times() bool {
	return it == TryStop5Times
}

func (it Request) IsInheritOnly() bool {
	return it == InheritOnly
}

func (it Request) IsInheritPlusOverride() bool {
	return it == InheritPlusOverride
}

// IsRestartOrReload  receiver. IsRestart() || receiver. IsReload()
func (it Request) IsRestartOrReload() bool {
	return it.IsRestart() || it.IsReload()
}

// IsAnySkipOnExist =>
// IsSkipOnExist, IsCreateOrSkipOnExist,
// IsUpdateOrSkipOnNonExist, IsDeleteOrSkipOnNonExist,
// IsDeleteOrSkipOnNonExist, IsDropOrSkipOnNonExist
func (it Request) IsAnySkipOnExist() bool {
	return it.IsSkipOnExist() ||
		it.IsCreateOrSkipOnExist() ||
		it.IsUpdateOrSkipOnNonExist() ||
		it.IsDeleteOrSkipOnNonExist() ||
		it.IsDropOrSkipOnNonExist()
}

// IsAnyApplyOnExist =>
// IsUpdateOnExist, IsDropOnExist,
func (it Request) IsAnyApplyOnExist() bool {
	return it.IsUpdateOnExist() ||
		it.IsDropOnExist()
}

// IsCrud returns true if Read, Update, Create, Delete, IsCreateOrUpdate
func (it Request) IsCrud() bool {
	return it.IsRead() ||
		it.IsCreate() ||
		it.IsCreateOrUpdate() ||
		it.IsUpdate() ||
		it.IsDelete()
}

// IsCrudSkip
//
// returns true if
// IsCreateOrSkipOnExist, IsUpdateOrSkipOnNonExist, IsDeleteOrSkipOnNonExist,
// IsDropOnExist, IsDropOrSkipOnNonExist,
func (it Request) IsCrudSkip() bool {
	return it.IsCreateOrSkipOnExist() ||
		it.IsUpdateOrSkipOnNonExist() ||
		it.IsDeleteOrSkipOnNonExist() ||
		it.IsDropOnExist() ||
		it.IsDropOrSkipOnNonExist()
}

// IsCrudOrSkip
//
// returns true if
// IsCrud || IsCrudSkip
func (it Request) IsCrudOrSkip() bool {
	return it.IsCrud() ||
		it.IsCrudSkip()
}

// IsAnyDrop
//
// returns true if
// IsDrop, IsDelete, IsDeleteOrSkipOnNonExist,
// IsDropOnExist, IsDropCreate, IsDropOrSkipOnNonExist
func (it Request) IsAnyDrop() bool {
	return it.IsDrop() ||
		it.IsDelete() ||
		it.IsDeleteOrSkipOnNonExist() ||
		it.IsDropOnExist() ||
		it.IsDropCreate() ||
		it.IsDropOrSkipOnNonExist()
}

// IsDropSafe
//
// returns true if
// IsDeleteOrSkipOnNonExist, IsDropOnExist,
// IsDropOrSkipOnNonExist
func (it Request) IsDropSafe() bool {
	return it.IsDeleteOrSkipOnNonExist() ||
		it.IsDropOnExist() ||
		it.IsDropOrSkipOnNonExist()
}

// IsAnyCreate
//
// returns true if
// IsCreate, IsCreateOrUpdate, IsCreateOrAppend,
// IsCreateOrPrepend, IsCreateOrSkipOnExist, IsDropCreate
func (it Request) IsAnyCreate() bool {
	return it.IsCreate() ||
		it.IsCreateOrUpdate() ||
		it.IsCreateOrAppend() ||
		it.IsCreateOrPrepend() ||
		it.IsCreateOrSkipOnExist() ||
		it.IsDropCreate()
}

// IsHttp
//
// returns true if
// IsGetHttp, IsPostHttp, IsPutHttp,
// IsDeleteHttp, IsPatchHttp
func (it Request) IsHttp() bool {
	return it.IsGetHttp() ||
		it.IsPostHttp() ||
		it.IsPutHttp() ||
		it.IsDeleteHttp() ||
		it.IsPatchHttp()
}

func (it Request) Name() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it Request) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.Value())
}

func (it Request) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(true, jsonUnmarshallingValue)
}

func (it Request) IsValidRange() bool {
	return BasicEnumImpl.IsValidRange(it.Value())
}

// IsInBetween edge case including the start, end
func (it Request) IsInBetween(
	start, end Request,
) bool {
	val := it.Value()

	return val >= start.Value() && val <= end.Value()
}

func (it Request) CurrentNotImpl(
	reference interface{},
	messages ...string,
) error {
	compiledMessage := strings.Join(messages, constants.Space)
	fullCompiled := it.String() +
		" : is not implemented. " +
		compiledMessage

	if reference == nil {
		return errcore.NotImplementedType.ErrorNoRefs(fullCompiled)
	}

	return errcore.NotImplementedType.Error(fullCompiled, reference)
}

func (it Request) NotSupportedErr(
	message string,
	reference interface{},
) error {
	return errcore.NotSupportedType.Error(
		message,
		reference)
}

// IsNotAnyOfReqs returns true only if none of these matches
func (it Request) IsNotAnyOfReqs(reqs ...Request) bool {
	if len(reqs) == 0 {
		return true
	}

	for _, req := range reqs {
		if req == it {
			return false
		}
	}

	return true
}

// IsAnyOfReqs returns true if current one is matching with any of it
func (it Request) IsAnyOfReqs(reqs ...Request) bool {
	if len(reqs) == 0 {
		return true
	}

	for _, req := range reqs {
		if req == it {
			return true
		}
	}

	return false
}

// GetStatusAnyOf returns status success true if current one is any of the given values.
func (it Request) GetStatusAnyOf(reqs ...Request) *ResultStatus {
	if len(reqs) == 0 {
		return &ResultStatus{
			IsSuccess:  true,
			IndexMatch: constants.InvalidNotFoundCase,
			Ranges:     reqs,
			Error:      nil,
		}
	}

	for i, req := range reqs {
		if req == it {
			return &ResultStatus{
				IsSuccess:  true,
				IndexMatch: i,
				Ranges:     reqs,
				Error:      nil,
			}
		}
	}

	errMsg := errcore.RangeNotMeet(
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
func (it Request) GetInBetweenStatus(start, end Request) *ResultStatus {
	isInBetween := it.IsInBetween(start, end)
	ranges := RangesInBetween(start, end)

	if isInBetween {
		return &ResultStatus{
			IsSuccess:  isInBetween,
			IndexMatch: it.ValueInt(),
			Ranges:     ranges,
			Error:      nil,
		}
	}

	errMsg := errcore.RangeNotMeet(
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

func (it Request) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it Request) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it Request) ValueByte() byte {
	return it.Value()
}

func (it Request) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it Request) Value() byte {
	return byte(it)
}

func (it Request) ValueInt() int {
	return int(it)
}

func (it Request) IsAnyOf(checkingItems ...byte) bool {
	return BasicEnumImpl.IsAnyOf(it.Value(), checkingItems...)
}

func (it Request) String() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it *Request) UnmarshalJSON(data []byte) error {
	dataConv, err := it.UnmarshallEnumToValue(data)

	if err == nil {
		*it = Request(dataConv)
	}

	return err
}

func (it Request) ToPtr() *Request {
	return &it
}

func (it *Request) ToSimple() Request {
	if it == nil {
		return Invalid
	}

	return *it
}

func (it Request) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.Value()), nil
}

func (it Request) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return &it
}

func (it *Request) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it Request) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return &it
}
