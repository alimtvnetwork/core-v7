package errcoreinf

import (
	"reflect"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/coreinterface/enuminf"
)

type ShouldBeMessager interface {
	Title() string
	Actual() interface{}
	Expected() interface{}
	GenericErrorCompiler

	IsTitleEqual(title string) bool
	IsEqualInterface(right ShouldBeMessager) bool

	CloneNewInterface() ShouldBeMessager
	ConcatNewInterface(another ShouldBeMessager) ShouldBeMessager
}

type ShouldBeChainCollectionDefiner interface {
	GenericErrorCompiler

	ListShouldBeChainCollectionDefiner() []ShouldBeMessager
	Strings() []string
}

type AnyShouldBer interface {
	AnyShouldBe(
		title string,
		actual, expected interface{},
	) BaseErrorOrCollectionWrapper
}

type checkerShouldBer interface {
	IsEmptyShouldBeTrue(
		actual coreinterface.IsEmptyChecker,
	) ShouldBeChainer

	IsEmptyShouldBeFalse(
		actual coreinterface.IsEmptyChecker,
	) ShouldBeChainer

	IsEnableAnyShouldBeTrue(
		actual coreinterface.IsEnableAnyChecker,
	) ShouldBeChainer

	IsEnableAllShouldBeTrue(
		actual coreinterface.IsEnableAllChecker,
	) ShouldBeChainer

	IsDisabledShouldBeTrue(
		actual coreinterface.IsDisabledChecker,
	) ShouldBeChainer

	IsDisableAllShouldBeTrue(
		actual coreinterface.IsDisableAllChecker,
	) ShouldBeChainer

	IsDisableAnyShouldBeTrue(
		actual coreinterface.IsDisableAnyChecker,
	) ShouldBeChainer
}

type stringsShouldBer interface {
	ShouldBe(
		actual, expected []string,
	) ShouldBeChainer

	ShouldBeOptions(
		compareTyper enuminf.StringCompareTyper,
		actual, expected []string,
	) ShouldBeChainer

	LengthShouldBe(
		actual []string,
		lengthExpected int,
	) ShouldBeChainer

	ShouldBeEqualDistinctOptions(
		compareTyper enuminf.StringCompareTyper,
		actual, expected []string,
	) ShouldBeChainer

	ShouldBeEqualRegardlessOrderOptions(
		compareTyper enuminf.StringCompareTyper,
		actual, expected []string,
	) ShouldBeChainer

	DistinctShouldBeOptions(
		compareTyper enuminf.CompareMethodsTyper,
		actual, expected []string,
	) ShouldBeChainer
}

type stringShouldBer interface {
	ShouldBe(
		actual, expected string,
	) ShouldBeChainer

	ShouldBeEmpty(
		actual string,
	) ShouldBeChainer

	ShouldBeWhitespace(
		actual string,
	) ShouldBeChainer

	ShouldStartsWith(
		startsWith,
		actual string,
	) ShouldBeChainer

	ShouldEndsWith(
		endsWith,
		actual string,
	) ShouldBeChainer

	ShouldContains(
		contains string,
		actual string,
	) ShouldBeChainer

	NotStartsWith(
		startsWith,
		actual string,
	) ShouldBeChainer

	NotEndsWith(
		endsWith,
		actual string,
	) ShouldBeChainer

	NotContains(
		contains string,
		actual string,
	) ShouldBeChainer

	ShouldBeOptions(
		compareTyper enuminf.StringCompareTyper,
		actual, expected string,
	) ShouldBeChainer

	ShouldBeDefined(
		title string,
		actual string,
	) ShouldBeChainer

	ShouldBeEqualByFunc(
		title string,
		actual, expected string,
		compareFunc func(actual, expected string) (isMatch bool),
	) ShouldBeChainer
}

type errorShouldBer interface {
	BaseErrShouldBeEmpty(
		title string,
		actual BaseErrorOrCollectionWrapper,
	) ShouldBeChainer

	ShouldBeDefined(
		title string,
		actual error,
	) ShouldBeChainer

	ShouldBeEmpty(
		title string,
		actual error,
	) ShouldBeChainer

	ErrorDefined(actual error) ShouldBeChainer
	ErrorEmpty(actual error) ShouldBeChainer
}

type ShouldBeChainer interface {
	On(isCollect bool) ShouldBeChainer
	OnString(actual, expected string) ShouldBeChainer
	OnNull(anyItem interface{}) ShouldBeChainer
	OnDefined(anyItem interface{}) ShouldBeChainer

	IsCompleted() bool
	IsFrozen() bool
	IsAddPossible() bool

	Title(title string) ShouldBeChainer

	JsonerShouldBe(
		title string,
		actual, expected corejson.Jsoner,
	) ShouldBeChainer

	JsonerShouldBeDefined(
		title string,
		actual corejson.Jsoner,
	) ShouldBeChainer

	IntegerShouldBeDefined(
		title string,
		actual int,
	) ShouldBeChainer

	ShouldBeEmptyString(
		title string,
		actual string,
	) ShouldBeChainer

	ShouldBeEmptyInteger(
		title string,
		actual int,
	) ShouldBeChainer

	ShouldBeEmptyByte(
		title string,
		actual int,
	) ShouldBeChainer

	ShouldBeFalse(
		title string,
		actual bool,
	) ShouldBeChainer

	ShouldBeTrue(
		title string,
		actual bool,
	) ShouldBeChainer

	JsonResultShouldBe(
		title string,
		actual, expected *corejson.Result,
	) ShouldBeChainer

	JsonResultShouldBeDefined(
		title string,
		actual *corejson.Result,
	) ShouldBeChainer

	JsonResultShouldHaveNoError(
		title string,
		actual *corejson.Result,
	) ShouldBeChainer

	JsonResultShouldHaveNoIssuesOrEmpty(
		title string,
		actual *corejson.Result,
	) ShouldBeChainer

	IntegerShouldBeGreater(
		title string,
		actual, expected int,
	) ShouldBeChainer

	IntegerShouldBeGreaterOrEqual(
		title string,
		actual, expected int,
	) ShouldBeChainer

	BytesShouldBe(
		title string,
		actual, expected []byte,
	) ShouldBeChainer

	BytesShouldBeDefined(
		title string,
		actual []byte,
	) ShouldBeChainer

	TypeShouldBe(
		title string,
		actual, expected reflect.Type,
	) ShouldBeChainer

	TypeShouldBeAnyOf(
		title string,
		actual reflect.Type,
		expectedTypes ...reflect.Type,
	) ShouldBeChainer

	ShouldBeSuccess(
		title string,
		actual coreinterface.IsSuccessValidator,
	) ShouldBeChainer

	ShouldBeFailed(
		title string,
		actual coreinterface.IsSuccessValidator,
	) ShouldBeChainer

	ShouldBeValid(
		title string,
		actual coreinterface.IsSuccessValidator,
	) ShouldBeChainer

	PointerShouldBe(
		title string,
		actual, expected interface{},
	) ShouldBeChainer

	IntegerShouldBe(
		title string,
		actual, expected int,
	) ShouldBeChainer

	IntegerShouldBeLessThan(
		title string,
		actual, expected int,
	) ShouldBeChainer

	IntegerShouldBeCompare(
		numberCompareTyper enuminf.CompareMethodsTyper,
		title string,
		actual, expected int,
	) ShouldBeChainer

	ByteShouldBe(
		title string,
		actual, expected byte,
	) ShouldBeChainer

	ChainerShouldBeEmpty(
		title string,
		actual ShouldBeChainer,
	) ShouldBeChainer

	LengthShouldBe(
		title string,
		actual coreinterface.LengthGetter,
		expected int,
	) ShouldBeChainer

	ShouldBeHaveItem(
		title string,
		actual coreinterface.LengthGetter,
	) ShouldBeChainer

	SimpleEnumShouldBe(
		title string,
		actual, expected enuminf.SimpleEnumer,
	) ShouldBeChainer

	BasicEnumShouldBe(
		title string,
		actual, expected enuminf.BasicEnumer,
	) ShouldBeChainer

	BasicEnumShouldBeInvalid(
		title string,
		actual enuminf.BasicEnumer,
	) ShouldBeChainer

	BasicEnumShouldBeDefined(
		title string,
		actual enuminf.BasicEnumer,
	) ShouldBeChainer

	BooleanShouldBe(
		title string,
		actual, expected bool,
	) ShouldBeChainer

	MapStringAnyShouldBe(
		title string,
		actual, expected map[string]interface{},
	) ShouldBeChainer

	Error() errorShouldBer

	ShouldBe(
		title string,
		actual, expected interface{},
	) ShouldBeChainer

	ShouldBeOn(
		isCollect bool,
		title string,
		actual, expected interface{},
	) ShouldBeChainer

	ShouldBeRegardlessOn(
		isCollect bool,
		title string,
		actual, expected interface{},
	) ShouldBeChainer

	ShouldBeOption(
		isRegardless bool,
		title string,
		actual, expected interface{},
	) ShouldBeChainer

	ShouldBeOptionOn(
		isCollect bool,
		isRegardless bool,
		title string,
		actual, expected interface{},
	) ShouldBeChainer

	ShouldBeRegardless(
		title string,
		actual, expected interface{},
	) ShouldBeChainer

	ShouldBeUsingFunc(
		title string,
		actual, expected interface{},
		compareFunc func(actual, expected interface{}) (isMatch bool),
	) ShouldBeChainer

	ShouldBeHaveNoPanicFunc(
		title string,
		actual, expected interface{},
		recoverPanicCompareFunc func(actual, expected interface{}) (isMatch bool),
	) ShouldBeChainer

	ShouldBeDefined(
		title string,
		actual interface{},
	) ShouldBeChainer

	ManyShouldBeDefined(
		title string,
		actualItems ...interface{},
	) ShouldBeChainer

	ShouldBeNull(
		title string,
		actual interface{},
	) ShouldBeChainer

	PointerShouldBeNull(
		title string,
		actual interface{},
	) ShouldBeChainer

	Equal(actual, expected interface{}) ShouldBeChainer
	EqualFunc(
		actual, expected interface{},
		compareFunc func(actual, expected interface{}) (isMatch bool),
	) ShouldBeChainer
	EqualOption(isRegardless bool, actual, expected interface{}) ShouldBeChainer
	EqualRegardless(actual, expected interface{}) ShouldBeChainer

	Checker(title string) checkerShouldBer
	Str(title string) stringShouldBer
	Strs(title string) stringsShouldBer

	Compile() BaseErrorOrCollectionWrapper
	CompileErr() error
	MustBeEmptier
	GenericErrorCompiler

	HandleError()
	CompileString() string
	CompileJson() string
	IsCollected() bool

	Append(anotherItems ...ShouldBeChainer) ShouldBeChainer
	AppendNew(anotherItems ...ShouldBeChainer) ShouldBeChainer

	LogOnIssues() (logged string)

	Strings() []string
}
