package filemode

type Variant string

//goland:noinspection ALL
const (
	AllRead                             Variant = "444"
	AllWrite                            Variant = "222"
	AllExecute                          Variant = "111"
	AllReadWrite                        Variant = "666"
	AllReadExecute                      Variant = "555"
	AllWriteExecute                     Variant = "333"
	OwnerAllReadWriteGroupOther         Variant = "755"
	ReadWriteOwnerReadGroupOther        Variant = "644"
	ReadWriteOwnerReadExecuteGroupOther Variant = "655"
	All                                 Variant = "777"
)

func (variant Variant) String() string {
	return string(variant)
}

func (variant Variant) ToWrapper() (Wrapper, error) {
	return NewUsingVariant(variant)
}
