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
		UbuntuDesktop:        "UbuntuDesktop",
		Centos:               "Centos",
		Centos7:              "Centos7",
		Centos8:              "Centos8",
		Centos9:              "Centos9",
		DebianServer:         "DebianServer",
		DebianDesktop:        "DebianDesktop",
		Docker:               "Docker",
		DockerUbuntuServer:   "DockerUbuntuServer",
		DockerUbuntuServer20: "DockerUbuntuServer20",
		DockerUbuntuServer21: "DockerUbuntuServer21",
		DockerCentos9:        "DockerCentos9",
		Android:              "Android",
	}

	RangesMap = map[string]Variation{
		"Unknown":              Unknown,
		"UbuntuServer":         UbuntuServer,
		"UbuntuServer18":       UbuntuServer18,
		"UbuntuServer19":       UbuntuServer19,
		"UbuntuServer20":       UbuntuServer20,
		"UbuntuServer21":       UbuntuServer21,
		"UbuntuDesktop":        UbuntuDesktop,
		"Centos":               Centos,
		"Centos7":              Centos7,
		"Centos8":              Centos8,
		"Centos9":              Centos9,
		"DebianServer":         DebianServer,
		"DebianDesktop":        DebianDesktop,
		"Docker":               Docker,
		"DockerUbuntuServer":   DockerUbuntuServer,
		"DockerUbuntuServer20": DockerUbuntuServer20,
		"DockerUbuntuServer21": DockerUbuntuServer21,
		"DockerCentos9":        DockerCentos9,
		"Android":              Android,
	}

	BasicEnumImpl = enumimpl.NewBasicByteUsingIndexedSlice(
		reflectinternal.TypeName(Unknown),
		Ranges[:])
)
