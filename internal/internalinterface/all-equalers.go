package internalinterface

type IdentifierWithEqualer interface {
	IdentifierGetter
	IsIdentifierEqualer
}

type IsTypeNameEqualer interface {
	IsTypeName(typeName string) bool
}

type IsCategoryNameEqualer interface {
	IsCategoryName(categoryName string) bool
}

type IsEntityTypeNameEqualer interface {
	IsEntityTypeName(entityName string) bool
}

type TypeNameWithEqualer interface {
	TypeNameGetter
	IsTypeNameEqualer
}

type CategoryNameWithEqualer interface {
	CategoryNamer
	IsCategoryNameEqualer
}

type EntityTypeNameWithEqualer interface {
	EntityTypeNamer
	IsEntityTypeNameEqualer
}

type IsNameEqualer interface {
	IsNameEqual(name string) bool
}

type NameWithNameEqualer interface {
	Namer
	IsNameEqualer
}
