package reflectmodel

type ReflectValue struct {
	TypeName     string
	FieldsNames  []string
	MethodsNames []string
	RawData      interface{}
}
