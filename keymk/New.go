package keymk

func New(
	option *Option,
	main string,
	starterKeyChains ...interface{},
) *Key {
	slice := make([]string, 0, len(starterKeyChains)+DefaultCap)

	key := &Key{
		option:    option,
		mainName:  main,
		keyChains: slice,
	}

	if len(starterKeyChains) > 0 {
		key.keyChains = appendAnyItemsWithBaseStrings(
			option.IsSkipEmptyEntry,
			key.keyChains,
			starterKeyChains)
	}

	return key
}
