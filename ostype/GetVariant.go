package ostype

func GetVariant(rawRuntimeGoos string) Variation {
	variant, has := OsStringToVariantMap[rawRuntimeGoos]

	if has {
		return variant
	}

	return Unknown
}
