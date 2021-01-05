package ostype

import "gitlab.com/evatix-go/core/osconsts"

func IsWindows(rawRuntimeGoos string) bool {
	return rawRuntimeGoos == osconsts.Windows
}
