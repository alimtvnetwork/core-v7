package internalinterface

type EntityTypeNamer interface {
	EntityTypeName() string
}

type CategoryNamer interface {
	CategoryName() string
}

type TableNamer interface {
	TableName() string
}

type Namer interface {
	Name() string
}
