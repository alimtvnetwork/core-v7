package corevalidator

type Condition struct {
	IsTrimCompare        bool
	IsUniqueWordOnly     bool
	IsNonEmptyWhitespace bool // Split by whitespace and then compare, don't keep whitespace into comparison
	IsSortStringsBySpace bool
}

func (it *Condition) IsSplitByWhitespace() bool {
	return it.IsUniqueWordOnly ||
		it.IsNonEmptyWhitespace ||
		it.IsSortStringsBySpace
}
