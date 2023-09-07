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

It makes our other go-packages DRY and concise. It was the first package in the auk-go ecosystem that is core of everything.

It was first designed for constants, later it got enhanced with:
- codestack
  - very powerful to deal with codestack
- chmodhelper
- enums
  - base logic for generating stuff
- coresort
  - sorting functionalities
- coremath
  - deals with integer, ..., all type - min, max
- coretests
  - deals with basic functionality for test
- corevalidator
  - deals with validation
- coreversion
  - deals with version (major, minor, patch -- data type)
- regexnew
  - lazyregex - doesn't compile until needed but only once (lock / non lock)
- converters (move from type to type)
- corecsv
- issetter
  - 3 / 4 phase enum treat as 3 / 4 phase selector (better than bool pointer)
- coredata
  - coreapi
  - coredynamic - dynamic data type
  - corejson
  - coreonce - data which generate once
  - corepayload - deals with enhance payloads
  - corerange - works with ranges
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

## License

[Evatix MIT License](/LICENSE)
