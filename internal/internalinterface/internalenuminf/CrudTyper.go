package internalenuminf

type CrudTyper interface {
	BasicEnumer
	ByteValuePlusEqualer

	IsNone() bool
	IsCreate() bool
	IsUpdate() bool
	IsRead() bool
	IsDelete() bool
	IsDrop() bool
	IsSkipOnExist() bool
	IsDropOnExist() bool
	IsCreateOrUpdate() bool
	IsCreateLogically() bool
	IsCreateOrUpdateLogically() bool
	IsDropLogically() bool
	IsUpdateOnExist() bool
	IsCrudOnlyLogically() bool
	IsNotCrudOnlyLogically() bool
	IsReadOrEditLogically() bool
	IsReadOrUpdateLogically() bool
	IsEditOrUpdateLogically() bool
	IsOnExistCheckLogically() bool
	IsOnExistOrSkipOnNonExistLogically() bool
	IsUpdateOrRemoveLogically() bool

	IsValidInvalidChecker
}
