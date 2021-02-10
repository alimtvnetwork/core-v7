package coredata

type IntegersDsc []int

func (p IntegersDsc) Len() int {
	if p == nil {
		return 0
	}

	return len(p)
}

func (p IntegersDsc) Less(i, j int) bool { return p[i] > p[j] }
func (p IntegersDsc) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
