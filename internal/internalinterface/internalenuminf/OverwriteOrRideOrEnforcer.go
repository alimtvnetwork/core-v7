package internalenuminf

type OverwriteOrRideOrEnforcer interface {
	BasicEnumer
	ByteValuePlusEqualer

	IsOverwrite() bool
	IsOverride() bool
	IsEnforce() bool

	IsOverrideOrOverwriteOrEnforce() bool
	IsNotOverrideOrOverwriteOrEnforce() bool
}
