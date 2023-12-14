package reflectmodel

import "reflect"

// FieldProcessor
//
// To get the value use
// Field.Interface() or field.Int() or field.String() depending on the type
type FieldProcessor struct {
	Name      string
	Index     int
	Field     reflect.StructField // To get value use Field.Interface()
	FieldType reflect.Type
}

func (it *FieldProcessor) IsFieldType(t reflect.Type) bool {
	return it != nil && it.FieldType == t
}

func (it *FieldProcessor) IsFieldKind(k reflect.Kind) bool {
	return it != nil && it.FieldType.Kind() == k
}
