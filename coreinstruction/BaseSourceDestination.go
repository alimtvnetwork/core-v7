package coreinstruction

type BaseSourceDestination struct {
	SourceDestination
}

func NewBaseSourceDestination(
	src, dst string,
) BaseSourceDestination {
	return BaseSourceDestination{
		SourceDestination{
			Source:      src,
			Destination: dst,
		},
	}
}
