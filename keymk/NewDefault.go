package keymk

func NewDefault(
	main string,
	starterKeyChains ...interface{},
) *Key {
	return New(JoinerOption, main, starterKeyChains...)
}
