package ostype

import "gitlab.com/evatix-go/core/osconsts"

func GetCurrentGroup() Group {
	return GetGroup(osconsts.CurrentOperatingSystem)
}
