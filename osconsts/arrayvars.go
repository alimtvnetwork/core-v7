package osconsts

//goland:noinspection ALL
var (
	X32Architectures = []string{
		"386",
		"arm",
		"armbe",
		"mips",
		"amd64p32",
		"mips64p32",
		"mips64p32le",
		"ppc",
		"riscv",
		"s390",
		"sparc",
	}

	X64Architectures = []string{
		"amd64",
		"arm64",
		"ppc64",
		"ppc64le",
		"mips64",
		"mips64le",
		"riscv64",
		"s390x",
		"wasm",
		"arm64be",
		"sparc64",
	}

	UnixGroups = []string{
		"android",
		"darwin",
		"dragonfly",
		"freebsd",
		"linux",
		"nacl",
		"netbsd",
		"openBSD",
		"plan9",
		"solaris",
	}
)
