package main

import (
	"fmt"

	"gitlab.com/auk-go/core/codegen"
	"gitlab.com/auk-go/core/codegen/codegentype"
	"gitlab.com/auk-go/core/codegen/fmtcodegentype"
	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type unitTestGenerator struct{}

func (it unitTestGenerator) Generate() {
	curFunc := reflectinternal.GetFunc.PascalFuncName

	generateFunc := codegen.GenerateFunc{
		Func:         curFunc,
		GenerateType: codegentype.Simple,
		FmtType:      fmtcodegentype.Default,
		TestCases: []coretestcases.CaseV1{
			{
				Title: "Some",
				ArrangeInput: args.One{
					First:  "someName",
					Expect: nil,
				},
				ActualInput:     nil,
				ExpectedInput:   nil,
				Additional:      nil,
				CustomFormat:    "",
				VerifyTypeOf:    nil,
				Parameters:      nil,
				IsEnable:        0,
				HasError:        false,
				HasPanic:        false,
				IsValidateError: false,
			},
		},
		Behaviours: []string{
			"Verification",
		},
		UnitTestRootPath:        codestack.Dir.CurDirJoin("unit-test"),
		IsGenerateSeparateCases: false,
		IsIncludeFunction:       true,
		IsOverwrite:             true,
	}

	wrap := args.NewFuncWrap.Single(curFunc)
	inArsgTypes := wrap.GetInArgsTypes()

	fmt.Println(inArsgTypes)

	err := generateFunc.Generate()

	errcore.HandleErr(err)
}
