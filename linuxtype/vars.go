package linuxtype

import (
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

var (
	Ranges = [...]string{
		Unknown:              "Unknown",
		UbuntuServer:         "UbuntuServer",
		UbuntuServer18:       "UbuntuServer18",
		UbuntuServer19:       "UbuntuServer19",
		UbuntuServer20:       "UbuntuServer20",
		UbuntuServer21:       "UbuntuServer21",
		UbuntuServer22:       "UbuntuServer22",
		UbuntuServer23:       "UbuntuServer23",
		Centos:               "Centos",
		Centos7:              "Centos7",
		Centos8:              "Centos8",
		Centos9:              "Centos9",
		Centos10:             "Centos10",
		Centos11:             "Centos11",
		Centos12:             "Centos12",
		CentosStream:         "CentosStream",
		DebianServer:         "DebianServer",
		DebianServer7:        "DebianServer7",
		DebianServer8:        "DebianServer8",
		DebianServer9:        "DebianServer9",
		DebianServer10:       "DebianServer10",
		DebianServer11:       "DebianServer11",
		DebianServer12:       "DebianServer12",
		DebianServer13:       "DebianServer13",
		DebianServer14:       "DebianServer14",
		DebianDesktop:        "DebianDesktop",
		Docker:               "Docker",
		DockerUbuntuServer:   "DockerUbuntuServer",
		DockerUbuntuServer18: "DockerUbuntuServer18",
		DockerUbuntuServer19: "DockerUbuntuServer19",
		DockerUbuntuServer20: "DockerUbuntuServer20",
		DockerUbuntuServer21: "DockerUbuntuServer21",
		DockerUbuntuServer22: "DockerUbuntuServer22",
		DockerCentos7:        "DockerCentos7",
		DockerCentos8:        "DockerCentos8",
		DockerCentos9:        "DockerCentos9",
		DockerCentos10:       "DockerCentos10",
		Android:              "Android",
		UbuntuDesktop:        "UbuntuDesktop",
	}

	BasicEnumImpl = enumimpl.NewBasicByteUsingIndexedSlice(
		reflectinternal.TypeName(Unknown),
		Ranges[:])
)
