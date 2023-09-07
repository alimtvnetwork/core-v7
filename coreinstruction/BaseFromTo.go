package coreinstruction

type BaseFromTo struct {
	FromTo
}

func NewBaseFromTo(
	from, to string,
) BaseFromTo {
	return BaseFromTo{
		FromTo{
			From: from,
			To:   to,
		},
	}
}
