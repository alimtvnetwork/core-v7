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
)
