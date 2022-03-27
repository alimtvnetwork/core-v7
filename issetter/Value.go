package issetter

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl/enumtype"
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/internal/csvinternal"
)

// Value
//
//  Used evaluate lazy boolean valuesNames.
//
// Values:
//  - Uninitialized Value = 0
//  - True          Value = 1
//  - False         Value = 2
//  - Unset         Value = 3
//  - Set           Value = 4
//  - Wildcard      Value = 5
type Value byte

const (
	Uninitialized Value = 0
	True          Value = 1
	False         Value = 2
	Unset         Value = 3
	Set           Value = 4
	Wildcard      Value = 5
)

func (it Value) AllNameValues() []string {
	slice := make([]string, len(valuesNames))

	for i := range valuesNames {
		slice[i] = Value(i).NameValue()
	}

	return slice
}

func (it Value) OnlySupportedErr(names ...string) error {
	if len(names) == 0 {
		return nil
	}

	hashset := toHashset(names...)
	var unsupportedNames []string

	for _, name := range valuesNames {
		_, has := hashset[name]

		if !has {
			unsupportedNames = append(unsupportedNames, name)
		}
	}

	if len(unsupportedNames) > 0 {
		return errors.New(csvinternal.StringsToStringDefault(unsupportedNames...) + " not supported")
	}

	return nil

}

func (it Value) OnlySupportedMsgErr(message string, names ...string) error {
	err := it.OnlySupportedErr(names...)

	if err == nil {
		return nil
	}

	return errors.New(message + err.Error())
}

func (it Value) ValueUInt16() uint16 {
	return uint16(it)
}

func (it Value) IntegerEnumRanges() []int {
	return integerRanges
}

func (it Value) MinMaxAny() (min, max interface{}) {
	return Min(), Max()
}

func (it Value) MinValueString() string {
	return Min().StringValue()
}

func (it Value) MaxValueString() string {
	return Max().StringValue()
}

func (it Value) MaxInt() int {
	return Max().ValueInt()
}

func (it Value) MinInt() int {
	return Min().ValueInt()
}

func (it Value) RangesDynamicMap() map[string]interface{} {
	return dynamicRangesMap
}

func (it Value) IsValueEqual(value byte) bool {
	return byte(it) == value
}

func (it Value) RangeNamesCsv() string {
	return RangeNamesCsv()
}

func (it Value) IsByteValueEqual(value byte) bool {
	return byte(it) == value
}

func (it Value) IsOn() bool {
	return trueMap[it]
}

func (it Value) IsOff() bool {
	return falseMap[it]
}

func (it Value) IsLater() bool {
	return it.IsUndefinedLogically()
}

// IsNo
//
//  Returns true if False or Unset
func (it Value) IsNo() bool {
	return falseMap[it]
}

// IsAsk
//
//  Returns true if Uninitialized or Wildcard
func (it Value) IsAsk() bool {
	return undefinedMap[it]
}

// IsIndeterminate
//
//  Returns true if Uninitialized or Wildcard
func (it Value) IsIndeterminate() bool {
	return undefinedMap[it]
}

// IsAccept
//
//  Returns true if True or Set
func (it Value) IsAccept() bool {
	return trueMap[it]
}

// IsReject
//
//  Returns true if False or Unset
func (it Value) IsReject() bool {
	return falseMap[it]
}

func (it Value) IsFailed() bool {
	return falseMap[it]
}

func (it Value) IsSuccess() bool {
	return trueMap[it]
}

// IsSkip
//
//  Returns true if Uninitialized or Wildcard
func (it Value) IsSkip() bool {
	return undefinedMap[it]
}

func (it Value) NameValue() string {
	return fmt.Sprintf(
		constants.EnumNameValueFormat,
		it.Name(),
		it.Value())
}

func (it Value) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Value) IsAnyNamesOf(names ...string) bool {
	for _, name := range names {
		if it.IsNameEqual(name) {
			return true
		}
	}

	return false
}

func (it Value) ToNumberString() string {
	return strconv.Itoa(it.ValueInt())
}

func (it Value) ValueByte() byte {
	return byte(it)
}

func (it Value) ValueInt() int {
	return int(it)
}

func (it Value) ValueInt8() int8 {
	return int8(it)
}

func (it Value) ValueInt16() int16 {
	return int16(it)
}

func (it Value) ValueInt32() int32 {
	return int32(it)
}

func (it Value) ValueString() string {
	return strconv.Itoa(it.ValueInt())
}

func (it Value) Format(format string) (compiled string) {
	newMap := map[string]string{
		"{type-name}": typeName,
		"{name}":      it.Name(),
		"{value}":     it.ValueString(),
	}

	for search, replacer := range newMap {
		format = strings.ReplaceAll(format, search, replacer)
	}

	return format
}

func (it Value) EnumType() enuminf.EnumTyper {
	return enumtype.Byte
}

func (it Value) Value() byte {
	return byte(it)
}

func (it Value) StringValue() string {
	return strconv.Itoa(it.ValueInt())
}

func (it Value) String() string {
	return valuesNames[it]
}

// IsTrue v == True
func (it Value) IsTrue() bool {
	return it == True
}

// IsFalse v == False
func (it Value) IsFalse() bool {
	return it == False
}

func (it Value) IsTrueOrSet() bool {
	return it == True || it == Set
}

// IsSet v == Set
func (it Value) IsSet() bool {
	return it == Set
}

// IsUnset v == Unset
func (it Value) IsUnset() bool {
	return it == Unset
}

func (it Value) HasInitialized() bool {
	return it != Uninitialized
}

func (it Value) HasInitializedAndSet() bool {
	return it == Set
}

func (it Value) HasInitializedAndTrue() bool {
	return it == True
}

func (it Value) IsWildcard() bool {
	return it == Wildcard
}

func (it Value) IsInit() bool {
	return it != Uninitialized
}

func (it Value) IsInitBoolean() bool {
	return it == True || it == False
}

func (it Value) IsDefinedBoolean() bool {
	return it == True || it == False
}

func (it Value) IsInitBooleanWild() bool {
	return it == True || it == False || it == Wildcard
}

func (it Value) IsInitSet() bool {
	return it == Set || it == Unset
}

func (it Value) IsInitSetWild() bool {
	return it == Set || it == Unset || it == Wildcard
}

func (it Value) IsYes() bool {
	return it == True
}

func (it Value) Boolean() bool {
	return it == True
}

func (it Value) IsOnLogically() bool {
	return it.IsInitialized() && trueMap[it]
}

func (it Value) IsOffLogically() bool {
	return it.IsInitialized() && falseMap[it]
}

func (it Value) IsAccepted() bool {
	return it.IsOnLogically()
}

func (it Value) IsRejected() bool {
	return it.IsOffLogically()
}

// IsDefinedLogically
//
// Not Uninitialized, Wildcard
func (it Value) IsDefinedLogically() bool {
	return !undefinedMap[it]
}

// IsUndefinedLogically
//
// Either Uninitialized, Wildcard
func (it Value) IsUndefinedLogically() bool {
	return undefinedMap[it]
}

func (it Value) IsInvalid() bool {
	return it == Uninitialized
}

func (it Value) IsValid() bool {
	return it != Uninitialized
}

func (it *Value) GetSetBoolOnInvalid(
	setterValue bool,
) bool {
	if it.IsDefinedBoolean() {
		return it.IsTrue()
	}

	*it = GetBool(setterValue)

	return it.IsTrue()
}

func (it *Value) GetSetBoolOnInvalidFunc(
	setterFunc func() bool,
) bool {
	if it.IsDefinedBoolean() {
		return it.IsTrue()
	}

	*it = GetBool(setterFunc())

	return it.IsTrue()
}

func (it Value) ToBooleanValue() Value {
	return convSetUnsetToTrueFalseMap[it]
}

func (it Value) ToSetUnsetValue() Value {
	return convTrueFalseToSetUnsetMap[it]
}

// LazyEvaluateBool
//
// Only execute evaluatorFunc if Uninitialized
// and then set True to self and returns t/f based on called or not
func (it *Value) LazyEvaluateBool(
	evaluatorFunc func(),
) (isCalled bool) {
	if it.IsDefinedBoolean() {
		return false
	}

	evaluatorFunc()
	*it = True

	return it.IsTrue()
}

// LazyEvaluateSet
//
// Only execute evaluatorFunc if Uninitialized
// and then set True to self and returns t/f based on called or not
func (it *Value) LazyEvaluateSet(
	evaluatorFunc func(),
) (isCalled bool) {
	if it.IsInitSet() {
		return false
	}

	evaluatorFunc()
	*it = Set

	return it.IsSet()
}

// IsWildcardOrBool
//
// if v.IsWildcard() then returns true regardless
//
// or else
//
// returns (isBool && v.IsTrue()) || (!isBool && v.IsFalse())
func (it Value) IsWildcardOrBool(isBool bool) bool {
	if it.IsWildcard() {
		return true
	}

	return isBool
}

func (it Value) ToByteCondition(trueVal, falseVal, invalid byte) byte {
	if it.IsTrue() {
		return trueVal
	}

	if it.IsFalse() {
		return falseVal
	}

	return invalid
}

func (it Value) ToByteConditionWithWildcard(wildcard, trueVal, falseVal, invalid byte) byte {
	if it.IsWildcard() {
		return wildcard
	}

	return it.ToByteCondition(trueVal, falseVal, invalid)
}

// WildcardApply
//
// if IsWildcard() || IsUnSetOrUninitialized() then
//
//      return inputVal
// else
//
//      return v. IsTrue()
func (it Value) WildcardApply(inputBool bool) bool {
	if it.IsWildcard() || it.IsUnSetOrUninitialized() {
		return inputBool
	}

	return it.IsTrue()
}

// WildcardValueApply
//
// if IsWildcard() || IsUnSetOrUninitialized() then
//
//      return inputVal
// else
//
//      return v. IsTrue()
func (it Value) WildcardValueApply(inputVal Value) bool {
	if it.IsWildcard() || it.IsUnSetOrUninitialized() {
		return inputVal.IsTrue()
	}

	return it.IsTrue()
}

// OrBool
//
// if IsWildcard() || IsUnSetOrUninitialized() then
//
//      return inputBool
// else
//
//      return v. IsTrue() || inputBool
func (it Value) OrBool(inputBool bool) bool {
	if it.IsWildcard() || it.IsUnSetOrUninitialized() {
		return inputBool
	}

	return it.IsTrue() || inputBool
}

// OrValue
//
// if IsWildcard() || IsUnSetOrUninitialized() then
//
//      return inputVal
// else
//
//      return v. IsTrue() || inputVal. IsTrue()
func (it Value) OrValue(inputVal Value) bool {
	if it.IsWildcard() || it.IsUnSetOrUninitialized() {
		return inputVal.IsTrue()
	}

	return it.IsTrue() || inputVal.IsTrue()
}

// AndBool
//
// if IsWildcard() || IsUnSetOrUninitialized() then
//
//      return inputVal
// else
//
//      return v. IsTrue() && inputBool
func (it Value) AndBool(inputBool bool) bool {
	if it.IsWildcard() || it.IsUnSetOrUninitialized() {
		return inputBool
	}

	return it.IsTrue() && inputBool
}

// And
//
// if IsWildcard() || IsUnSetOrUninitialized() then
//
//      return inputVal
// else
//
//      return GetBool(v. IsTrue() && inputVal. IsTrue())
func (it Value) And(inputVal Value) Value {
	if it.IsWildcard() || it.IsUnSetOrUninitialized() {
		return inputVal
	}

	return GetBool(it.IsTrue() && inputVal.IsTrue())
}

// IsUninitialized v == Uninitialized
func (it Value) IsUninitialized() bool {
	return it == Uninitialized
}

func (it Value) IsInitialized() bool {
	return it != Uninitialized
}

// IsUnSetOrUninitialized v == Uninitialized || v == Unset
func (it Value) IsUnSetOrUninitialized() bool {
	return it == Uninitialized || it == Unset
}

// IsNegative v == Uninitialized || v == Unset || v == False
func (it Value) IsNegative() bool {
	return it == Uninitialized || it == Unset || it == False
}

// IsPositive v == True || v == Set
func (it Value) IsPositive() bool {
	return it == True || it == Set
}

// IsBetween val >= start &&  val <= end
func (it Value) IsBetween(start, end byte) bool {
	val := it.Value()

	return val >= start && val <= end
}

// IsBetweenInt val >= start &&  val <= end
func (it Value) IsBetweenInt(start, end int) bool {
	val := it.Value()

	return val >= byte(start) && val <= byte(end)
}

// Add v + n
func (it Value) Add(n byte) Value {
	return Value(it.Value() + n)
}

func (it Value) Is(n Value) bool {
	return it.Value() == n.Value()
}

func (it Value) IsEqual(n byte) bool {
	return it.Value() == n
}

// IsGreater v.Value() > n
func (it Value) IsGreater(n byte) bool {
	return it.Value() > n
}

// IsGreaterEqual v.Value() >= n
func (it Value) IsGreaterEqual(n byte) bool {
	return it.Value() >= n
}

// IsLess v.Value() < n
func (it Value) IsLess(n byte) bool {
	return it.Value() < n
}

// IsLessEqual v.Value() <= n
func (it Value) IsLessEqual(n byte) bool {
	return it.Value() <= n
}

func (it Value) IsEqualInt(n int) bool {
	return it.Value() == byte(n)
}

// IsGreaterInt v.Value() > n
func (it Value) IsGreaterInt(n int) bool {
	return it.Value() > byte(n)
}

// IsGreaterEqualInt v.Value() >= n
func (it Value) IsGreaterEqualInt(n int) bool {
	return it.Value() >= byte(n)
}

// IsLessInt v.Value() < n
func (it Value) IsLessInt(n int) bool {
	return it.Value() < byte(n)
}

// IsLessEqualInt v.Value() <= n
func (it Value) IsLessEqualInt(n int) bool {
	return it.Value() <= byte(n)
}

func (it Value) PanicOnOutOfRange(n byte, msg string) {
	if IsOutOfRange(n) {
		panic(msg)
	}
}

func (it Value) GetErrorOnOutOfRange(n byte, msg string) error {
	if IsOutOfRange(n) {
		return errors.New(msg)
	}

	return nil
}

func (it Value) Name() string {
	return valuesToNameMap[it]
}

func (it Value) YesNoMappedValue() string {
	if it.IsUninitialized() {
		return constants.EmptyString
	}

	if it.IsTrueOrSet() {
		return Yes
	}

	return No
}

func (it Value) YesNoLowercaseName() string {
	return lowerCaseYesNoNames[it]
}

func (it Value) YesNoName() string {
	return yesNoNames[it]
}

func (it Value) TrueFalseName() string {
	return trueFalseNames[it]
}

func (it Value) OnOffLowercaseName() string {
	return lowerCaseOnOffNames[it]
}

func (it Value) OnOffName() string {
	return onOffNames[it]
}

func (it Value) TrueFalseLowercaseName() string {
	return trueFalseLowerNames[it]
}

func (it Value) SetUnsetLowercaseName() string {
	return setUnsetLowerNames[it]
}

func (it Value) MarshalJSON() ([]byte, error) {
	return valuesToJsonBytesMap[it], nil
}

func (it *Value) UnmarshalJSON(data []byte) error {
	if data == nil {
		return defaulterr.UnmarshallingFailedDueToNilOrEmpty
	}

	str := string(data)
	val, has := jsonValuesMap[str]

	if !has {
		//goland:noinspection SpellCheckingInspection
		return errors.New(
			"UnmarshalJSON failed , cannot map " +
				str +
				" to issetter.Value")
	}

	*it = val

	return nil
}

func (it Value) Serialize() ([]byte, error) {
	return it.MarshalJSON()
}

func (it Value) TypeName() string {
	return typeName
}

func (it Value) IsAnyValuesEqual(
	anyByteValues ...byte,
) bool {
	for _, value := range anyByteValues {
		if it.Value() == value {
			return true
		}
	}

	return false
}

func (it Value) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	err := it.UnmarshalJSON(jsonUnmarshallingValue)

	return it.ValueByte(), err
}

func (it Value) Deserialize(
	jsonBytes []byte,
) (Value, error) {
	currentVal, err := it.UnmarshallEnumToValue(jsonBytes)

	if err != nil {
		return Uninitialized, err
	}

	return Value(currentVal), err
}

func (it Value) MaxByte() byte {
	return Wildcard.ValueByte()
}

func (it Value) MinByte() byte {
	return Uninitialized.ValueByte()
}

func (it Value) RangesByte() []byte {
	panic("not implemented, later, todo")
}

func (it Value) ToPtr() *Value {
	return &it
}
