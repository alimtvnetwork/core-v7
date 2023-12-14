package codegen

import (
	"gitlab.com/auk-go/core/coreutils/stringutil"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

var (
	NewCodeOutput = newCodeOutputCreator{}

	vars = unitVariables{
		PackageName:      "$packageName",
		NewPackages:      "$newPackages",
		FuncName:         "$FuncName",
		ArrangeType:      "$ArrangeType",
		LinesPossible:    "$linesPossible",
		ActArgsSetup:     "$actArgsSetup",
		InArgs:           "$inArgs",
		OutArgs:          "$outArgs",
		FmtJoin:          "$fmtJoin",
		FmtOutputs:       "$fmtOutputs",
		Behaviour:        "$Behaviour",
		TestCaseName:     "$testCaseName",
		DirectFuncInvoke: "$directFuncInvoke",
		Title:            "$title",
		ArrangeSetup:     "$arrangeSetup",
		ExpectedLines:    "$expectedLines",
		CaseItem:         "$caseItem",
		TestCases:        "$testCases",
		VerifyTypeOf:     "$VerifyTypeOf",
		VariablesSetup:   "$variablesSetup",
		workFunc:         "WorkFunc",
		expect:           "Expect",
		inputExpectedVar: "input.Expect",
	}

	templateReplacerFunc = stringutil.
				ReplaceTemplate.
				DirectKeyUsingMapTrim

	pascalCaseFunc = convertinteranl.Util.String.PascalCase
	camelCaseFunc  = convertinteranl.Util.String.CamelCase
)
