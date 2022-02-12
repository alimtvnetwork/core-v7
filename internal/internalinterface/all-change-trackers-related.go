package internalinterface

type CountStateTracker interface {
	IsSameStateUsingCount(
		currentCount int,
	) bool
	IsSameState() bool
	IsSuccessValidator
	HasChangesChecker
}

type HasChangesChecker interface {
	HasChanges() bool
}

type DynamicDiffChangesGetter interface {
	DynamicDiffChanges() ([]byte, error)
}

type DynamicChangeStateDetector interface {
	IsDynamicEqual(rawBytes []byte) bool
	DynamicDiffChangesGetter
	HasChangesChecker
	IsSuccessValidator
	ChangesLogger
	MustChangesLogger
}

type ChangesLogger interface {
	LogChanges() error
}

type MustChangesLogger interface {
	LogChangesMust()
}

type AcceptRejectOrSkipper interface {
	IsAccept() bool
	IsReject() bool
	IsSkip() bool
}

type YesNoAsker interface {
	IsYes() bool
	IsNo() bool
	IsAsk() bool
	IsIndeterminate() bool
}

type YesNoAcceptRejecter interface {
	YesNoAsker
	AcceptRejectOrSkipper
}
