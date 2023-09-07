package internalenuminf

type OnOffLater interface {
	BasicEnumer
	ByteValuePlusEqualer

	IsOn() bool
	IsOff() bool
	IsLater() bool

	OnOffName() string
	OnOffLowercaseName() string
}
