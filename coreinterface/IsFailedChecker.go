package coreinterface

type IsFailedChecker interface {
	// IsFailed has error or any other issues, or alias for HasIssues or HasError
	IsFailed() bool
}
