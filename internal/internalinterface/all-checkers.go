package internalinterface

type IsIdentifierEqualer interface {
	IsIdentifier(identifier string) bool
}

type IsIdEqualer interface {
	IsId(id string) bool
}

type IsIdUnsignedIntegerEqualer interface {
	IsId(id uint) bool
}

type HasErrorChecker interface {
	HasError() bool
}

type IsValidChecker interface {
	// IsValid similar or alias for IsSuccessChecker
	IsValid() bool
}

type IsInvalidChecker interface {
	IsInvalid() bool
}
type IsSuccessChecker interface {
	// IsSuccess No error
	IsSuccess() bool
}

type IsFailedChecker interface {
	// IsFailed has error or any other issues, or alias for HasIssues or HasError
	IsFailed() bool
}

type IsSuccessValidator interface {
	IsValidChecker
	IsSuccessChecker
	IsFailedChecker
}

type IsEmptyChecker interface {
	IsEmpty() bool
}

type IsEmptyErrorChecker interface {
	IsEmptyError() bool
}

type IsErrorEqualsChecker interface {
	IsErrorEquals(err error) bool
	IsErrorMessageEqual(msg string) bool
	IsErrorMessage(msg string, isCaseSensitive bool) bool
	IsErrorMessageContains(
		msg string,
		isCaseSensitive bool,
	) bool
}

type HasErrorOrHasAnyErrorChecker interface {
	HasError() bool
	HasAnyError() bool
}
