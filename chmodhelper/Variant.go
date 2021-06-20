package chmodhelper

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
	X111                                Variant = "111"
	X222                                Variant = "222"
	X333                                Variant = "333"
	X444                                Variant = "444"
	X555                                Variant = "555"
	X644                                Variant = "644"
	X655                                Variant = "655"
	X666                                Variant = "666"
	X677                                Variant = "677"
	X711                                Variant = "711"
	X722                                Variant = "722"
	X733                                Variant = "733"
	X744                                Variant = "744"
	X755                                Variant = "755"
	X766                                Variant = "766"
	X777                                Variant = "777"
)

func (variant *Variant) String() string {
	return string(*variant)
}

func (variant *Variant) ExpandOctalByte() (r7, w7, x7 byte) {
	return ExpandCharRwx(string(*variant))
}

func (variant *Variant) ToWrapper() (RwxWrapper, error) {
	return NewUsingVariant(*variant)
}

func (variant *Variant) ToWrapperPtr() (*RwxWrapper, error) {
	return NewUsingVariantPtr(*variant)
}
