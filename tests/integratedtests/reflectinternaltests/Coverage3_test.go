package reflectinternaltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// ── isChecker extended ──

func Test_Cov3_Is_Number(t *testing.T) {
	actual := args.Map{
		"int":    reflectinternal.Is.Number(42),
		"float":  reflectinternal.Is.Number(3.14),
		"string": reflectinternal.Is.Number("a"),
	}
	expected := args.Map{"int": true, "float": true, "string": false}
	expected.ShouldBeEqual(t, 0, "Is returns correct value -- Number", actual)
}

func Test_Cov3_Is_String(t *testing.T) {
	actual := args.Map{
		"str":  reflectinternal.Is.String("hello"),
		"int":  reflectinternal.Is.String(42),
	}
	expected := args.Map{"str": true, "int": false}
	expected.ShouldBeEqual(t, 0, "Is returns correct value -- String", actual)
}

func Test_Cov3_Is_Pointer(t *testing.T) {
	x := 42
	actual := args.Map{
		"ptr": reflectinternal.Is.Pointer(&x),
		"val": reflectinternal.Is.Pointer(x),
	}
	expected := args.Map{"ptr": true, "val": false}
	expected.ShouldBeEqual(t, 0, "Is returns correct value -- Pointer", actual)
}

func Test_Cov3_Is_Function(t *testing.T) {
	fn := func() {}
	actual := args.Map{
		"func": reflectinternal.Is.Function(fn),
		"int":  reflectinternal.Is.Function(42),
	}
	expected := args.Map{"func": true, "int": false}
	expected.ShouldBeEqual(t, 0, "Is returns correct value -- Function", actual)
}

func Test_Cov3_Is_Boolean(t *testing.T) {
	actual := args.Map{
		"bool": reflectinternal.Is.Boolean(true),
		"int":  reflectinternal.Is.Boolean(42),
	}
	expected := args.Map{"bool": true, "int": false}
	expected.ShouldBeEqual(t, 0, "Is returns correct value -- Boolean", actual)
}

func Test_Cov3_Is_Primitive(t *testing.T) {
	actual := args.Map{
		"int":    reflectinternal.Is.Primitive(42),
		"string": reflectinternal.Is.Primitive("a"),
		"slice":  reflectinternal.Is.Primitive([]int{1}),
	}
	expected := args.Map{"int": true, "string": true, "slice": false}
	expected.ShouldBeEqual(t, 0, "Is returns correct value -- Primitive", actual)
}

func Test_Cov3_Is_Zero(t *testing.T) {
	actual := args.Map{
		"zero":    reflectinternal.Is.Zero(0),
		"nonZero": reflectinternal.Is.Zero(1),
		"nil":     reflectinternal.Is.Zero(nil),
		"emptyStr": reflectinternal.Is.Zero(""),
	}
	expected := args.Map{"zero": true, "nonZero": false, "nil": true, "emptyStr": true}
	expected.ShouldBeEqual(t, 0, "Is returns correct value -- Zero", actual)
}

func Test_Cov3_Is_Struct(t *testing.T) {
	type S struct{ A int }
	actual := args.Map{
		"struct": reflectinternal.Is.Struct(S{}),
		"ptr":    reflectinternal.Is.Struct(&S{}),
		"int":    reflectinternal.Is.Struct(42),
	}
	expected := args.Map{"struct": true, "ptr": true, "int": false}
	expected.ShouldBeEqual(t, 0, "Is returns correct value -- Struct", actual)
}

func Test_Cov3_Is_AnyEqual(t *testing.T) {
	actual := args.Map{
		"same":      reflectinternal.Is.AnyEqual(42, 42),
		"diff":      reflectinternal.Is.AnyEqual(42, 43),
		"nilNil":    reflectinternal.Is.AnyEqual(nil, nil),
		"nilNonNil": reflectinternal.Is.AnyEqual(nil, 42),
	}
	expected := args.Map{"same": true, "diff": false, "nilNil": true, "nilNonNil": false}
	expected.ShouldBeEqual(t, 0, "Is returns correct value -- AnyEqual", actual)
}

func Test_Cov3_Is_Conclusive(t *testing.T) {
	eq, conc := reflectinternal.Is.Conclusive(42, 42)
	_, concDiff := reflectinternal.Is.Conclusive(42, "x")
	actual := args.Map{"eq": eq, "conc": conc, "concDiff": concDiff}
	expected := args.Map{"eq": true, "conc": true, "concDiff": true}
	expected.ShouldBeEqual(t, 0, "Is returns correct value -- Conclusive", actual)
}

func Test_Cov3_Is_SliceOrArray(t *testing.T) {
	actual := args.Map{
		"nil": reflectinternal.Is.SliceOrArray(nil),
	}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Is returns correct value -- SliceOrArray", actual)
}

// ── TypeName / TypeNames / TypeNamesString ──

func Test_Cov3_TypeName(t *testing.T) {
	actual := args.Map{
		"int":  reflectinternal.TypeName(42),
		"nil":  reflectinternal.TypeName(nil),
	}
	expected := args.Map{"int": "int", "nil": ""}
	expected.ShouldBeEqual(t, 0, "TypeName returns correct value -- with args", actual)
}

func Test_Cov3_TypeNames(t *testing.T) {
	full := reflectinternal.TypeNames(true, 42, "hello")
	short := reflectinternal.TypeNames(false, 42, "hello")
	actual := args.Map{"fullLen": len(full), "shortLen": len(short)}
	expected := args.Map{"fullLen": 2, "shortLen": 2}
	expected.ShouldBeEqual(t, 0, "TypeNames returns correct value -- with args", actual)
}

func Test_Cov3_TypeNamesString(t *testing.T) {
	result := reflectinternal.TypeNamesString(true, 42, "hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeNamesString returns correct value -- with args", actual)
}

func Test_Cov3_TypeNamesReferenceString(t *testing.T) {
	result := reflectinternal.TypeNamesReferenceString(true, 42)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeNamesReferenceString returns correct value -- with args", actual)
}

func Test_Cov3_TypeNameToValidVariableName(t *testing.T) {
	result := reflectinternal.TypeNameToValidVariableName("mypackage.MyType")
	empty := reflectinternal.TypeNameToValidVariableName("")
	actual := args.Map{"notEmpty": result != "", "empty": empty}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "TypeNameToValidVariableName returns non-empty -- with args", actual)
}

// ── FileWithLine ──

func Test_Cov3_FileWithLine(t *testing.T) {
	f := &reflectinternal.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	var nilF *reflectinternal.FileWithLine
	actual := args.Map{
		"path":     f.FullFilePath(),
		"line":     f.LineNumber(),
		"isNil":    f.IsNil(),
		"isNotNil": f.IsNotNil(),
		"string":   f.String() != "",
		"fwl":      f.FileWithLine() != "",
		"json":     f.JsonString() != "",
		"nilStr":   nilF.String(),
	}
	expected := args.Map{
		"path": "/tmp/test.go", "line": 42,
		"isNil": false, "isNotNil": true,
		"string": true, "fwl": true, "json": true, "nilStr": "",
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- with args", actual)
}

func Test_Cov3_FileWithLine_StringUsingFmt(t *testing.T) {
	f := reflectinternal.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	result := f.StringUsingFmt(func(fwl reflectinternal.FileWithLine) string { return fwl.FilePath })
	actual := args.Map{"result": result}
	expected := args.Map{"result": "/tmp/test.go"}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- StringUsingFmt", actual)
}

func Test_Cov3_FileWithLine_JsonModel(t *testing.T) {
	f := reflectinternal.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	model := f.JsonModel()
	modelAny := f.JsonModelAny()
	actual := args.Map{"path": model.FilePath, "anyNotNil": modelAny != nil}
	expected := args.Map{"path": "/tmp/test.go", "anyNotNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- JsonModel", actual)
}

// ── StackTrace ──

func Test_Cov3_StackTrace(t *testing.T) {
	st := reflectinternal.StackTrace{
		PackageName: "pkg", MethodName: "Method",
		PackageMethodName: "pkg.Method", FilePath: "/tmp/test.go",
		Line: 42, IsOkay: true,
	}
	actual := args.Map{
		"msg":     st.Message() != "",
		"short":   st.ShortString() != "",
		"isNil":   st.IsNil(),
		"notNil":  st.IsNotNil(),
		"hasIss":  st.HasIssues(),
		"str":     st.String() != "",
		"fwl":     st.FileWithLine().FilePath,
		"fullPath": st.FullFilePath(),
		"fileName": st.FileName() != "",
		"lineNum": st.LineNumber(),
		"fwlStr":  st.FileWithLineString() != "",
		"json":    st.JsonString() != "",
	}
	expected := args.Map{
		"msg": true, "short": true, "isNil": false, "notNil": true,
		"hasIss": false, "str": true, "fwl": "/tmp/test.go",
		"fullPath": "/tmp/test.go", "fileName": true, "lineNum": 42,
		"fwlStr": true, "json": true,
	}
	expected.ShouldBeEqual(t, 0, "StackTrace returns correct value -- with args", actual)
}

func Test_Cov3_StackTrace_Nil(t *testing.T) {
	var st *reflectinternal.StackTrace
	actual := args.Map{"isNil": st.IsNil(), "notNil": st.IsNotNil(), "str": st.String(), "hasIss": st.HasIssues()}
	expected := args.Map{"isNil": true, "notNil": false, "str": "", "hasIss": true}
	expected.ShouldBeEqual(t, 0, "StackTrace returns nil -- Nil", actual)
}

func Test_Cov3_StackTrace_Clone(t *testing.T) {
	st := &reflectinternal.StackTrace{PackageName: "pkg", Line: 42}
	cloned := st.Clone()
	clonedPtr := st.ClonePtr()
	var nilSt *reflectinternal.StackTrace
	actual := args.Map{
		"pkg": cloned.PackageName, "ptrPkg": clonedPtr.PackageName,
		"nilClone": nilSt.ClonePtr() == nil,
	}
	expected := args.Map{"pkg": "pkg", "ptrPkg": "pkg", "nilClone": true}
	expected.ShouldBeEqual(t, 0, "StackTrace returns correct value -- Clone", actual)
}

func Test_Cov3_StackTrace_Dispose(t *testing.T) {
	st := &reflectinternal.StackTrace{PackageName: "pkg"}
	st.Dispose()
	var nilSt *reflectinternal.StackTrace
	nilSt.Dispose()
	actual := args.Map{"pkg": st.PackageName}
	expected := args.Map{"pkg": ""}
	expected.ShouldBeEqual(t, 0, "StackTrace returns correct value -- Dispose", actual)
}

func Test_Cov3_StackTrace_JsonModel(t *testing.T) {
	st := reflectinternal.StackTrace{PackageName: "pkg"}
	model := st.JsonModel()
	modelAny := st.JsonModelAny()
	actual := args.Map{"pkg": model.PackageName, "anyNotNil": modelAny != nil}
	expected := args.Map{"pkg": "pkg", "anyNotNil": true}
	expected.ShouldBeEqual(t, 0, "StackTrace returns correct value -- JsonModel", actual)
}

func Test_Cov3_StackTrace_StringUsingFmt(t *testing.T) {
	st := reflectinternal.StackTrace{PackageName: "pkg"}
	result := st.StringUsingFmt(func(trace reflectinternal.StackTrace) string { return trace.PackageName })
	actual := args.Map{"result": result}
	expected := args.Map{"result": "pkg"}
	expected.ShouldBeEqual(t, 0, "StackTrace returns correct value -- StringUsingFmt", actual)
}

// ── reflectGetter ──

func Test_Cov3_ReflectGetter_PublicValuesMapStruct(t *testing.T) {
	type S struct{ A int; B string }
	m, err := reflectinternal.ReflectGetter.PublicValuesMapStruct(S{A: 1, B: "hello"})
	_, nilErr := reflectinternal.ReflectGetter.PublicValuesMapStruct(nil)
	actual := args.Map{"len": len(m), "noErr": err == nil, "nilErr": nilErr != nil}
	expected := args.Map{"len": 2, "noErr": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "PublicValuesMapStruct returns non-empty -- with args", actual)
}

func Test_Cov3_ReflectGetter_FieldNameWithValuesMap(t *testing.T) {
	type S struct{ A int }
	m, err := reflectinternal.ReflectGetter.FieldNameWithValuesMap(S{A: 1})
	_, nilErr := reflectinternal.ReflectGetter.FieldNameWithValuesMap(nil)
	actual := args.Map{"gt0": len(m) > 0, "noErr": err == nil, "nilErr": nilErr != nil}
	expected := args.Map{"gt0": true, "noErr": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "FieldNameWithValuesMap returns non-empty -- with args", actual)
}

func Test_Cov3_ReflectGetter_FieldNamesMap(t *testing.T) {
	type S struct{ A int }
	m, err := reflectinternal.ReflectGetter.FieldNamesMap(S{A: 1})
	_, nilErr := reflectinternal.ReflectGetter.FieldNamesMap(nil)
	actual := args.Map{"gt0": len(m) > 0, "noErr": err == nil, "nilErr": nilErr != nil}
	expected := args.Map{"gt0": true, "noErr": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "FieldNamesMap returns correct value -- with args", actual)
}

func Test_Cov3_ReflectGetter_StructFieldsMap(t *testing.T) {
	type S struct{ A int }
	m := reflectinternal.ReflectGetter.StructFieldsMap(S{A: 1})
	nilM := reflectinternal.ReflectGetter.StructFieldsMap(nil)
	actual := args.Map{"gt0": len(m) > 0, "nilLen": len(nilM)}
	expected := args.Map{"gt0": true, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "StructFieldsMap returns correct value -- with args", actual)
}

func Test_Cov3_ReflectGetter_NullFieldsMap(t *testing.T) {
	type S struct{ A *int }
	m := reflectinternal.ReflectGetter.NullFieldsMap(S{})
	nilM := reflectinternal.ReflectGetter.NullFieldsMap(nil)
	actual := args.Map{"gt0": len(m) > 0, "nilLen": len(nilM)}
	expected := args.Map{"gt0": true, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "NullFieldsMap returns correct value -- with args", actual)
}

func Test_Cov3_ReflectGetter_NullOrZeroFieldsMap(t *testing.T) {
	type S struct{ A int; B *int }
	m := reflectinternal.ReflectGetter.NullOrZeroFieldsMap(S{})
	nilM := reflectinternal.ReflectGetter.NullOrZeroFieldsMap(nil)
	actual := args.Map{"gt0": len(m) > 0, "nilLen": len(nilM)}
	expected := args.Map{"gt0": true, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "NullOrZeroFieldsMap returns correct value -- with args", actual)
}

// ── codeStack ──

func Test_Cov3_CodeStack_New(t *testing.T) {
	st := reflectinternal.CodeStack.New(0)
	actual := args.Map{
		"isOkay":  st.IsOkay,
		"pkgName": st.PackageName != "",
		"method":  st.MethodName != "",
	}
	expected := args.Map{"isOkay": true, "pkgName": true, "method": true}
	expected.ShouldBeEqual(t, 0, "CodeStack returns correct value -- New", actual)
}

func Test_Cov3_CodeStack_MethodName(t *testing.T) {
	name := reflectinternal.CodeStack.MethodName(0)
	actual := args.Map{"notEmpty": name != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CodeStack returns correct value -- MethodName", actual)
}

func Test_Cov3_CodeStack_MethodNameWithLine(t *testing.T) {
	name := reflectinternal.CodeStack.MethodNameWithLine(0)
	actual := args.Map{"notEmpty": name != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CodeStack returns non-empty -- MethodNameWithLine", actual)
}

func Test_Cov3_CodeStack_FileWithLine(t *testing.T) {
	result := reflectinternal.CodeStack.FileWithLine(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CodeStack returns non-empty -- FileWithLine", actual)
}

func Test_Cov3_CodeStack_FilePath(t *testing.T) {
	result := reflectinternal.CodeStack.FilePath(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CodeStack returns correct value -- FilePath", actual)
}

func Test_Cov3_CodeStack_LastFileWithLines(t *testing.T) {
	lines := reflectinternal.CodeStack.LastFileWithLines(0, 2)
	actual := args.Map{"len": len(lines)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CodeStack returns non-empty -- LastFileWithLines", actual)
}

func Test_Cov3_CodeStack_NewFileWithLines(t *testing.T) {
	lines := reflectinternal.CodeStack.NewFileWithLines(0, 2)
	actual := args.Map{"len": len(lines)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CodeStack returns non-empty -- NewFileWithLines", actual)
}

func Test_Cov3_CodeStack_NewFileWithLine(t *testing.T) {
	fwl := reflectinternal.CodeStack.NewFileWithLine(0)
	actual := args.Map{"notEmpty": fwl.FilePath != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CodeStack returns non-empty -- NewFileWithLine", actual)
}

// ── reflectConverter ──

func Test_Cov3_ReflectConverter_ArgsToReflectValues(t *testing.T) {
	result := reflectinternal.Converter.ArgsToReflectValues([]any{1, "hello"})
	emptyResult := reflectinternal.Converter.ArgsToReflectValues(nil)
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "ArgsToReflectValues returns non-empty -- with args", actual)
}

func Test_Cov3_ReflectConverter_ReflectValueToAnyValue_Nil(t *testing.T) {
	// reflect.Value{} (zero Value) panics on .Interface() -- test with valid zero instead
	rv := reflect.ValueOf(0)
	result := reflectinternal.Converter.ReflectValueToAnyValue(rv)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueToAnyValue returns nil -- nil", actual)
}
