package ostype

import "gitlab.com/auk-go/core/osconsts"

func GetCurrentGroup() Group {
	return GetGroup(osconsts.CurrentOperatingSystem)
}
