package internalenuminf

import "fmt"

type LinuxTyper interface {
	BasicEnumer
	ByteValuePlusEqualer

	IsUnknown() bool
	IsUbuntuServer() bool
	IsUbuntuServer18() bool
	IsUbuntuServer19() bool
	IsUbuntuServer20() bool
	IsUbuntuServer21() bool
	IsUbuntuDesktop() bool
	IsCentos() bool
	IsCentos7() bool
	IsCentos8() bool
	IsCentos9() bool
	IsDebianServer() bool
	IsDebianDesktop() bool
	IsDocker() bool
	IsDockerUbuntuServer() bool
	IsDockerUbuntuServer20() bool
	IsDockerUbuntuServer21() bool
	IsDockerCentos9() bool
	IsAndroid() bool
	fmt.Stringer
}
