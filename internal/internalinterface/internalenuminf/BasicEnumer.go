package internalenuminf

type BasicEnumer interface {
	enumNameStinger
	nameValuer
	IsNameEqualer
	toNumberStringer
	IsValidInvalidChecker
	RangeNamesCsvGetter
}
