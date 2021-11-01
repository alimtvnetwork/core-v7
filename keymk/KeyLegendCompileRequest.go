package keymk

type KeyLegendCompileRequest struct {
	StateName,
	UserId,
	GroupId,
	ItemId string
}

func (it KeyLegendCompileRequest) NewKeyLegend(
	option *Option,
	legendName LegendName,
	isAttachLegend bool,
	rootName,
	packageName,
	stateName string,
) *KeyWithLegend {
	return NewKeyWithLegend(
		option,
		legendName,
		isAttachLegend,
		rootName,
		packageName,
		it.GroupId,
		stateName)
}

func (it KeyLegendCompileRequest) NewKeyLegendDefaults(
	rootName,
	packageName,
	stateName string,
) *KeyWithLegend {
	return NewKeyWithLegend(
		JoinerOption,
		ShortLegends,
		false,
		rootName,
		packageName,
		it.GroupId,
		stateName)
}
