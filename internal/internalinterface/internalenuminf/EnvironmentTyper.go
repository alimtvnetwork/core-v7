package internalenuminf

type EnvironmentTyper interface {
	BasicEnumer
	ByteValuePlusEqualer

	NameValue() string
	KeyName() string
	CurlyKeyName() string

	ToNumberString() string
	IsValidInvalidChecker
	IsUninitialized() bool
	IsInitialized() bool

	IsDevelopment() bool
	IsDevelopment1() bool
	IsDevelopment2() bool
	IsAnyDevelopment() bool

	IsTest() bool
	IsTest1() bool
	IsTest2() bool
	IsTestEnvLogically() bool
	IsAnyTestEnv() bool
	IsNotTestEnvLogically() bool

	IsDevEnvLogically() bool
	IsNotDevEnvLogically() bool

	IsProdEnvLogically() bool
	IsNotProdEnvLogically() bool

	IsProduction() bool
	IsProduction1() bool
	IsProduction2() bool
	IsAnyProduction() bool

	VersionNumber() int
}
