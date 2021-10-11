package devenv

import "gitlab.com/evatix-go/core/enums/envtype"

type Info struct {
	IsVerbose, IsLog, IsWarning, IsDebug bool
	EnvName                              envtype.Variant
}
