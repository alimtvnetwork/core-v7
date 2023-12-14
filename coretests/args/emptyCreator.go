package args

type emptyCreator struct{}

func (it emptyCreator) Map() Map {
	return map[string]interface{}{}
}

func (it emptyCreator) FuncWrap() *FuncWrap {
	return &FuncWrap{
		isInvalid: true,
	}
}

func (it emptyCreator) FuncMap() FuncMap {
	return map[string]FuncWrap{}
}
