package internalenuminf

type LogLevelTyper interface {
	BasicEnumer
	ByteValuePlusEqualer

	IsLevelNum(level int) bool
	IsLevel(level byte) bool
	IsLevel1() bool
	IsLevel2() bool
	IsLevel3() bool
	IsLevel4() bool
	IsLevel5() bool
	IsLevel6() bool
	IsLevel7() bool
	IsLevel8() bool
	IsLevel9() bool
	IsLevel10() bool
}
