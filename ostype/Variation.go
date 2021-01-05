package ostype

import "gitlab.com/evatix-go/core/msgtype"

type Variation byte

const (
	Windows Variation = iota
	Linux
	DarwinOrMacOs
	JavaScript
	FreeBsd
	NetBsd
	OpenBsd
	DragonFly
	Android
	Unknown
)

func (variation Variation) IsByte(another byte) bool {
	return variation == Variation(another)
}

func (variation Variation) Is(another Variation) bool {
	return variation == another
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
