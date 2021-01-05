package ostype

import "gitlab.com/evatix-go/core/osconsts"

func IsUnixGroup(rawRuntimeGoos string) bool {
	isUnixGroup, has := osconsts.UnixGroupsMap[rawRuntimeGoos]

	if has && isUnixGroup {
		return UnixGroup.IsUnix()
	}

	return false
}
