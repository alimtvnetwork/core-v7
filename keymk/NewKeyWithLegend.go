package keymk

// NewKeyWithLegend
//
// Chain Sequence (Root-Package-Group-State-User-Item)
func NewKeyWithLegend(
	option *Option,
	legendName LegendName,
	isAttachLegendNames bool,
	rootName,
	packageName,
	group,
	stateName string,
) *KeyWithLegend {
	keyWithLegend := &KeyWithLegend{
		option:              option,
		LegendName:          legendName,
		rootName:            rootName,
		packageName:         packageName,
		groupName:           group,
		stateName:           stateName,
		isAttachLegendNames: isAttachLegendNames,
	}

	return keyWithLegend
}
