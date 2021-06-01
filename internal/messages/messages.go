package messages

const (
	HyphenedRwxRwxRwxLengthMustBe10               = "`-rwxrwxrwx` length must be 10. Reference : https://ss64.com/bash/chmod.html"
	RwxRwxRwxLengthMustBe9                        = "`rwxrwxrwx` length must be 9."
	ModeCharShouldBeAllNumbersAndWithin0To7       = "mode char should be all digits and under 0 to 7"
	DynamicFailedToParseToFloat64BecauseNull      = "dynamic datatype failed to parse to float64 because it is nil."
	FailedToCompileChmodhelperVarWrapperToWrapper = "Failed to compile chmodhelper.VarWrapper to failedToExecute.Wrapper."
	FailedToGetFileModeRwx                        = "Failed to get the existing filemode (Rwx...)."
)
