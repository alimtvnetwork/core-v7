# `Core` Intro

![Use Package logo](UseLogo)

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

## Issues

- [Create your issues](https://gitlab.com/evatix-go/core/-/issues)

## Notes

## Contributors

## License

[Evatix MIT License](/LICENSE)
