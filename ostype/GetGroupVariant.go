package ostype

func GetGroupVariant() GroupVariant {
	group := GetCurrentGroup()
	variant := GetCurrentVariant()

	return GroupVariant{
		Group:     group,
		Variation: variant,
	}
}
