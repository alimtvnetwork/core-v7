package corestr

type OnCompleteCharCollectionMap func(charCollection *CharCollectionMap)
type OnCompleteCharHashsetMap func(charHashset *CharHashsetMap)
type IsStringFilter func(str string) (result string, isKeep bool)
type IsKeyAnyValueFilter func(pair KeyAnyValuePair) (result string, isKeep bool)
type IsKeyValueFilter func(pair KeyValuePair) (result string, isKeep bool)
type IsStringPointerFilter func(stringPointer *string) (result *string, isKeep bool)
