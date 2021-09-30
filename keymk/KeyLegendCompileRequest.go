package keymk

type KeyLegendCompileRequest struct {
	UserId,
	GroupId,
	ItemId string
}

func (it KeyLegendCompileRequest) NewKeyLegend(
	option *Option,
	legendName LegendName,
	isAttachLegend bool,
	rootName,
	packageName string,
) *KeyWithLegend {
	return NewKeyWithLegend(
		option,
		legendName,
		isAttachLegend,
		rootName,
		packageName,
		it.GroupId)
}

func (it KeyLegendCompileRequest) NewKeyLegendDefaults(
	rootName,
	packageName string,
) *KeyWithLegend {
	return NewKeyWithLegend(
		JoinerOption,
		ShortLegends,
		false,
		rootName,
		packageName,
		it.GroupId)
}
