package osconsts

var (
	IsX32Architecture = isX32Bit()
	IsX64Architecture = !IsX32Architecture
	windowsCDrive     = WindowsCDrive
)
