package coredata

type Integers []int

func (p Integers) Len() int {
	if p == nil {
		return 0
	}

	return len(p)
}

func (p Integers) Less(i, j int) bool { return p[i] < p[j] }
func (p Integers) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
