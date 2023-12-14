package codegen

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/codegen/codegentype"
	"gitlab.com/auk-go/core/codegen/fmtcodegentype"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/coreindexes"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/internal/pathinternal"
	"gitlab.com/auk-go/core/internal/reflectinternal"
	"gitlab.com/auk-go/core/isany"
	"gitlab.com/auk-go/core/iserror"
)

type GenerateFunc struct {
	Func                    interface{}
	Struct                  interface{}
	FuncOverrideCall        string
	GenerateType            codegentype.Variant
	FmtType                 fmtcodegentype.Variant
	TestCases               []coretestcases.CaseV1
	Behaviours              corestr.SimpleSlice
	UnitTestRootPath        string
	OverridingTestPkgName   string
	IsGenerateSeparateCases bool
	IsIncludeFunction       bool
	IsOverwrite             bool
	packageHeader           corestr.SimpleStringOnce
	funcWrap                *args.FuncWrap
	setupVariable           *variablesSetup
}

func (it GenerateFunc) Function() interface{} {
	return it.Func
}

func (it GenerateFunc) CurStruct() interface{} {
	return it.Struct
}

func (it GenerateFunc) GenType() codegentype.Variant {
	return it.GenerateType
}

func (it GenerateFunc) JoinFormatType() fmtcodegentype.Variant {
	return it.FmtType
}

func (it GenerateFunc) Cases() []coretestcases.CaseV1 {
	return it.TestCases
}

func (it GenerateFunc) CurBehaviours() corestr.SimpleSlice {
	return it.Behaviours
}

func (it GenerateFunc) CurFuncOverrideCall() interface{} {
	return it.FuncOverrideCall
}

func (it GenerateFunc) IsFunctionIncluded() bool {
	return it.IsIncludeFunction
}

func (it GenerateFunc) Generate() error {
	codeOutput := it.GenerateCodeOutput()

	return codeOutput.Write().CompiledError()
}

func (it GenerateFunc) GenerateCodeOutput() *CodeOutput {
	toWrap := it.FuncWrap()

	if toWrap.IsInvalid() {
		return NewCodeOutput.Invalid(toWrap.InvalidError())
	}

	inArgs, inArgsErr := it.InArgs()

	if iserror.Defined(inArgsErr) {
		return NewCodeOutput.Invalid(inArgsErr)
	}

	outArgs, outArgsErr := it.OutArgs()

	if iserror.Defined(outArgsErr) {
		return NewCodeOutput.Invalid(outArgsErr)
	}

	funcName := toWrap.GetFuncName()
	firstArrangeTypeName := it.FirstArrangeTypeName()
	fmtOutputs, fmtErr := it.generateFmtOutputs(
		fmtJoiner,
		funcName,
		"",
		outArgs,
		inArgs,
	)

	if iserror.Defined(fmtErr) {
		return NewCodeOutput.Invalid(fmtErr)
	}

	funcTemplateReplacer := map[string]string{
		vars.FuncName:         funcName,
		vars.ArrangeType:      firstArrangeTypeName,
		vars.LinesPossible:    totalSliceLength,
		vars.InArgs:           inArgs.Join(ArgsJoiner),
		vars.OutArgs:          outArgs.Join(ArgsJoiner),
		vars.FmtJoin:          it.FmtJoin(),
		vars.FmtOutputs:       fmtOutputs.Join(fmtJoiner),
		vars.DirectFuncInvoke: it.DirectFuncInvokeName(),
	}

	unitTests, unitErr := it.UnitTests(
		inArgs,
		outArgs,
		funcTemplateReplacer,
	)

	if iserror.Defined(unitErr) {
		return NewCodeOutput.Invalid(unitErr)
	}

	unitTestCode := unitTests.JoinLine()
	optimizeHeaderBasedOnCode := it.GetOptimizePackageHeader(unitTestCode)

	finalUnitTest := stringslice.Joins(
		constants.NewLineUnix,
		optimizeHeaderBasedOnCode,
		"",
		unitTestCode,
		"",
	)

	testCaseCompiled, testCaseErr := it.TestCasesCompiled()

	return &CodeOutput{
		UnitTest:   finalUnitTest,
		TestCase:   testCaseCompiled,
		StructName: it.StructName(),
		FuncName:   funcName,
		Error:      testCaseErr,
		FileWriter: it.fileWriter(it.TestPkgName()),
	}
}

func (it GenerateFunc) UnitTests(
	inArgs,
	outArgs *corestr.SimpleSlice,
	tempMap map[string]string,
) (*corestr.SimpleSlice, error) {
	totalBehaviours := len(it.Behaviours)
	testsSlice := corestr.
		New.
		SimpleSlice.
		Cap(totalBehaviours)

	if totalBehaviours == 0 {
		return testsSlice, errors.New("must set behaviours it cannot be empty")
	}

	funcName := it.FuncName()

	for _, behaviour := range it.Behaviours {
		fmtOutputs, fmtErr := it.generateFmtOutputs(
			fmtJoiner,
			funcName,
			vars.inputExpectedVar,
			outArgs,
			inArgs,
		)

		tempMap[vars.FmtOutputs] = fmtOutputs.Join(fmtJoiner)
		tempMap[vars.Behaviour] = behaviour
		tempMap[vars.TestCaseName] = it.TestCaseName(
			totalBehaviours,
			funcName,
			behaviour,
		)
		tempMap[vars.VariablesSetup] = it.CompiledVariablesSetup()
		if iserror.Defined(fmtErr) {
			return testsSlice, fmtErr
		}

		unitTest := it.ReplaceTemplate(
			funcTemplate,
			tempMap,
		)

		testsSlice.Add(unitTest)
	}

	return testsSlice, nil
}

func (it GenerateFunc) TestCaseName(
	totalBehaviours int,
	funcName,
	behaviour string,
) string {
	if totalBehaviours == 1 {
		return camelCaseFunc(
			fmt.Sprintf(
				"%sTestCases",
				funcName,
			),
		)
	}

	return camelCaseFunc(
		fmt.Sprintf(
			"%sTestCases%s",
			funcName,
			pascalCaseFunc(behaviour),
		),
	)
}

func (it GenerateFunc) PackageHeader() (testPkgName string, packageHeader string) {
	testPkgName = it.TestPkgName()

	if it.packageHeader.IsDefined() {
		return testPkgName, it.packageHeader.String()
	}

	newPackagesLines := it.AllPackages()
	packagesTemplate := map[string]string{
		"$packageName": testPkgName,
		"$newPackages": newPackagesLines,
	}

	packageHeader = it.ReplaceTemplate(
		testPkgHeaderTemplate,
		packagesTemplate,
	)

	return testPkgName, it.packageHeader.GetSetOnce(packageHeader)
}

func (it GenerateFunc) fileWriter(unitTestPackageName string) *chmodhelper.SimpleFileReaderWriter {
	finalUnitTestPath := it.unitTestRootPath(unitTestPackageName)

	return chmodhelper.
		New.
		SimpleFileReaderWriter.
		Options(
			true,
			true,
			true,
			finalUnitTestPath,
		)
}

func (it GenerateFunc) unitTestRootPath(unitTestPackageName string) string {
	return pathinternal.Join(
		it.UnitTestRootPath,
		unitTestPackageName,
		"x.go", // we are writing to the parent dir
	)
}

func (it GenerateFunc) FirstArrangeTypeName() string {
	if len(it.TestCases) == 0 {
		return ""
	}

	return convertinteranl.AnyTo.TypeName(
		it.TestCases[0].ArrangeInput,
	)
}

func (it GenerateFunc) FirstTestCase() *coretestcases.CaseV1 {
	if len(it.TestCases) == 0 {
		return nil
	}

	return &it.TestCases[0]
}

func (it GenerateFunc) AllPackages() string {
	arrangePkgPaths := it.ArrangePackages()

	newPackages := corestr.
		New.
		SimpleSlice.
		Hashset(arrangePkgPaths).
		Add(it.FuncWrap().PkgPath()).
		WrapDoubleQuote()

	newPackagesLines := newPackages.JoinLine()

	return newPackagesLines
}

func (it GenerateFunc) FirstArrangeType() *reflect.Type {
	if len(it.TestCases) == 0 {
		return nil
	}

	rt := reflect.TypeOf(
		it.TestCases[0].ArrangeInput,
	)

	return &rt
}

func (it GenerateFunc) ArrangeReflectTypes() []reflect.Type {
	var results []reflect.Type

	reducerFunc := reflectinternal.Looper.ReducePointerDefault

	for _, testCase := range it.TestCases {
		r := reducerFunc(testCase)

		if r.IsInvalid() {
			continue
		}

		results = append(
			results,
			r.FinalReflectVal.Type(),
		)
	}

	return results
}

func (it GenerateFunc) ArrangePackages() *corestr.Hashset {
	allReflectTypes := it.ArrangeReflectTypes()

	pks := corestr.New.Hashset.Cap(len(allReflectTypes))

	for _, reflectType := range allReflectTypes {
		pks.Add(reflectType.PkgPath())
	}

	return pks
}

func (it GenerateFunc) TestPkgName() string {
	return it.FuncWrap().PkgNameOnly() + "tests"
}

func (it GenerateFunc) FuncWrap() *args.FuncWrap {
	if it.funcWrap != nil {
		return it.funcWrap
	}

	it.funcWrap = args.
		NewFuncWrap.
		Default(it.Func)

	return it.funcWrap
}

func (it GenerateFunc) generateVariablesSetup() *corestr.SimpleSlice {
	return nil
}

func (it GenerateFunc) FmtJoin() string {
	return it.FmtType.Fmt()
}

func (it GenerateFunc) generateFmtOutputs(
	joiner string,
	funcName string,
	expected string,
	outArs, inArgs *corestr.SimpleSlice,
) (*corestr.SimpleSlice, error) {
	slice := corestr.New.SimpleSlice.Cap(20)

	switch it.FmtType {
	case fmtcodegentype.Default: // "%d : %s -> %s",
		outArgsString := outArs.Join(joiner)
		inArgsString := inArgs.Join(joiner)
		slice.Add(inArgsString)
		slice.Add(outArgsString)

		return slice, nil
	case fmtcodegentype.WithFunction: // "%d : %s(%s) -> %s | %s",
		outArgsString := outArs.Join(joiner)
		inArgsString := inArgs.Join(joiner)
		slice.Add(funcName)
		slice.Add(inArgsString)
		slice.Add(outArgsString)
		slice.Add(expected)

		return slice, nil
	}

	return slice, it.FmtType.OnlySupportedMsgErr(
		"only supported",
		fmtcodegentype.Default.Name(),
		fmtcodegentype.WithFunction.Name(),
	)
}

func (it GenerateFunc) FuncName() string {
	funcWrap := it.FuncWrap()

	if funcWrap.IsInvalid() {
		return ""
	}

	return funcWrap.GetFuncName()
}

// OutArgs
//
//	Aka returns Args
//
// - if one then return "result" only
// - Or else, result1, result2 ...
func (it GenerateFunc) OutArgs() (*corestr.SimpleSlice, error) {
	funcWrap := it.FuncWrap()

	if funcWrap.IsInvalid() {
		return it.emptySlice(), errors.New("func wrap is invalid - return args")
	}

	length := funcWrap.ReturnLength()
	slice := corestr.New.SimpleSlice.Cap(length)

	if length == 1 {
		return slice.Add("result"), nil
	}

	for i := 0; i < length; i++ {
		slice.AppendFmt("result%d", i+1)
	}

	return slice, nil
}

// InArgs
//
// - if one then return "result" only
// - Or else, result1, result2 ...
func (it GenerateFunc) InArgs() (*corestr.SimpleSlice, error) {
	funcWrap := it.FuncWrap()

	if funcWrap.IsInvalid() {
		return it.emptySlice(), errors.New("func wrap is invalid - return args")
	}

	length := funcWrap.ArgsCount()
	slice := corestr.New.SimpleSlice.Cap(length)

	if length == 0 {
		return slice, nil
	}

	return &it.VariablesSetup().inArgsNames, nil
}

// VariableName
//
// variable.First or variable.Second ... based on index.
func (it GenerateFunc) VariableName(parentVar string, index int) string {
	return parentVar + "." + it.indexByName(index)
}

func (it GenerateFunc) indexByName(index int) string {
	return coreindexes.NameByIndex(index)
}

func (it GenerateFunc) emptySlice() *corestr.SimpleSlice {
	return corestr.Empty.SimpleSlice()
}

func (it GenerateFunc) DirectFuncInvokeName() string {
	if len(it.FuncOverrideCall) > 0 {
		return it.FuncOverrideCall
	}

	return it.FuncWrap().FuncDirectInvokeName()
}

func (it GenerateFunc) StructName() string {
	if isany.Null(it.Struct) {
		return ""
	}

	return reflectinternal.TypeName(it.Struct)
}

func (it GenerateFunc) ReplaceTemplate(
	format string,
	replacerMap map[string]string,
) string {
	if len(format) == 0 {
		return ""
	}

	return templateReplacerFunc(
		format,
		replacerMap,
	)
}

func (it GenerateFunc) TestCasesCompiled() (string, error) {
	caseGenerator := testCaseGenerator{
		baseGenerator: it.AsBaseGenerator(),
	}

	return caseGenerator.Compile()
}

func (it *GenerateFunc) VariablesSetup() *variablesSetup {
	if it.setupVariable != nil {
		return it.setupVariable
	}

	generator := variablesGenerator{
		baseGenerator: it.AsBaseGenerator(),
	}

	vs := generator.Generate()
	it.setupVariable = &vs

	return it.setupVariable
}

func (it GenerateFunc) CompiledVariablesSetup() string {
	return it.VariablesSetup().CompiledSetupLine()
}

func (it GenerateFunc) AsBaseGenerator() BaseGenerator {
	return it
}

func (it GenerateFunc) GetOptimizePackageHeader(code string) string {
	_, packageHeader := it.PackageHeader()

	headerLines := corestr.New.SimpleSlice.SplitLines(packageHeader)
	isImportStarted := false
	var removeIndexes []int

	for i, h := range headerLines.List() {
		h = strings.TrimSpace(h)
		if !isImportStarted && strings.HasPrefix(h, "import") {
			isImportStarted = true

			continue
		}

		if !isImportStarted {
			continue
		}

		if h == ")" || h == "" {
			continue
		}

		// after import
		_, pkgName := GetPkgName(h)
		pkgNameNext := pkgName + "."

		if !strings.Contains(code, pkgNameNext) {
			removeIndexes = append(removeIndexes, i)
		}
	}

	lines, err := headerLines.RemoveIndexes(removeIndexes...)
	errcore.HandleErr(err)

	return lines.JoinLine()
}
