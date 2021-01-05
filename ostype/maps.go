package ostype

import "gitlab.com/evatix-go/core/osconsts"

var (
	OsVariantToStringMap = map[Variation]string{
		Windows:       osconsts.Windows,
		Linux:         osconsts.Linux,
		DarwinOrMacOs: osconsts.DarwinOrMacOs,
		JavaScript:    osconsts.JavaScript,
		FreeBsd:       osconsts.FreeBsd,
		NetBsd:        osconsts.NetBsd,
		OpenBsd:       osconsts.OpenBsd,
		DragonFly:     osconsts.DragonFly,
		Android:       osconsts.Android,
		Unknown:       osconsts.Unknown,
	}

	OsStringToVariantMap = map[string]Variation{
		osconsts.Windows:       Windows,
		osconsts.Linux:         Linux,
		osconsts.DarwinOrMacOs: DarwinOrMacOs,
		osconsts.JavaScript:    JavaScript,
		osconsts.FreeBsd:       FreeBsd,
		osconsts.NetBsd:        NetBsd,
		osconsts.OpenBsd:       OpenBsd,
		osconsts.DragonFly:     DragonFly,
		osconsts.Android:       Android,
		osconsts.Unknown:       Unknown,
	}
)
