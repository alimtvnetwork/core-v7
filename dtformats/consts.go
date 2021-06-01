package dtformats

type Layout string

const (
	ANSIC       Layout = "Mon Jan _2 15:04:05 2006"
	UnixDate    Layout = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    Layout = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      Layout = "02 Jan 06 15:04 MST"
	RFC822Z     Layout = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      Layout = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     Layout = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    Layout = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     Layout = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano Layout = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     Layout = "3:04PM"
	Stamp       Layout = "Jan _2 15:04:05"
	StampMilli  Layout = "Jan _2 15:04:05.000"
	StampMicro  Layout = "Jan _2 15:04:05.000000"
	StampNano   Layout = "Jan _2 15:04:05.000000000"
)

func (layout Layout) Value() string {
	return string(layout)
}

func (layout Layout) Is(format string) bool {
	return string(layout) == format
}

func (layout Layout) IsTimeOnly() bool {
	return layout == Kitchen
}

func (layout Layout) IsTimeFocused() bool {
	return layout == Stamp ||
		layout == StampMilli ||
		layout == StampMicro ||
		layout == Kitchen ||
		layout == StampNano
}

func (layout Layout) IsDateTime() bool {
	return layout == ANSIC ||
		layout == UnixDate ||
		layout == RubyDate ||
		layout == RFC822 ||
		layout == RFC822Z ||
		layout == RFC850 ||
		layout == RFC1123 ||
		layout == RFC1123Z ||
		layout == RFC3339 ||
		layout == RFC3339Nano
}

func (layout Layout) HasTimeZone() bool {
	return layout == UnixDate ||
		layout == RubyDate ||
		layout == RFC822 ||
		layout == RFC822Z ||
		layout == RFC850 ||
		layout == RFC1123 ||
		layout == RFC1123Z ||
		layout == RFC3339 ||
		layout == RFC3339Nano
}
