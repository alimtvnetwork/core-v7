package enumtype

var (
	rangesMap = [...]string{
		Invalid:           "Invalid",
		Boolean:           "Boolean",
		Byte:              "Byte",
		UnsignedInteger16: "UnsignedInteger16",
		UnsignedInteger32: "UnsignedInteger32",
		UnsignedInteger64: "UnsignedInteger64",
		Integer8:          "Integer8",
		Integer16:         "Integer16",
		Integer32:         "Integer32",
		Integer64:         "Integer64",
		Integer:           "Integer",
		String:            "String",
	}

	stringToVariantMap = map[string]Variant{
		"Invalid":           Invalid,
		"Boolean":           Boolean,
		"Byte":              Byte,
		"UnsignedInteger16": UnsignedInteger16,
		"UnsignedInteger32": UnsignedInteger32,
		"UnsignedInteger64": UnsignedInteger64,
		"Integer8":          Integer8,
		"Integer16":         Integer16,
		"Integer32":         Integer32,
		"Integer64":         Integer64,
		"Integer":           Integer,
		"String":            String,
	}

	unSignedMap = map[Variant]bool{
		Byte:              true,
		UnsignedInteger16: true,
		UnsignedInteger32: true,
		UnsignedInteger64: true,
	}

	numbersMap = map[Variant]bool{
		Byte:              true,
		UnsignedInteger16: true,
		UnsignedInteger32: true,
		UnsignedInteger64: true,
		Integer8:          true,
		Integer16:         true,
		Integer32:         true,
		Integer64:         true,
		Integer:           true,
	}

	integersMap = map[Variant]bool{
		Integer8:  true,
		Integer16: true,
		Integer32: true,
		Integer64: true,
		Integer:   true,
	}
)
