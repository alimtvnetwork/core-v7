package osconsts

import (
	"runtime"

	"gitlab.com/evatix-go/core/constants"
)

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
	MacOs                     = "mac"
	IOs                       = "ios"
	Ubuntu                    = "ubuntu"
	Unknown                   = "Unknown"
	NewLine                   = constants.NewLine
	PathSeparator             = constants.PathSeparator
	CurrentOperatingSystem    = runtime.GOOS
	CurrentSystemArchitecture = runtime.GOARCH
	IsWindows                 = CurrentOperatingSystem == Windows
	IsLinux                   = CurrentOperatingSystem == Linux
	IsDarwinOrMacOs           = CurrentOperatingSystem == DarwinOrMacOs
	IsUbuntu                  = CurrentOperatingSystem == Ubuntu
	IsUnixGroup               = !IsWindows
)
