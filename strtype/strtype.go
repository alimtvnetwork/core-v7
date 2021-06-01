package strtype

type Variant string

func (v Variant) Value() string {
	return string(v)
}

func (v Variant) StringValue() string {
	return string(v)
}

// Add v + n
func (v Variant) Add(n string) Variant {
	return Variant(v.Value() + n)
}

func (v Variant) Is(n Variant) bool {
	return v.Value() == n.Value()
}

func (v Variant) IsEqual(n string) bool {
	return v.Value() == n
}

// IsGreater v.Value() > n
func (v Variant) IsGreater(n string) bool {
	return v.Value() > n
}

// IsGreaterEqual v.Value() >= n
func (v Variant) IsGreaterEqual(n string) bool {
	return v.Value() >= n
}

// IsLess v.Value() < n
func (v Variant) IsLess(n string) bool {
	return v.Value() < n
}

// IsLessEqual v.Value() <= n
func (v Variant) IsLessEqual(n string) bool {
	return v.Value() <= n
}
