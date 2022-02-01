package issetter

import (
	"errors"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/defaulterr"
)

// Value
//
//  Used evaluate lazy boolean values.
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

func (it Value) Value() byte {
	return byte(it)
}

func (it Value) StringValue() string {
	return string(it)
}

func (it Value) String() string {
	return values[it]
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

// LazyEvaluateBool
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
