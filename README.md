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

## Why `constants?`

It makes our other go-packages DRY and concise.

## Examples

`Code Smaples`

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

## Issues

- [Create your issues](https://gitlab.com/evatix-go/core/-/issues)

## Notes

## Contributors

## License

[Evatix MIT License](/LICENSE)
