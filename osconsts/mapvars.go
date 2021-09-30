package osconsts

var (
	X32ArchitecturesMap = map[string]bool{
		"386":         true,
		"arm":         true,
		"armbe":       true,
		"mips":        true,
		"amd64p32":    true,
		"mips64p32":   true,
		"mips64p32le": true,
		"ppc":         true,
		"riscv":       true,
		"s390":        true,
		"sparc":       true,
	}

	X64ArchitecturesMap = map[string]bool{
		"amd64":    true,
		"arm64":    true,
		"ppc64":    true,
		"ppc64le":  true,
		"mips64":   true,
		"mips64le": true,
		"riscv64":  true,
		"s390x":    true,
		"wasm":     true,
		"arm64be":  true,
		"sparc64":  true,
	}

	UnixGroupsMap = map[string]bool{
		"android":   true,
		"darwin":    true,
		"dragonfly": true,
		"freebsd":   true,
		"linux":     true,
		"nacl":      true,
		"netbsd":    true,
		"openBSD":   true,
		"plan9":     true,
		"solaris":   true,
	}
)
