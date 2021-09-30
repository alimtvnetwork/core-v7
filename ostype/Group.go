package ostype

type Group byte

const (
	WindowsGroup Group = iota
	UnixGroup
	AndroidGroup
	UnknownGroup
)

func (group Group) Is(another Group) bool {
	return group == another
}

func (group Group) IsWindows() bool {
	return group == WindowsGroup
}

func (group Group) IsUnix() bool {
	return group == UnixGroup
}

func (group Group) IsAndroid() bool {
	return group == AndroidGroup
}

func (group Group) IsUnknown() bool {
	return group == UnknownGroup
}

func (group Group) Value() byte {
	return byte(group)
}

func (group Group) Byte() byte {
	return byte(group)
}

func (group Group) String() string {
	return osGroups[group]
}
