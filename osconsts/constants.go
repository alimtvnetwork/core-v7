package osconsts

import (
	"gitlab.com/auk-go/core/internal/osconstsinternal"
)

// GOOS values https://stackoverflow.com/a/20728862
//
//goland:noinspection ALL
const (
	LsbCommand                = osconstsinternal.LsbCommand
	Android                   = osconstsinternal.Android
	DarwinOrMacOs             = osconstsinternal.DarwinOrMacOs
	DragonFly                 = osconstsinternal.DragonFly
	FreeBsd                   = osconstsinternal.FreeBsd
	Linux                     = osconstsinternal.Linux
	Nacl                      = osconstsinternal.Nacl
	NetBsd                    = osconstsinternal.NetBsd
	OpenBsd                   = osconstsinternal.OpenBsd
	Plan9                     = osconstsinternal.Plan9
	Solaris                   = osconstsinternal.Solaris
	Windows                   = osconstsinternal.Windows
	Unknown                   = osconstsinternal.Unknown
	Any                       = osconstsinternal.Any
	Illumos                   = osconstsinternal.Illumos
	IOs                       = osconstsinternal.IOs
	Aix                       = osconstsinternal.Aix
	NewLine                   = osconstsinternal.NewLine
	PathSeparator             = osconstsinternal.PathSeparator
	CurrentOperatingSystem    = osconstsinternal.CurrentOperatingSystem
	CurrentSystemArchitecture = osconstsinternal.CurrentSystemArchitecture
	IsWindows                 = osconstsinternal.IsWindows
	IsLinux                   = osconstsinternal.IsLinux
	IsDarwinOrMacOs           = osconstsinternal.IsDarwinOrMacOs
	IsPlan9                   = osconstsinternal.IsPlan9
	IsSolaris                 = osconstsinternal.IsSolaris
	IsFreebsd                 = osconstsinternal.IsFreebsd
	IsNetBsd                  = osconstsinternal.IsNetBsd
	IsOpenBsd                 = osconstsinternal.IsOpenBsd
	IsDragonFly               = osconstsinternal.IsDragonFly
	IsNacl                    = osconstsinternal.IsNacl
	IsUnixGroup               = osconstsinternal.IsUnixGroup
	WindowsCDrive             = osconstsinternal.WindowsCDrive
	LinuxHome                 = osconstsinternal.LinuxHome
	LinuxHomeSlash            = osconstsinternal.LinuxHomeSlash
	LinuxBin                  = osconstsinternal.LinuxBin
	LinuxPermanentTemp        = osconstsinternal.LinuxPermanentTemp // https://prnt.sc/gW0DA5d4jt6R, unix : /var/tmp/
)
