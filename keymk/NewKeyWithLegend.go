package keymk

// NewKeyWithLegend
//
// Chain Sequence (Root-Package-Group-User-Item)
func NewKeyWithLegend(
	option *Option,
	legendName LegendName,
	isAttachLegendNames bool,
	rootName,
	packageName,
	group string,
) *KeyWithLegend {
	keyWithLegend := &KeyWithLegend{
		option:              option,
		LegendName:          legendName,
		rootName:            rootName,
		packageName:         packageName,
		groupName:           group,
		isAttachLegendNames: isAttachLegendNames,
	}

	return keyWithLegend
}
