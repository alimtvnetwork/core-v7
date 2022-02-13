package coreinterface

import "gitlab.com/evatix-go/core/internal/internalinterface"

type ChangesCommitter interface {
	HasChangesChecker
	ChangeAccepter
	ChangeRejecter
	RemindLaterChangeSkipper
	Commit(option AcceptRejectOrSkipper) error
	CommitMust(option AcceptRejectOrSkipper)
}

type ChangeAccepter interface {
	AcceptChanges() error
	AcceptChangesMust()
}

type ChangeRejecter interface {
	RejectChanges() error
	RejectChangesMust()
}

type RemindLaterChangeSkipper interface {
	SkipChangesRemindLater() error
	SkipChangesRemindLaterMust()
}

type CountStateTracker interface {
	internalinterface.CountStateTracker
}

type DynamicDiffChangesGetter interface {
	internalinterface.DynamicDiffChangesGetter
}

type HasChangesChecker interface {
	internalinterface.HasChangesChecker
}

type DynamicChangeStateDetector interface {
	internalinterface.DynamicChangeStateDetector
}

type ChangesLogger interface {
	internalinterface.ChangesLogger
}

type MustChangesLogger interface {
	internalinterface.MustChangesLogger
}

type YesNoAsker interface {
	internalinterface.YesNoAsker
}

type AcceptRejectOrSkipper interface {
	internalinterface.AcceptRejectOrSkipper
}

type YesNoAcceptRejecter interface {
	internalinterface.YesNoAcceptRejecter
}
