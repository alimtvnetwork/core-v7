package enumimpl

func UnsupportedNames(
	allNames []string,
	supportedNames ...string,
) []string {
	unsupportedNames := make(
		[]string,
		0,
		len(allNames)-len(supportedNames)+1)
	supportedNamesHashset := toHashset(supportedNames...)

	for _, name := range allNames {
		_, has := supportedNamesHashset[name]

		if !has {
			unsupportedNames = append(
				unsupportedNames,
				name)
		}
	}

	return supportedNames
}
