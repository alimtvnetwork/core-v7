package constants

//goland:noinspection ALL
var (
	// Copied from golang strings
	AsciiSpace = [256]uint8{
		TabByte:            One,
		LineFeedUnixByte:   One,
		TabVByte:           One,
		FormFeedByte:       One,
		CarriageReturnByte: One,
		SpaceByte:          One,
		0x85:               One, // reference : https://bit.ly/2JWdIoj
		0xA0:               One, // reference : https://bit.ly/2JWdIoj
	}

	// FormFeed \f is also marked as newline here.
	AsciiNewLinesChars = [256]uint8{
		LineFeedUnix:   One,
		FormFeed:       One,
		CarriageReturn: One,
	}

	SpecialChars = [256]uint8{
		'!': One,
		'@': One,
		'#': One,
		'$': One,
		'%': One,
		'^': One,
		'&': One,
		'*': One,
		'(': One,
		')': One,
	}

	BracketChars = [256]uint8{
		'[': One,
		']': One,
		'{': One,
		'}': One,
		'(': One,
		')': One,
		'<': One,
		'>': One,
	}

	EmptyStrings       []string
	EmptyPtrStrings    []*string
	EmptyInts          []int
	EmptyBytes         []byte
	EmptyFloats        []float32
	EmptyFloat64s      []float64
	EmptyInterfaces    []any
	EmptyIntToIntsMap  map[int][]int
	EmptyIntToBytesMap map[int][]byte
	EmptyStringMap     map[string]string
	EmptyStrToIntsMap  map[string][]int
	EmptyStrToBytesMap map[string][]byte
	EmptyStringsMap    map[string][]string

	EmptyIntToPtrIntsMap  map[int]*[]int
	EmptyIntToPtrBytesMap map[int]*[]byte
	EmptyStrToPtrIntsMap  map[string]*[]int
	EmptyStrToPtrBytesMap map[string]*[]byte
	EmptyPtrStringsMap    map[string]*[]string
)
