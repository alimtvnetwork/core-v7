package coredata

type PointerStrings []*string

func (p PointerStrings) Len() int {
	if p == nil {
		return 0
	}

	return len(p)
}

func (p PointerStrings) Less(i, j int) bool { return *p[i] < *p[j] }
func (p PointerStrings) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
