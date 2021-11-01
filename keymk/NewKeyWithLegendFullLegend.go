package keymk

// NewKeyWithLegendFullLegend
//
// Chain Sequence (Root-Package-Group-State-User-Item)
func NewKeyWithLegendFullLegend(
	option *Option,
	rootName,
	packageName,
	group string,
) *KeyWithLegend {
	keyWithLegend := &KeyWithLegend{
		option:              option,
		LegendName:          FullLegends,
		rootName:            rootName,
		packageName:         packageName,
		groupName:           group,
		isAttachLegendNames: true,
	}

	return keyWithLegend
}
