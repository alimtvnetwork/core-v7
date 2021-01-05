package osarchs

import (
	"gitlab.com/evatix-go/core/osconsts"
)

func Get(rawGoArch string) Architecture {
	_, isIn32Bit := osconsts.X32ArchitecturesMap[rawGoArch]

	if isIn32Bit {
		return X32
	}

	_, is64Bit := osconsts.X64ArchitecturesMap[rawGoArch]

	if is64Bit {
		return X64
	}

	return Unknown
}
