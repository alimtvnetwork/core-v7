package coreinterface

// NameValuer
//
//   should be a combined string output using name[value]
//
// Sample :
//  - "`EnumName[EnumValInteger]" -> `EnumName[2]`
type NameValuer interface {
	// NameValue
	//
	//   should be a combined string output using name[value]
	//
	// Sample :
	//  - "`EnumName[EnumValInteger]" -> `EnumName[2]`
	NameValue() string
}
