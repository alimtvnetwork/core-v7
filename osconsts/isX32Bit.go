package osconsts

func isX32Bit() bool {
	value32Bit, has32Bit := X32ArchitecturesMap[CurrentSystemArchitecture]

	return has32Bit && value32Bit
}
