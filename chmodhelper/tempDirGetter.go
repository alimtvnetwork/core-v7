package chmodhelper

import "gitlab.com/auk-go/core/internal/osconstsinternal"

type tempDirGetter struct{}

// TempDefault
//
// Example:
//  - unix    : /tmp
//  - windows : %temp%
func (it tempDirGetter) TempDefault() string {
	return TempDirDefault
}

// TempPermanent
//
// Windows:
//  - c:\\windows\\temp
//
// unix:
//  - /var/tmp/
//
// Reference:
//  - Why "/var/tmp/" : https://prnt.sc/gW0DA5d4jt6R
func (it tempDirGetter) TempPermanent() string {
	if osconstsinternal.IsWindows {
		return osconstsinternal.WindowsPermanentTemp
	}

	// unix
	return osconstsinternal.LinuxPermanentTemp
}

func (it tempDirGetter) TempOption(isPermanent bool) string {
	if isPermanent {
		return it.TempPermanent()
	}

	return it.TempDefault()
}
