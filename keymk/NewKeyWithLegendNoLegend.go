package keymk

// NewKeyWithLegendNoLegend
//
// Chain Sequence (Root-Package-Group-State-User-Item)
func NewKeyWithLegendNoLegend(
	option *Option,
	rootName,
	packageName,
	group string,
) *KeyWithLegend {
	keyWithLegend := &KeyWithLegend{
		option:              option,
		rootName:            rootName,
		packageName:         packageName,
		groupName:           group,
		isAttachLegendNames: false,
	}

	return keyWithLegend
}
