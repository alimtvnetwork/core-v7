package ostype

import "gitlab.com/evatix-go/core/osconsts"

func GetCurrentVariant() Variation {
	return GetVariant(osconsts.CurrentOperatingSystem)
}
