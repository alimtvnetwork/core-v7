package osconsts

import (
	"runtime"

	"gitlab.com/evatix-go/core/constants"
)

// GOOS values https://stackoverflow.com/a/20728862
//goland:noinspection ALL
const (
	LsbCommand                = "lsb_release"
	Windows                   = "windows"
	Android                   = "android"
	DarwinOrMacOs             = "darwin"
	Linux                     = "linux"
	DragonFly                 = "dragonfly"
	FreeBsd                   = "freebsd"
	JavaScript                = "js"
	NetBsd                    = "netbsd"
	OpenBsd                   = "openbsd"
	Debian                    = "debian"
	Plan9                     = "plan9"
	MacOs                     = "mac"
	IOs                       = "ios"
	Ubuntu                    = "ubuntu"
	Solaris                   = "solaris"
	Freebsd                   = "freebsd"
	Nacl                      = "nacl"
	Unknown                   = "Unknown"
	NewLine                   = constants.NewLine
	PathSeparator             = constants.PathSeparator
	CurrentOperatingSystem    = runtime.GOOS
	CurrentSystemArchitecture = runtime.GOARCH
	IsWindows                 = CurrentOperatingSystem == Windows
	IsLinux                   = CurrentOperatingSystem == Linux
	IsDarwinOrMacOs           = CurrentOperatingSystem == DarwinOrMacOs
	IsUbuntu                  = CurrentOperatingSystem == Ubuntu
	IsPlan9                   = CurrentOperatingSystem == Plan9
	IsSolaris                 = CurrentOperatingSystem == Solaris
	IsFreebsd                 = CurrentOperatingSystem == Freebsd
	IsDebian                  = CurrentOperatingSystem == Debian
	IsNetBsd                  = CurrentOperatingSystem == NetBsd
	IsOpenBsd                 = CurrentOperatingSystem == OpenBsd
	IsDragonFly               = CurrentOperatingSystem == DragonFly
	IsNacl                    = CurrentOperatingSystem == Nacl
	IsUnixGroup               = !IsWindows
	WindowsCDrive             = "C:\\"
)
