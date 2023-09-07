package entityinf

type EntityWriter interface {
	EntityWrite(record RecordEntityDefiner) error
	EntityWriteMust(record RecordEntityDefiner)
}
