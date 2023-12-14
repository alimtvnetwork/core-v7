# `Core` Intro

![Use Package logo](assets/core-250.png)

All common core infrastructure and constants combined package.

## Git Clone

`git clone https://gitlab.com/auk-go/core.git`

### Prerequisites

- Update git to latest 2.29
- Update or install the latest of Go 1.17.8
- Either add your ssh key to your gitlab account
- Or, use your access token to clone it.

## Installation

`go get gitlab.com/auk-go/core`

## Why `core?`

It makes our other go-packages DRY and concise. It was the first package in the auk-go ecosystem that is core of
everything.

It was first designed for constants, later it got enhanced with:

- [codestack](/codestack)
    - very powerful to deal with codestack
- chmodhelper
- [enums](/enums)
    - base logic for generating stuff
- coresort
    - sorting functionalities
- [coremath](/coremath)
    - deals with integer, ..., all type - min, max
- coretests
    - deals with basic functionality for test
- [corevalidator](/corevalidator)
    - [LineValidator](/corevalidator/LineValidator.go)
        - Deals with validation
        - Helps to do integration tests
- [coreversion](/coreversion/Version.go) - deals with all kinds of versioning stuff.
    - deals with version (major, minor, patch -- data type)
- [regexnew](/regexnew)
    - [LazyRegex](/regexnew/LazyRegex.go) - doesn't compile until needed but only once (lock / non lock)
    - [CreateLock](/regexnew/CreateLock.go) - Create Lazy lock in a loop or running process.
    - [Create](/regexnew/Create.go) - Create - use create to get lazyregex when created in global vars
- converters (move from type to type)
- [corecsv](/corecsv)
- [errorcore](/errcore)
- coreutils
    - [stringutil](/coreutils/stringutil)
        - [AnyToTypeString](/coreutils/stringutil/AnyToTypeString.go)
        - [AnyToString](/coreutils/stringutil/AnyToString.go)
        - [ClonePtr](/coreutils/stringutil/ClonePtr.go)
        - [FirstChar](/coreutils/stringutil/FirstChar.go)
        - [IsAnyEndsWith](/coreutils/stringutil/IsAnyEndsWith.go)
        - [IsAnyStartsWith](/coreutils/stringutil/IsAnyStartsWith.go)
        - [IsBlank](/coreutils/stringutil/IsBlank.go)
        - [IsBlankPtr](/coreutils/stringutil/IsBlankPtr.go)
        - [IsContains](/coreutils/stringutil/IsContains.go)
        - [IsContainsPtr](/coreutils/stringutil/IsContainsPtr.go)
        - [IsContainsPtrSimple](/coreutils/stringutil/IsContainsPtrSimple.go)
        - [IsDefined](/coreutils/stringutil/IsDefined.go)
        - [IsDefinedPtr](/coreutils/stringutil/IsDefinedPtr.go)
        - [IsEmpty](/coreutils/stringutil/IsEmpty.go)
        - [IsEmptyOrWhitespace](/coreutils/stringutil/IsEmptyOrWhitespace.go)
        - [IsEmptyOrWhitespacePtr](/coreutils/stringutil/IsEmptyOrWhitespacePtr.go)
        - [IsEmptyPtr](/coreutils/stringutil/IsEmptyPtr.go)
        - [IsEnds](/coreutils/stringutil/IsEnds.go)
        - [IsEndsChar](/coreutils/stringutil/IsEndsChar.go)
        - [IsEndsRune](/coreutils/stringutil/IsEndsRune.go)
        - [IsEndsWith](/coreutils/stringutil/IsEndsWith.go)
        - [IsNotEmpty](/coreutils/stringutil/IsNotEmpty.go)
        - [IsNullOrEmptyPtr](/coreutils/stringutil/IsNullOrEmptyPtr.go)
        - [IsStarts](/coreutils/stringutil/IsStarts.go)
        - [IsStartsAndEndsWithChar](/coreutils/stringutil/IsStartsAndEndsChar.go)
        - [IsStartsAndEndsWith](/coreutils/stringutil/IsStartsAndEndsWith.go)
        - [IsStartsChar](/coreutils/stringutil/IsStartsChar.go)
        - [IsEndsRune](/coreutils/stringutil/IsEndsRune.go)
        - [IsStartsWith](/coreutils/stringutil/IsStartsWith.go)
        - [MaskLine](/coreutils/stringutil/MaskLine.go)
        - [MaskLines](/coreutils/stringutil/MaskLines.go)
        - [MaskTrimLine](/coreutils/stringutil/MaskTrimLine.go)
        - [MaskTrimLines](/coreutils/stringutil/MaskTrimLines.go)
        - [RemoveMany](/coreutils/stringutil/RemoveMany.go)
        - [RemoveManuBySplitting](/coreutils/stringutil/RemoveManyBySplitting.go)
        - [SafeClonePtr](/coreutils/stringutil/SafeClonePtr.go)
        - [SafeSubstring](/coreutils/stringutil/SafeSubstring.go)
        - [SafeSubstringEnds](/coreutils/stringutil/SafeSubstringEnds.go)
        - [SafeSubstringStarts](/coreutils/stringutil/SafeSubstringStarts.go)
        - [SplitContentsByWhitespaceConditions](/coreutils/stringutil/SplitContentsByWhitespaceConditions.go)
        - [SplitFirstLast](/coreutils/stringutil/SplitFirstLast.go)
        - [SplitLeftRight](/coreutils/stringutil/SplitLeftRight.go)
        - [SplitLeftRightsTrims](/coreutils/stringutil/SplitLeftRightsTrims.go)
        - [SplitLeftRightTrimmed](/coreutils/stringutil/SplitLeftRightTrimmed.go)
        - [SplitLeftRightType](/coreutils/stringutil/SplitLeftRightType.go)
        - [SplitLeftRightTypeTrimmed](/coreutils/stringutil/SplitLeftRightTypeTrimmed.go)
        - [ToBool](/coreutils/stringutil/ToBool.go)
        - [ToByte](/coreutils/stringutil/ToByte.go)
        - [ToByteDefault](/coreutils/stringutil/ToByteDefault.go)
        - [ToInt](/coreutils/stringutil/ToInt.go)
        - [ToInt8](/coreutils/stringutil/ToInt8.go)
        - [ToInt8Def](/coreutils/stringutil/ToInt8Def.go)
        - [ToInt16](/coreutils/stringutil/ToInt16.go)
        - [ToInt16Default](/coreutils/stringutil/ToInt16Default.go)
        - [ToInt32](/coreutils/stringutil/ToInt32.go)
        - [ToInt32Def](/coreutils/stringutil/ToInt32Def.go)
        - [ToIntDef](/coreutils/stringutil/ToIntDef.go)
        - [ToIntDefault](/coreutils/stringutil/ToIntDefault.go)
        - [ToIntUsingRegexMatch](/coreutils/stringutil/ToIntUsingRegexMatch.go)
        - [ToUint16Default](/coreutils/stringutil/ToUint16Default.go)
        - [ToUint32Default](/coreutils/stringutil/ToUint32Default.go)
        - [ReplaceTemplate](/coreutils/stringutil/replaceTemplate.go)
            - CurlyOne
            - Curly
            - CurlyTwo
            - DirectOne
            - DirectTwoItem
            - CurlyTwoItem
            - DirectKeyUsingMap
            - CurlyKeyUsingMap
            - UsingMapOptions
            - UsingNamerMapOptions
            - UsingStringerMapOptions
            - UsingWrappedTemplate
            - UsingBracketsWrappedTemplate
            - UsingQuotesWrappedTemplate
            - UsingValueTemplate
            - UsingValueWithFieldsTemplate
- [issetter.value](/issetter/Value.go) - 4 Valued Booleans, instead using bool* use this.
    - 4 Valued Booleans
        - One Approach
            - Uninitialized
            - True
            - False
            - Wildcard / Any
        - Or,
            - Uninitialized
            - Set
            - Unset
            - Wildcard / Any
- [coredata](/coredata)
    - [coreapi](/coredata/coreapi)
    - [coredynamic](/coredata/coredynamic) - dynamic data type
        - [AnyCollection](/coredata/coredynamic/AnyCollection.go) - deals with collection of interface means any
        - AnySliceValToInterfacesAsync
        - AnyToReflectVal
        - AnyTypeMapToMapStringAny
        - BytesConverter
        - CastedResult
        - CastTo
        - Dynamic
        - DynamicCollection
        - DynamicStatus
        - [funcs](/coredata/coredynamic/funcs.go) - smart converter functions
        - IsAnyTypesOf
        - KeyVal
        - KeyValCollection
        - LeftRight
        - LengthOfReflect
        - MapAnyItemsDiff
        - MapAsKeyValSlice
        - MapKeysStringSlice
        - MapKeysStringSliceAnyMust
        - MapKeysStringSliceAnySorted
        - MapKeysStringSliceAnySortedMust
        - MustBeAcceptedTypes
        - NotAcceptedTypesErr
        - PonterOrNonPointer
        - PointerOrNonPointerUsingReflectValue
        - ReflectInterfaceVal
        - ReflectKindValidation
        - ReflectSetFromTo
        - ReflectTypeValidation
        - ReflectValToInterfaces
        - ReflectValToInterfacesAsync
        - ReflectValToInterfacesUsingProcessor
        - SafeTypeName
        - SafeZeroSet
        - SimpleRequest
        - SimpleResult
        - SliceItemsAsStrings
        - SliceItemsAsStringsAny
        - SliceItemsAsStringsAnyMust
        - SliceItemsProcessorAsStrings
        - SliceItemsSimpleProcessorAsStrings
        - Type
        - TypeMustBeSame
        - TypeName
        - TypeNames
        - TypeNamesReferenceString
        - TypeNamesString
        - TypeNamesStringUsingReflectType
        - TypeNamesUsingReflectType
        - TypeNotEqualErr
        - TypeSameStatus
        - TypeIndexOf
        - TypeStatus
        - ValueStatus
        - ZeroSet
        - ZeroSetAny
    - [corejson](/coredata/corejson)
        - Result
        - ResultCollection
        - [SerializerLogic (corejson.Serialize.Method)](/coredata/corejson/serializerLogic.go)
            - corejson.Serialize.StringsApply
            - corejson.Serialize.Apply
            - corejson.Serialize.FromBytes
            - corejson.Serialize.FromStrings
            - corejson.Serialize.FromStringsSpread
            - corejson.Serialize.FromString
            - corejson.Serialize.FromInteger
            - corejson.Serialize.FromInteger64
            - corejson.Serialize.FromBool
            - corejson.Serialize.FromIntegers
            - corejson.Serialize.FromStringer
            - corejson.Serialize.UsingAnyPtr
            - corejson.Serialize.UsingAny
            - corejson.Serialize.Raw
            - corejson.Serialize.Marshal
            - corejson.Serialize.ApplyMust
            - corejson.Serialize.ToBytesMust
            - corejson.Serialize.ToSafeBytesMust
            - corejson.Serialize.ToSafeBytesSwallowErr
            - corejson.Serialize.ToBytesSwallowErr
            - corejson.Serialize.ToBytesErr
            - corejson.Serialize.ToString
            - corejson.Serialize.ToStringMust
            - corejson.Serialize.ToStringErr
            - corejson.Serialize.ToPrettyStringErr
            - corejson.Serialize.ToPrettyStringIncludingErr
        - [DeserializerLogic (corejson.Deserialize.Method)](/coredata/corejson/deserializerLogic.go)
            - corejson.Deserialize.StringsApply
            - corejson.Deserialize.Apply
            - corejson.Deserialize.UsingStringPtr
            - corejson.Deserialize.UsingError
            - corejson.Deserialize.UsingErrorWhichJsonResult
            - corejson.Deserialize.UsingResult
            - corejson.Deserialize.ApplyMust
            - corejson.Deserialize.UsingString
            - corejson.Deserialize.FromString
            - corejson.Deserialize.FromTo
            - corejson.Deserialize.MapAnyToPointer
            - corejson.Deserialize.UsingStringOption
            - corejson.Deserialize.UsingStringIgnoreEmpty
            - corejson.Deserialize.UsingBytes
            - corejson.Deserialize.UsingBytesPointerMust
            - corejson.Deserialize.UsingBytesIf
            - corejson.Deserialize.UsingBytesPointerIf
            - corejson.Deserialize.UsingBytesPointer
            - corejson.Deserialize.UsingBytesMust
            - corejson.Deserialize.UsingSafeBytesMust
            - corejson.Deserialize.AnyToFieldsMap
            - corejson.Deserialize.UsingSerializerTo
            - corejson.Deserialize.UsingSerializerFuncTo
            - corejson.Deserialize.UsingDeserializerToOption
            - corejson.Deserialize.UsingDeserializerDefined
            - corejson.Deserialize.UsingDeserializerFuncDefined
            - corejson.Deserialize.UsingJsonerToAny
    - [coreonce - data which generate once](/coredata/coreonce)
    - [corepayload - deals with enhance payloads](/coredata/corepayload)
    - [corerange - works with ranges](/coredata/corerange)
    - [ostype](/ostype)
    - corestr - string related core functionalities and data-types
        - hashmap
        - hashset - lock features for `map[string]bool`
        - Collection - List like C# / ArrayList like Java - enhance APIs
        - CollectionsOfCollection
        - LinkedList
        - HashDiff
        - CloneSlice
        - CharCollectionMap - `map[byte]*Collection`
        - HashsetsCollection
        - KeyValueCollection
        - LeftRight - Left, Right string
        - LeftMiddleRight
        - SimpleStringOnce
        - ValidValue
        - ValidValues
        - ValueStatus
        - AnyToString

## Examples Json - Deserializer

```go=
    type Example struct {
		A       string
		B       int
		SomeMap map[string]string
	}

	exampleFrom := &Example{
		A:       "Something",
		B:       1,
		SomeMap: map[string]string{},
	}

	exampleTo := &Example{}

	err := corejson.Deserialize.FromTo(
		exampleFrom,
		exampleTo)
		
	// exampleTo will be a copy of public fields From exampleFrom
	// exampleFrom := &Example{
	// 	A:       "Something",
	// 	B:       1,
	// 	SomeMap: map[string]string{},
	// }
```

## Examples

```go=
// substituting functions as ternary operator
fmt.Println(conditional.Int(true, 2, 7)) // 2
fmt.Println(conditional.Int(false, 2, 7)) // 7

// making collection from array of strings
stringValues := []string{"hello", "world", "something"}
collectionPtr1 := corestr.NewCollectionPtrUsingStrings(&stringValues, constants.Zero)
fmt.Println(collectionPtr1)
/* outputs:
   - hello
   - world
   - something
*/

// different methods of collection
fmt.Println(collectionPtr1.Length()) // 3
fmt.Println(collectionPtr1.IsEmpty()) // false

// adding more element including empty string
collectionPtr2 := collectionPtr1.AddsLock("else")
fmt.Println(collectionPtr2.Length()) // 4

// checking equality
fmt.Println(collectionPtr1.IsEqualsPtr(collectionPtr2)) // true

// creating CharCollectionMap using collection
sampleMap := collectionPtr1.CharCollectionPtrMap()
fmt.Println(sampleMap)

// methods on CharCollectionMap
fmt.Println(sampleMap.Length()) // 4
fmt.Println(sampleMap.AllLengthsSum()) // 4
fmt.Println(sampleMap.Clear()) // prints: # Summary of `*corestr.CharCollectionMap`, Length ("0") - Sequence `1`
otherMap := sampleMap.Add("another")
fmt.Println(otherMap)
/* prints:
   # Summary of `*corestr.CharCollectionMap`, Length ("1") - Sequence `1`
         1 . `a` has `1` items.
   ## Items of `a`
         - another
*/

// declaring an empty hashset of length 2 and calling methods on it
newHashSet := corestr.NewHashset(2)
fmt.Println(newHashSet.Length()) // 2
fmt.Println(newHashSet.IsEmpty()) // true
fmt.Println(newHashSet.Items()) // &map[]

// adding items to hashset
strPtr := "new"
newHashSet.AddPtr(&strPtr)
fmt.Println(newHashSet.Items()) // &map[new:true]

// adding map to hashset
newHashSet.AddItemsMap(&map[string]bool{"hi": true, "no": false})
fmt.Println(newHashSet.Items()) // &map[hi:true new:true]

// math operations: getting the larger/smaller value from two given values
fmt.Println(coremath.MaxByte('e', 'i')) // 105 which represents 'i' in ASCII
fmt.Println(coremath.MinByte(23, 5))    // 5

// initializing issetter value
isSetterValue := issetter.False // initializing as false
fmt.Println(isSetterValue.HasInitialized()) // true
fmt.Println(isSetterValue.Value()) // 2
fmt.Println(isSetterValue.IsPositive()) // false

// sorting strings
fruits := []string{"banana", "mango", "apple"}
fmt.Println(strsort.Quick(&fruits)) // &[apple banana mango]
fmt.Println(strsort.QuickDsc(&fruits)) // &[mango banana apple]

// converting pointer strings to strings
mile := "mile"
km := "km"
measures := []*string{&mile, &km}
fmt.Println(converters.PointerStringsToStrings(&measures)) // &[mile km]
fmt.Printf("Type %T", converters.PointerStringsToStrings(&measures)) // Type *[]string

// comparing two int arays
Values := []int{1, 2, 3, 4}
OtherValues := []int{5, 6, 7, 8}
fmt.Println(corecompare.IntArray(Values, OtherValues)) // false
```

## Acknowledgement

Any other packages used

## Links

- [go - Calling a method on a nil struct pointer doesn't panic.](https://t.ly/aTp0)
- [Array of pointers to JSON - Stack Overflow](https://stackoverflow.com/questions/28101966/array-of-pointers-to-json)
- [Json Parsing of Array Pointers](https://play.golang.org/p/zTuMLBgGWk)
- [Go Slice Tricks Cheat Sheet
  ](https://ueokande.github.io/go-slice-tricks/)
- [SliceTricks · golang/go Wiki
  ](https://github.com/golang/go/wiki/SliceTricks)
- [ueokande/go-slice-tricks: Cheat Sheet for Go Slice Tricks](https://github.com/ueokande/go-slice-tricks)
- [Quick Sort in Go (Golang) - golangprograms.com](https://t.ly/pDyj)
    - [Sorting using golang lib](https://play.golang.org/p/sJ8a464USeV)
    - [Pointer Strings Sort](https://play.golang.org/p/8V8YYdQrO6q)
- [Golang Array process issue without copying (!Important)](https://play.golang.org/p/GvdJMPmCStz)
- [Linked List | Set 2 (Inserting a node) - GeeksforGeeks](https://t.ly/MMaY)
- [Go Data Structures: Linked List](https://t.ly/QLLy)
- [System info](https://github.com/zcalusic/sysinfo)
  -[Stackoverflow Centos detect](https://stackoverflow.com/a/65207574)

### Regex Patterns

#### Path RegEx Patterns

* [java - Regex pattern to validate Linux folder path - Stack Overflow](https://stackoverflow.com/questions/55069650/regex-pattern-to-validate-linux-folder-path/55070259)
* [regex - What is the most correct regular expression for a UNIX file path? - Stack Overflow](https://stackoverflow.com/questions/537772/what-is-the-most-correct-regular-expression-for-a-unix-file-path)
* [java - Regular expression to validate windows and linux path with extension - Stack Overflow](https://stackoverflow.com/questions/44289075/regular-expression-to-validate-windows-and-linux-path-with-extension)
* [javascript - Regex windows path validator - Stack Overflow](https://stackoverflow.com/questions/51494579/regex-windows-path-validator/51504254)
* [Path Regex fix for all OS](https://t.ly/1JuS)

## Issues

- [Create your issues](https://gitlab.com/auk-go/core/-/issues)

## Notes

## Contributors

- [Md. Alim Ul Karim](https://www.google.com/search?q=Alim+Ul+Karim)

## License

[MIT License](/LICENSE)
