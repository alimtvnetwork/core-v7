package ostype

import "gitlab.com/auk-go/core/osconsts"

var (
	osTypesStrings = [...]string{
		Any:           osconsts.Any,
		Unknown:       osconsts.Unknown,
		Windows:       osconsts.Windows,
		Linux:         osconsts.Linux,
		DarwinOrMacOs: osconsts.DarwinOrMacOs,
		FreeBsd:       osconsts.FreeBsd,
		NetBsd:        osconsts.NetBsd,
		OpenBsd:       osconsts.OpenBsd,
		DragonFly:     osconsts.DragonFly,
		Android:       osconsts.Android,
		Plan9:         osconsts.Plan9,
		Solaris:       osconsts.Solaris,
		Nacl:          osconsts.Nacl,
		Illumos:       osconsts.Illumos,
		IOs:           osconsts.IOs,
		Aix:           osconsts.Aix,
	}

	OsVariantToStringMap = map[Variation]string{
		Any:           osconsts.Any,
		Unknown:       osconsts.Unknown,
		Windows:       osconsts.Windows,
		Linux:         osconsts.Linux,
		DarwinOrMacOs: osconsts.DarwinOrMacOs,
		FreeBsd:       osconsts.FreeBsd,
		NetBsd:        osconsts.NetBsd,
		OpenBsd:       osconsts.OpenBsd,
		DragonFly:     osconsts.DragonFly,
		Android:       osconsts.Android,
		Plan9:         osconsts.Plan9,
		Solaris:       osconsts.Solaris,
		Nacl:          osconsts.Nacl,
		Illumos:       osconsts.Illumos,
		IOs:           osconsts.IOs,
		Aix:           osconsts.Aix,
	}

	OsStringToVariantMap = map[string]Variation{
		osconsts.Any:           Any,
		osconsts.Unknown:       Unknown,
		osconsts.Windows:       Windows,
		osconsts.Linux:         Linux,
		osconsts.DarwinOrMacOs: DarwinOrMacOs,
		osconsts.FreeBsd:       FreeBsd,
		osconsts.NetBsd:        NetBsd,
		osconsts.OpenBsd:       OpenBsd,
		osconsts.DragonFly:     DragonFly,
		osconsts.Android:       Android,
		osconsts.Plan9:         Plan9,
		osconsts.Solaris:       Solaris,
		osconsts.Nacl:          Nacl,
		osconsts.Illumos:       Illumos,
		osconsts.IOs:           IOs,
		osconsts.Aix:           Aix,
	}
)
