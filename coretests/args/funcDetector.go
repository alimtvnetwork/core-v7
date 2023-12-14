package args

type funcDetector struct{}

func (it funcDetector) GetFuncWrap(i interface{}) *FuncWrap {
	switch v := i.(type) {
	case Map:
		return v.FuncWrap()
	case *FuncWrap:
		return v
	case FuncWrapGetter:
		return v.FuncWrap()
	case ArgsMapper:
		return v.FuncWrap()
	default:
		return NewFuncWrap.Default(i)
	}
}
