package intsort

import (
	"sort"

	"gitlab.com/evatix-go/core/coredata"
)

// Warning: Data gets mutated inside.
//
// Reference : https://play.golang.org/p/_OkY82E2kO9
func QuickPtr(pointerStringsIn *[]*int) *[]*int {
	pointerStrings := coredata.PointerIntegers(*pointerStringsIn)
	sort.Sort(pointerStrings)

	return pointerStringsIn
}

// Warning: Data gets mutated inside.
//
// Reference : https://play.golang.org/p/sJ8a464USeV
func Quick(intsPointerIn *[]int) *[]int {
	sort.Ints(*intsPointerIn)

	return intsPointerIn
}

// Warning: Data gets mutated inside.
//
// Reference : https://play.golang.org/p/_OkY82E2kO9
func QuickDscPtr(pointerStringsIn *[]*int) *[]*int {
	pointerStringsDsc := coredata.PointerIntegersDsc(*pointerStringsIn)
	sort.Sort(pointerStringsDsc)

	return pointerStringsIn
}

// Warning: Data gets mutated inside.
//
// Reference : https://play.golang.org/p/sJ8a464USeV
func QuickDsc(intsPointerIn *[]int) *[]int {
	pointerStringsDsc := coredata.IntegersDsc(*intsPointerIn)
	sort.Sort(pointerStringsDsc)

	return intsPointerIn
}
