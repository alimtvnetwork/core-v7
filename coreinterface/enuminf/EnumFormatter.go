package enuminf

// EnumFormatter
//
//	Outputs name and
//	value by given format.
//
// sample-format :
//   - "Enum of {type-name} - {name} - {value}"
//
// sample-format-output :
//   - "Enum of EnumFullName - Invalid - 0"
//
// Key-Meaning :
//   - {type-name} : represents type-name string
//   - {name}      : represents name string
//   - {value}     : represents value string
type EnumFormatter interface {
	// Format
	//
	//  Outputs name and
	//  value by given format.
	//
	// sample-format :
	//  - "Enum of {type-name} - {name} - {value}"
	//
	// sample-format-output :
	//  - "Enum of EnumFullName - Invalid - 0"
	//
	// Key-Meaning :
	//  - {type-name} : represents type-name string
	//  - {name}      : represents name string
	//  - {value}     : represents value string
	Format(format string) (compiled string)
}
