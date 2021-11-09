package corestr

type newSimpleStringOnceCreator struct{}

func (it *newSimpleStringOnceCreator) Any(
	isIncludeFieldNames bool,
	value interface{},
	isInitialize bool,
) SimpleStringOnce {
	toString := AnyToString(
		isIncludeFieldNames,
		value)

	return SimpleStringOnce{
		value:        toString,
		isInitialize: isInitialize,
	}
}

func (it *newSimpleStringOnceCreator) Uninitialized(
	value string,
) SimpleStringOnce {
	return SimpleStringOnce{
		value:        value,
		isInitialize: false,
	}
}

func (it *newSimpleStringOnceCreator) Init(
	value string,
) SimpleStringOnce {
	return SimpleStringOnce{
		value:        value,
		isInitialize: true,
	}
}

func (it *newSimpleStringOnceCreator) InitPtr(
	value string,
) *SimpleStringOnce {
	return &SimpleStringOnce{
		value:        value,
		isInitialize: true,
	}
}

func (it *newSimpleStringOnceCreator) Create(
	value string,
	isInitialize bool,
) SimpleStringOnce {
	return SimpleStringOnce{
		value:        value,
		isInitialize: isInitialize,
	}
}

func (it *newSimpleStringOnceCreator) CreatePtr(
	value string,
	isInitialize bool,
) *SimpleStringOnce {
	return &SimpleStringOnce{
		value:        value,
		isInitialize: isInitialize,
	}
}

func (it *newSimpleStringOnceCreator) Empty() SimpleStringOnce {
	return SimpleStringOnce{}
}
