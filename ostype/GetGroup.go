package ostype

import (
	"gitlab.com/evatix-go/core/osconsts"
)

// GetGroup rawRuntimeGoos = runtime.GOOS
func GetGroup(rawRuntimeGoos string) Group {
	if rawRuntimeGoos == osconsts.Windows {
		return WindowsGroup
	}

	if rawRuntimeGoos == osconsts.Android {
		return AndroidGroup
	}

	isUnixGroup, has := osconsts.UnixGroupsMap[rawRuntimeGoos]

	if has && isUnixGroup {
		return UnixGroup
	}

	return UnknownGroup
}
