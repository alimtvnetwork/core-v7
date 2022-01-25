package keymk

import "gitlab.com/evatix-go/core/constants"

type newKeyCreator struct{}

func (it *newKeyCreator) Create(
	option *Option,
	main string,
) *Key {
	key := &Key{
		option:   option,
		mainName: main,
		keyChains: make(
			[]string,
			0,
			constants.Capacity7),
	}

	return key
}

func (it *newKeyCreator) PathTemplate(
	root string,
	starterKeyChains ...interface{},
) *Key {
	return it.All(
		CurlyBracePathJoinerOption,
		root,
		starterKeyChains...)
}

func (it *newKeyCreator) PathTemplateDefault(
	starterKeyChains ...interface{},
) *Key {
	return it.All(
		CurlyBracePathJoinerOption,
		constants.PathRootTemplate,
		starterKeyChains...)
}

func (it *newKeyCreator) PathTemplatePrefixRelativeIdDefault() *Key {
	return it.All(
		CurlyBracePathJoinerOption,
		root,
		prefix,
		relative,
		id)
}

func (it *newKeyCreator) PathTemplatePrefixRelativeIdFileDefault() *Key {
	return it.All(
		CurlyBracePathJoinerOption,
		root,
		prefix,
		relative,
		id,
		constants.FileKeyword)
}

func (it *newKeyCreator) All(
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

func (it *newKeyCreator) AllStrings(
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

func (it *newKeyCreator) StringsWithOptions(
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

func (it *newKeyCreator) Parenthesis(
	main string,
	starterKeyChains ...interface{},
) *Key {
	return it.All(
		ParenthesisJoinerOption,
		main,
		starterKeyChains...)
}

func (it *newKeyCreator) ParenthesisStrings(
	main string,
	starterKeyChains ...string,
) *Key {
	return it.AllStrings(
		ParenthesisJoinerOption,
		main,
		starterKeyChains...)
}

func (it *newKeyCreator) Curly(
	main string,
	starterKeyChains ...interface{},
) *Key {
	return it.All(
		CurlyBraceJoinerOption,
		main,
		starterKeyChains...)
}

func (it *newKeyCreator) CurlyStrings(
	main string,
	starterKeyChains ...string,
) *Key {
	return it.AllStrings(
		CurlyBraceJoinerOption,
		main,
		starterKeyChains...)
}

func (it *newKeyCreator) SquareBrackets(
	main string,
	starterKeyChains ...interface{},
) *Key {
	return it.All(
		BracketJoinerOption,
		main,
		starterKeyChains...)
}

func (it *newKeyCreator) SquareBracketsStrings(
	main string,
	starterKeyChains ...string,
) *Key {
	return it.AllStrings(
		BracketJoinerOption,
		main,
		starterKeyChains...)
}

func (it *newKeyCreator) Default(
	main string,
	starterKeyChains ...interface{},
) *Key {
	return it.All(
		JoinerOption,
		main,
		starterKeyChains...)
}

func (it *newKeyCreator) DefaultStrings(
	main string,
	starterKeyChains ...string,
) *Key {
	return it.Create(
		JoinerOption,
		main).
		AppendChainStrings(starterKeyChains...)
}

func (it *newKeyCreator) DefaultMain(
	main string,
) *Key {
	return it.Create(
		JoinerOption,
		main)
}

func (it *newKeyCreator) OptionMain(
	option *Option,
	main string,
) *Key {
	return it.Create(
		option,
		main)
}
