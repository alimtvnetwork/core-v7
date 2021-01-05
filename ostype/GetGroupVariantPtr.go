package ostype

func GetGroupVariantPtr() *GroupVariant {
	group := GetCurrentGroup()
	variant := GetCurrentVariant()

	return &GroupVariant{
		Group:     group,
		Variation: variant,
	}
}
