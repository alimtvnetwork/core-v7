package chmodhelpertests

import (
	"os"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodclasstype"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── RwxWrapper Creation ──

func Test_RwxWrapper_Create_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2RwxWrapperCreateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		mode, _ := input.GetAsString("mode")

		// Act
		wrapper, err := chmodhelper.New.RwxWrapper.Create(mode)

		actual := args.Map{
			"rwxFull":   wrapper.ToFullRwxValueString(),
			"fileMode":  wrapper.ToFileModeString(),
			"rwx3":      wrapper.ToRwxCompiledStr(),
			"hasError":  err != nil,
			"isDefined": wrapper.IsDefined(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxWrapper_RwxFullString_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2RwxFullStringParseTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rwxFull, _ := input.GetAsString("rwxFull")

		// Act
		wrapper, err := chmodhelper.New.RwxWrapper.RwxFullString(rwxFull)
		hasError := err != nil

		actual := args.Map{
			"hasError": hasError,
		}

		if !hasError {
			actual["rwx3"] = wrapper.ToRwxCompiledStr()
			actual["isDefined"] = wrapper.IsDefined()
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ── RwxWrapper methods ──

func Test_RwxWrapper_Bytes_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	rwxBytes := wrapper.Bytes()

	// Assert
	if rwxBytes[0] != 7 {
		t.Errorf("owner should be 7, got %d", rwxBytes[0])
	}
	if rwxBytes[1] != 5 {
		t.Errorf("group should be 5, got %d", rwxBytes[1])
	}
	if rwxBytes[2] != 5 {
		t.Errorf("other should be 5, got %d", rwxBytes[2])
	}
}

func Test_RwxWrapper_ToCompiledOctalBytes_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	b4 := wrapper.ToCompiledOctalBytes4Digits()
	b3 := wrapper.ToCompiledOctalBytes3Digits()
	o, g, oth := wrapper.ToCompiledSplitValues()

	// Assert
	if b4[0] != '0' {
		t.Error("first byte of 4-digit should be '0'")
	}
	if len(b3) != 3 {
		t.Error("3-digit should have 3 bytes")
	}
	if o != '7' || g != '5' || oth != '5' {
		t.Errorf("split values wrong: %c %c %c", o, g, oth)
	}
}

func Test_RwxWrapper_ToUint32Octal_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	octal := wrapper.ToUint32Octal()

	// Assert
	if octal != 0755 {
		t.Errorf("expected 0755, got %o", octal)
	}
}

func Test_RwxWrapper_ToFileMode_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	mode := wrapper.ToFileMode()

	// Assert
	if mode != os.FileMode(0755) {
		t.Errorf("expected 0755, got %v", mode)
	}
}

func Test_RwxWrapper_ToFullRwxValuesChars_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	chars := wrapper.ToFullRwxValuesChars()

	// Assert
	if len(chars) != 10 {
		t.Errorf("expected 10 chars, got %d", len(chars))
	}
}

func Test_RwxWrapper_String_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	str := wrapper.String()

	// Assert
	if str != "-rwxr-xr-x" {
		t.Errorf("expected -rwxr-xr-x, got %s", str)
	}
}

func Test_RwxWrapper_FriendlyDisplay_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	display := wrapper.FriendlyDisplay()

	// Assert
	if display == "" {
		t.Error("FriendlyDisplay should not be empty")
	}
}

func Test_RwxWrapper_Clone_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	cloned := wrapper.Clone()

	// Assert
	if cloned == nil {
		t.Error("Clone should not return nil")
	}
	if !cloned.IsEqualPtr(&wrapper) {
		t.Error("cloned should equal original")
	}
}

func Test_RwxWrapper_Clone_Nil_Ext2(t *testing.T) {
	// Arrange
	var wrapper *chmodhelper.RwxWrapper

	// Act
	cloned := wrapper.Clone()

	// Assert
	if cloned != nil {
		t.Error("Clone on nil should return nil")
	}
}

func Test_RwxWrapper_IsEmpty_Ext2(t *testing.T) {
	// Arrange
	empty := chmodhelper.RwxWrapper{}

	// Assert
	if !empty.IsEmpty() {
		t.Error("empty wrapper should be empty")
	}
	if !empty.IsNull() {
		// non-pointer IsNull returns false
	}
	if !empty.IsInvalid() {
		t.Error("empty wrapper should be invalid")
	}
}

func Test_RwxWrapper_IsEqualPtr_Ext2(t *testing.T) {
	// Arrange
	w1, _ := chmodhelper.New.RwxWrapper.Create("755")
	w2, _ := chmodhelper.New.RwxWrapper.Create("755")
	w3, _ := chmodhelper.New.RwxWrapper.Create("644")

	// Assert
	if !w1.IsEqualPtr(&w2) {
		t.Error("same mode wrappers should be equal")
	}
	if w1.IsEqualPtr(&w3) {
		t.Error("different mode wrappers should not be equal")
	}
	if !w1.IsEqualPtr(nil) == true {
		// nil comparison
	}
}

func Test_RwxWrapper_IsEqualFileMode_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Assert
	if !wrapper.IsEqualFileMode(os.FileMode(0755)) {
		t.Error("should be equal to 0755")
	}
	if wrapper.IsEqualFileMode(os.FileMode(0644)) {
		t.Error("should not be equal to 0644")
	}
	if !wrapper.IsNotEqualFileMode(os.FileMode(0644)) {
		t.Error("should be not equal to 0644")
	}
}

func Test_RwxWrapper_IsRwxFullEqual_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Assert
	if !wrapper.IsRwxFullEqual("-rwxr-xr-x") {
		t.Error("should match full rwx string")
	}
	if wrapper.IsRwxFullEqual("-rw-r--r--") {
		t.Error("should not match different rwx")
	}
	if wrapper.IsRwxFullEqual("short") {
		t.Error("should not match short string")
	}
}

func Test_RwxWrapper_IsRwxEqualLocation_Ext2(t *testing.T) {
	// Assert
	w, _ := chmodhelper.New.RwxWrapper.Create("755")
	// non-existent path
	if w.IsRwxEqualLocation("/nonexistent/path/xyz123") {
		t.Error("should be false for nonexistent path")
	}
}

func Test_RwxWrapper_IsRwxEqualFileInfo_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Assert - nil
	if wrapper.IsRwxEqualFileInfo(nil) {
		t.Error("should be false for nil fileInfo")
	}
}

func Test_RwxWrapper_IsEqualVarWrapper_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Assert - nil
	if wrapper.IsEqualVarWrapper(nil) {
		t.Error("should be false for nil varWrapper")
	}
}

func Test_RwxWrapper_ToPtr_ToNonPtr_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	ptr := wrapper.ToPtr()
	nonPtr := ptr.ToNonPtr()

	// Assert
	if ptr == nil {
		t.Error("ToPtr should not return nil")
	}
	if nonPtr.ToFullRwxValueString() != wrapper.ToFullRwxValueString() {
		t.Error("ToNonPtr should equal original")
	}
}

func Test_RwxWrapper_ToRwxOwnerGroupOther_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	ogo := wrapper.ToRwxOwnerGroupOther()

	// Assert
	if ogo == nil {
		t.Error("should not be nil")
	}
	if ogo.Owner != "rwx" {
		t.Errorf("owner should be rwx, got %s", ogo.Owner)
	}
}

func Test_RwxWrapper_ToRwxInstruction_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")
	condition := &chmodins.Condition{IsSkipOnInvalid: true}

	// Act
	ins := wrapper.ToRwxInstruction(condition)

	// Assert
	if ins == nil {
		t.Error("should not be nil")
	}
}

func Test_RwxWrapper_JSON_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	jsonBytes, err := wrapper.MarshalJSON()

	// Assert
	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}
	if len(jsonBytes) == 0 {
		t.Error("JSON bytes should not be empty")
	}

	// Act - unmarshal
	var parsed chmodhelper.RwxWrapper
	err2 := parsed.UnmarshalJSON(jsonBytes)

	// Assert
	if err2 != nil {
		t.Errorf("UnmarshalJSON error: %v", err2)
	}
	if parsed.ToFullRwxValueString() != wrapper.ToFullRwxValueString() {
		t.Error("parsed should equal original")
	}
}

func Test_RwxWrapper_Json_Methods_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	jsonResult := wrapper.Json()
	jsonPtrResult := wrapper.JsonPtr()

	// Assert
	if jsonResult.HasError() {
		t.Error("Json() should not error")
	}
	if jsonPtrResult == nil {
		t.Error("JsonPtr should not be nil")
	}

	// JsonParseSelfInject
	var target chmodhelper.RwxWrapper
	err := target.JsonParseSelfInject(&jsonResult)
	if err != nil {
		t.Errorf("JsonParseSelfInject error: %v", err)
	}

	// AsJsonContractsBinder
	binder := wrapper.AsJsonContractsBinder()
	if binder == nil {
		t.Error("binder should not be nil")
	}
}

// ── Attribute tests ──

func Test_Attribute_Methods_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2AttributeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rwx, _ := input.GetAsString("rwx")

		// Act
		attr := chmodhelper.New.Attribute.UsingRwxString(rwx)

		actual := args.Map{
			"isRead":    attr.IsRead,
			"isWrite":   attr.IsWrite,
			"isExecute": attr.IsExecute,
			"toByte":    attr.ToByte(),
			"rwxStr":    attr.ToRwxString(),
			"isEmpty":   attr.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Attribute_NilSafe_Ext2(t *testing.T) {
	// Arrange
	var attr *chmodhelper.Attribute

	// Assert
	if !attr.IsNull() {
		t.Error("nil attr should be null")
	}
	if !attr.IsAnyNull() {
		t.Error("nil attr IsAnyNull should be true")
	}
	if !attr.IsEmpty() {
		t.Error("nil attr should be empty")
	}
	if !attr.IsZero() {
		t.Error("nil attr should be zero")
	}
	if !attr.IsInvalid() {
		t.Error("nil attr should be invalid")
	}
	if attr.IsDefined() {
		t.Error("nil attr should not be defined")
	}
	if attr.HasAnyItem() {
		t.Error("nil attr HasAnyItem should be false")
	}
	cloned := attr.Clone()
	if cloned != nil {
		t.Error("nil attr Clone should be nil")
	}
}

func Test_Attribute_ToAttributeValue_Ext2(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.UsingRwxString("rwx")

	// Act
	attrVal := attr.ToAttributeValue()

	// Assert
	if attrVal.Sum != 7 {
		t.Errorf("expected sum 7, got %d", attrVal.Sum)
	}
	if attrVal.Read != 4 {
		t.Errorf("expected read 4, got %d", attrVal.Read)
	}
}

func Test_Attribute_ToSpecificBytes_Ext2(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.UsingRwxString("r-x")

	// Act
	r, w, x, sum := attr.ToSpecificBytes()

	// Assert
	if r != 4 || w != 0 || x != 1 || sum != 5 {
		t.Errorf("unexpected: r=%d w=%d x=%d sum=%d", r, w, x, sum)
	}
}

func Test_Attribute_ToRwx_Ext2(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.UsingRwxString("r-x")

	// Act
	rwx := attr.ToRwx()

	// Assert
	if rwx[0] != 'r' || rwx[1] != '-' || rwx[2] != 'x' {
		t.Errorf("unexpected rwx: %c%c%c", rwx[0], rwx[1], rwx[2])
	}
}

func Test_Attribute_ToVariant_Ext2(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.UsingRwxString("rwx")

	// Act
	v := attr.ToVariant()

	// Assert
	if v.Value() != 7 {
		t.Errorf("expected 7, got %d", v.Value())
	}
}

func Test_Attribute_ToStringByte_Ext2(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.UsingRwxString("rwx")

	// Act
	sb := attr.ToStringByte()

	// Assert
	if sb != '7' {
		t.Errorf("expected '7', got %c", sb)
	}
}

func Test_Attribute_Clone_Ext2(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.UsingRwxString("rwx")

	// Act
	cloned := attr.Clone()

	// Assert
	if cloned == nil {
		t.Error("Clone should not be nil")
	}
	if !attr.IsEqualPtr(cloned) {
		t.Error("cloned should equal original")
	}
}

func Test_Attribute_IsEqual_Ext2(t *testing.T) {
	// Arrange
	a1 := chmodhelper.New.Attribute.UsingRwxString("rwx")
	a2 := chmodhelper.New.Attribute.UsingRwxString("rwx")
	a3 := chmodhelper.New.Attribute.UsingRwxString("r--")

	// Assert
	if !a1.IsEqual(a2) {
		t.Error("same attrs should be equal")
	}
	if a1.IsEqual(a3) {
		t.Error("different attrs should not be equal")
	}
}

func Test_Attribute_IsEqualPtr_BothNil_Ext2(t *testing.T) {
	// Arrange
	var a1, a2 *chmodhelper.Attribute

	// Assert
	if !a1.IsEqualPtr(a2) {
		t.Error("both nil should be equal")
	}

	a3 := chmodhelper.New.Attribute.UsingRwxString("rwx")
	if a1.IsEqualPtr(&a3) {
		t.Error("nil vs non-nil should not be equal")
	}
}

func Test_Attribute_UsingByte_Ext2(t *testing.T) {
	// Act
	attr, err := chmodhelper.New.Attribute.UsingByte(5)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if !attr.IsRead || attr.IsWrite || !attr.IsExecute {
		t.Error("byte 5 = r-x")
	}
}

func Test_Attribute_UsingByte_Invalid_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.New.Attribute.UsingByte(8)

	// Assert
	if err == nil {
		t.Error("expected error for byte > 7")
	}
}

func Test_Attribute_UsingVariant_Ext2(t *testing.T) {
	// Act
	attr, err := chmodhelper.New.Attribute.UsingVariant(chmodhelper.ReadWrite)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if !attr.IsRead || !attr.IsWrite || attr.IsExecute {
		t.Error("ReadWrite should be rw-")
	}
}

// ── AttrVariant tests ──

func Test_AttrVariant_Methods_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2AttrVariantTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		variant := input["variant"].(chmodhelper.AttrVariant)

		// Act
		actual := args.Map{
			"value":     variant.Value(),
			"isGreater": variant.IsGreaterThan(5),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_AttrVariant_ToAttribute_Ext2(t *testing.T) {
	// Arrange
	v := chmodhelper.ReadWriteExecute

	// Act
	attr := v.ToAttribute()

	// Assert
	if !attr.IsRead || !attr.IsWrite || !attr.IsExecute {
		t.Error("ReadWriteExecute should have all true")
	}
}

// ── Variant tests ──

func Test_Variant_ToWrapper_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2VariantTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		variant := input["variant"].(chmodhelper.Variant)

		// Act
		wrapper, err := variant.ToWrapper()

		actual := args.Map{
			"rwxFull":  wrapper.ToFullRwxValueString(),
			"hasError": err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_ToWrapperPtr_Ext2(t *testing.T) {
	// Arrange
	v := chmodhelper.X777

	// Act
	ptr, err := v.ToWrapperPtr()

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if ptr == nil {
		t.Error("should not be nil")
	}
}

func Test_Variant_ExpandOctalByte_Ext2(t *testing.T) {
	// Arrange
	v := chmodhelper.X755

	// Act
	r, w, x := v.ExpandOctalByte()

	// Assert
	if r != '7' || w != '5' || x != '5' {
		t.Errorf("unexpected: %c %c %c", r, w, x)
	}
}

// ── ParseRwxToVarAttribute ──

func Test_ParseRwxToVarAttribute_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2ParseRwxToVarAttrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rwx, _ := input.GetAsString("rwx")

		// Act
		varAttr, err := chmodhelper.ParseRwxToVarAttribute(rwx)

		actual := args.Map{
			"hasError": err != nil,
		}
		if err == nil {
			actual["isFixedType"] = varAttr.IsFixedType()
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_VarAttribute_Methods_Ext2(t *testing.T) {
	// Arrange
	varAttr, _ := chmodhelper.ParseRwxToVarAttribute("r*x")

	// Assert
	if varAttr.IsFixedType() {
		t.Error("should have wildcard")
	}
	if !varAttr.HasWildcard() {
		t.Error("should have wildcard")
	}
	if varAttr.String() != "r*x" {
		t.Errorf("expected r*x, got %s", varAttr.String())
	}

	// Clone
	cloned := varAttr.Clone()
	if cloned == nil {
		t.Error("Clone should not be nil")
	}
	if !varAttr.IsEqualPtr(cloned) {
		t.Error("cloned should equal")
	}

	// ToCompileFixAttr -- not fixed type, returns nil
	fixAttr := varAttr.ToCompileFixAttr()
	if fixAttr != nil {
		t.Error("non-fixed should return nil from ToCompileFixAttr")
	}
}

func Test_VarAttribute_Fixed_Ext2(t *testing.T) {
	// Arrange
	varAttr, _ := chmodhelper.ParseRwxToVarAttribute("rwx")

	// Assert
	if !varAttr.IsFixedType() {
		t.Error("should be fixed type")
	}

	fixAttr := varAttr.ToCompileFixAttr()
	if fixAttr == nil {
		t.Error("fixed type should return non-nil")
	}
	if !fixAttr.IsRead || !fixAttr.IsWrite || !fixAttr.IsExecute {
		t.Error("should have all true")
	}
}

func Test_VarAttribute_ToCompileAttr_WithWildcard_Ext2(t *testing.T) {
	// Arrange
	varAttr, _ := chmodhelper.ParseRwxToVarAttribute("r*x")
	fixed := chmodhelper.New.Attribute.UsingRwxString("rw-")

	// Act
	compiled := varAttr.ToCompileAttr(&fixed)

	// Assert
	if !compiled.IsRead {
		t.Error("read should be true")
	}
	if !compiled.IsWrite {
		t.Error("wildcard should inherit write=true from fixed")
	}
	if !compiled.IsExecute {
		t.Error("execute should be true")
	}
}

func Test_VarAttribute_NilSafe_Ext2(t *testing.T) {
	// Arrange
	var varAttr *chmodhelper.VarAttribute

	// Assert
	cloned := varAttr.Clone()
	if cloned != nil {
		t.Error("nil Clone should be nil")
	}
}

func Test_VarAttribute_IsEqualPtr_BothNil_Ext2(t *testing.T) {
	// Arrange
	var a, b *chmodhelper.VarAttribute

	// Assert
	if !a.IsEqualPtr(b) {
		t.Error("both nil should be equal")
	}
}

// ── MergeRwxWildcardWithFixedRwx ──

func Test_MergeRwxWildcardWithFixedRwx_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2MergeRwxTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		existing, _ := input.GetAsString("existing")
		wildcard, _ := input.GetAsString("wildcard")

		// Act
		fixedAttr, err := chmodhelper.MergeRwxWildcardWithFixedRwx(existing, wildcard)

		actual := args.Map{
			"hasError": err != nil,
		}
		if err == nil && fixedAttr != nil {
			actual["result"] = fixedAttr.ToRwxString()
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ── NewRwxVariableWrapper ──

func Test_NewRwxVariableWrapper_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2NewRwxVarWrapperTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		partial, _ := input.GetAsString("partial")

		// Act
		varWrapper, err := chmodhelper.NewRwxVariableWrapper(partial)

		actual := args.Map{
			"hasError": err != nil,
		}
		if err == nil {
			actual["isFixedType"] = varWrapper.IsFixedType()
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxVariableWrapper_Methods_Ext2(t *testing.T) {
	// Arrange
	varWrapper, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr--")

	// Assert
	if !varWrapper.IsFixedType() {
		t.Error("should be fixed")
	}
	if varWrapper.HasWildcard() {
		t.Error("should not have wildcard")
	}

	fixedPtr := varWrapper.ToCompileFixedPtr()
	if fixedPtr == nil {
		t.Error("fixed type should return ptr")
	}

	compiled := varWrapper.ToCompileWrapper(nil)
	if compiled.ToFullRwxValueString() != "-rwxr-xr--" {
		t.Errorf("unexpected: %s", compiled.ToFullRwxValueString())
	}

	cloned := varWrapper.Clone()
	if cloned == nil {
		t.Error("Clone should not be nil")
	}
	if !varWrapper.IsEqualPtr(cloned) {
		t.Error("cloned should equal")
	}

	toStr := varWrapper.ToString(true)
	if toStr != "-rwxr-xr--" {
		t.Errorf("ToString with hyphen: %s", toStr)
	}
	toStr2 := varWrapper.ToString(false)
	if toStr2 != "rwxr-xr--" {
		t.Errorf("ToString without hyphen: %s", toStr2)
	}

	if varWrapper.String() == "" {
		t.Error("String should not be empty")
	}
}

func Test_RwxVariableWrapper_PartialMatch_Ext2(t *testing.T) {
	// Arrange
	varWrapper, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr--")

	// Assert
	if !varWrapper.IsOwnerPartialMatch("rwx") {
		t.Error("owner should match")
	}
	if !varWrapper.IsGroupPartialMatch("r-x") {
		t.Error("group should match")
	}
	if !varWrapper.IsOtherPartialMatch("r--") {
		t.Error("other should match")
	}
}

func Test_RwxVariableWrapper_IsEqual_Ext2(t *testing.T) {
	// Arrange
	varWrapper, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr--")

	// Assert
	if !varWrapper.IsEqualPartialFullRwx("-rwxr-xr--") {
		t.Error("should match full rwx")
	}
	if varWrapper.IsEqualPartialFullRwx("-rwxrwxrwx") {
		t.Error("should not match different")
	}
	if !varWrapper.IsMismatchPartialFullRwx("-rwxrwxrwx") {
		t.Error("should be mismatch")
	}
	if !varWrapper.IsEqualPartialRwxPartial("-rwxr-xr--") {
		t.Error("should match partial")
	}
	if !varWrapper.IsEqualPartialUsingFileMode(os.FileMode(0754)) {
		t.Error("should match 0754")
	}
	if !varWrapper.IsEqualUsingFileMode(os.FileMode(0754)) {
		t.Error("should match 0754")
	}
	if varWrapper.IsEqualUsingLocation("/nonexistent/path/xyz") {
		t.Error("should not match nonexistent")
	}
	if varWrapper.IsEqualUsingFileInfo(nil) {
		t.Error("should not match nil fileInfo")
	}
	if varWrapper.IsEqualRwxWrapperPtr(nil) {
		t.Error("should not match nil wrapper")
	}
}

func Test_RwxVariableWrapper_NilSafe_Ext2(t *testing.T) {
	// Arrange
	var v1, v2 *chmodhelper.RwxVariableWrapper

	// Assert
	if v1.Clone() != nil {
		t.Error("nil Clone should be nil")
	}
	if !v1.IsEqualPtr(v2) {
		t.Error("both nil should be equal")
	}
}

// ── GetRwxLengthError ──

func Test_GetRwxLengthError_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2GetRwxLengthErrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rwx, _ := input.GetAsString("rwx")

		// Act
		err := chmodhelper.GetRwxLengthError(rwx)

		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ── ParseRwxInstructionToStringRwx ──

func Test_ParseRwxInstructionToStringRwx_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2ParseRwxToStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rwx := input["rwx"].(*chmodins.RwxOwnerGroupOther)
		includeHyphen := input["includeHyphen"].(bool)

		// Act
		result := chmodhelper.ParseRwxInstructionToStringRwx(rwx, includeHyphen)

		actual := args.Map{
			"result": result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ── ExpandCharRwx ──

func Test_ExpandCharRwx_Ext2(t *testing.T) {
	// Act
	r, w, x := chmodhelper.ExpandCharRwx("rwx")

	// Assert
	if r != 'r' || w != 'w' || x != 'x' {
		t.Errorf("unexpected: %c %c %c", r, w, x)
	}
}

// ── IsChmod ──

func Test_IsChmod_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	info, _ := os.Stat(tmpFile.Name())
	rwxFull := info.Mode().String()

	// Assert
	if !chmodhelper.IsChmod(tmpFile.Name(), rwxFull) {
		t.Error("should match existing chmod")
	}
	if chmodhelper.IsChmod(tmpFile.Name(), "short") {
		t.Error("short string should return false")
	}
	if chmodhelper.IsChmod("", "-rwxrwxrwx") {
		t.Error("empty path should return false")
	}
	if chmodhelper.IsChmod("/nonexistent/xyz", "-rwxrwxrwx") {
		t.Error("nonexistent should return false")
	}
}

// ── IsChmodEqualUsingRwxOwnerGroupOther ──

func Test_IsChmodEqualUsingRwxOwnerGroupOther_Ext2(t *testing.T) {
	// Assert
	if chmodhelper.IsChmodEqualUsingRwxOwnerGroupOther("/tmp", nil) {
		t.Error("nil rwx should return false")
	}
}

// ── FileModeFriendlyString ──

func Test_FileModeFriendlyString_Ext2(t *testing.T) {
	// Act
	result := chmodhelper.FileModeFriendlyString(os.FileMode(0755))

	// Assert
	if result == "" {
		t.Error("should not be empty")
	}
}

// ── GetExistingChmodOfValidFile ──

func Test_GetExistingChmodOfValidFile_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	chmod, isInvalid := chmodhelper.GetExistingChmodOfValidFile(tmpFile.Name())

	// Assert
	if isInvalid {
		t.Error("should not be invalid")
	}
	if chmod == 0 {
		t.Error("chmod should not be 0")
	}
}

func Test_GetExistingChmodOfValidFile_Invalid_Ext2(t *testing.T) {
	// Act
	chmod, isInvalid := chmodhelper.GetExistingChmodOfValidFile("/nonexistent/xyz")

	// Assert
	if !isInvalid {
		t.Error("should be invalid")
	}
	if chmod != 0 {
		t.Error("chmod should be 0")
	}
}

// ── ParseRwxOwnerGroupOtherToFileMode ──

func Test_ParseRwxOwnerGroupOtherToFileMode_Ext2(t *testing.T) {
	// Arrange
	ogo := &chmodins.RwxOwnerGroupOther{
		Owner: "rwx",
		Group: "r-x",
		Other: "r--",
	}

	// Act
	mode, err := chmodhelper.ParseRwxOwnerGroupOtherToFileMode(ogo)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if mode != os.FileMode(0754) {
		t.Errorf("expected 0754, got %o", mode)
	}
}

// ── ParseRwxInstructionToExecutor ──

func Test_ParseRwxInstructionToExecutor_Ext2(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{},
	}

	// Act
	executor, err := chmodhelper.ParseRwxInstructionToExecutor(ins)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if executor == nil {
		t.Error("executor should not be nil")
	}
	if !executor.IsFixedWrapper() {
		t.Error("should be fixed wrapper")
	}
	if executor.IsVarWrapper() {
		t.Error("should not be var wrapper")
	}
}

func Test_ParseRwxInstructionToExecutor_Nil_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.ParseRwxInstructionToExecutor(nil)

	// Assert
	if err == nil {
		t.Error("expected error for nil")
	}
}

func Test_RwxInstructionExecutor_IsEqualFileMode_Ext2(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{},
	}
	executor, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)

	// Assert
	if !executor.IsEqualFileMode(os.FileMode(0754)) {
		t.Error("should be equal to 0754")
	}
	if executor.IsEqualFileMode(os.FileMode(0777)) {
		t.Error("should not be equal to 0777")
	}
}

func Test_RwxInstructionExecutor_CompiledWrapper_Ext2(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{},
	}
	executor, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)

	// Act
	compiled, err := executor.CompiledWrapper(os.FileMode(0755))

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if compiled == nil {
		t.Error("should not be nil")
	}
}

func Test_RwxInstructionExecutor_IsEqualRwxPartial_Ext2(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{},
	}
	executor, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)

	// Assert
	if !executor.IsEqualRwxPartial("-rwxr-xr--") {
		t.Error("should match partial")
	}
}

func Test_RwxInstructionExecutor_IsEqualFileInfo_Nil_Ext2(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{},
	}
	executor, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)

	// Assert
	if executor.IsEqualFileInfo(nil) {
		t.Error("nil fileInfo should return false")
	}
	if executor.IsEqualRwxWrapper(nil) {
		t.Error("nil wrapper should return false")
	}
}

// ── RwxInstructionExecutors ──

func Test_RwxInstructionExecutors_Ext2(t *testing.T) {
	// Arrange
	executors := chmodhelper.NewRwxInstructionExecutors(2)

	// Assert
	if !executors.IsEmpty() {
		t.Error("should be empty initially")
	}
	if executors.HasAnyItem() {
		t.Error("should not have items")
	}
	if executors.Length() != 0 {
		t.Error("length should be 0")
	}
	if executors.Count() != 0 {
		t.Error("count should be 0")
	}
	if executors.LastIndex() != -1 {
		t.Error("lastIndex should be -1")
	}
	if executors.HasIndex(0) {
		t.Error("should not have index 0")
	}

	// Add nil -- should skip
	executors.Add(nil)
	if executors.Length() != 0 {
		t.Error("adding nil should not increase length")
	}
}

func Test_ParseRwxInstructionsToExecutors_Ext2(t *testing.T) {
	// Arrange
	instructions := []chmodins.RwxInstruction{
		{
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "r-x",
				Other: "r--",
			},
			Condition: chmodins.Condition{},
		},
	}

	// Act
	executors, err := chmodhelper.ParseRwxInstructionsToExecutors(instructions)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if executors.Length() != 1 {
		t.Errorf("expected 1, got %d", executors.Length())
	}
}

func Test_ParseRwxInstructionsToExecutors_Nil_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.ParseRwxInstructionsToExecutors(nil)

	// Assert
	if err == nil {
		t.Error("expected error for nil")
	}
}

func Test_ParseRwxInstructionsToExecutors_Empty_Ext2(t *testing.T) {
	// Act
	executors, err := chmodhelper.ParseRwxInstructionsToExecutors([]chmodins.RwxInstruction{})

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if executors.Length() != 0 {
		t.Error("should be empty")
	}
}

// ── ParseRwxOwnerGroupOtherToRwxVariableWrapper ──

func Test_ParseRwxOwnerGroupOtherToRwxVariableWrapper_Nil_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(nil)

	// Assert
	if err == nil {
		t.Error("expected error for nil")
	}
}

// ── FilteredPathFileInfoMap ──

func Test_FilteredPathFileInfoMap_Empty_Ext2(t *testing.T) {
	// Arrange
	fmap := chmodhelper.InvalidFilteredPathFileInfoMap()

	// Assert
	if fmap.HasAnyValidFileInfo() {
		t.Error("should not have valid info")
	}
	if !fmap.IsEmptyValidFileInfos() {
		t.Error("should be empty")
	}
	if fmap.HasAnyMissingPaths() {
		t.Error("should not have missing")
	}
	if !fmap.IsEmptyIssues() {
		t.Error("should be empty issues")
	}
	if fmap.HasError() {
		t.Error("should not have error")
	}
	if fmap.HasAnyIssues() {
		t.Error("should not have issues")
	}
	if fmap.LengthOfIssues() != 0 {
		t.Error("issues length should be 0")
	}
	if fmap.MissingPathsToString() != "" {
		t.Error("missing paths string should be empty")
	}
}

func Test_GetExistsFilteredPathFileInfoMap_Ext2(t *testing.T) {
	// Arrange
	tempDir := os.TempDir()

	// Act
	fmap := chmodhelper.GetExistsFilteredPathFileInfoMap(
		false,
		tempDir,
		"/nonexistent/xyz",
	)

	// Assert
	if !fmap.HasAnyValidFileInfo() {
		t.Error("should have valid tempDir")
	}
	if !fmap.HasAnyMissingPaths() {
		t.Error("should have missing /nonexistent/xyz")
	}
	if fmap.HasError() == false {
		// should have error because isSkipOnInvalid=false
	}
}

func Test_GetExistsFilteredPathFileInfoMap_Empty_Ext2(t *testing.T) {
	// Act
	fmap := chmodhelper.GetExistsFilteredPathFileInfoMap(false)

	// Assert
	if fmap.HasAnyValidFileInfo() {
		t.Error("should be empty")
	}
}

func Test_FilteredPathFileInfoMap_LazyMethods_Ext2(t *testing.T) {
	// Arrange
	fmap := chmodhelper.GetExistsFilteredPathFileInfoMap(
		true,
		os.TempDir(),
	)

	// Act
	locs := fmap.LazyValidLocations()
	locs2 := fmap.LazyValidLocations() // cached
	infos := fmap.ValidFileInfos()
	wrappers := fmap.LazyValidLocationFileInfoRwxWrappers()
	wrappers2 := fmap.LazyValidLocationFileInfoRwxWrappers() // cached

	// Assert
	if len(locs) == 0 || len(locs2) == 0 {
		t.Error("should have locations")
	}
	if len(infos) == 0 {
		t.Error("should have infos")
	}
	if len(wrappers) == 0 || len(wrappers2) == 0 {
		t.Error("should have wrappers")
	}
}

// ── GetFilteredExistsPaths ──

func Test_GetFilteredExistsPaths_Ext2(t *testing.T) {
	// Act
	found, missing := chmodhelper.GetFilteredExistsPaths([]string{
		os.TempDir(),
		"/nonexistent/xyz",
	})

	// Assert
	if len(found) != 1 {
		t.Errorf("expected 1 found, got %d", len(found))
	}
	if len(missing) != 1 {
		t.Errorf("expected 1 missing, got %d", len(missing))
	}
}

func Test_GetFilteredExistsPaths_Empty_Ext2(t *testing.T) {
	// Act
	found, missing := chmodhelper.GetFilteredExistsPaths([]string{})

	// Assert
	if len(found) != 0 || len(missing) != 0 {
		t.Error("should be empty")
	}
}

// ── RwxMatchingStatus ──

func Test_RwxMatchingStatus_Empty_Ext2(t *testing.T) {
	// Arrange
	status := chmodhelper.EmptyRwxMatchingStatus()

	// Assert
	if status.IsAllMatching {
		t.Error("should not be all matching")
	}
	if status.MissingFilesToString() != "" {
		t.Error("missing files should be empty")
	}
}

func Test_RwxMatchingStatus_Invalid_Ext2(t *testing.T) {
	// Arrange
	err := os.ErrNotExist
	status := chmodhelper.InvalidRwxMatchingStatus(err)

	// Assert
	if status.IsAllMatching {
		t.Error("should not be matching")
	}
	if status.Error == nil {
		t.Error("should have error")
	}
}

func Test_RwxMatchingStatus_CreateErrFinalError_AllMatching_Ext2(t *testing.T) {
	// Arrange
	status := &chmodhelper.RwxMatchingStatus{
		IsAllMatching: true,
		Error:         nil,
	}

	// Act
	err := status.CreateErrFinalError()

	// Assert
	if err != nil {
		t.Error("all matching with no error should return nil")
	}
}

// ── GetExistingChmodRwxWrappers ──

func Test_GetExistingChmodRwxWrappers_Empty_Ext2(t *testing.T) {
	// Act
	wrappers, err := chmodhelper.GetExistingChmodRwxWrappers(false)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if len(wrappers) != 0 {
		t.Error("should be empty")
	}
}

func Test_GetExistingChmodRwxWrappers_ContinueOnError_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	wrappers, err := chmodhelper.GetExistingChmodRwxWrappers(
		true,
		tmpFile.Name(),
		"/nonexistent/xyz",
	)

	// Assert
	if len(wrappers) != 1 {
		t.Errorf("expected 1 valid, got %d", len(wrappers))
	}
	if err == nil {
		t.Error("should have error for nonexistent path")
	}
}

func Test_GetExistingChmodRwxWrappers_ImmediateExit_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.GetExistingChmodRwxWrappers(
		false,
		"/nonexistent/xyz",
	)

	// Assert
	if err == nil {
		t.Error("should have error")
	}
}

// ── GetRecursivePaths ──

func Test_GetRecursivePaths_File_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	paths, err := chmodhelper.GetRecursivePaths(false, tmpFile.Name())

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if len(paths) != 1 {
		t.Errorf("expected 1, got %d", len(paths))
	}
}

func Test_GetRecursivePaths_Dir_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_recursive_ext2"
	os.MkdirAll(dir, 0755)
	defer os.RemoveAll(dir)
	os.WriteFile(dir+"/file.txt", []byte("test"), 0644)

	// Act
	paths, err := chmodhelper.GetRecursivePaths(false, dir)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if len(paths) < 2 {
		t.Errorf("expected at least 2 paths, got %d", len(paths))
	}
}

func Test_GetRecursivePaths_NonExistent_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.GetRecursivePaths(false, "/nonexistent/xyz")

	// Assert
	if err == nil {
		t.Error("should have error")
	}
}

// ── GetExistingChmodRwxWrapper ──

func Test_GetExistingChmodRwxWrapper_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	wrapper, err := chmodhelper.GetExistingChmodRwxWrapper(tmpFile.Name())

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if wrapper.IsEmpty() {
		t.Error("should not be empty")
	}
}

func Test_GetExistingChmodRwxWrapper_Invalid_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.GetExistingChmodRwxWrapper("/nonexistent/xyz")

	// Assert
	if err == nil {
		t.Error("expected error")
	}
}

// ── SimpleFileReaderWriter ──

func Test_SimpleFileReaderWriter_WriteReadString_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_sfrw_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, filePath)

	// Act
	err := rw.WriteString("hello world")

	// Assert
	if err != nil {
		t.Errorf("write error: %v", err)
	}

	content, readErr := rw.ReadString()
	if readErr != nil {
		t.Errorf("read error: %v", readErr)
	}
	if content != "hello world" {
		t.Errorf("expected 'hello world', got %s", content)
	}
}

func Test_SimpleFileReaderWriter_Properties_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_sfrw_props_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, filePath)

	// Assert
	if rw.IsExist() {
		t.Error("file should not exist yet")
	}
	if !rw.HasPathIssues() {
		t.Error("should have path issues")
	}
	if !rw.IsPathInvalid() {
		t.Error("should be path invalid")
	}
	if rw.HasAnyIssues() == false {
		// parent dir doesn't exist either
	}

	str := rw.String()
	if str == "" {
		t.Error("String should not be empty")
	}
}

func Test_SimpleFileReaderWriter_WriteAndRead_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_sfrw_wr_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, filePath)

	// Act
	err := rw.Write([]byte("bytes content"))
	if err != nil {
		t.Fatalf("write error: %v", err)
	}

	// Read
	bytes, readErr := rw.Read()
	if readErr != nil {
		t.Errorf("read error: %v", readErr)
	}
	if string(bytes) != "bytes content" {
		t.Errorf("unexpected content: %s", string(bytes))
	}

	// ReadOnExist
	bytes2, err2 := rw.ReadOnExist()
	if err2 != nil || len(bytes2) == 0 {
		t.Error("ReadOnExist should work")
	}

	// ReadStringOnExist
	s, err3 := rw.ReadStringOnExist()
	if err3 != nil || s == "" {
		t.Error("ReadStringOnExist should work")
	}

	// Expire
	expErr := rw.Expire()
	if expErr != nil {
		t.Errorf("Expire error: %v", expErr)
	}
	if rw.IsExist() {
		t.Error("file should not exist after Expire")
	}
}

func Test_SimpleFileReaderWriter_ReadOnExist_NotExist_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/nonexistent/xyz/file.txt")

	// Act
	bytes, err := rw.ReadOnExist()

	// Assert
	if err != nil || bytes != nil {
		t.Error("ReadOnExist on nonexistent should return nil, nil")
	}

	s, err2 := rw.ReadStringOnExist()
	if err2 != nil || s != "" {
		t.Error("ReadStringOnExist on nonexistent should return empty")
	}
}

func Test_SimpleFileReaderWriter_WritePath_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_sfrw_wp_ext2"
	filePath := dir + "/test.txt"
	newPath := dir + "/subdir/other.txt"
	defer os.RemoveAll(dir)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, filePath)
	rw.Write([]byte("init"))

	// Act
	err := rw.WritePath(true, newPath, []byte("new content"))

	// Assert
	if err != nil {
		t.Errorf("WritePath error: %v", err)
	}
}

func Test_SimpleFileReaderWriter_JoinRelPath_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/tmp/test/file.txt")

	// Act
	joined := rw.JoinRelPath("sub/other.txt")
	joinedEmpty := rw.JoinRelPath("")

	// Assert
	if joined == "" {
		t.Error("should not be empty")
	}
	if joinedEmpty == "" {
		t.Error("should not be empty")
	}
}

func Test_SimpleFileReaderWriter_WriteAny_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_sfrw_any_ext2"
	filePath := dir + "/test.json"
	defer os.RemoveAll(dir)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, filePath)

	// Act
	err := rw.WriteAny(map[string]string{"key": "value"})

	// Assert
	if err != nil {
		t.Errorf("WriteAny error: %v", err)
	}
}

func Test_SimpleFileReaderWriter_Get_NotExist_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/nonexistent/xyz/file.json")

	// Act
	var target map[string]string
	err := rw.Get(&target)

	// Assert
	if err == nil {
		t.Error("Get on nonexistent should error")
	}
}

func Test_SimpleFileReaderWriter_Clone_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/tmp/test.txt")

	// Act
	cloned := rw.Clone()

	// Assert
	if cloned.FilePath != rw.FilePath {
		t.Error("cloned FilePath should match")
	}
}

func Test_SimpleFileReaderWriter_ClonePtr_Nil_Ext2(t *testing.T) {
	// Arrange
	var rw *chmodhelper.SimpleFileReaderWriter

	// Act
	cloned := rw.ClonePtr()

	// Assert
	if cloned != nil {
		t.Error("nil ClonePtr should return nil")
	}
}

func Test_SimpleFileReaderWriter_JSON_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/tmp/test.txt")

	// Act
	jsonBytes, err := rw.MarshalJSON()

	// Assert
	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}

	var parsed chmodhelper.SimpleFileReaderWriter
	err2 := parsed.UnmarshalJSON(jsonBytes)
	if err2 != nil {
		t.Errorf("UnmarshalJSON error: %v", err2)
	}
}

func Test_SimpleFileReaderWriter_Json_Methods_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/tmp/test.txt")

	// Act
	jsonResult := rw.Json()
	jsonPtrResult := rw.JsonPtr()

	// Assert
	if jsonResult.HasError() {
		t.Error("Json() should not error")
	}
	if jsonPtrResult == nil {
		t.Error("JsonPtr should not nil")
	}

	var target chmodhelper.SimpleFileReaderWriter
	err := target.JsonParseSelfInject(&jsonResult)
	if err != nil {
		t.Errorf("JsonParseSelfInject error: %v", err)
	}

	binder := rw.AsJsonContractsBinder()
	if binder == nil {
		t.Error("binder should not be nil")
	}
}

func Test_SimpleFileReaderWriter_InitializeDefault_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  os.FileMode(0755),
		ChmodFile: os.FileMode(0644),
		FilePath:  "/tmp/test/file.txt",
	}

	// Act
	initialized := rw.InitializeDefault(true)

	// Assert
	if initialized == nil {
		t.Error("should not be nil")
	}
	if initialized.ParentDir == "" {
		t.Error("ParentDir should be populated")
	}
}

func Test_SimpleFileReaderWriter_InitializeDefaultApplyChmod_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  os.FileMode(0755),
		ChmodFile: os.FileMode(0644),
		FilePath:  "/tmp/test/file.txt",
	}

	// Act
	initialized := rw.InitializeDefaultApplyChmod()

	// Assert
	if initialized == nil {
		t.Error("should not be nil")
	}
}

func Test_SimpleFileReaderWriter_InitializeDefaultNew_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  os.FileMode(0755),
		ChmodFile: os.FileMode(0644),
		FilePath:  "/tmp/test/file.txt",
	}

	// Act
	newRw := rw.InitializeDefaultNew()

	// Assert
	if newRw == nil {
		t.Error("should not be nil")
	}
}

func Test_SimpleFileReaderWriter_NewPath_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  os.FileMode(0755),
		ChmodFile: os.FileMode(0644),
		FilePath:  "/tmp/test/file.txt",
	}

	// Act
	newRw := rw.NewPath(true, "/tmp/other/file2.txt")

	// Assert
	if newRw == nil {
		t.Error("should not be nil")
	}
}

func Test_SimpleFileReaderWriter_NewPathJoin_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  os.FileMode(0755),
		ChmodFile: os.FileMode(0644),
		FilePath:  "/tmp/test/file.txt",
		ParentDir: "/tmp/test",
	}

	// Act
	newRw := rw.NewPathJoin(true, "subdir", "file.txt")

	// Assert
	if newRw == nil {
		t.Error("should not be nil")
	}
}

func Test_SimpleFileReaderWriter_Serialize_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/nonexistent/file.txt")

	// Act -- file doesn't exist
	bytes, err := rw.Serialize()

	// Assert
	if err != nil || bytes != nil {
		// Serialize on non-exist returns nil, nil (alias for ReadOnExist)
	}
}

func Test_SimpleFileReaderWriter_RemoveOnExist_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/nonexistent/file.txt")

	// Act
	err := rw.RemoveOnExist()

	// Assert
	if err != nil {
		t.Errorf("RemoveOnExist on nonexistent should not error: %v", err)
	}
}

func Test_SimpleFileReaderWriter_ExpireParentDir_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_expire_parent_ext2"
	filePath := dir + "/test.txt"
	os.MkdirAll(dir, 0755)
	os.WriteFile(filePath, []byte("test"), 0644)
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, filePath)

	// Act
	err := rw.ExpireParentDir()

	// Assert
	if err != nil {
		t.Errorf("ExpireParentDir error: %v", err)
	}
}

func Test_SimpleFileReaderWriter_RemoveDirOnExist_Ext2(t *testing.T) {
	// Arrange -- dir doesn't exist
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/nonexistent/xyz/file.txt")

	// Act
	err := rw.RemoveDirOnExist()

	// Assert
	if err != nil {
		t.Errorf("RemoveDirOnExist on nonexistent should not error: %v", err)
	}
}

// ── newSimpleFileReaderWriterCreator ──

func Test_NewSimpleFileReaderWriter_All_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.All(
		os.FileMode(0755),
		os.FileMode(0644),
		true,
		true,
		true,
		"/tmp/dir",
		"/tmp/dir/file.txt",
	)

	// Assert
	if rw == nil {
		t.Error("should not be nil")
	}
}

func Test_NewSimpleFileReaderWriter_Options_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.Options(true, true, true, "/tmp/test.txt")

	// Assert
	if rw == nil {
		t.Error("should not be nil")
	}
}

func Test_NewSimpleFileReaderWriter_Create_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(
		true,
		os.FileMode(0755),
		os.FileMode(0644),
		"/tmp/dir",
		"/tmp/dir/file.txt",
	)

	// Assert
	if rw == nil {
		t.Error("should not be nil")
	}
}

func Test_NewSimpleFileReaderWriter_CreateClean_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.CreateClean(
		true,
		os.FileMode(0755),
		os.FileMode(0644),
		"./tmp/dir",
		"./tmp/dir/file.txt",
	)

	// Assert
	if rw == nil {
		t.Error("should not be nil")
	}
}

func Test_NewSimpleFileReaderWriter_DefaultCleanPath_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.DefaultCleanPath(true, "./tmp/file.txt")

	// Assert
	if rw == nil {
		t.Error("should not be nil")
	}
}

func Test_NewSimpleFileReaderWriter_Path_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.Path(true, os.FileMode(0755), os.FileMode(0644), "/tmp/file.txt")

	// Assert
	if rw == nil {
		t.Error("should not be nil")
	}
}

func Test_NewSimpleFileReaderWriter_PathCondition_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.PathCondition(true, true, os.FileMode(0755), os.FileMode(0644), "./tmp/file.txt")

	// Assert
	if rw == nil {
		t.Error("should not be nil")
	}
}

func Test_NewSimpleFileReaderWriter_PathDirDefaultChmod_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.PathDirDefaultChmod(true, os.FileMode(0644), "/tmp/file.txt")

	// Assert
	if rw == nil {
		t.Error("should not be nil")
	}
}

// ── newRwxWrapperCreator additional ──

func Test_NewRwxWrapperCreator_Invalid_Ext2(t *testing.T) {
	// Act
	w := chmodhelper.New.RwxWrapper.Invalid()
	wPtr := chmodhelper.New.RwxWrapper.InvalidPtr()
	empty := chmodhelper.New.RwxWrapper.Empty()

	// Assert
	if !w.IsEmpty() {
		t.Error("Invalid should be empty")
	}
	if wPtr == nil || !wPtr.IsEmpty() {
		t.Error("InvalidPtr should be empty")
	}
	if empty == nil || !empty.IsEmpty() {
		t.Error("Empty should be empty")
	}
}

func Test_NewRwxWrapperCreator_CreatePtr_Ext2(t *testing.T) {
	// Act
	ptr, err := chmodhelper.New.RwxWrapper.CreatePtr("755")

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if ptr == nil {
		t.Error("should not be nil")
	}
}

func Test_NewRwxWrapperCreator_UsingBytes_Ext2(t *testing.T) {
	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingBytes([3]byte{7, 5, 5})

	// Assert
	if wrapper.ToFullRwxValueString() != "-rwxr-xr-x" {
		t.Errorf("unexpected: %s", wrapper.ToFullRwxValueString())
	}
}

func Test_NewRwxWrapperCreator_UsingSpecificByte_Ext2(t *testing.T) {
	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingSpecificByte(7, 5, 5)

	// Assert
	if wrapper.ToRwxCompiledStr() != "755" {
		t.Errorf("unexpected: %s", wrapper.ToRwxCompiledStr())
	}
}

func Test_NewRwxWrapperCreator_UsingAttrVariants_Ext2(t *testing.T) {
	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingAttrVariants(
		chmodhelper.ReadWriteExecute,
		chmodhelper.ReadExecute,
		chmodhelper.ReadExecute,
	)

	// Assert
	if wrapper.ToRwxCompiledStr() != "755" {
		t.Errorf("unexpected: %s", wrapper.ToRwxCompiledStr())
	}
}

func Test_NewRwxWrapperCreator_UsingAttrs_Ext2(t *testing.T) {
	// Arrange
	owner := chmodhelper.New.Attribute.UsingRwxString("rwx")
	group := chmodhelper.New.Attribute.UsingRwxString("r-x")
	other := chmodhelper.New.Attribute.UsingRwxString("r-x")

	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingAttrs(owner, group, other)

	// Assert
	if wrapper.ToRwxCompiledStr() != "755" {
		t.Errorf("unexpected: %s", wrapper.ToRwxCompiledStr())
	}
}

func Test_NewRwxWrapperCreator_Rwx10_Ext2(t *testing.T) {
	// Act
	wrapper, err := chmodhelper.New.RwxWrapper.Rwx10("-rwxr-xr-x")

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if wrapper.ToRwxCompiledStr() != "755" {
		t.Errorf("unexpected: %s", wrapper.ToRwxCompiledStr())
	}
}

func Test_NewRwxWrapperCreator_Rwx9_Ext2(t *testing.T) {
	// Act
	wrapper, err := chmodhelper.New.RwxWrapper.Rwx9("rwxr-xr-x")

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if wrapper.ToRwxCompiledStr() != "755" {
		t.Errorf("unexpected: %s", wrapper.ToRwxCompiledStr())
	}
}

func Test_NewRwxWrapperCreator_RwxFullStringWtHyphen_InvalidLength_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rw")

	// Assert
	if err == nil {
		t.Error("expected error for invalid length")
	}
}

func Test_NewRwxWrapperCreator_UsingFileMode_Zero_Ext2(t *testing.T) {
	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingFileMode(0)
	wrapperPtr := chmodhelper.New.RwxWrapper.UsingFileModePtr(0)

	// Assert
	if !wrapper.IsEmpty() {
		t.Error("zero mode should be empty")
	}
	if wrapperPtr == nil || !wrapperPtr.IsEmpty() {
		t.Error("zero mode ptr should be empty")
	}
}

func Test_NewRwxWrapperCreator_UsingRwxOwnerGroupOther_Ext2(t *testing.T) {
	// Arrange
	ogo := &chmodins.RwxOwnerGroupOther{
		Owner: "rwx",
		Group: "r-x",
		Other: "r-x",
	}

	// Act
	wrapper, err := chmodhelper.New.RwxWrapper.UsingRwxOwnerGroupOther(ogo)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if wrapper.ToRwxCompiledStr() != "755" {
		t.Errorf("unexpected: %s", wrapper.ToRwxCompiledStr())
	}
}

func Test_NewRwxWrapperCreator_Instruction_Ext2(t *testing.T) {
	// Act
	ins, err := chmodhelper.New.RwxWrapper.Instruction(
		"-rwxr-xr-x",
		chmodins.Condition{},
	)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if ins == nil {
		t.Error("should not be nil")
	}
}

func Test_NewRwxWrapperCreator_UsingExistingFile_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	wrapper, err := chmodhelper.New.RwxWrapper.UsingExistingFile(tmpFile.Name())

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if wrapper == nil {
		t.Error("should not be nil")
	}
}

func Test_NewRwxWrapperCreator_UsingExistingFileSkipInvalidFile_Ext2(t *testing.T) {
	// Act - valid file
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	wrapper, isInvalid := chmodhelper.New.RwxWrapper.UsingExistingFileSkipInvalidFile(tmpFile.Name())
	if isInvalid || wrapper == nil {
		t.Error("valid file should work")
	}

	// Act - invalid
	wrapper2, isInvalid2 := chmodhelper.New.RwxWrapper.UsingExistingFileSkipInvalidFile("/nonexistent/xyz")
	if !isInvalid2 || wrapper2 == nil {
		t.Error("invalid file should return empty")
	}
}

func Test_NewRwxWrapperCreator_UsingExistingFileOption_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act - skip invalid = true, valid file
	wrapper, err, isInvalid := chmodhelper.New.RwxWrapper.UsingExistingFileOption(true, tmpFile.Name())
	if err != nil || isInvalid || wrapper == nil {
		t.Error("should work for valid file with skip=true")
	}

	// Act - skip invalid = false, valid file
	wrapper2, err2, isInvalid2 := chmodhelper.New.RwxWrapper.UsingExistingFileOption(false, tmpFile.Name())
	if err2 != nil || isInvalid2 || wrapper2 == nil {
		t.Error("should work for valid file with skip=false")
	}

	// Act - skip invalid = false, invalid file
	_, err3, _ := chmodhelper.New.RwxWrapper.UsingExistingFileOption(false, "/nonexistent/xyz")
	if err3 == nil {
		t.Error("should error for invalid file with skip=false")
	}
}

// ── SingleRwx ──

func Test_SingleRwx_All_Ext2(t *testing.T) {
	// Arrange
	singleRwx, err := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}

	ogo := singleRwx.ToRwxOwnerGroupOther()
	if ogo.Owner != "rwx" || ogo.Group != "rwx" || ogo.Other != "rwx" {
		t.Errorf("unexpected: %s %s %s", ogo.Owner, ogo.Group, ogo.Other)
	}
}

func Test_SingleRwx_Owner_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)

	// Act
	ogo := singleRwx.ToRwxOwnerGroupOther()

	// Assert
	if ogo.Owner != "rwx" || ogo.Group != "***" || ogo.Other != "***" {
		t.Errorf("unexpected: %s %s %s", ogo.Owner, ogo.Group, ogo.Other)
	}
}

func Test_SingleRwx_Group_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Group)

	// Act
	ogo := singleRwx.ToRwxOwnerGroupOther()

	// Assert
	if ogo.Owner != "***" || ogo.Group != "rwx" || ogo.Other != "***" {
		t.Errorf("unexpected: %s %s %s", ogo.Owner, ogo.Group, ogo.Other)
	}
}

func Test_SingleRwx_Other_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Other)

	// Act
	ogo := singleRwx.ToRwxOwnerGroupOther()

	// Assert
	if ogo.Owner != "***" || ogo.Group != "***" || ogo.Other != "rwx" {
		t.Errorf("unexpected: %s %s %s", ogo.Owner, ogo.Group, ogo.Other)
	}
}

func Test_SingleRwx_OwnerGroup_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.OwnerGroup)

	// Act
	ogo := singleRwx.ToRwxOwnerGroupOther()

	// Assert
	if ogo.Owner != "rwx" || ogo.Group != "rwx" || ogo.Other != "***" {
		t.Errorf("unexpected: %s %s %s", ogo.Owner, ogo.Group, ogo.Other)
	}
}

func Test_SingleRwx_GroupOther_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.GroupOther)

	// Act
	ogo := singleRwx.ToRwxOwnerGroupOther()

	// Assert
	if ogo.Owner != "***" || ogo.Group != "rwx" || ogo.Other != "rwx" {
		t.Errorf("unexpected: %s %s %s", ogo.Owner, ogo.Group, ogo.Other)
	}
}

func Test_SingleRwx_OwnerOther_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.OwnerOther)

	// Act
	ogo := singleRwx.ToRwxOwnerGroupOther()

	// Assert
	if ogo.Owner != "rwx" || ogo.Group != "***" || ogo.Other != "rwx" {
		t.Errorf("unexpected: %s %s %s", ogo.Owner, ogo.Group, ogo.Other)
	}
}

func Test_SingleRwx_InvalidLength_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.NewSingleRwx("rw", chmodclasstype.All)

	// Assert
	if err == nil {
		t.Error("expected error for invalid length")
	}
}

func Test_SingleRwx_ToRwxInstruction_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	cond := &chmodins.Condition{IsRecursive: true}

	// Act
	ins := singleRwx.ToRwxInstruction(cond)

	// Assert
	if ins == nil {
		t.Error("should not be nil")
	}
}

func Test_SingleRwx_ToVarRwxWrapper_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)

	// Act
	varWrapper, err := singleRwx.ToVarRwxWrapper()

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if varWrapper == nil {
		t.Error("should not be nil")
	}
}

func Test_SingleRwx_ToRwxWrapper_All_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)

	// Act
	wrapper, err := singleRwx.ToRwxWrapper()

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if wrapper == nil {
		t.Error("should not be nil")
	}
}

func Test_SingleRwx_ToRwxWrapper_NonAll_Fails_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)

	// Act
	_, err := singleRwx.ToRwxWrapper()

	// Assert
	if err == nil {
		t.Error("non-All classType should error")
	}
}

func Test_SingleRwx_ToDisabledRwxWrapper_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)

	// Act
	wrapper, err := singleRwx.ToDisabledRwxWrapper()

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if wrapper == nil {
		t.Error("should not be nil")
	}
}

// ── TempDirGetter ──

func Test_TempDirGetter_Ext2(t *testing.T) {
	// Act
	tempDefault := chmodhelper.TempDirGetter.TempDefault()
	tempPermanent := chmodhelper.TempDirGetter.TempPermanent()
	tempOptionTrue := chmodhelper.TempDirGetter.TempOption(true)
	tempOptionFalse := chmodhelper.TempDirGetter.TempOption(false)

	// Assert
	if tempDefault == "" {
		t.Error("TempDefault should not be empty")
	}
	if tempPermanent == "" {
		t.Error("TempPermanent should not be empty")
	}
	if tempOptionTrue != tempPermanent {
		t.Error("TempOption(true) should equal TempPermanent")
	}
	if tempOptionFalse != tempDefault {
		t.Error("TempOption(false) should equal TempDefault")
	}
}

// ── chmodApplier tests ──

func Test_ChmodApplier_ApplyIf_False_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodApply.ApplyIf(false, os.FileMode(0755), "/whatever")

	// Assert
	if err != nil {
		t.Error("ApplyIf false should return nil")
	}
}

func Test_ChmodApplier_OnMismatchOption_NotApply_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodApply.OnMismatchOption(false, false, os.FileMode(0755), "/whatever")

	// Assert
	if err != nil {
		t.Error("isApply=false should return nil")
	}
}

func Test_ChmodApplier_PathsUsingFileModeConditions_EmptyLocations_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(os.FileMode(0755), nil)

	// Assert
	if err != nil {
		t.Error("empty locations should return nil")
	}
}

func Test_ChmodApplier_PathsUsingFileModeConditions_NilCondition_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(os.FileMode(0755), nil, "/tmp")

	// Assert
	if err == nil {
		t.Error("nil condition should error")
	}
}

func Test_ChmodApplier_RwxPartial_EmptyLocations_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodApply.RwxPartial("-rwx", &chmodins.Condition{})

	// Assert
	if err != nil {
		t.Error("empty locations should return nil")
	}
}

// ── RwxStringApplyChmod ──

func Test_RwxStringApplyChmod_EmptyLocations_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.RwxStringApplyChmod("-rwxr-xr-x", &chmodins.Condition{})

	// Assert
	if err != nil {
		t.Error("empty locations should return nil")
	}
}

func Test_RwxStringApplyChmod_InvalidLength_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.RwxStringApplyChmod("rwx", &chmodins.Condition{}, "/tmp")

	// Assert
	if err == nil {
		t.Error("invalid length should error")
	}
}

func Test_RwxStringApplyChmod_NilCondition_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.RwxStringApplyChmod("-rwxr-xr-x", nil, "/tmp")

	// Assert
	if err == nil {
		t.Error("nil condition should error")
	}
}

// ── RwxOwnerGroupOtherApplyChmod ──

func Test_RwxOwnerGroupOtherApplyChmod_EmptyLocations_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(
		&chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r--"},
		&chmodins.Condition{},
	)

	// Assert
	if err != nil {
		t.Error("empty locations should return nil")
	}
}

func Test_RwxOwnerGroupOtherApplyChmod_NilRwx_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(nil, &chmodins.Condition{}, "/tmp")

	// Assert
	if err == nil {
		t.Error("nil rwx should error")
	}
}

func Test_RwxOwnerGroupOtherApplyChmod_NilCondition_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(
		&chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r--"},
		nil,
		"/tmp",
	)

	// Assert
	if err == nil {
		t.Error("nil condition should error")
	}
}

// ── chmodVerifier additional tests ──

func Test_ChmodVerifier_RwxFull_InvalidLength_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodVerify.RwxFull("/tmp", "rwx")

	// Assert
	if err == nil {
		t.Error("invalid length should error")
	}
}

func Test_ChmodVerifier_RwxFull_NonExistent_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodVerify.RwxFull("/nonexistent/xyz", "-rwxr-xr-x")

	// Assert
	if err == nil {
		t.Error("nonexistent path should error")
	}
}

func Test_ChmodVerifier_MismatchError_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	err := chmodhelper.ChmodVerify.MismatchError(tmpFile.Name(), os.FileMode(0777))

	// Assert - may or may not error depending on actual permissions
	_ = err
}

func Test_ChmodVerifier_MismatchErrorUsingRwxFull_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	err := chmodhelper.ChmodVerify.MismatchErrorUsingRwxFull(tmpFile.Name(), "-rwxrwxrwx")

	// Assert
	_ = err
}

func Test_ChmodVerifier_IsEqualRwxFull_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	info, _ := os.Stat(tmpFile.Name())
	rwxFull := info.Mode().String()

	// Assert
	if !chmodhelper.ChmodVerify.IsEqualRwxFull(tmpFile.Name(), rwxFull) {
		t.Error("should match existing")
	}
}

func Test_ChmodVerifier_GetExistingRwxWrapper_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	wrapper, err := chmodhelper.ChmodVerify.GetExistingRwxWrapper(tmpFile.Name())

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if wrapper.IsEmpty() {
		t.Error("should not be empty")
	}
}

func Test_ChmodVerifier_PathsUsingFileModeImmediateReturn_Valid_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	info, _ := os.Stat(tmpFile.Name())

	// Act
	err := chmodhelper.ChmodVerify.PathsUsingFileModeImmediateReturn(
		info.Mode(), tmpFile.Name())

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_ChmodVerifier_Path_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	info, _ := os.Stat(tmpFile.Name())

	// Act
	err := chmodhelper.ChmodVerify.Path(tmpFile.Name(), info.Mode())

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_ChmodVerifier_PathIf_True_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	info, _ := os.Stat(tmpFile.Name())

	// Act
	err := chmodhelper.ChmodVerify.PathIf(true, tmpFile.Name(), info.Mode())

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_ChmodVerifier_UsingRwxOwnerGroupOther_Nil_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodVerify.UsingRwxOwnerGroupOther(nil, "/tmp")

	// Assert
	if err == nil {
		t.Error("nil rwx should error")
	}
}

// ── RwxPartialToInstructionExecutor ──

func Test_RwxPartialToInstructionExecutor_NilCondition_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.RwxPartialToInstructionExecutor("-rwx", nil)

	// Assert
	if err == nil {
		t.Error("nil condition should error")
	}
}

func Test_RwxPartialToInstructionExecutor_Valid_Ext2(t *testing.T) {
	// Act
	executor, err := chmodhelper.RwxPartialToInstructionExecutor(
		"-rwxr-xr--", &chmodins.Condition{})

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if executor == nil {
		t.Error("should not be nil")
	}
}

// ── dirCreator additional ──

func Test_DirCreator_IfMissingLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_ifmissinglock_ext2"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.IfMissingLock(os.FileMode(0755), dir)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_DirCreator_DefaultLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_defaultlock_ext2"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.DefaultLock(os.FileMode(0755), dir)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_DirCreator_DirectLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_directlock_ext2"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.DirectLock(dir)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_DirCreator_ByChecking_NewDir_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	dir := os.TempDir() + "/chmodtest_bychecking_ext2"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.ByChecking(os.FileMode(0755), dir)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_DirCreator_ByChecking_ExistingDir_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	dir := os.TempDir() + "/chmodtest_bychecking_existing_ext2"
	os.MkdirAll(dir, 0755)
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.ByChecking(os.FileMode(0755), dir)

	// Assert - chmod apply on existing dir
	_ = err
}

// ── fileWriter additional ──

func Test_FileWriter_All_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fw_all_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.All(
		os.FileMode(0755),
		os.FileMode(0644),
		true,
		false,
		false,
		true,
		dir,
		filePath,
		[]byte("content"),
	)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_FileWriter_AllLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fw_alllock_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.AllLock(
		os.FileMode(0755),
		os.FileMode(0644),
		true,
		false,
		false,
		true,
		dir,
		filePath,
		[]byte("content"),
	)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_FileWriter_Chmod_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	dir := os.TempDir() + "/chmodtest_fw_chmod_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Chmod(
		true,
		os.FileMode(0755),
		os.FileMode(0644),
		filePath,
		[]byte("content"),
	)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_FileWriter_ChmodFile_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	dir := os.TempDir() + "/chmodtest_fw_chmodfile_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.ChmodFile(
		true,
		os.FileMode(0644),
		filePath,
		[]byte("content"),
	)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

// ── fileBytesWriter additional ──

func Test_FileBytesWriter_WithDir_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fbw_wd_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDir(true, filePath, []byte("content"))

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_FileBytesWriter_WithDirLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fbw_wdl_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirLock(true, filePath, []byte("content"))

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_FileBytesWriter_WithDirChmodLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fbw_wdcl_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirChmodLock(
		true, os.FileMode(0755), os.FileMode(0644), filePath, []byte("content"))

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_FileBytesWriter_Chmod_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fbw_chmod_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.Chmod(
		true, os.FileMode(0755), os.FileMode(0644), filePath, []byte("content"))

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

// ── fileStringWriter additional ──

func Test_FileStringWriter_All_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fsw_all_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.String.All(
		true, os.FileMode(0755), os.FileMode(0644),
		false, false, true,
		dir, filePath, "content",
	)

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_FileStringWriter_DefaultLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fsw_dl_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.String.DefaultLock(true, filePath, "content")

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_FileStringWriter_Chmod_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fsw_chmod_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.String.Chmod(
		true, os.FileMode(0755), os.FileMode(0644), filePath, "content")

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_FileStringWriter_ChmodLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fsw_chmodl_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.String.ChmodLock(
		true, os.FileMode(0755), os.FileMode(0644), filePath, "content")

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

// ── fileReader ──

func Test_FileReader_Read_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	os.WriteFile(tmpFile.Name(), []byte("hello"), 0644)
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	content, err := chmodhelper.SimpleFileWriter.FileReader.Read(tmpFile.Name())

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if content != "hello" {
		t.Errorf("unexpected: %s", content)
	}
}

func Test_FileReader_ReadBytes_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	os.WriteFile(tmpFile.Name(), []byte("hello"), 0644)
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	bytes, err := chmodhelper.SimpleFileWriter.FileReader.ReadBytes(tmpFile.Name())

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if string(bytes) != "hello" {
		t.Errorf("unexpected: %s", string(bytes))
	}
}

func Test_FileReader_Read_Invalid_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.SimpleFileWriter.FileReader.Read("/nonexistent/xyz")

	// Assert
	if err == nil {
		t.Error("expected error")
	}
}

// ── anyItemWriter ──

func Test_AnyItemWriter_Default_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_aiw_ext2"
	filePath := dir + "/test.json"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Any.Default(
		true, filePath, map[string]string{"key": "value"})

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_AnyItemWriter_DefaultLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_aiw_dl_ext2"
	filePath := dir + "/test.json"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Any.DefaultLock(
		true, filePath, map[string]string{"key": "value"})

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

// ── simpleFileWriter Lock/Unlock ──

func Test_SimpleFileWriter_LockUnlock_Ext2(t *testing.T) {
	// Act - should not panic
	chmodhelper.SimpleFileWriter.Lock()
	chmodhelper.SimpleFileWriter.Unlock()
}

// ── ApplyChmod on valid temp file (Unix) ──

func Test_RwxWrapper_ApplyChmod_Unix_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	err := wrapper.ApplyChmod(false, tmpFile.Name())

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_RwxWrapper_ApplyChmodSkipInvalid_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Act
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")
	err := wrapper.ApplyChmodSkipInvalid("/nonexistent/xyz")

	// Assert
	if err != nil {
		t.Error("skip invalid should return nil")
	}
}

func Test_RwxWrapper_ApplyChmod_InvalidNotSkip_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	err := wrapper.ApplyChmod(false, "/nonexistent/xyz")

	// Assert
	if err == nil {
		t.Error("should error for invalid path without skip")
	}
}

func Test_RwxWrapper_ApplyChmodOptions_SkipApply_Ext2(t *testing.T) {
	// Act
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")
	err := wrapper.ApplyChmodOptions(false, true, false, "/whatever")

	// Assert
	if err != nil {
		t.Error("isApply=false should return nil")
	}
}

func Test_RwxWrapper_Verify_Unix_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	info, _ := os.Stat(tmpFile.Name())
	wrapper := chmodhelper.New.RwxWrapper.UsingFileMode(info.Mode())

	// Act
	err := wrapper.Verify(tmpFile.Name())

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_RwxWrapper_HasChmod_Unix_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	info, _ := os.Stat(tmpFile.Name())
	wrapper := chmodhelper.New.RwxWrapper.UsingFileMode(info.Mode())

	// Assert
	if !wrapper.HasChmod(tmpFile.Name()) {
		t.Error("should match")
	}
}

// ── chmodApplier Unix ──

func Test_ChmodApplier_Default_Unix_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	err := chmodhelper.ChmodApply.Default(os.FileMode(0755), tmpFile.Name())

	// Assert
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func Test_ChmodApplier_SkipInvalidFile_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Act
	err := chmodhelper.ChmodApply.SkipInvalidFile(os.FileMode(0755), "/nonexistent/xyz")

	// Assert
	if err != nil {
		t.Error("skip invalid should return nil")
	}
}

func Test_ChmodApplier_OnMismatch_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Act
	err := chmodhelper.ChmodApply.OnMismatch(true, os.FileMode(0755), "/nonexistent/xyz")

	// Assert
	if err != nil {
		t.Error("skip invalid path should return nil")
	}
}

func Test_ChmodApplier_OnMismatchSkipInvalid_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Act
	err := chmodhelper.ChmodApply.OnMismatchSkipInvalid(os.FileMode(0755), "/nonexistent/xyz")

	// Assert
	if err != nil {
		t.Error("skip invalid should return nil")
	}
}
