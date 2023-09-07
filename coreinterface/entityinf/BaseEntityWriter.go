package entityinf

type BaseEntityWriter interface {
	BaseEntityWrite(record BaseRecordEntityDefiner) error
	BaseEntityWriteMust(record BaseRecordEntityDefiner)
}
