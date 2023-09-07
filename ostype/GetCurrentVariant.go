package ostype

import "gitlab.com/auk-go/core/osconsts"

func GetCurrentVariant() Variation {
	return GetVariant(osconsts.CurrentOperatingSystem)
}
