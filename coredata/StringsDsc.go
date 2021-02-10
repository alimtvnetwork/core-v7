package coredata

type StringsDsc []string

func (p StringsDsc) Len() int {
	if p == nil {
		return 0
	}

	return len(p)
}

func (p StringsDsc) Less(i, j int) bool { return p[i] > p[j] }
func (p StringsDsc) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
