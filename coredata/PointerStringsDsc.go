package coredata

type PointerStringsDsc []*string

func (p PointerStringsDsc) Len() int {
	if p == nil {
		return 0
	}

	return len(p)
}

func (p PointerStringsDsc) Less(i, j int) bool { return *p[i] > *p[j] }
func (p PointerStringsDsc) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
