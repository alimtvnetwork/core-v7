package args

import (
	"fmt"

	"gitlab.com/auk-go/core/coreinterface"
)

type FuncWrapGetter interface {
	FuncWrap() *FuncWrap
}

type FuncNumber interface {
	GetWorkFunc() interface{}
	coreinterface.FuncByIndexParameter
	FuncWrapGetter
}

type FuncNamer interface {
	GetWorkFunc() interface{}
	coreinterface.FuncByNameParameter
	FuncWrapGetter
}

type OneParameter interface {
	ArgBaseContractsBinder
	AsArgBaseContractsBinder
	coreinterface.OneParameter
}

type OneFuncParameter interface {
	ArgFuncContractsBinder
	AsArgFuncContractsBinder
	OneParameter
	FuncNumber
}

type TwoParameter interface {
	ArgBaseContractsBinder
	OneParameter
	coreinterface.TwoParameter
}

type TwoFuncParameter interface {
	OneFuncParameter
	TwoParameter
	FuncNumber
}

type ThreeParameter interface {
	TwoParameter
	coreinterface.ThreeParameter
}

type ThreeFuncParameter interface {
	TwoFuncParameter
	ThreeParameter
	FuncNumber
}

type FourParameter interface {
	ThreeParameter
	coreinterface.FourthParameter
}

type FourFuncParameter interface {
	ThreeFuncParameter
	FourParameter
	FuncNumber
}

type FifthParameter interface {
	FourParameter
	coreinterface.FifthParameter
}

type FifthFuncParameter interface {
	FourFuncParameter
	FifthParameter
	FuncNumber
}

type SixthParameter interface {
	FifthParameter
	coreinterface.SixthParameter
}

type SixthFuncParameter interface {
	FifthFuncParameter
	SixthParameter
	FuncNumber
}

type ArgsMapper interface {
	ArgBaseContractsBinder

	coreinterface.FirstItemGetter
	coreinterface.ExpectGetter
	HasFirst() bool
	coreinterface.HasExpectChecker
	coreinterface.ValidArgsGetter
	coreinterface.SliceGetter
	coreinterface.ByIndexGetter
	coreinterface.UptoSixthItemGetter

	FuncNamer
}

type FuncWrapper interface {
	coreinterface.FuncWrapContractsBinder
	InvalidError() error
	IsEqual(
		another *FuncWrap,
	) bool
	IsNotEqual(
		another *FuncWrap,
	) bool
}

type HasFirstChecker interface {
	HasFirst() bool
}

type ArgBaseContractsBinder interface {
	coreinterface.FirstItemGetter
	coreinterface.ExpectGetter
	HasFirstChecker

	coreinterface.HasExpectChecker
	coreinterface.ValidArgsGetter
	coreinterface.SliceGetter
	coreinterface.ByIndexGetter

	coreinterface.ArgsCountGetter

	fmt.Stringer
}

type ArgFuncContractsBinder interface {
	ArgBaseContractsBinder
	FuncNumber
}

type AsArgBaseContractsBinder interface {
	AsArgBaseContractsBinder() ArgBaseContractsBinder
}

type AsArgFuncContractsBinder interface {
	AsArgFuncContractsBinder() ArgFuncContractsBinder
}

type ArgFuncNameContractsBinder interface {
	ArgBaseContractsBinder
	FuncNamer
}

type AsArgFuncNameContractsBinder interface {
	AsArgFuncNameContractsBinder() ArgFuncNameContractsBinder
}

type AsOneFuncParameter interface {
	AsOneFuncParameter() OneFuncParameter
}

type AsTwoFuncParameter interface {
	AsTwoFuncParameter() TwoFuncParameter
}

type AsThreeFuncParameter interface {
	AsThreeFuncParameter() ThreeFuncParameter
}

type AsFourFuncParameter interface {
	AsFourFuncParameter() FourFuncParameter
}

type AsFifthFuncParameter interface {
	AsFifthFuncParameter() FifthFuncParameter
}

type AsSixthFuncParameter interface {
	AsSixthFuncParameter() SixthFuncParameter
}

type AsOneParameter interface {
	AsOneParameter() OneParameter
}

type AsTwoParameter interface {
	AsTwoParameter() TwoParameter
}

type AsThreeParameter interface {
	AsThreeParameter() ThreeParameter
}

type AsFourParameter interface {
	AsFourParameter() FourParameter
}

type AsFifthParameter interface {
	AsFifthParameter() FifthParameter
}

type AsSixthParameter interface {
	AsSixthParameter() SixthParameter
}
