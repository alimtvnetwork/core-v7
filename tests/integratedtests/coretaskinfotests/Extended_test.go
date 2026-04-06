package coretaskinfotests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretaskinfo"
)

// ============================================================================
// Info: Core Identity
// ============================================================================

func Test_Info_Name_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("myName", "desc", "http://url")
	if info.Name() != "myName" {
		t.Errorf("expected 'myName', got '%s'", info.Name())
	}
}

func Test_Info_NilName_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	if info.Name() != "" {
		t.Error("nil Name should return empty")
	}
}

func Test_Info_IsDefined_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	if !info.IsDefined() {
		t.Error("should be defined")
	}
}

func Test_Info_NilIsDefined_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	if info.IsDefined() {
		t.Error("nil should not be defined")
	}
}

func Test_Info_HasAnyName_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	if !info.HasAnyName() {
		t.Error("should have name")
	}
}

func Test_Info_IsName_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("myName", "d", "u")
	if !info.IsName("myName") {
		t.Error("should match name")
	}
	if info.IsName("other") {
		t.Error("should not match other")
	}
}

func Test_Info_IsEmpty_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	if !info.IsEmpty() {
		t.Error("nil should be empty")
	}
}

func Test_Info_HasAnyItem_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	if !info.HasAnyItem() {
		t.Error("should have item")
	}
}

func Test_Info_Options_Ext(t *testing.T) {
	// Plain.Default does not set ExcludeOptions, so Options() returns nil.
	// This is correct production behavior — nil means "no exclusions".
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	opts := info.Options()
	if opts != nil {
		t.Error("Options should be nil for Default (no ExcludeOptions set)")
	}
}

func Test_Info_NilOptions_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	opts := info.Options()
	if opts == nil {
		t.Error("nil Options should return empty options")
	}
}

// ============================================================================
// Info: Clone / ToPtr / ToNonPtr
// ============================================================================

func Test_Info_Clone_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	cloned := info.Clone()
	if cloned.RootName != "n" {
		t.Error("Clone should preserve RootName")
	}
}

func Test_Info_ClonePtr_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	cloned := info.ClonePtr()
	if cloned == nil || cloned.RootName != "n" {
		t.Error("ClonePtr should preserve RootName")
	}
}

func Test_Info_NilClonePtr_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	if info.ClonePtr() != nil {
		t.Error("nil ClonePtr should return nil")
	}
}

func Test_Info_ToPtr_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	ptr := info.ToPtr()
	if ptr == nil {
		t.Error("ToPtr should not be nil")
	}
}

func Test_Info_ToNonPtr_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	nonPtr := info.ToNonPtr()
	if nonPtr.RootName != "n" {
		t.Error("ToNonPtr should preserve RootName")
	}
}

// ============================================================================
// Info: SetSecure / SetPlain
// ============================================================================

func Test_Info_SetSecure_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	result := info.SetSecure()
	if !result.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Info_NilSetSecure_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	result := info.SetSecure()
	if result == nil || !result.IsSecure() {
		t.Error("nil SetSecure should return secure info")
	}
}

func Test_Info_SetPlain_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.Default("n", "d", "u")
	result := info.SetPlain()
	if !result.IsPlainText() {
		t.Error("should be plain text")
	}
}

func Test_Info_NilSetPlain_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	result := info.SetPlain()
	if result == nil {
		t.Error("nil SetPlain should return info")
	}
}

// ============================================================================
// Info: Getters
// ============================================================================

func Test_Info_IsInclude_Getters_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.AllUrlExamples(
		"n", "d", "http://url", "http://hint", "http://err", "ex1",
	)
	info.SingleExample = "chain"
	info.ExampleUrl = "http://exurl"

	if !info.IsIncludeRootName() {
		t.Error("should include root name")
	}
	if !info.IsIncludeDescription() {
		t.Error("should include description")
	}
	if !info.IsIncludeUrl() {
		t.Error("should include url")
	}
	if !info.IsIncludeHintUrl() {
		t.Error("should include hint url")
	}
	if !info.IsIncludeErrorUrl() {
		t.Error("should include error url")
	}
	if !info.IsIncludeExampleUrl() {
		t.Error("should include example url")
	}
	if !info.IsIncludeSingleExample() {
		t.Error("should include single example")
	}
	if !info.IsIncludeExamples() {
		t.Error("should include examples")
	}
	if !info.IsIncludeAdditionalErrorWrap() {
		t.Error("should include additional error wrap")
	}
}

func Test_Info_NilIsInclude_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	if info.IsIncludeRootName() {
		t.Error("nil should not include root name")
	}
	if !info.IsIncludeAdditionalErrorWrap() {
		t.Error("nil should include additional error wrap")
	}
}

func Test_Info_IsSecure_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.Default("n", "d", "u")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
	if !info.IsExcludePayload() {
		t.Error("should exclude payload")
	}
	if info.IsPlainText() {
		t.Error("should not be plain text")
	}
}

func Test_Info_IsPlainText_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	if !info.IsPlainText() {
		t.Error("should be plain text")
	}
	if !info.IsIncludePayloads() {
		t.Error("should include payloads")
	}
}

// ============================================================================
// Info: Safe* getters
// ============================================================================

func Test_Info_SafeGetters_Ext(t *testing.T) {
	info := &coretaskinfo.Info{
		RootName:    "n",
		Description: "d",
		Url:         "u",
		HintUrl:     "h",
		ErrorUrl:    "e",
		ExampleUrl:  "ex",
	}
	if info.SafeName() != "n" {
		t.Error("SafeName mismatch")
	}
	if info.SafeDescription() != "d" {
		t.Error("SafeDescription mismatch")
	}
	if info.SafeUrl() != "u" {
		t.Error("SafeUrl mismatch")
	}
	if info.SafeHintUrl() != "h" {
		t.Error("SafeHintUrl mismatch")
	}
	if info.SafeErrorUrl() != "e" {
		t.Error("SafeErrorUrl mismatch")
	}
	if info.SafeExampleUrl() != "ex" {
		t.Error("SafeExampleUrl mismatch")
	}
	if info.SafeChainingExample() != "ex" {
		t.Error("SafeChainingExample mismatch")
	}
}

func Test_Info_NilSafeGetters_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	if info.SafeName() != "" {
		t.Error("nil SafeName should be empty")
	}
	if info.SafeDescription() != "" {
		t.Error("nil SafeDescription should be empty")
	}
	if info.SafeUrl() != "" {
		t.Error("nil SafeUrl should be empty")
	}
	if info.SafeHintUrl() != "" {
		t.Error("nil SafeHintUrl should be empty")
	}
	if info.SafeErrorUrl() != "" {
		t.Error("nil SafeErrorUrl should be empty")
	}
	if info.SafeExampleUrl() != "" {
		t.Error("nil SafeExampleUrl should be empty")
	}
	if info.SafeChainingExample() != "" {
		t.Error("nil SafeChainingExample should be empty")
	}
}

// ============================================================================
// Info: Has* checks
// ============================================================================

func Test_Info_HasChecks_Ext(t *testing.T) {
	info := &coretaskinfo.Info{
		RootName:    "n",
		Description: "d",
		Url:         "u",
		HintUrl:     "h",
		ErrorUrl:    "e",
		ExampleUrl:  "ex",
		SingleExample: "se",
		Examples:    []string{"e1"},
		ExcludeOptions: &coretaskinfo.ExcludingOptions{IsExcludeRootName: true},
	}
	if !info.HasRootName() {
		t.Error("should have root name")
	}
	if !info.HasDescription() {
		t.Error("should have description")
	}
	if !info.HasUrl() {
		t.Error("should have url")
	}
	if !info.HasHintUrl() {
		t.Error("should have hint url")
	}
	if !info.HasErrorUrl() {
		t.Error("should have error url")
	}
	if !info.HasExampleUrl() {
		t.Error("should have example url")
	}
	if !info.HasChainingExample() {
		t.Error("should have chaining example")
	}
	if !info.HasExamples() {
		t.Error("should have examples")
	}
	if !info.HasExcludeOptions() {
		t.Error("should have exclude options")
	}
}

// ============================================================================
// Info: IsEmpty* checks
// ============================================================================

func Test_Info_IsEmptyChecks_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	if !info.IsEmptyName() {
		t.Error("nil should be empty name")
	}
	if !info.IsEmptyDescription() {
		t.Error("nil should be empty description")
	}
	if !info.IsEmptyUrl() {
		t.Error("nil should be empty url")
	}
	if !info.IsEmptyHintUrl() {
		t.Error("nil should be empty hint url")
	}
	if !info.IsEmptyErrorUrl() {
		t.Error("nil should be empty error url")
	}
	if !info.IsEmptyExampleUrl() {
		t.Error("nil should be empty example url")
	}
	if !info.IsEmptySingleExample() {
		t.Error("nil should be empty single example")
	}
	if !info.IsEmptyExamples() {
		t.Error("nil should be empty examples")
	}
	if !info.IsEmptyExcludeOptions() {
		t.Error("nil should be empty exclude options")
	}
}

// ============================================================================
// Info: IsExclude* checks
// ============================================================================

func Test_Info_IsExcludeChecks_Ext(t *testing.T) {
	info := &coretaskinfo.Info{
		ExcludeOptions: &coretaskinfo.ExcludingOptions{
			IsExcludeRootName:            true,
			IsExcludeDescription:         true,
			IsExcludeUrl:                 true,
			IsExcludeHintUrl:             true,
			IsExcludeErrorUrl:            true,
			IsExcludeAdditionalErrorWrap: true,
			IsExcludeExampleUrl:          true,
			IsExcludeSingleExample:       true,
			IsExcludeExamples:            true,
		},
	}
	if !info.IsExcludeRootName() {
		t.Error("should exclude root name")
	}
	if !info.IsExcludeDescription() {
		t.Error("should exclude description")
	}
	if !info.IsExcludeUrl() {
		t.Error("should exclude url")
	}
	if !info.IsExcludeHintUrl() {
		t.Error("should exclude hint url")
	}
	if !info.IsExcludeErrorUrl() {
		t.Error("should exclude error url")
	}
	if !info.IsExcludeAdditionalErrorWrap() {
		t.Error("should exclude additional error wrap")
	}
	if !info.IsExcludeExampleUrl() {
		t.Error("should exclude example url")
	}
	if !info.IsExcludeSingleExample() {
		t.Error("should exclude single example")
	}
	if !info.IsExcludeExamples() {
		t.Error("should exclude examples")
	}
}

// ============================================================================
// Info: JSON methods
// ============================================================================

func Test_Info_JsonString_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	result := info.JsonString()
	if result == "" {
		t.Error("JsonString should return non-empty")
	}
}

func Test_Info_NilJsonString_Ext(t *testing.T) {
	var info coretaskinfo.Info
	result := info.JsonString()
	if result != "" {
		t.Error("nil/zero JsonString should return empty")
	}
}

func Test_Info_PrettyJsonString_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	result := info.PrettyJsonString()
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_Info_NilPrettyJsonString_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	result := info.PrettyJsonString()
	if result != "" {
		t.Error("nil should return empty")
	}
}

func Test_Info_String_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	result := info.String()
	if result == "" {
		t.Error("String should return non-empty")
	}
}

func Test_Info_NilString_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	result := info.String()
	if result != "" {
		t.Error("nil String should return empty")
	}
}

func Test_Info_Serialize_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	bytes, err := info.Serialize()
	if err != nil {
		t.Errorf("Serialize error: %v", err)
	}
	if len(bytes) == 0 {
		t.Error("should return non-empty bytes")
	}
}

func Test_Info_ExamplesAsString_Ext(t *testing.T) {
	info := &coretaskinfo.Info{Examples: []string{"e1", "e2"}}
	result := info.ExamplesAsString()
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_Info_NilExamplesAsString_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	result := info.ExamplesAsString()
	if result != "" {
		t.Error("nil should return empty")
	}
}

func Test_Info_ExamplesAsSlice_Ext(t *testing.T) {
	info := &coretaskinfo.Info{Examples: []string{"e1", "e2"}}
	result := info.ExamplesAsSlice()
	if result == nil {
		t.Error("should not be nil")
	}
}

func Test_Info_NilExamplesAsSlice_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	result := info.ExamplesAsSlice()
	if result == nil {
		t.Error("nil should return empty slice")
	}
}

func Test_Info_AsJsonContractsBinder_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	binder := info.AsJsonContractsBinder()
	if binder == nil {
		t.Error("should not be nil")
	}
}

func Test_Info_JsonStringMust_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	result := info.JsonStringMust()
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_Info_LazyMapPrettyJsonString_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	result := info.LazyMapPrettyJsonString()
	if result == "" {
		t.Error("should return non-empty")
	}
}

// ============================================================================
// Info: Map methods
// ============================================================================

func Test_Info_Map_Ext(t *testing.T) {
	info := &coretaskinfo.Info{
		RootName:      "n",
		Description:   "d",
		Url:           "u",
		HintUrl:       "h",
		ErrorUrl:      "e",
		ExampleUrl:    "ex",
		SingleExample: "se",
		Examples:      []string{"e1"},
	}
	m := info.Map()
	if len(m) == 0 {
		t.Error("should have entries")
	}
}

func Test_Info_NilMap_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	m := info.Map()
	if len(m) != 0 {
		t.Error("nil map should be empty")
	}
}

func Test_Info_MapWithPayload_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	m := info.MapWithPayload([]byte("payload"))
	if m == nil {
		t.Error("should not be nil")
	}
}

func Test_Info_LazyMap_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	m1 := info.LazyMap()
	m2 := info.LazyMap() // cached
	if len(m1) != len(m2) {
		t.Error("lazy map should be cached")
	}
}

func Test_Info_NilLazyMap_Ext(t *testing.T) {
	var info *coretaskinfo.Info
	m := info.LazyMap()
	if len(m) != 0 {
		t.Error("nil lazy map should be empty")
	}
}

func Test_Info_PrettyJsonStringWithPayloads_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	result := info.PrettyJsonStringWithPayloads([]byte("payload"))
	if result == "" {
		t.Error("should return non-empty")
	}
}

func Test_Info_LazyMapWithPayload_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	m := info.LazyMapWithPayload([]byte("payload"))
	if m == nil {
		t.Error("should not be nil")
	}
}

func Test_Info_MapWithPayloadAsAny_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	m := info.MapWithPayloadAsAny("test-payload")
	if m == nil {
		t.Error("should not be nil")
	}
}

func Test_Info_LazyMapWithPayloadAsAny_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	m := info.LazyMapWithPayloadAsAny("test-payload")
	if m == nil {
		t.Error("should not be nil")
	}
}

func Test_Info_SecureMapWithPayloadAsAny_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.Default("n", "d", "u")
	m := info.MapWithPayloadAsAny("test-payload")
	if m == nil {
		t.Error("should not be nil")
	}
}

func Test_Info_SecureLazyMapWithPayloadAsAny_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.Default("n", "d", "u")
	m := info.LazyMapWithPayloadAsAny("test-payload")
	if m == nil {
		t.Error("should not be nil")
	}
}

// ============================================================================
// ExcludingOptions
// ============================================================================

func Test_ExcludingOptions_SetSecure_Ext(t *testing.T) {
	opts := &coretaskinfo.ExcludingOptions{}
	result := opts.SetSecure()
	if !result.IsSafeSecureText() {
		t.Error("should be secure")
	}
}

func Test_ExcludingOptions_NilSetSecure_Ext(t *testing.T) {
	var opts *coretaskinfo.ExcludingOptions
	result := opts.SetSecure()
	if result == nil || !result.IsSafeSecureText() {
		t.Error("nil SetSecure should return secure")
	}
}

func Test_ExcludingOptions_SetPlainText_Ext(t *testing.T) {
	opts := &coretaskinfo.ExcludingOptions{IsSecureText: true}
	result := opts.SetPlainText()
	if result.IsSafeSecureText() {
		t.Error("should not be secure after SetPlainText")
	}
}

func Test_ExcludingOptions_NilSetPlainText_Ext(t *testing.T) {
	var opts *coretaskinfo.ExcludingOptions
	result := opts.SetPlainText()
	if result == nil {
		t.Error("nil SetPlainText should return options")
	}
}

func Test_ExcludingOptions_IsEmpty_Ext(t *testing.T) {
	opts := &coretaskinfo.ExcludingOptions{}
	if !opts.IsEmpty() {
		t.Error("default should be empty")
	}
	opts.IsExcludeRootName = true
	if opts.IsEmpty() {
		t.Error("should not be empty after setting a flag")
	}
}

func Test_ExcludingOptions_NilIsEmpty_Ext(t *testing.T) {
	var opts *coretaskinfo.ExcludingOptions
	if !opts.IsEmpty() {
		t.Error("nil should be empty")
	}
}

func Test_ExcludingOptions_Clone_Ext(t *testing.T) {
	opts := coretaskinfo.ExcludingOptions{IsExcludeRootName: true}
	cloned := opts.Clone()
	if !cloned.IsExcludeRootName {
		t.Error("Clone should preserve flags")
	}
}

func Test_ExcludingOptions_ClonePtr_Ext(t *testing.T) {
	opts := &coretaskinfo.ExcludingOptions{IsExcludeRootName: true}
	cloned := opts.ClonePtr()
	if cloned == nil || !cloned.IsExcludeRootName {
		t.Error("ClonePtr should preserve flags")
	}
}

func Test_ExcludingOptions_NilClonePtr_Ext(t *testing.T) {
	var opts *coretaskinfo.ExcludingOptions
	result := opts.ClonePtr()
	if result == nil {
		t.Error("nil ClonePtr should return empty options")
	}
}

func Test_ExcludingOptions_ToPtr_Ext(t *testing.T) {
	opts := coretaskinfo.ExcludingOptions{IsExcludeRootName: true}
	ptr := opts.ToPtr()
	if ptr == nil || !ptr.IsExcludeRootName {
		t.Error("ToPtr should preserve flags")
	}
}

func Test_ExcludingOptions_ToNonPtr_Ext(t *testing.T) {
	opts := coretaskinfo.ExcludingOptions{IsExcludeRootName: true}
	nonPtr := opts.ToNonPtr()
	if !nonPtr.IsExcludeRootName {
		t.Error("ToNonPtr should preserve flags")
	}
}

func Test_ExcludingOptions_IsInclude_All_Ext(t *testing.T) {
	var opts *coretaskinfo.ExcludingOptions
	if !opts.IsIncludeRootName() {
		t.Error("nil should include root name")
	}
	if !opts.IsIncludeDescription() {
		t.Error("nil should include description")
	}
	if !opts.IsIncludeUrl() {
		t.Error("nil should include url")
	}
	if !opts.IsIncludeHintUrl() {
		t.Error("nil should include hint url")
	}
	if !opts.IsIncludeErrorUrl() {
		t.Error("nil should include error url")
	}
	if !opts.IsIncludeExampleUrl() {
		t.Error("nil should include example url")
	}
	if !opts.IsIncludeSingleExample() {
		t.Error("nil should include single example")
	}
	if !opts.IsIncludeExamples() {
		t.Error("nil should include examples")
	}
	if !opts.IsIncludeAdditionalErrorWrap() {
		t.Error("nil should include additional error wrap")
	}
	if !opts.IsIncludePayloads() {
		t.Error("nil should include payloads")
	}
}

func Test_ExcludingOptions_IsSafe_All_Ext(t *testing.T) {
	var opts *coretaskinfo.ExcludingOptions
	if opts.IsSafeExcludeRootName() {
		t.Error("nil should not exclude")
	}
	if opts.IsSafeExcludeDescription() {
		t.Error("nil should not exclude")
	}
	if opts.IsSafeExcludeUrl() {
		t.Error("nil should not exclude")
	}
	if opts.IsSafeExcludeErrorUrl() {
		t.Error("nil should not exclude")
	}
	if opts.IsSafeExcludeAdditionalErrorWrap() {
		t.Error("nil should not exclude")
	}
	if opts.IsSafeExcludeHintUrl() {
		t.Error("nil should not exclude")
	}
	if opts.IsSafeExcludeExampleUrl() {
		t.Error("nil should not exclude")
	}
	if opts.IsSafeExcludeSingleExample() {
		t.Error("nil should not exclude")
	}
	if opts.IsSafeExcludeExamples() {
		t.Error("nil should not exclude")
	}
	if opts.IsSafeSecureText() {
		t.Error("nil should not be secure")
	}
}

// ============================================================================
// newInfoCreator methods
// ============================================================================

func Test_NewInfo_Default_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Default("n", "d", "u")
	if info.RootName != "n" {
		t.Error("should set root name")
	}
}

func Test_NewInfo_Examples_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Examples("n", "d", "u", "e1", "e2")
	if len(info.Examples) != 2 {
		t.Error("should have 2 examples")
	}
}

func Test_NewInfo_Create_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Create(true, "n", "d", "u", "h", "e", "ex", "se", "e1")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_NewInfo_SecureCreate_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.SecureCreate("n", "d", "u", "h", "e", "ex", "se", "e1")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_NewInfo_PlainCreate_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.PlainCreate("n", "d", "u", "h", "e", "ex", "se", "e1")
	if info.IsSecure() {
		t.Error("should not be secure")
	}
}

// ============================================================================
// newInfoPlainTextCreator - remaining methods
// ============================================================================

func Test_Plain_NameDescUrl_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.NameDescUrl("n", "d", "u")
	if info.RootName != "n" {
		t.Error("should set name")
	}
}

func Test_Plain_NameDescUrlExamples_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.NameDescUrlExamples("n", "d", "u", "e1")
	if len(info.Examples) != 1 {
		t.Error("should have 1 example")
	}
}

func Test_Plain_NewNameDescUrlErrorUrl_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.NewNameDescUrlErrorUrl("n", "d", "u", "eu")
	if info.ErrorUrl != "eu" {
		t.Error("should set error url")
	}
}

func Test_Plain_NameDescUrlErrUrlExamples_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.NameDescUrlErrUrlExamples("n", "d", "u", "eu", "e1")
	if info.ErrorUrl != "eu" || len(info.Examples) != 1 {
		t.Error("should set error url and examples")
	}
}

func Test_Plain_NameDescExamples_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.NameDescExamples("n", "d", "e1")
	if len(info.Examples) != 1 {
		t.Error("should have 1 example")
	}
}

func Test_Plain_Examples_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.Examples("n", "d", "e1")
	if len(info.Examples) != 1 {
		t.Error("should have 1 example")
	}
}

func Test_Plain_NameUrlExamples_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.NameUrlExamples("n", "u", "e1")
	if info.Url != "u" {
		t.Error("should set url")
	}
}

func Test_Plain_UrlExamples_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.UrlExamples("u", "e1")
	if info.Url != "u" {
		t.Error("should set url")
	}
}

func Test_Plain_ExamplesOnly_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.ExamplesOnly("e1")
	if len(info.Examples) != 1 {
		t.Error("should have 1 example")
	}
}

func Test_Plain_UrlOnly_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.UrlOnly("u")
	if info.Url != "u" {
		t.Error("should set url")
	}
}

func Test_Plain_ErrorUrlOnly_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.ErrorUrlOnly("eu")
	if info.ErrorUrl != "eu" {
		t.Error("should set error url")
	}
}

func Test_Plain_HintUrlOnly_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.HintUrlOnly("hu")
	if info.HintUrl != "hu" {
		t.Error("should set hint url")
	}
}

func Test_Plain_DescHintUrlOnly_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.DescHintUrlOnly("d", "hu")
	if info.Description != "d" || info.HintUrl != "hu" {
		t.Error("should set desc and hint url")
	}
}

func Test_Plain_NameHintUrlOnly_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.NameHintUrlOnly("n", "hu")
	if info.RootName != "n" || info.HintUrl != "hu" {
		t.Error("should set name and hint url")
	}
}

func Test_Plain_SingleExampleOnly_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.SingleExampleOnly("se")
	if info.SingleExample != "se" {
		t.Error("should set single example")
	}
}

func Test_Plain_AllUrl_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.AllUrl("n", "d", "u", "hu", "eu")
	if info.HintUrl != "hu" || info.ErrorUrl != "eu" {
		t.Error("should set all urls")
	}
}

func Test_Plain_UrlSingleExample_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.UrlSingleExample("n", "d", "u", "se")
	if info.SingleExample != "se" {
		t.Error("should set single example")
	}
}

func Test_Plain_SingleExample_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.SingleExample("n", "d", "se")
	if info.SingleExample != "se" {
		t.Error("should set single example")
	}
}

func Test_Plain_ExampleUrl_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.ExampleUrl("n", "d", "exu", "se")
	if info.ExampleUrl != "exu" {
		t.Error("should set example url")
	}
}

func Test_Plain_ExampleUrlSingleExample_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Plain.ExampleUrlSingleExample("n", "d", "exu", "se")
	if info.ExampleUrl != "exu" || info.SingleExample != "se" {
		t.Error("should set both")
	}
}

// ============================================================================
// newInfoSecureTextCreator - remaining methods
// ============================================================================

func Test_Secure_NameDescUrl_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.NameDescUrl("n", "d", "u")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_NameDescUrlExamples_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.NameDescUrlExamples("n", "d", "u", "e1")
	if !info.IsSecure() || len(info.Examples) != 1 {
		t.Error("should be secure with examples")
	}
}

func Test_Secure_NewNameDescUrlErrorUrl_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.NewNameDescUrlErrorUrl("n", "d", "u", "eu")
	if !info.IsSecure() || info.ErrorUrl != "eu" {
		t.Error("should be secure with error url")
	}
}

func Test_Secure_NameDescUrlErrUrlExamples_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.NameDescUrlErrUrlExamples("n", "d", "u", "eu", "e1")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_NameDescExamples_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.NameDescExamples("n", "d", "e1")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_Examples_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.Examples("n", "d", "e1")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_ExamplesOnly_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.ExamplesOnly("e1")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_UrlOnly_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.UrlOnly("u")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_ErrorUrlOnly_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.ErrorUrlOnly("eu")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_HintUrlOnly_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.HintUrlOnly("hu")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_DescHintUrlOnly_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.DescHintUrlOnly("d", "hu")
	if info.Description != "d" || info.HintUrl != "hu" {
		t.Error("should set desc and hint url")
	}
}

func Test_Secure_NameHintUrlOnly_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.NameHintUrlOnly("n", "hu")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_SingleExampleOnly_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.SingleExampleOnly("se")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_AllUrlExamples_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.AllUrlExamples("n", "d", "u", "hu", "eu", "e1")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_AllUrl_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.AllUrl("n", "d", "u", "hu", "eu")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_UrlSingleExample_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.UrlSingleExample("n", "d", "u", "se")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_SingleExample_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.SingleExample("n", "d", "se")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_ExampleUrl_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.ExampleUrl("n", "d", "exu", "se")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_ExampleUrlSingleExample_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.ExampleUrlSingleExample("n", "d", "exu", "se")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

func Test_Secure_NewExampleUrlSecure_Ext(t *testing.T) {
	info := coretaskinfo.New.Info.Secure.NewExampleUrlSecure("n", "d", "exu", "se")
	if !info.IsSecure() {
		t.Error("should be secure")
	}
}

// ============================================================================
// newInfoCreator: Deserialized / DeserializedUsingJsonResult
// ============================================================================

func Test_NewInfo_Deserialized_Ext(t *testing.T) {
	original := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	bytes, _ := original.Serialize()
	result, err := coretaskinfo.New.Info.Deserialized(bytes)
	if err != nil {
		t.Errorf("Deserialized error: %v", err)
	}
	if result.RootName != "n" {
		t.Error("should preserve root name")
	}
}

func Test_NewInfo_DeserializedUsingJsonResult_Ext(t *testing.T) {
	original := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	jsonResult := original.JsonPtr()
	result, err := coretaskinfo.New.Info.DeserializedUsingJsonResult(jsonResult)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if result.RootName != "n" {
		t.Error("should preserve root name")
	}
}

func Test_Info_Deserialize_Ext(t *testing.T) {
	original := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	var target coretaskinfo.Info
	err := original.Deserialize(&target)
	if err != nil {
		t.Errorf("Deserialize error: %v", err)
	}
	if target.RootName != "n" {
		t.Error("should preserve root name")
	}
}

func Test_Info_JsonParseSelfInject_Ext(t *testing.T) {
	original := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	jsonResult := original.JsonPtr()
	var target coretaskinfo.Info
	err := target.JsonParseSelfInject(jsonResult)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}
