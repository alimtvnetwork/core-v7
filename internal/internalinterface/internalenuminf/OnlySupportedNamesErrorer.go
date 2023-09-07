package internalenuminf

type OnlySupportedNamesErrorer interface {
	OnlySupportedErr(names ...string) error
	OnlySupportedMsgErr(message string, names ...string) error
}
