package keymk

func NewKeyWithLegendShortLegend(
	option *Option,
	rootName,
	packageName,
	group string,
) *KeyWithLegend {
	keyWithLegend := &KeyWithLegend{
		option:              option,
		LegendName:          ShortLegends,
		rootName:            rootName,
		packageName:         packageName,
		groupName:           group,
		isAttachLegendNames: true,
	}

	return keyWithLegend
}
