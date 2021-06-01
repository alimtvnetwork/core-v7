# `Core` Intro

![Use Package logo](https://gitlab.com/evatix-go/core/uploads/486811aa7446cd43b17ff167ceaf90d1/core-250.png)

All common core infrastructure and constants combined package.

## Git Clone

`git clone https://gitlab.com/evatix-go/core.git`

### 2FA enabled, for linux

`git clone https://[YourGitLabUserName]:[YourGitlabAcessTokenGenerateFromGitlabsTokens]@gitlab.com/evatix-go/core.git`

### Prerequisites

- Update git to latest 2.29
- Update or install the latest of Go 1.15.2
- Either add your ssh key to your gitlab account
- Or, use your access token to clone it.

## Installation

`go get gitlab.com/evatix-go/core`

### Go get issue for private package

- Update git to 2.29
- Enable go modules. (Windows : `go env -w GO111MODULE=on`, Unix : `export GO111MODULE=on`)
- Add `gitlab.com/evatix-go` to go env private

To set for Windows:

`go env -w GOPRIVATE=[AddExistingOnes;]gitlab.com/evatix-go`

To set for Unix:

`expoort GOPRIVATE=[AddExistingOnes;]gitlab.com/evatix-go`

## Why `core?`

It makes our other go-packages DRY and concise.

## Examples Videos

- [Core Basics Intro](https://drive.google.com/file/d/1CA4817zaehhWqgtAGI2UH7Tojtngcyjw/view)
- [Core Usage Video](https://drive.google.com/file/d/1kwC_3R-QIZE1pNK_9F7hFdYuGB0CSGYh/view?usp=sharing)

## Examples

```go
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

- [go - Calling a method on a nil struct pointer doesn't panic. Why not? - Stack Overflow](https://stackoverflow.com/questions/42238624/calling-a-method-on-a-nil-struct-pointer-doesnt-panic-why-not)
- [Array of pointers to JSON - Stack Overflow](https://stackoverflow.com/questions/28101966/array-of-pointers-to-json)
- [Json Parsing of Array Pointers](https://play.golang.org/p/zTuMLBgGWk)
- [Go Slice Tricks Cheat Sheet](https://ueokande.github.io/go-slice-tricks/)
- [SliceTricks · golang/go Wiki](https://github.com/golang/go/wiki/SliceTricks)
- [ueokande/go-slice-tricks: Cheat Sheet for Go Slice Tricks](https://github.com/ueokande/go-slice-tricks)
- [Quick Sort in Go (Golang) - golangprograms.com](https://www.golangprograms.com/golang-program-for-implementation-of-quick-sort.html)
    - [Sorting using golang lib](https://play.golang.org/p/sJ8a464USeV)
    - [Pointer Strings Sort](https://play.golang.org/p/8V8YYdQrO6q)
- [Golang Array process issue without copying (!Important)](https://play.golang.org/p/GvdJMPmCStz)
- [Linked List | Set 2 (Inserting a node) - GeeksforGeeks](https://www.geeksforgeeks.org/linked-list-set-2-inserting-a-node/)
- [Go Data Structures: Linked List](https://flaviocopes.com/golang-data-structure-linked-list/)
- [System info](https://github.com/zcalusic/sysinfo)
  -[Stackoverflow Centos detect](https://stackoverflow.com/a/65207574)

## Issues

- [Create your issues](https://gitlab.com/evatix-go/core/-/issues)

## Notes

## Contributors

## License

[Evatix MIT License](/LICENSE)
