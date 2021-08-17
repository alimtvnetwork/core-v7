package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coreinstruction"
	"gitlab.com/evatix-go/core/enums/stringcompareas"
)

func main() {
	// fmt.Println(corerange.StartEndSimpleString{
	// 	Start: "1",
	// 	End:   "2",
	// })
	//
	// a, _ := issetter.Wildcard.MarshalJSON()
	// val2 := issetter.Value(0)
	//
	// fmt.Println(val2.UnmarshalJSON(a))
	// fmt.Println(val2)
	//
	// lineValidator := corevalidator.LineValidator{
	// 	LineNumber: corevalidator.LineNumber{
	// 		LineNumber: 2,
	// 	},
	// 	TextValidator: corevalidator.TextValidator{
	// 		Search: "   alim      alim 2 alim 3                 ",
	// 		ValidatorCoreCondition: corevalidator.ValidatorCoreCondition{
	// 			IsTrimCompare:        false,
	// 			IsUniqueWordOnly:     false,
	// 			IsNonEmptyWhitespace: false,
	// 			IsSortStringsBySpace: false,
	// 		},
	// 		SearchAs: stringcompareas.Equal,
	// 	},
	// }
	//
	// params := corevalidator.ValidatorParamsBase{
	// 	CaseIndex:                         0,
	// 	IsIgnoreCompareOnActualInputEmpty: false,
	// 	IsAttachUserInputs:                false,
	// 	IsCaseSensitive:                   false,
	// }
	// err := lineValidator.VerifyError(
	// 	&params,
	// 	-1,
	// 	"alim      alim 2 alim 4",
	// )
	//
	// // fmt.Println(err)
	//
	// lines := msgtype.ErrorToSplitLines(err)
	// fmt.Println(err)
	//
	// sliceValidator := corevalidator.SliceValidator{
	// 	InputLines: lines,
	// 	ComparingLines: []string{
	// 		"----------------------",
	// 		"2 )\tExpectation failed: Using Method `\"Equal\"`",
	// 		"        Content-Processed:`\"alim      alim 2 alim 4\"`",
	// 		"     SearchTerm-Processed:`\"alim      alim 2 alim 3\"`",
	// 		"Additional:`corevalidator.TextValidator{Search:\"   alim      alim 2 alim 3                 \", " +
	// 			"IsTrimCompare:true, IsSplitByWhitespace:false, " +
	// 			"IsNonEmptyWhitespace:true, IsSortStringsBySpace:false",
	// 	},
	// 	ValidatorCoreCondition: corevalidator.ValidatorCoreCondition{
	// 		IsTrimCompare:        true,
	// 		IsNonEmptyWhitespace: false,
	// 		IsSortStringsBySpace: false,
	// 	},
	// 	CompareAs: stringcompareas.Equal,
	// }
	//
	// params2 := corevalidator.ValidatorParamsBase{
	// 	CaseIndex:                         0,
	// 	IsIgnoreCompareOnActualInputEmpty: false,
	// 	IsAttachUserInputs:                true,
	// 	IsCaseSensitive:                   true,
	// }

	// fmt.Println(sliceValidator.IsValid(false))
	// fmt.Println(sliceValidator.
	// 	AllVerifyErrorExceptLast(
	// 		&params2,
	// 	))

	ins := coreinstruction.NewStringCompare(
		stringcompareas.EndsWith,
		false,
		"hellO",
		"none hello")

	fmt.Println(ins.VerifyError())

	// fmt.Println(stringcompareas.NotEqual.VerifyErrorCaseSensitive(
	// "abc","abc"))

	// mapKeys := map[string]string{
	// 	"Hello":"",
	// 	"Hello2":"",
	// 	"Hello3":"",
	// 	"Hello4":"",
	// 	"Hello5":"",
	// 	"Hello6":"",
	// 	"Hello7":"",
	// }
	//
	// slice2 := []corevalidator.LineValidator{
	// 	{
	// 		LineNumber: corevalidator.LineNumber{},
	// 		TextValidator: corevalidator.TextValidator{
	// 			Search:                        "wdwdwddw",
	// 			IsTrimCompare:                 false,
	// 			IsSplitByWhitespace: false,
	// 			IsNonEmptyWhitespace:          false,
	// 			IsSortStringsBySpace:          false,
	// 			SearchAs:                      0,
	// 		},
	// 	},
	// }

	// keys, err5 := coredynamic.SliceItemsAsStringsAny(slice2)
	//
	// fmt.Println(keys, err5)
	// gherkins := coretests.SimpleGherkins{
	// 	Feature: "ft",
	// 	Given:   "g",
	// 	When:    "w",
	// 	Then:    "t",
	// 	Expect:  "ww",
	// 	Actual:  "rrr",
	// }
	// fmt.Println(gherkins.String())
	// fmt.Println(gherkins.GetWithExpectation(2))
	//
	// fmt.Println(sliceValidator.AllVerifyError(0, true))

	// rwx, err := chmodhelper.NewRwxVariableWrapper("-rwx-*-r*x")
	// fmt.Println(err)
	// fmt.Println(rwx.String())
	//
	// rs := rwx.IsEqualPartialRwxPartial("-rwx-w-r*x")
	// fmt.Println(rs)
	//
	// fmt.Println(msgtype.ExpectingSimpleNoType("Alim", "Rwx", "wrx"))
	//
	// rwxIns := chmodins.RwxInstruction{
	// 	RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
	// 		Owner: "rwx",
	// 		Group: "rwx",
	// 		Other: "r--",
	// 	},
	// 	ValidatorCoreCondition: chmodins.ValidatorCoreCondition{
	// 		IsSkipOnInvalid:  false,
	// 		IsContinueOnError: false,
	// 		IsRecursive:       false,
	// 	},
	// }
	//
	// executor, err := chmodhelper.ParseRwxInstructionToExecutor(&rwxIns)
	//
	// msgtype.SimpleHandleErr(err, "")
	//
	// locations := []string{
	// 	"/temp/core/test-cases-2",
	// 	"/temp/core/test-cases-3s",
	// 	"/temp/core/test-cases-3x",
	// 	"/temp/core/test-cases-3",
	// }
	//
	// err2 := chmodhelper.VerifyChmodLocationsUsingPartialRwx(
	// 	true, true,
	// 	"-rwxrwx",
	// 	locations)
	//
	// err3 := executor.VerifyRwxModifiersDirect(
	// 	false,
	// 	locations...)
	//
	// msgtype.SimpleHandleErrMany("", err2)
}

func PrintCollection(collection *corestr.Collection) {
	fmt.Println(collection.GetPagedCollection(3).String())

	fmt.Print("\n\nTake 5:\n\n")
	fmt.Println(collection.Take(5))
	fmt.Print("\n\n Skip 2:\n\n")
	fmt.Println(collection.Skip(2))
	fmt.Print("\n\n Skip 0:\n\n")
	fmt.Println(collection.Skip(0))
	fmt.Print("\n\n Take 0:\n\n")
	fmt.Println(collection.Take(0))
	fmt.Print("\n\n Skip(5).Take(2):\n\n")
	fmt.Println(collection.Skip(5).Take(2))

}

func PrintCollectionPtr(collectionPtr *corestr.CollectionPtr) {
	fmt.Println(collectionPtr.GetPagedCollection(3).String())

	fmt.Print("\n\nTake 5:\n\n")
	fmt.Println(collectionPtr.Take(5))
	fmt.Print("\n\n Skip 2:\n\n")
	fmt.Println(collectionPtr.Skip(2))
	fmt.Print("\n\n Skip 0:\n\n")
	fmt.Println(collectionPtr.Skip(0))
	fmt.Print("\n\n Take 0:\n\n")
	fmt.Println(collectionPtr.Take(0))
	fmt.Print("\n\n Skip(5).Take(2):\n\n")
	fmt.Println(collectionPtr.Skip(5).Take(2))
}
