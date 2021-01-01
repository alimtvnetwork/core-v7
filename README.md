# `Constants` Intro

![Use Package logo](UseLogo)

All common constants package.

## Git Clone

`git clone https://gitlab.com/evatix-go/constants.git`

### 2FA enabled, for linux

`git clone https://[YourGitLabUserName]:[YourGitlabAcessTokenGenerateFromGitlabsTokens]@gitlab.com/evatix-go/constants.git`

### Prerequisites

- Update git to latest 2.29
- Update or install the latest of Go 1.15.2
- Either add your ssh key to your gitlab account
- Or, use your access token to clone it.

## Installation

`go get YourModuleName`

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

## Issues

- [Create your issues](https://gitlab.com/evatix-go/constants/-/issues)

## Notes

## Contributors

## License

[Evatix MIT License](/LICENSE)
