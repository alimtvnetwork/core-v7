package corejson

type newMapResultsCreator struct{}

func (it newMapResultsCreator) Empty() *MapResults {
	return &MapResults{
		Items: map[string]Result{},
	}
}

func (it newMapResultsCreator) UsingCap(
	addCapacity int,
) *MapResults {
	return &MapResults{
		Items: make(
			map[string]Result,
			addCapacity),
	}
}

func (it newMapResultsCreator) UsingKeyAnyItems(
	addCapacity int,
	keyAnyItems ...KeyAny,
) *MapResults {
	length := addCapacity + len(keyAnyItems)

	if length == 0 || len(keyAnyItems) == 0 {
		return it.UsingCap(length)
	}

	collection := it.UsingCap(length + addCapacity)

	return collection.
		AddKeyAnyItems(keyAnyItems...)
}

func (it newMapResultsCreator) UsingMapOptions(
	isClone, isDeepClone bool,
	addCapacity int,
	mapResults map[string]Result,
) *MapResults {
	if len(mapResults) == 0 {
		return it.UsingCap(
			addCapacity)
	}

	hasNoChange :=
		addCapacity == 0 &&
			!isClone &&
			!isDeepClone

	if hasNoChange {
		return &MapResults{
			Items: mapResults,
		}
	}

	additionalCapacity :=
		len(mapResults) +
			addCapacity

	finalMapResults := it.UsingCap(
		additionalCapacity)

	return finalMapResults.AddMapResultsUsingCloneOption(
		isClone,
		isDeepClone,
		mapResults)
}

func (it newMapResultsCreator) UsingMapPlusCap(
	addCapacity int,
	mapResults map[string]Result,
) *MapResults {
	if len(mapResults) == 0 {
		return it.UsingCap(
			addCapacity)
	}

	return it.UsingMapOptions(
		false,
		false,
		addCapacity,
		mapResults)
}

func (it newMapResultsCreator) UsingMapPlusCapClone(
	addCapacity int,
	mapResults map[string]Result,
) *MapResults {
	if len(mapResults) == 0 {
		return it.UsingCap(
			addCapacity)
	}

	return it.UsingMapOptions(
		true,
		false,
		addCapacity,
		mapResults)
}

func (it newMapResultsCreator) UsingMapPlusCapDeepClone(
	addCapacity int,
	mapResults map[string]Result,
) *MapResults {
	if len(mapResults) == 0 {
		return it.UsingCap(
			addCapacity)
	}

	return it.UsingMapOptions(
		true,
		true,
		addCapacity,
		mapResults)
}

func (it newMapResultsCreator) UsingMap(
	mapResults map[string]Result,
) *MapResults {
	if len(mapResults) == 0 {
		return it.Empty()
	}

	return it.UsingMapOptions(
		false,
		false,
		0,
		mapResults)
}

func (it newMapResultsCreator) UsingMapAnyItemsPlusCap(
	addCapacity int,
	mapAnyItems map[string]interface{},
) *MapResults {
	if len(mapAnyItems) == 0 {
		return it.UsingCap(addCapacity)
	}

	collection := it.UsingCap(
		addCapacity + len(mapAnyItems))

	return collection.
		AddMapAnyItems(mapAnyItems)
}

func (it newMapResultsCreator) UsingMapAnyItems(
	mapAnyItems map[string]interface{},
) *MapResults {
	return it.UsingMapAnyItemsPlusCap(
		0,
		mapAnyItems)
}

func (it newMapResultsCreator) UsingKeyWithResultsPlusCap(
	addCapacity int,
	keyWithResults ...KeyWithResult,
) *MapResults {
	if keyWithResults == nil {
		return it.UsingCap(addCapacity)
	}

	mapResults := it.UsingCap(
		addCapacity + len(keyWithResults))

	return mapResults.
		AddKeysWithResults(keyWithResults...)
}

func (it newMapResultsCreator) UsingKeyWithResults(
	keyWithResults ...KeyWithResult,
) *MapResults {
	return it.UsingKeyWithResultsPlusCap(
		0,
		keyWithResults...)
}

func (it newMapResultsCreator) UsingKeyJsonersPlusCap(
	addCapacity int,
	keyWithJsoners ...KeyWithJsoner,
) *MapResults {
	if keyWithJsoners == nil {
		return it.UsingCap(addCapacity)
	}

	mapResults := it.UsingCap(
		addCapacity + len(keyWithJsoners))

	return mapResults.
		AddKeysWithJsoners(keyWithJsoners...)
}

func (it newMapResultsCreator) UsingKeyJsoners(
	keyWithJsoners ...KeyWithJsoner,
) *MapResults {
	return it.UsingKeyJsonersPlusCap(
		0,
		keyWithJsoners...)
}
