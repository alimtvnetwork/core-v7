package filemode

// 1 - Execute true
// 2 - Write true
// 3 - Write + Execute true
// 4 - Read true
// 5 - Read + Execute true
// 6 - Read + Write true
// 7 - Read + Write + Execute all true
type AttrVariant byte

//goland:noinspection ALL
const (
	Execute          AttrVariant = 1
	Write            AttrVariant = 2
	WriteExecute     AttrVariant = 3
	Read             AttrVariant = 4
	ReadExecute      AttrVariant = 5
	ReadWrite        AttrVariant = 6
	ReadWriteExecute AttrVariant = 7
)

// v > byte(attrVariant)
func (attrVariant AttrVariant) IsGreaterThan(v byte) bool {
	return v > byte(attrVariant)
}

func (attrVariant AttrVariant) String() string {
	return string(attrVariant)
}

func (attrVariant AttrVariant) Value() byte {
	return byte(attrVariant)
}

func (attrVariant AttrVariant) ToAttribute() Attribute {
	return NewAttributeUsingVariant(attrVariant)
}
