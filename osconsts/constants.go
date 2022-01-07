package osconsts

import (
	"runtime"

	"gitlab.com/evatix-go/core/constants"
)

// GOOS values https://stackoverflow.com/a/20728862
//goland:noinspection ALL
const (
	LsbCommand                = "lsb_release"
	Android                   = "android"
	DarwinOrMacOs             = "darwin"
	DragonFly                 = "dragonfly"
	FreeBsd                   = "freebsd"
	Linux                     = "linux"
	Nacl                      = "nacl"
	NetBsd                    = "netbsd"
	OpenBsd                   = "openbsd"
	Plan9                     = "plan9"
	Solaris                   = "solaris"
	Windows                   = "windows"
	Unknown                   = "Unknown"
	Any                       = "Any"
	Illumos                   = "illumos"
	IOs                       = "ios"
	Aix                       = "aix"
	NewLine                   = constants.NewLine
	PathSeparator             = constants.PathSeparator
	CurrentOperatingSystem    = runtime.GOOS
	CurrentSystemArchitecture = runtime.GOARCH
	IsWindows                 = CurrentOperatingSystem == Windows
	IsLinux                   = CurrentOperatingSystem == Linux
	IsDarwinOrMacOs           = CurrentOperatingSystem == DarwinOrMacOs
	IsPlan9                   = CurrentOperatingSystem == Plan9
	IsSolaris                 = CurrentOperatingSystem == Solaris
	IsFreebsd                 = CurrentOperatingSystem == FreeBsd
	IsNetBsd                  = CurrentOperatingSystem == NetBsd
	IsOpenBsd                 = CurrentOperatingSystem == OpenBsd
	IsDragonFly               = CurrentOperatingSystem == DragonFly
	IsNacl                    = CurrentOperatingSystem == Nacl
	IsUnixGroup               = !IsWindows
	WindowsCDrive             = "C:\\"
	LinuxHome                 = "/home"
	LinuxHomeSlash            = "/home/"
	LinuxBin                  = "/bin"
)
