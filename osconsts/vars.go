package osconsts

//goland:noinspection ALL
var (
	IsX32Architecture = isX32Bit()
	IsX64Architecture = !IsX32Architecture
	windowsCDrive     = WindowsCDrive
)
