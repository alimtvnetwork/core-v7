package keymk

func NewStrings(
	option *Option,
	main string,
	starterKeyChains ...string,
) *Key {
	slice := make([]string, 0, len(starterKeyChains)+DefaultCap)

	key := &Key{
		option:    option,
		mainName:  main,
		keyChains: slice,
	}

	if len(starterKeyChains) > 0 {
		key.AppendChainStrings(starterKeyChains...)
	}

	return key
}
