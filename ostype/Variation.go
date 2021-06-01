package ostype

import (
	"gitlab.com/evatix-go/core/msgtype"
)

type Variation byte

// https://stackoverflow.com/a/50117892 | https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63
// go tool dist list
const (
	Any Variation = iota
	Windows
	Linux
	DarwinOrMacOs
	JavaScript
	FreeBsd
	NetBsd
	OpenBsd
	DragonFly
	Android
	Plan9
	Solaris
	Nacl
	Unknown
)

func (variation Variation) IsByte(another byte) bool {
	return variation == Variation(another)
}

func (variation Variation) IsAnyOperatingSystem() bool {
	return Any == variation
}

func (variation Variation) Is(other Variation) bool {
	return other == variation
}

func (variation Variation) IsAnyMatch(others ...Variation) bool {
	for _, other := range others {
		if other == variation {
			return true
		}
	}

	return false
}

func (variation Variation) IsStringsMatchAny(others ...string) bool {
	for _, other := range others {
		otherVariant := GetVariant(other)

		if otherVariant == variation {
			return true
		}
	}

	return false
}

// IsPossibleUnixGroup variation != Windows
func (variation Variation) IsPossibleUnixGroup() bool {
	return variation != Windows
}

func (variation Variation) IsLinuxOrMac() bool {
	return variation == Linux || variation == DarwinOrMacOs
}

func (variation Variation) Group() Group {
	if variation == Windows {
		return WindowsGroup
	}

	if variation == Android {
		return AndroidGroup
	}

	return UnixGroup
}

func (variation Variation) IsActualGroupUnix() bool {
	return variation.Group().IsUnix()
}

func (variation Variation) IsWindows() bool {
	return variation == Windows
}

func (variation Variation) IsLinux() bool {
	return variation == Linux
}

func (variation Variation) IsDarwinOrMacOs() bool {
	return variation == DarwinOrMacOs
}

func (variation Variation) IsJavaScript() bool {
	return variation == JavaScript
}

func (variation Variation) IsFreeBsd() bool {
	return variation == FreeBsd
}

func (variation Variation) IsNetBsd() bool {
	return variation == NetBsd
}

func (variation Variation) IsOpenBsd() bool {
	return variation == NetBsd
}

func (variation Variation) IsDragonFly() bool {
	return variation == DragonFly
}

func (variation Variation) Value() byte {
	return byte(variation)
}

func (variation Variation) String() string {
	osName, has := OsVariantToStringMap[variation]

	if has {
		return osName
	}

	msg := msgtype.UnsupportedCategory.Combine(
		"os type pkg: variant not supported.",
		string(variation))

	panic(msg)
}
