package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// Result — comprehensive coverage
// ═══════════════════════════════════════════

func Test_Cov5_Result_Map(t *testing.T) {
	r := corejson.NewPtr("hello")
	m := r.Map()
	actual := args.Map{"hasBytes": m["Bytes"] != "", "hasType": m["Type"] != ""}
	expected := args.Map{"hasBytes": true, "hasType": true}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- Map", actual)
}

func Test_Cov5_Result_Map_WithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail"), TypeName: "test"}
	m := r.Map()
	actual := args.Map{"hasErr": m["Error"] != ""}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Result returns error -- Map with error", actual)
}

func Test_Cov5_Result_Map_Nil(t *testing.T) {
	var r *corejson.Result
	m := r.Map()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Result returns nil -- Map nil", actual)
}

func Test_Cov5_Result_SafeString(t *testing.T) {
	r := corejson.NewPtr("hello")
	s := r.SafeString()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- SafeString", actual)
}

func Test_Cov5_Result_Length(t *testing.T) {
	r := corejson.NewPtr("hello")
	var nilR *corejson.Result
	actual := args.Map{"len": r.Length(), "nilLen": nilR.Length()}
	expected := args.Map{"len": r.Length(), "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- Length", actual)
}

func Test_Cov5_Result_ErrorString(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail")}
	rOk := corejson.NewPtr("ok")
	actual := args.Map{"errStr": r.ErrorString() != "", "okStr": rOk.ErrorString()}
	expected := args.Map{"errStr": true, "okStr": ""}
	expected.ShouldBeEqual(t, 0, "Result returns error -- ErrorString", actual)
}

func Test_Cov5_Result_IsErrorEqual(t *testing.T) {
	e := errors.New("fail")
	r := &corejson.Result{Bytes: []byte("x"), Error: e}
	actual := args.Map{
		"equal":    r.IsErrorEqual(errors.New("fail")),
		"notEqual": r.IsErrorEqual(errors.New("other")),
		"nilBoth":  corejson.NewPtr("ok").IsErrorEqual(nil),
		"nilOne":   r.IsErrorEqual(nil),
	}
	expected := args.Map{"equal": true, "notEqual": false, "nilBoth": true, "nilOne": false}
	expected.ShouldBeEqual(t, 0, "Result returns error -- IsErrorEqual", actual)
}

func Test_Cov5_Result_String(t *testing.T) {
	r := corejson.NewPtr("hello")
	s := r.NonPtr().String()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- String", actual)
}

func Test_Cov5_Result_SafeNonIssueBytes(t *testing.T) {
	r := corejson.NewPtr("hello")
	b := r.SafeNonIssueBytes()
	rErr := &corejson.Result{Error: errors.New("e")}
	bErr := rErr.SafeNonIssueBytes()
	actual := args.Map{"len": len(b) > 0, "errLen": len(bErr)}
	expected := args.Map{"len": true, "errLen": 0}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- SafeNonIssueBytes", actual)
}

func Test_Cov5_Result_Values(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{
		"valuesLen": len(r.Values()) > 0,
		"safeLen":   len(r.SafeValues()) > 0,
	}
	expected := args.Map{"valuesLen": true, "safeLen": true}
	expected.ShouldBeEqual(t, 0, "Result returns non-empty -- Values", actual)
}

func Test_Cov5_Result_Raw(t *testing.T) {
	r := corejson.NewPtr("hello")
	b, err := r.Raw()
	rs, rsErr := r.RawString()
	actual := args.Map{
		"bLen": len(b) > 0, "noErr": err == nil,
		"rsLen": len(rs) > 0, "rsNoErr": rsErr == nil,
	}
	expected := args.Map{"bLen": true, "noErr": true, "rsLen": true, "rsNoErr": true}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- Raw", actual)
}

func Test_Cov5_Result_Raw_Nil(t *testing.T) {
	var r *corejson.Result
	_, err := r.Raw()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Result returns nil -- Raw nil", actual)
}

func Test_Cov5_Result_RawPrettyString(t *testing.T) {
	r := corejson.NewPtr(map[string]string{"a": "1"})
	s, err := r.RawPrettyString()
	actual := args.Map{"notEmpty": s != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- RawPrettyString", actual)
}

func Test_Cov5_Result_RawErrString(t *testing.T) {
	r := corejson.NewPtr("hello")
	b, errMsg := r.RawErrString()
	actual := args.Map{"bLen": len(b) > 0, "errMsg": errMsg}
	expected := args.Map{"bLen": true, "errMsg": ""}
	expected.ShouldBeEqual(t, 0, "Result returns error -- RawErrString", actual)
}

func Test_Cov5_Result_MeaningfulError(t *testing.T) {
	r := corejson.NewPtr("hello")
	rErr := &corejson.Result{Error: errors.New("fail"), Bytes: []byte("x")}
	rEmpty := &corejson.Result{}
	var nilR *corejson.Result
	actual := args.Map{
		"noErr":    r.MeaningfulError() == nil,
		"hasErr":   rErr.MeaningfulError() != nil,
		"emptyErr": rEmpty.MeaningfulError() != nil,
		"nilErr":   nilR.MeaningfulError() != nil,
		"errMsg":   r.MeaningfulErrorMessage(),
	}
	expected := args.Map{
		"noErr": true, "hasErr": true, "emptyErr": true, "nilErr": true, "errMsg": "",
	}
	expected.ShouldBeEqual(t, 0, "Result returns error -- MeaningfulError", actual)
}

func Test_Cov5_Result_PrettyJsonStringOrErrString(t *testing.T) {
	r := corejson.NewPtr("hello")
	rErr := &corejson.Result{Error: errors.New("fail"), Bytes: []byte("x")}
	var nilR *corejson.Result
	actual := args.Map{
		"pretty":  r.PrettyJsonStringOrErrString() != "",
		"errStr":  rErr.PrettyJsonStringOrErrString() != "",
		"nilStr":  nilR.PrettyJsonStringOrErrString() != "",
	}
	expected := args.Map{"pretty": true, "errStr": true, "nilStr": true}
	expected.ShouldBeEqual(t, 0, "Result returns error -- PrettyJsonStringOrErrString", actual)
}

func Test_Cov5_Result_IsEmpty(t *testing.T) {
	r := corejson.NewPtr("hello")
	var nilR *corejson.Result
	actual := args.Map{
		"isEmpty":    r.IsEmpty(),
		"isEmptyErr": r.IsEmptyError(),
		"isAnyNull":  r.IsAnyNull(),
		"hasItems":   r.HasSafeItems(),
		"hasBytes":   r.HasBytes(),
		"hasJson":    r.HasJson(),
		"isEmptyJ":   r.IsEmptyJson(),
		"nilEmpty":   nilR.IsEmpty(),
		"nilAnyNull": nilR.IsAnyNull(),
		"hasAny":     r.HasAnyItem(),
	}
	expected := args.Map{
		"isEmpty": false, "isEmptyErr": true, "isAnyNull": false,
		"hasItems": true, "hasBytes": true, "hasJson": true,
		"isEmptyJ": false, "nilEmpty": true, "nilAnyNull": true,
		"hasAny": true,
	}
	expected.ShouldBeEqual(t, 0, "Result returns empty -- IsEmpty methods", actual)
}

func Test_Cov5_Result_Clone(t *testing.T) {
	r := corejson.NewPtr("hello")
	cloned := r.Clone(true)
	clonedShallow := r.Clone(false)
	clonedPtr := r.ClonePtr(true)
	var nilR *corejson.Result
	clonedIf := r.NonPtr().CloneIf(true, true)
	noClone := r.NonPtr().CloneIf(false, false)
	actual := args.Map{
		"clonedLen":    cloned.Length() > 0,
		"shallowLen":   clonedShallow.Length() > 0,
		"cpLen":        clonedPtr.Length() > 0,
		"nilClone":     nilR.ClonePtr(true) == nil,
		"cloneIfLen":   clonedIf.Length() > 0,
		"noCloneLen":   noClone.Length() > 0,
	}
	expected := args.Map{
		"clonedLen": true, "shallowLen": true, "cpLen": true,
		"nilClone": true, "cloneIfLen": true, "noCloneLen": true,
	}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- Clone", actual)
}

func Test_Cov5_Result_CloneError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail"), Bytes: []byte("x")}
	rOk := corejson.NewPtr("ok")
	actual := args.Map{
		"hasCloneErr": r.CloneError() != nil,
		"noCloneErr":  rOk.CloneError() == nil,
	}
	expected := args.Map{"hasCloneErr": true, "noCloneErr": true}
	expected.ShouldBeEqual(t, 0, "Result returns error -- CloneError", actual)
}

func Test_Cov5_Result_PtrNonPtr(t *testing.T) {
	r := corejson.NewPtr("hello")
	np := r.NonPtr()
	p := np.Ptr()
	tp := np.ToPtr()
	tnp := np.ToNonPtr()
	actual := args.Map{
		"npLen": np.Length() > 0,
		"pNN":   p != nil,
		"tpNN":  tp != nil,
		"tnpLen": tnp.Length() > 0,
	}
	expected := args.Map{"npLen": true, "pNN": true, "tpNN": true, "tnpLen": true}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- Ptr/NonPtr", actual)
}

func Test_Cov5_Result_NonPtr_Nil(t *testing.T) {
	var r *corejson.Result
	np := r.NonPtr()
	actual := args.Map{"hasErr": np.Error != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Result returns nil -- NonPtr nil", actual)
}

func Test_Cov5_Result_IsEqual(t *testing.T) {
	r1 := corejson.New("hello")
	r2 := corejson.New("hello")
	r3 := corejson.New("world")
	actual := args.Map{
		"equal":    r1.IsEqual(r2),
		"notEqual": r1.IsEqual(r3),
	}
	expected := args.Map{"equal": true, "notEqual": false}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- IsEqual", actual)
}

func Test_Cov5_Result_IsEqualPtr(t *testing.T) {
	r1 := corejson.NewPtr("hello")
	r2 := corejson.NewPtr("hello")
	r3 := corejson.NewPtr("world")
	var nilR *corejson.Result
	actual := args.Map{
		"equal":    r1.IsEqualPtr(r2),
		"notEqual": r1.IsEqualPtr(r3),
		"samePtr":  r1.IsEqualPtr(r1),
		"nilBoth":  nilR.IsEqualPtr(nil),
		"nilOne":   nilR.IsEqualPtr(r1),
	}
	expected := args.Map{
		"equal": true, "notEqual": false, "samePtr": true,
		"nilBoth": true, "nilOne": false,
	}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- IsEqualPtr", actual)
}

func Test_Cov5_Result_Serialize(t *testing.T) {
	r := corejson.NewPtr("hello")
	b, err := r.Serialize()
	var nilR *corejson.Result
	_, nilErr := nilR.Serialize()
	rErr := &corejson.Result{Error: errors.New("fail")}
	_, errErr := rErr.Serialize()
	actual := args.Map{
		"bLen": len(b) > 0, "noErr": err == nil,
		"nilErr": nilErr != nil, "errErr": errErr != nil,
	}
	expected := args.Map{"bLen": true, "noErr": true, "nilErr": true, "errErr": true}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- Serialize", actual)
}

func Test_Cov5_Result_JsonModel(t *testing.T) {
	r := corejson.NewPtr("hello")
	jm := r.JsonModel()
	jma := r.JsonModelAny()
	var nilR *corejson.Result
	nilJM := nilR.JsonModel()
	actual := args.Map{
		"jmLen":     jm.Length() > 0,
		"jmaNN":     jma != nil,
		"nilJMErr":  nilJM.Error != nil,
	}
	expected := args.Map{"jmLen": true, "jmaNN": true, "nilJMErr": true}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- JsonModel", actual)
}

func Test_Cov5_Result_BytesError(t *testing.T) {
	r := corejson.NewPtr("hello")
	be := r.BytesError()
	var nilR *corejson.Result
	actual := args.Map{
		"beNN":    be != nil,
		"bLen":    len(be.Bytes) > 0,
		"nilBE":   nilR.BytesError() == nil,
	}
	expected := args.Map{"beNN": true, "bLen": true, "nilBE": true}
	expected.ShouldBeEqual(t, 0, "Result returns error -- BytesError", actual)
}

func Test_Cov5_Result_Dispose(t *testing.T) {
	r := corejson.NewPtr("hello")
	r.Dispose()
	var nilR *corejson.Result
	nilR.Dispose() // should not panic
	actual := args.Map{"bytes": r.Bytes == nil, "err": r.Error == nil}
	expected := args.Map{"bytes": true, "err": true}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- Dispose", actual)
}

func Test_Cov5_Result_CombineErrorWithRef(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail"), Bytes: []byte("x")}
	s := r.CombineErrorWithRefString("ref1", "ref2")
	e := r.CombineErrorWithRefError("ref1")
	rOk := corejson.NewPtr("ok")
	actual := args.Map{
		"sLen":      len(s) > 0,
		"eNN":       e != nil,
		"okEmpty":   rOk.CombineErrorWithRefString("ref"),
		"okNilErr":  rOk.CombineErrorWithRefError("ref") == nil,
	}
	expected := args.Map{"sLen": true, "eNN": true, "okEmpty": "", "okNilErr": true}
	expected.ShouldBeEqual(t, 0, "Result returns error -- CombineErrorWithRef", actual)
}

// ═══════════════════════════════════════════
// Serializer
// ═══════════════════════════════════════════

func Test_Cov5_Serialize_Various(t *testing.T) {
	sr := corejson.Serialize.FromString("hello")
	si := corejson.Serialize.FromInteger(42)
	si64 := corejson.Serialize.FromInteger64(100)
	sb := corejson.Serialize.FromBool(true)
	sints := corejson.Serialize.FromIntegers([]int{1, 2})
	sfb := corejson.Serialize.FromBytes([]byte(`"hi"`))
	sfs := corejson.Serialize.FromStrings([]string{"a", "b"})
	sfss := corejson.Serialize.FromStringsSpread("a", "b")
	ssa := corejson.Serialize.StringsApply([]string{"a"})
	actual := args.Map{
		"srHas": sr.HasBytes(), "siHas": si.HasBytes(), "si64Has": si64.HasBytes(),
		"sbHas": sb.HasBytes(), "sintsHas": sints.HasBytes(), "sfbHas": sfb.HasBytes(),
		"sfsHas": sfs.HasBytes(), "sfssHas": sfss.HasBytes(), "ssaHas": ssa.HasBytes(),
	}
	expected := args.Map{
		"srHas": true, "siHas": true, "si64Has": true,
		"sbHas": true, "sintsHas": true, "sfbHas": true,
		"sfsHas": true, "sfssHas": true, "ssaHas": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize returns correct value -- various", actual)
}

func Test_Cov5_Serialize_UsingAny(t *testing.T) {
	r := corejson.Serialize.UsingAny("hello")
	rp := corejson.Serialize.UsingAnyPtr("hello")
	actual := args.Map{"rHas": r.HasBytes(), "rpHas": rp.HasBytes()}
	expected := args.Map{"rHas": true, "rpHas": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns correct value -- UsingAny", actual)
}

func Test_Cov5_Serialize_ToString(t *testing.T) {
	s := corejson.Serialize.ToString("hello")
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns correct value -- ToString", actual)
}

func Test_Cov5_Serialize_ToBytes(t *testing.T) {
	b, err := corejson.Serialize.ToBytesErr("hello")
	bSwallow := corejson.Serialize.ToBytesSwallowErr("hello")
	bSafe := corejson.Serialize.ToSafeBytesSwallowErr("hello")
	actual := args.Map{
		"bLen": len(b) > 0, "noErr": err == nil,
		"swallowLen": len(bSwallow) > 0, "safeLen": len(bSafe) > 0,
	}
	expected := args.Map{"bLen": true, "noErr": true, "swallowLen": true, "safeLen": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns correct value -- ToBytes", actual)
}

func Test_Cov5_Serialize_ToStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToStringErr("hello")
	actual := args.Map{"notEmpty": s != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns error -- ToStringErr", actual)
}

func Test_Cov5_Serialize_ToPrettyStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToPrettyStringErr(map[string]string{"a": "1"})
	actual := args.Map{"notEmpty": s != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns error -- ToPrettyStringErr", actual)
}

func Test_Cov5_Serialize_ToPrettyStringIncludingErr(t *testing.T) {
	s := corejson.Serialize.ToPrettyStringIncludingErr("hello")
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns error -- ToPrettyStringIncludingErr", actual)
}

func Test_Cov5_Serialize_Pretty(t *testing.T) {
	s := corejson.Serialize.Pretty(map[string]string{"a": "1"})
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns correct value -- Pretty", actual)
}

// ═══════════════════════════════════════════
// ResultsCollection
// ═══════════════════════════════════════════

func Test_Cov5_ResultsCollection_Basic(t *testing.T) {
	rc := corejson.Empty.ResultsCollection()
	r1 := corejson.New("a")
	r2 := corejson.New("b")
	rc.Items = append(rc.Items, r1, r2)
	actual := args.Map{
		"len":       rc.Length(),
		"lastIdx":   rc.LastIndex(),
		"isEmpty":   rc.IsEmpty(),
		"hasAny":    rc.HasAnyItem(),
		"firstNN":   rc.FirstOrDefault() != nil,
		"lastNN":    rc.LastOrDefault() != nil,
	}
	expected := args.Map{
		"len": 2, "lastIdx": 1, "isEmpty": false, "hasAny": true,
		"firstNN": true, "lastNN": true,
	}
	expected.ShouldBeEqual(t, 0, "ResultsCollection returns correct value -- basic", actual)
}

func Test_Cov5_ResultsCollection_Empty(t *testing.T) {
	rc := corejson.Empty.ResultsCollection()
	actual := args.Map{
		"firstNil": rc.FirstOrDefault() == nil,
		"lastNil":  rc.LastOrDefault() == nil,
	}
	expected := args.Map{"firstNil": true, "lastNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection returns empty -- empty", actual)
}

// ═══════════════════════════════════════════
// MapResults
// ═══════════════════════════════════════════

func Test_Cov5_MapResults_Basic(t *testing.T) {
	mr := corejson.Empty.MapResults()
	r := corejson.New("val")
	mr.Items["key1"] = r
	actual := args.Map{
		"len":     mr.Length(),
		"lastIdx": mr.LastIndex(),
		"isEmpty": mr.IsEmpty(),
		"hasAny":  mr.HasAnyItem(),
		"hasErr":  mr.HasError(),
		"getNN":   mr.GetByKey("key1") != nil,
		"getNil":  mr.GetByKey("missing") == nil,
	}
	expected := args.Map{
		"len": 1, "lastIdx": 0, "isEmpty": false, "hasAny": true,
		"hasErr": false, "getNN": true, "getNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MapResults returns correct value -- basic", actual)
}

func Test_Cov5_MapResults_AddSkipOnNil(t *testing.T) {
	mr := corejson.Empty.MapResults()
	r := corejson.NewPtr("val")
	mr.AddSkipOnNil("key1", r)
	mr.AddSkipOnNil("key2", nil)
	actual := args.Map{"len": mr.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapResults returns nil -- AddSkipOnNil", actual)
}

func Test_Cov5_MapResults_AllErrors(t *testing.T) {
	mr := corejson.Empty.MapResults()
	mr.Items["ok"] = corejson.New("val")
	mr.Items["fail"] = corejson.Result{Error: errors.New("err")}
	errs, hasAny := mr.AllErrors()
	actual := args.Map{"len": len(errs), "hasAny": hasAny}
	expected := args.Map{"len": 1, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "MapResults returns error -- AllErrors", actual)
}

func Test_Cov5_MapResults_HasError(t *testing.T) {
	mr := corejson.Empty.MapResults()
	mr.Items["fail"] = corejson.Result{Error: errors.New("err")}
	actual := args.Map{"hasErr": mr.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapResults returns error -- HasError", actual)
}
