package errcore

type CountStateChangeTracker struct {
	lengthGetter
	initLength int
}

func NewCountStateChangeTracker(
	lengthGetter lengthGetter,
) CountStateChangeTracker {
	return CountStateChangeTracker{
		lengthGetter: lengthGetter,
		initLength:   lengthGetter.Length(),
	}
}

func (it CountStateChangeTracker) IsSameStateUsingCount(
	currentCount int,
) bool {
	return currentCount == it.initLength
}

func (it CountStateChangeTracker) IsSameState() bool {
	return it.lengthGetter.Length() == it.initLength
}

func (it CountStateChangeTracker) IsValid() bool {
	return it.lengthGetter.Length() == it.initLength
}

func (it CountStateChangeTracker) IsSuccess() bool {
	return it.lengthGetter.Length() == it.initLength
}

func (it CountStateChangeTracker) IsFailed() bool {
	return it.lengthGetter.Length() != it.initLength
}

func (it CountStateChangeTracker) HasChanges() bool {
	return it.lengthGetter.Length() != it.initLength
}
