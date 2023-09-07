package internalenuminf

type PrivilegeTyper interface {
	BasicEnumer
	ByteValuePlusEqualer

	IsNone() bool
	IsAll() bool
	IsSelect() bool
	IsRead() bool
	IsReadOrSelect() bool
	IsInsert() bool
	IsCreate() bool
	IsUpdate() bool
	IsAlter() bool
	IsRenameOrChange() bool
	IsDelete() bool
	// IsDrop
	//
	//  Refers to table drop and delete refers to delete records
	IsDrop() bool
	IsExecute() bool
	IsEvent() bool
	IsCreateView() bool
	IsIndex() bool
	IsLockTables() bool
	IsReferences() bool
	IsTrigger() bool
	IsShowView() bool
	IsAllOrValue(value byte) bool
	IsCreateOrUpdate() bool
	IsInsertOrUpdate() bool
	IsCreateOrUpdateOrInsertLogically() bool
	IsDropLogically() bool
	IsCrudOnlyLogically() bool
	IsReadOrEditLogically() bool
	IsReadOrUpdateLogically() bool
	IsEditOrUpdateLogically() bool
	IsUpdateOrRemoveLogically() bool
}
