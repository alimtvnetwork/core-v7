package regexnew

import (
	"errors"
	"regexp"
	"strings"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// Consolidated internal tests for unexported symbols in regexnew.
// These tests MUST remain in the source package because they access
// unexported types/fields/functions: lazyRegexMap, prettyJson,
// regExMatchValidationError, LazyRegex.pattern, LazyRegex.compiler.
//
// Source: Coverage1_Iteration11_test.go, Coverage2_Iteration17_test.go
// ══════════════════════════════════════════════════════════════════════════════

func Test_RX_LazyRegexMap_StateAndLockMethods(t *testing.T) {
	var nilMap *lazyRegexMap

	if !nilMap.IsEmpty() || !nilMap.IsEmptyLock() {
		t.Fatal("nil map should be empty")
	}

	if nilMap.HasAnyItem() || nilMap.HasAnyItemLock() {
		t.Fatal("nil map should not have items")
	}

	if nilMap.Length() != 0 || nilMap.LengthLock() != 0 {
		t.Fatal("nil map length should be zero")
	}

	m := &lazyRegexMap{items: map[string]*LazyRegex{}}
	if !m.IsEmpty() {
		t.Fatal("new map should be empty")
	}

	first, existed := m.CreateOrExisting(`^i11\\d+$`)
	if existed || first == nil {
		t.Fatal("expected creation path")
	}

	if !m.Has(`^i11\\d+$`) || !m.HasLock(`^i11\\d+$`) {
		t.Fatal("expected map has key")
	}

	if !m.HasAnyItem() || !m.HasAnyItemLock() {
		t.Fatal("expected map has item")
	}

	if m.Length() == 0 || m.LengthLock() == 0 {
		t.Fatal("expected non-zero length")
	}

	second, existedAgain := m.CreateOrExisting(`^i11\\d+$`)
	if !existedAgain || first != second {
		t.Fatal("expected existing pointer on second call")
	}

	third, existedLock := m.CreateOrExistingLock(`^i11-lock$`)
	if existedLock || third == nil {
		t.Fatal("expected lock create path")
	}

	fourth, existedLockIf := m.CreateOrExistingLockIf(true, `^i11-lock-if$`)
	if existedLockIf || fourth == nil {
		t.Fatal("expected lock-if create path")
	}

	fifth, existedNoLockIf := m.CreateOrExistingLockIf(false, `^i11-no-lock-if$`)
	if existedNoLockIf || fifth == nil {
		t.Fatal("expected no-lock-if create path")
	}
}

func Test_RX_LazyRegexMap_CreateLazyRegex_CustomCompiler(t *testing.T) {
	m := &lazyRegexMap{items: map[string]*LazyRegex{}}
	custom := m.createLazyRegex(`^z+$`, func(pattern string) (*regexp.Regexp, error) {
		return regexp.Compile(pattern)
	})

	if custom == nil || custom.pattern != `^z+$` {
		t.Fatal("expected custom lazy regex")
	}

	compiled, err := custom.Compile()
	if err != nil || compiled == nil {
		t.Fatal("expected custom compiler success")
	}
}

func Test_RX_PrettyJson_And_MatchErrorBranches(t *testing.T) {
	if prettyJson(nil) != "" {
		t.Fatal("nil prettyJson should be empty")
	}

	marshalFail := struct {
		Fn func()
	}{Fn: func() {}}
	if prettyJson(marshalFail) != "" {
		t.Fatal("marshal-fail prettyJson should be empty")
	}

	errCompile := regExMatchValidationError("", "abc", errors.New("bad-regex"), nil)
	if errCompile == nil || !strings.Contains(errCompile.Error(), "compile failed") {
		t.Fatal("expected compile-failed message")
	}

	errNilRegex := regExMatchValidationError("^x$", "abc", nil, nil)
	if errNilRegex == nil || !strings.Contains(errNilRegex.Error(), "invalid cannot match") {
		t.Fatal("expected nil-regex message")
	}

	errNoMatch := regExMatchValidationError("^x$", "abc", nil, regexp.MustCompile(`^x$`))
	if errNoMatch == nil || !strings.Contains(errNoMatch.Error(), "doesn't match") {
		t.Fatal("expected mismatch message")
	}
}

func Test_RX_LazyRegex_IsMatchBytes_InvalidPatternBranch(t *testing.T) {
	lr := &LazyRegex{
		pattern:  "[",
		compiler: CreateLock,
	}

	if lr.IsMatchBytes([]byte("anything")) {
		t.Fatal("invalid pattern should not match bytes")
	}

	if !lr.IsFailedMatchBytes([]byte("anything")) {
		t.Fatal("invalid pattern should fail match bytes")
	}
}


// Covers: Create, CreateLock, CreateLockIf, CreateMust, CreateMustLockIf,
// CreateApplicableLock, NewMustLock, IsMatchLock, IsMatchFailed,
// MatchError, MatchErrorLock, MatchUsingFuncErrorLock,
// MatchUsingCustomizeErrorFuncLock, newCreator, newLazyRegexCreator,
// LazyRegex extended methods, regexes-compiled vars

func Test_Create_New(t *testing.T) {
	r, err := Create(`^abc$`)
	if err != nil || r == nil {
		t.Fatal("expected valid regex")
	}
}

func Test_Create_Cached(t *testing.T) {
	r1, _ := Create(`^i17cached$`)
	r2, _ := Create(`^i17cached$`)
	if r1 != r2 {
		t.Fatal("expected same cached pointer")
	}
}

func Test_Create_Invalid(t *testing.T) {
	_, err := Create(`[invalid`)
	if err == nil {
		t.Fatal("expected error for invalid pattern")
	}
}

func Test_CreateLock(t *testing.T) {
	r, err := CreateLock(`^i17lock$`)
	if err != nil || r == nil {
		t.Fatal("expected valid")
	}
}

func Test_CreateLockIf_WithLock(t *testing.T) {
	r, err := CreateLockIf(true, `^i17lockif$`)
	if err != nil || r == nil {
		t.Fatal("expected valid")
	}
}

func Test_CreateLockIf_WithoutLock(t *testing.T) {
	r, err := CreateLockIf(false, `^i17nolockif$`)
	if err != nil || r == nil {
		t.Fatal("expected valid")
	}
}

func Test_CreateMust_Valid(t *testing.T) {
	r := CreateMust(`^i17must$`)
	if r == nil {
		t.Fatal("expected valid")
	}
}

func Test_CreateMust_Cached(t *testing.T) {
	r1 := CreateMust(`^i17mustcached$`)
	r2 := CreateMust(`^i17mustcached$`)
	if r1 != r2 {
		t.Fatal("expected same cached pointer")
	}
}

func Test_CreateMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	CreateMust(`[invalid`)
}

func Test_CreateMustLockIf_WithLock(t *testing.T) {
	r := CreateMustLockIf(true, `^i17mustlockif$`)
	if r == nil {
		t.Fatal("expected valid")
	}
}

func Test_CreateMustLockIf_WithoutLock(t *testing.T) {
	r := CreateMustLockIf(false, `^i17mustnolockif$`)
	if r == nil {
		t.Fatal("expected valid")
	}
}

func Test_CreateApplicableLock_Valid(t *testing.T) {
	r, err, isApplicable := CreateApplicableLock(`^i17appl$`)
	if err != nil || r == nil || !isApplicable {
		t.Fatal("expected applicable")
	}
}

func Test_CreateApplicableLock_Invalid(t *testing.T) {
	_, err, isApplicable := CreateApplicableLock(`[invalid`)
	if err == nil || isApplicable {
		t.Fatal("expected not applicable")
	}
}

func Test_NewMustLock(t *testing.T) {
	r := NewMustLock(`^i17newmustlock$`)
	if r == nil {
		t.Fatal("expected valid")
	}
}

func Test_IsMatchLock_Match(t *testing.T) {
	if !IsMatchLock(`^\d+$`, "123") {
		t.Fatal("expected match")
	}
}

func Test_IsMatchLock_NoMatch(t *testing.T) {
	if IsMatchLock(`^\d+$`, "abc") {
		t.Fatal("expected no match")
	}
}

func Test_IsMatchLock_InvalidRegex(t *testing.T) {
	if IsMatchLock(`[invalid`, "abc") {
		t.Fatal("expected false for invalid regex")
	}
}

func Test_IsMatchFailed_Match(t *testing.T) {
	if IsMatchFailed(`^\d+$`, "123") {
		t.Fatal("expected not failed for match")
	}
}

func Test_IsMatchFailed_NoMatch(t *testing.T) {
	if !IsMatchFailed(`^\d+$`, "abc") {
		t.Fatal("expected failed for no match")
	}
}

func Test_MatchError_Match(t *testing.T) {
	err := MatchError(`^\d+$`, "123")
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_MatchError_NoMatch(t *testing.T) {
	err := MatchError(`^\d+$`, "abc")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_MatchError_InvalidPattern(t *testing.T) {
	err := MatchError(`[invalid`, "abc")
	if err == nil {
		t.Fatal("expected error for invalid pattern")
	}
}

func Test_MatchErrorLock_Match(t *testing.T) {
	err := MatchErrorLock(`^\d+$`, "123")
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_MatchErrorLock_NoMatch(t *testing.T) {
	err := MatchErrorLock(`^\d+$`, "abc")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_MatchUsingFuncErrorLock_Match(t *testing.T) {
	fn := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	err := MatchUsingFuncErrorLock(`^\d+$`, "123", fn)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_MatchUsingFuncErrorLock_NoMatch(t *testing.T) {
	fn := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	err := MatchUsingFuncErrorLock(`^\d+$`, "abc", fn)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_MatchUsingCustomizeErrorFuncLock_Match(t *testing.T) {
	fn := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	err := MatchUsingCustomizeErrorFuncLock(`^\d+$`, "123", fn, nil)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_MatchUsingCustomizeErrorFuncLock_NoMatch_DefaultErr(t *testing.T) {
	fn := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	err := MatchUsingCustomizeErrorFuncLock(`^\d+$`, "abc", fn, nil)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_MatchUsingCustomizeErrorFuncLock_NoMatch_CustomErr(t *testing.T) {
	customSentinel := errors.New("custom-sentinel-error")
	fn := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	customErr := func(pattern, comparing string, compileErr error, regEx *regexp.Regexp) error {
		return customSentinel
	}
	err := MatchUsingCustomizeErrorFuncLock(`^\d+$`, "abc", fn, customErr)
	if err != customSentinel {
		t.Fatal("expected custom error")
	}
}

// newCreator tests
func Test_NewCreator_Lazy(t *testing.T) {
	lr := New.Lazy(`^i17-new-lazy$`)
	if lr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_NewCreator_LazyLock(t *testing.T) {
	lr := New.LazyLock(`^i17-new-lazylock$`)
	if lr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_NewCreator_Default(t *testing.T) {
	r, err := New.Default(`^i17-new-default$`)
	if err != nil || r == nil {
		t.Fatal("expected valid")
	}
}

func Test_NewCreator_DefaultLock(t *testing.T) {
	r, err := New.DefaultLock(`^i17-new-defaultlock$`)
	if err != nil || r == nil {
		t.Fatal("expected valid")
	}
}

func Test_NewCreator_DefaultLockIf(t *testing.T) {
	r, err := New.DefaultLockIf(true, `^i17-new-defaultlockif$`)
	if err != nil || r == nil {
		t.Fatal("expected valid")
	}
}

func Test_NewCreator_DefaultApplicableLock(t *testing.T) {
	r, err, isAppl := New.DefaultApplicableLock(`^i17-new-appllock$`)
	if err != nil || r == nil || !isAppl {
		t.Fatal("expected applicable")
	}
}

// newLazyRegexCreator tests
func Test_LazyRegexCreator_TwoLock(t *testing.T) {
	first, second := New.LazyRegex.TwoLock(`^i17-two1$`, `^i17-two2$`)
	if first == nil || second == nil {
		t.Fatal("expected both non-nil")
	}
}

func Test_LazyRegexCreator_ManyUsingLock_Empty(t *testing.T) {
	m := New.LazyRegex.ManyUsingLock()
	if len(m) != 0 {
		t.Fatal("expected empty map")
	}
}

func Test_LazyRegexCreator_ManyUsingLock_Valid(t *testing.T) {
	m := New.LazyRegex.ManyUsingLock(`^i17-many1$`, `^i17-many2$`)
	if len(m) != 2 {
		t.Fatal("expected 2 items")
	}
}

func Test_LazyRegexCreator_AllPatternsMap(t *testing.T) {
	m := New.LazyRegex.AllPatternsMap()
	if m == nil {
		t.Fatal("expected non-nil map")
	}
}

func Test_LazyRegexCreator_NewLockIf_Lock(t *testing.T) {
	lr := New.LazyRegex.NewLockIf(true, `^i17-lockif$`)
	if lr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_LazyRegexCreator_NewLockIf_NoLock(t *testing.T) {
	lr := New.LazyRegex.NewLockIf(false, `^i17-nolockif$`)
	if lr == nil {
		t.Fatal("expected non-nil")
	}
}

// LazyRegex extended methods
func Test_LazyRegex_IsNull(t *testing.T) {
	var lr *LazyRegex
	if !lr.IsNull() {
		t.Fatal("expected null")
	}
}

func Test_LazyRegex_IsDefined(t *testing.T) {
	lr := New.Lazy(`^i17-defined$`)
	if !lr.IsDefined() {
		t.Fatal("expected defined")
	}
}

func Test_LazyRegex_IsUndefined_Nil(t *testing.T) {
	var lr *LazyRegex
	if !lr.IsUndefined() {
		t.Fatal("expected undefined")
	}
}

func Test_LazyRegex_IsApplicable(t *testing.T) {
	lr := New.Lazy(`^i17-applicable$`)
	if !lr.IsApplicable() {
		t.Fatal("expected applicable")
	}
}

func Test_LazyRegex_IsApplicable_Nil(t *testing.T) {
	var lr *LazyRegex
	if lr.IsApplicable() {
		t.Fatal("expected not applicable for nil")
	}
}

func Test_LazyRegex_IsApplicable_Invalid(t *testing.T) {
	lr := &LazyRegex{pattern: "[", compiler: CreateLock}
	if lr.IsApplicable() {
		t.Fatal("expected not applicable for invalid")
	}
}

func Test_LazyRegex_IsApplicable_Cached(t *testing.T) {
	lr := New.Lazy(`^i17-appl-cached$`)
	lr.Compile() // compile first
	if !lr.IsApplicable() {
		t.Fatal("expected applicable from cache")
	}
}

func Test_LazyRegex_Compile_Undefined(t *testing.T) {
	lr := &LazyRegex{}
	_, err := lr.Compile()
	if err == nil {
		t.Fatal("expected error for undefined")
	}
}

func Test_LazyRegex_CompileMust_Success(t *testing.T) {
	lr := New.Lazy(`^i17-compilemust$`)
	r := lr.CompileMust()
	if r == nil {
		t.Fatal("expected valid")
	}
}

func Test_LazyRegex_CompileMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	lr := &LazyRegex{pattern: "[", compiler: CreateLock}
	lr.CompileMust()
}

func Test_LazyRegex_OnRequiredCompiled_Nil(t *testing.T) {
	var lr *LazyRegex
	err := lr.OnRequiredCompiled()
	if err == nil {
		t.Fatal("expected error for nil")
	}
}

func Test_LazyRegex_OnRequiredCompiled_AlreadyCompiled(t *testing.T) {
	lr := New.Lazy(`^i17-onreq$`)
	lr.Compile()
	err := lr.OnRequiredCompiled()
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func Test_LazyRegex_OnRequiredCompiledMust_Success(t *testing.T) {
	lr := New.Lazy(`^i17-onreqmust$`)
	lr.OnRequiredCompiledMust()
}

func Test_LazyRegex_OnRequiredCompiledMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	var lr *LazyRegex
	lr.OnRequiredCompiledMust()
}

func Test_LazyRegex_HasError_Valid(t *testing.T) {
	lr := New.Lazy(`^i17-haserror$`)
	if lr.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_LazyRegex_HasError_Invalid(t *testing.T) {
	lr := &LazyRegex{pattern: "[", compiler: CreateLock}
	if !lr.HasError() {
		t.Fatal("expected error")
	}
}

func Test_LazyRegex_HasAnyIssues_Nil(t *testing.T) {
	var lr *LazyRegex
	if !lr.HasAnyIssues() {
		t.Fatal("expected issues for nil")
	}
}

func Test_LazyRegex_HasAnyIssues_Valid(t *testing.T) {
	lr := New.Lazy(`^i17-issues$`)
	if lr.HasAnyIssues() {
		t.Fatal("expected no issues")
	}
}

func Test_LazyRegex_IsInvalid_Nil(t *testing.T) {
	var lr *LazyRegex
	if !lr.IsInvalid() {
		t.Fatal("expected invalid for nil")
	}
}

func Test_LazyRegex_IsInvalid_Valid(t *testing.T) {
	lr := New.Lazy(`^i17-invalid$`)
	if lr.IsInvalid() {
		t.Fatal("expected valid")
	}
}

func Test_LazyRegex_CompiledError(t *testing.T) {
	lr := New.Lazy(`^i17-comperr$`)
	err := lr.CompiledError()
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func Test_LazyRegex_Error(t *testing.T) {
	lr := New.Lazy(`^i17-err$`)
	err := lr.Error()
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func Test_LazyRegex_MustBeSafe_Success(t *testing.T) {
	lr := New.Lazy(`^i17-mustbesafe$`)
	lr.MustBeSafe()
}

func Test_LazyRegex_MustBeSafe_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	lr := &LazyRegex{pattern: "[", compiler: CreateLock}
	lr.MustBeSafe()
}

func Test_LazyRegex_String_Nil(t *testing.T) {
	var lr *LazyRegex
	if lr.String() != "" {
		t.Fatal("expected empty for nil")
	}
}

func Test_LazyRegex_String_Valid(t *testing.T) {
	lr := New.Lazy(`^i17-str$`)
	if lr.String() != `^i17-str$` {
		t.Fatal("unexpected string")
	}
}

func Test_LazyRegex_FullString_Nil(t *testing.T) {
	var lr *LazyRegex
	if lr.FullString() != "" {
		t.Fatal("expected empty for nil")
	}
}

func Test_LazyRegex_FullString_Valid(t *testing.T) {
	lr := New.Lazy(`^i17-fullstr$`)
	s := lr.FullString()
	if s == "" {
		t.Fatal("expected non-empty full string")
	}
}

func Test_LazyRegex_Pattern_Nil(t *testing.T) {
	var lr *LazyRegex
	if lr.Pattern() != "" {
		t.Fatal("expected empty for nil")
	}
}

func Test_LazyRegex_Pattern_Valid(t *testing.T) {
	lr := New.Lazy(`^i17-pattern$`)
	if lr.Pattern() != `^i17-pattern$` {
		t.Fatal("unexpected pattern")
	}
}

func Test_LazyRegex_MatchError_Match(t *testing.T) {
	lr := New.Lazy(`^\d+$`)
	err := lr.MatchError("123")
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_LazyRegex_MatchError_NoMatch(t *testing.T) {
	lr := New.Lazy(`^\d+$`)
	err := lr.MatchError("abc")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_LazyRegex_MatchUsingFuncError_Match(t *testing.T) {
	lr := New.Lazy(`^\d+$`)
	fn := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	err := lr.MatchUsingFuncError("123", fn)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_LazyRegex_MatchUsingFuncError_NoMatch(t *testing.T) {
	lr := New.Lazy(`^\d+$`)
	fn := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	err := lr.MatchUsingFuncError("abc", fn)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_LazyRegex_IsMatch(t *testing.T) {
	lr := New.Lazy(`^\d+$`)
	if !lr.IsMatch("123") {
		t.Fatal("expected match")
	}
	if lr.IsMatch("abc") {
		t.Fatal("expected no match")
	}
}

func Test_LazyRegex_IsMatchBytes(t *testing.T) {
	lr := New.Lazy(`^\d+$`)
	if !lr.IsMatchBytes([]byte("123")) {
		t.Fatal("expected match")
	}
}

func Test_LazyRegex_IsFailedMatch(t *testing.T) {
	lr := New.Lazy(`^\d+$`)
	if lr.IsFailedMatch("123") {
		t.Fatal("expected not failed for match")
	}
	if !lr.IsFailedMatch("abc") {
		t.Fatal("expected failed for no match")
	}
}

func Test_LazyRegex_IsFailedMatch_Invalid(t *testing.T) {
	lr := &LazyRegex{pattern: "[", compiler: CreateLock}
	if !lr.IsFailedMatch("abc") {
		t.Fatal("expected failed for invalid pattern")
	}
}

func Test_LazyRegex_IsFailedMatchBytes(t *testing.T) {
	lr := New.Lazy(`^\d+$`)
	if lr.IsFailedMatchBytes([]byte("123")) {
		t.Fatal("expected not failed")
	}
	if !lr.IsFailedMatchBytes([]byte("abc")) {
		t.Fatal("expected failed")
	}
}

func Test_LazyRegex_FirstMatchLine_Match(t *testing.T) {
	lr := New.Lazy(`(\d+)`)
	match, isInvalid := lr.FirstMatchLine("abc123def")
	if isInvalid || match == "" {
		t.Fatal("expected valid match")
	}
}

func Test_LazyRegex_FirstMatchLine_NoMatch(t *testing.T) {
	lr := New.Lazy(`^zzz$`)
	match, isInvalid := lr.FirstMatchLine("abc")
	if !isInvalid || match != "" {
		t.Fatal("expected invalid match")
	}
}

func Test_LazyRegex_FirstMatchLine_Invalid(t *testing.T) {
	lr := &LazyRegex{pattern: "[", compiler: CreateLock}
	match, isInvalid := lr.FirstMatchLine("abc")
	if !isInvalid || match != "" {
		t.Fatal("expected invalid for bad pattern")
	}
}

// Test pre-compiled regex vars
func Test_CompiledRegexVars(t *testing.T) {
	if WhitespaceFinderRegex == nil {
		t.Fatal("expected non-nil")
	}
	if HashCommentWithSpaceOptionalRegex == nil {
		t.Fatal("expected non-nil")
	}
	if WhitespaceOrPipeFinderRegex == nil {
		t.Fatal("expected non-nil")
	}
	if DollarIdentifierRegex == nil {
		t.Fatal("expected non-nil")
	}
	if PercentIdentifierRegex == nil {
		t.Fatal("expected non-nil")
	}
	if PrettyNameRegex == nil {
		t.Fatal("expected non-nil")
	}
	if ExactIdFieldMatchingRegex == nil {
		t.Fatal("expected non-nil")
	}
	if ExactVersionIdFieldMatchingRegex == nil {
		t.Fatal("expected non-nil")
	}
	if UbuntuNameCheckerRegex == nil {
		t.Fatal("expected non-nil")
	}
	if CentOsNameCheckerRegex == nil {
		t.Fatal("expected non-nil")
	}
	if RedHatNameCheckerRegex == nil {
		t.Fatal("expected non-nil")
	}
	if FirstNumberAnyWhereCheckerRegex == nil {
		t.Fatal("expected non-nil")
	}
	if WindowsVersionNumberCheckerRegex == nil {
		t.Fatal("expected non-nil")
	}
}
