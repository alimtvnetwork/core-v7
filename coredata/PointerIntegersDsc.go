package coredata

type PointerIntegersDsc []*int

func (p PointerIntegersDsc) Len() int {
	if p == nil {
		return 0
	}

	return len(p)
}

func (p PointerIntegersDsc) Less(i, j int) bool { return *p[i] > *p[j] }
func (p PointerIntegersDsc) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
