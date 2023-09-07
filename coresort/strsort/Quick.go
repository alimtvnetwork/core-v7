package strsort

import (
	"sort"

	"gitlab.com/auk-go/core/coredata"
)

// QuickPtr Warning: Data gets mutated inside.
//
// Reference : https://play.golang.org/p/_OkY82E2kO9
func QuickPtr(pointerStringsIn *[]*string) *[]*string {
	pointerStrings := coredata.PointerStrings(*pointerStringsIn)
	sort.Sort(pointerStrings)

	return pointerStringsIn
}

// Quick Warning: Data gets mutated inside.
//
// Reference : https://play.golang.org/p/sJ8a464USeV
func Quick(stringsPointerIn *[]string) *[]string {
	sort.Strings(*stringsPointerIn)

	return stringsPointerIn
}

// QuickDscPtr Warning: Data gets mutated inside.
//
// Reference : https://play.golang.org/p/_OkY82E2kO9
func QuickDscPtr(pointerStringsIn *[]*string) *[]*string {
	pointerStringsDsc := coredata.PointerStringsDsc(*pointerStringsIn)
	sort.Sort(pointerStringsDsc)

	return pointerStringsIn
}

// QuickDsc Warning: Data gets mutated inside.
//
// Reference : https://play.golang.org/p/sJ8a464USeV
func QuickDsc(stringsPointerIn *[]string) *[]string {
	pointerStringsDsc := coredata.StringsDsc(*stringsPointerIn)
	sort.Sort(pointerStringsDsc)

	return stringsPointerIn
}
