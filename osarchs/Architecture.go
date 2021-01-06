package osarchs

type Architecture byte

var (
	architectures = []string{
		"X32",
		"X64",
		"Unknown",
	}
)

const (
	X32 Architecture = iota
	X64
	Unknown
)

func (arch Architecture) Value() byte {
	return byte(arch)
}

func (arch Architecture) ValueInt() int {
	return int(arch)
}

func (arch Architecture) String() string {
	return architectures[arch]
}
