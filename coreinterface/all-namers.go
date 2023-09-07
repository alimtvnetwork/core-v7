package coreinterface

type EntityTypeNamer interface {
	EntityTypeName() string
}

type CategoryNamer interface {
	CategoryName() string
}

type TableNamer interface {
	TableName() string
}

type ToNamer interface {
	Name() string
}
