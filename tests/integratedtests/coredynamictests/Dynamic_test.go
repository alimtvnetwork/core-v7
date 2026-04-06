package coredynamictests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Test: Constructors — NewDynamicValid
// ==========================================================================

func Test_Dynamic_Constructor_NewDynamicValid(t *testing.T) {
	tc := dynamicConstructorNewDynamicValidTestCase
	d := refNewDynamicValid("hello")

	actual := args.Map{
		"isValid":   d.IsValid(),
		"dataValue": fmt.Sprintf("%v", d.Value()),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Constructors — NewDynamic invalid
// ==========================================================================

func Test_Dynamic_Constructor_NewDynamic_Invalid(t *testing.T) {
	tc := dynamicConstructorNewDynamicInvalidTestCase
	d := refNewDynamic(nil, false)

	actual := args.Map{
		"isValid": d.IsValid(),
		"isNull":  d.IsInvalid(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Constructors — InvalidDynamic
// ==========================================================================

func Test_Dynamic_Constructor_InvalidDynamic(t *testing.T) {
	tc := dynamicConstructorInvalidDynamicTestCase
	d := refInvalidDynamic()

	actual := args.Map{
		"isValid": d.IsValid(),
		"isNull":  d.IsNull(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Constructors — InvalidDynamicPtr
// ==========================================================================

func Test_Dynamic_Constructor_InvalidDynamicPtr(t *testing.T) {
	tc := dynamicConstructorInvalidDynamicPtrTestCase
	d := refInvalidDynamicPtr()

	actual := args.Map{
		"isNotNilPtr": d != nil,
		"isValid":     d.IsValid(),
		"isNull":      d.IsNull(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Constructors — NewDynamicPtr
// ==========================================================================

func Test_Dynamic_Constructor_NewDynamicPtr(t *testing.T) {
	tc := dynamicConstructorNewDynamicPtrTestCase
	d := refNewDynamicPtr(42, true)

	actual := args.Map{
		"isNotNilPtr": d != nil,
		"isValid":     d.IsValid(),
		"dataValue":   fmt.Sprintf("%v", d.Value()),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Clone
// ==========================================================================

func Test_Dynamic_Clone(t *testing.T) {
	tc := dynamicCloneTestCase
	original := refNewDynamicValid("data")
	cloned := original.Clone()

	actual := args.Map{
		"clonedValue":   fmt.Sprintf("%v", cloned.Value()),
		"isIndependent": cloned.IsValid(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// Note: ClonePtr nil receiver test migrated to NilReceiver_test.go using CaseNilSafe pattern.

func Test_Dynamic_ClonePtr_Valid(t *testing.T) {
	tc := dynamicClonePtrValidTestCase
	original := refNewDynamicPtr("data", true)
	cloned := original.ClonePtr()

	actual := args.Map{
		"isNotNilPtr": cloned != nil,
		"clonedValue": fmt.Sprintf("%v", cloned.Value()),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Dynamic_NonPtr(t *testing.T) {
	tc := dynamicNonPtrTestCase
	d := refNewDynamicValid("x")

	np := d.NonPtr()
	actLines := []string{fmt.Sprintf("%v", np.Value())}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: Type Checks — Data/Value equality
// ==========================================================================

func Test_Dynamic_DataValueEquality(t *testing.T) {
	tc := dynamicDataValueEqualityTestCase
	d := refNewDynamicValid(99)

	actual := args.Map{
		"dataValue":       fmt.Sprintf("%v", d.Data()),
		"dataEqualsValue": d.Data() == d.Value(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Type Checks — String non-empty
// ==========================================================================

func Test_Dynamic_StringNonEmpty(t *testing.T) {
	tc := dynamicStringNonEmptyTestCase
	d := refNewDynamicValid("hello")

	actLines := []string{fmt.Sprintf("%v", d.String() != "")}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: Type Checks — IsPointer with pointer data
// ==========================================================================

func Test_Dynamic_IsPointer_WithPointerData(t *testing.T) {
	tc := dynamicIsPointerTestCase
	val := 42
	d := refNewDynamicValid(&val)

	actLines := []string{fmt.Sprintf("%v", refIsPointer.Call(&d))}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: Type Checks — Bool method ref checks (uniform)
// ==========================================================================

func Test_Dynamic_TypeChecks_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicTypeCheckTestCases {
		input := tc.ArrangeInput.(dynamicBoolCheckInput)
		d := createDynamicFromBoolCheck(input)

		actLines := []string{fmt.Sprintf("%v", input.CheckRef.Call(d))}

		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================================================
// Test: IsStruct
// ==========================================================================

func Test_Dynamic_IsStruct_True(t *testing.T) {
	type sample struct{ Name string }

	tc := dynamicIsStructTrueTestCase
	d := refNewDynamicValid(sample{Name: "test"})

	actLines := []string{fmt.Sprintf("%v", refIsStruct.Call(&d))}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_Dynamic_IsStruct_False(t *testing.T) {
	tc := dynamicIsStructFalseTestCase
	d := refNewDynamicValid(5)

	actLines := []string{fmt.Sprintf("%v", refIsStruct.Call(&d))}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: Length
// ==========================================================================

func Test_Dynamic_Length_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicLengthTestCases {
		input := tc.ArrangeInput.(dynamicInputMap)
		d := createDynamicFromInputMap(input)

		actLines := []string{fmt.Sprintf("%d", d.Length())}

		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================================================
// Test: ValueInt
// ==========================================================================

func Test_Dynamic_ValueInt_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueIntTestCases {
		input := tc.ArrangeInput.(dynamicInputMap)
		d := refNewDynamicValid(input.InputData)

		actLines := []string{fmt.Sprintf("%d", d.ValueInt())}

		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================================================
// Test: ValueBool
// ==========================================================================

func Test_Dynamic_ValueBool_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueBoolTestCases {
		input := tc.ArrangeInput.(dynamicInputMap)
		d := refNewDynamicValid(input.InputData)

		actLines := []string{fmt.Sprintf("%v", d.ValueBool())}

		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================================================
// Test: ValueString
// ==========================================================================

func Test_Dynamic_ValueString_Direct(t *testing.T) {
	tc := dynamicValueStringDirectTestCase
	d := refNewDynamicValid("hello")

	actLines := []string{d.ValueString()}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_Dynamic_ValueString_NonString(t *testing.T) {
	tc := dynamicValueStringNonStringTestCase
	d := refNewDynamicValid(42)

	actLines := []string{fmt.Sprintf("%v", d.ValueString() != "")}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_Dynamic_ValueString_Nil(t *testing.T) {
	tc := dynamicValueStringNilTestCase
	d := refNewDynamic(nil, true)

	actLines := []string{fmt.Sprintf("%v", d.ValueString() == "")}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: ValueStrings
// ==========================================================================

func Test_Dynamic_ValueStrings_Slice(t *testing.T) {
	tc := dynamicValueStringsSliceTestCase
	d := refNewDynamicValid([]string{"a", "b"})
	result := d.ValueStrings()

	actual := args.Map{
		"item0": result[0],
		"item1": result[1],
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Dynamic_ValueStrings_NonSlice(t *testing.T) {
	tc := dynamicValueStringsNonSliceTestCase
	d := refNewDynamicValid(42)
	result := d.ValueStrings()

	actLines := []string{fmt.Sprintf("%v", result == nil)}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: ValueUInt
// ==========================================================================

func Test_Dynamic_ValueUInt_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueUIntTestCases {
		input := tc.ArrangeInput.(dynamicInputMap)
		d := refNewDynamicValid(input.InputData)

		actLines := []string{fmt.Sprintf("%d", d.ValueUInt())}

		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================================================
// Test: ValueInt64
// ==========================================================================

func Test_Dynamic_ValueInt64_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicValueInt64TestCases {
		input := tc.ArrangeInput.(dynamicInputMap)
		d := refNewDynamicValid(input.InputData)

		actLines := []string{fmt.Sprintf("%d", d.ValueInt64())}

		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================================================
// Test: Bytes
// ==========================================================================

func Test_Dynamic_Bytes_Valid(t *testing.T) {
	tc := dynamicBytesValidTestCase
	d := refNewDynamicValid([]byte("raw"))
	raw, ok := d.Bytes()

	actual := args.Map{
		"hasBytes": ok,
		"content":  string(raw),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Dynamic_Bytes_NonBytes(t *testing.T) {
	tc := dynamicBytesNonBytesTestCase
	d := refNewDynamicValid("str")
	_, ok := d.Bytes()

	actLines := []string{fmt.Sprintf("%v", ok)}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// Note: Bytes nil receiver test migrated to NilReceiver_test.go using CaseNilSafe pattern.

// ==========================================================================
// Test: IntDefault
// ==========================================================================

func Test_Dynamic_IntDefault_Valid(t *testing.T) {
	tc := dynamicIntDefaultValidTestCase
	d := refNewDynamicValid(42)
	val, ok := d.IntDefault(0)

	actual := args.Map{
		"isValid":  ok,
		"intValue": val,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Dynamic_IntDefault_NilData(t *testing.T) {
	tc := dynamicIntDefaultNilTestCase
	d := refNewDynamic(nil, true)
	val, ok := d.IntDefault(99)

	actual := args.Map{
		"isValid":      ok,
		"defaultValue": val,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: ValueNullErr
// ==========================================================================

// Note: ValueNullErr nil receiver test migrated to NilReceiver_test.go using CaseNilSafe pattern.

func Test_Dynamic_ValueNullErr_NullData(t *testing.T) {
	tc := dynamicValueNullErrNullDataTestCase
	d := refNewDynamic(nil, true)

	actLines := []string{fmt.Sprintf("%v", d.ValueNullErr() != nil)}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_Dynamic_ValueNullErr_ValidData(t *testing.T) {
	tc := dynamicValueNullErrValidTestCase
	d := refNewDynamicValid("ok")

	actLines := []string{fmt.Sprintf("%v", d.ValueNullErr() != nil)}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: Reflect — ReflectKind
// ==========================================================================

func Test_Dynamic_ReflectKind_String(t *testing.T) {
	tc := dynamicReflectKindStringTestCase
	d := refNewDynamicValid("text")

	actLines := []string{d.ReflectKind().String()}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_Dynamic_ReflectKind_Int(t *testing.T) {
	tc := dynamicReflectKindIntTestCase
	d := refNewDynamicValid(42)

	actLines := []string{d.ReflectKind().String()}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: Reflect — IsReflectKind
// ==========================================================================

func Test_Dynamic_IsReflectKindMatch(t *testing.T) {
	tc := dynamicIsReflectKindMatchTestCase
	d := refNewDynamicValid("x")

	actLines := []string{fmt.Sprintf("%v", d.IsReflectKind(reflect.String))}

	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_Dynamic_IsReflectKindMismatch(t *testing.T) {
	tc := dynamicIsReflectKindMismatchTestCase
	d := refNewDynamicValid("x")

	actLines := []string{fmt.Sprintf("%v", d.IsReflectKind(reflect.Int))}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: Reflect — ReflectTypeName
// ==========================================================================

func Test_Dynamic_ReflectTypeName(t *testing.T) {
	tc := dynamicReflectTypeNameTestCase
	d := refNewDynamicValid("text")

	actLines := []string{fmt.Sprintf("%v", d.ReflectTypeName() != "")}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: Reflect — ReflectType
// ==========================================================================

func Test_Dynamic_ReflectType(t *testing.T) {
	tc := dynamicReflectTypeTestCase
	d := refNewDynamicValid(42)

	actLines := []string{fmt.Sprintf("%v", d.ReflectType() == reflect.TypeOf(42))}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: Reflect — IsReflectTypeOf
// ==========================================================================

func Test_Dynamic_IsReflectTypeOf(t *testing.T) {
	tc := dynamicIsReflectTypeOfTestCase
	d := refNewDynamicValid("hello")

	actLines := []string{fmt.Sprintf("%v", d.IsReflectTypeOf(reflect.TypeOf("")))}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: ReflectValue (cached)
// ==========================================================================

func Test_Dynamic_ReflectValue_Verification(t *testing.T) {
	tc := dynamicReflectValueCachedTestCase
	d := refNewDynamicPtr(42, true)

	rv1 := d.ReflectValue()
	rv2 := d.ReflectValue()

	actual := args.Map{
		"isCached":       rv1 == rv2,
		"extractedValue": int(rv1.Int()),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Loop — Iterate
// ==========================================================================

func Test_Dynamic_Loop_Iterate(t *testing.T) {
	tc := dynamicLoopIterateTestCase
	d := refNewDynamicValid([]string{"a", "b", "c"})
	collected := make([]string, 0, 3)
	called := d.Loop(func(index int, item any) bool {
		collected = append(collected, item.(string))

		return false
	})

	actual := args.Map{
		"didLoop": called,
		"item0":   collected[0],
		"item1":   collected[1],
		"item2":   collected[2],
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Loop — Invalid
// ==========================================================================

func Test_Dynamic_Loop_Invalid(t *testing.T) {
	tc := dynamicLoopInvalidTestCase
	d := refInvalidDynamicPtr()
	called := d.Loop(func(index int, item any) bool { return false })

	actLines := []string{fmt.Sprintf("%v", called)}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: Loop — Break
// ==========================================================================

func Test_Dynamic_Loop_Break(t *testing.T) {
	tc := dynamicLoopBreakTestCase
	d := refNewDynamicValid([]int{1, 2, 3, 4})
	count := 0
	d.Loop(func(index int, item any) bool {
		count++

		return index == 1
	})

	actLines := []string{fmt.Sprintf("%d", count)}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: ItemAccess — ItemUsingIndex
// ==========================================================================

func Test_Dynamic_ItemUsingIndex(t *testing.T) {
	tc := dynamicItemUsingIndexTestCase
	d := refNewDynamicValid([]string{"a", "b"})

	actual := args.Map{
		"item0": fmt.Sprintf("%v", d.ItemUsingIndex(0)),
		"item1": fmt.Sprintf("%v", d.ItemUsingIndex(1)),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: ItemAccess — ItemUsingKey
// ==========================================================================

func Test_Dynamic_ItemUsingKey(t *testing.T) {
	tc := dynamicItemUsingKeyTestCase
	d := refNewDynamicValid(map[string]int{"k": 42})

	actLines := []string{fmt.Sprintf("%v", d.ItemUsingKey("k"))}

	tc.ShouldBeEqual(t, 0, actLines...)
}

// ==========================================================================
// Test: IsStructStringNullOrEmpty
// ==========================================================================

func Test_Dynamic_StructStringNullOrEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range dynamicStructStringNullOrEmptyTestCases {
		input := tc.ArrangeInput.(dynamicBoolCheckInput)
		d := createDynamicFromBoolCheck(input)

		actLines := []string{fmt.Sprintf("%v", input.CheckRef.Call(d))}

		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// =============================================================================
// Helpers
// =============================================================================

func createDynamicFromInputMap(input dynamicInputMap) *coredynamic.Dynamic {
	d := refNewDynamic(input.InputData, input.IsValid)

	return &d
}

func createDynamicFromBoolCheck(input dynamicBoolCheckInput) *coredynamic.Dynamic {
	d := refNewDynamic(input.InputData, input.IsValid)

	return &d
}
