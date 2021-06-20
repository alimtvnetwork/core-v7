package scripttype

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
)

var (
	scriptTypeStringRanges = [...]string{
		Uninitialized: "Uninitialized",
		Shell:         "Shell",
		Bash:          "Bash",
		Perl:          "Perl",
		Python:        "Python",
		Python2:       "Python2",
		Python3:       "Python3",
		CLang:         "CLang",
		MakeScript:    "MakeScript",
		Powershell:    "Powershell",
		Cmd:           "Cmd",
	}

	scriptTypeRanges = [...]Variant{
		Uninitialized: Uninitialized,
		Shell:         Shell,
		Bash:          Bash,
		Perl:          Perl,
		Python:        Python,
		Python2:       Python2,
		Python3:       Python3,
		CLang:         CLang,
		MakeScript:    MakeScript,
		Powershell:    Powershell,
		Cmd:           Cmd,
	}

	RangesMap = map[Variant]*ScriptDefault{
		Uninitialized: {
			ScriptType:    Uninitialized,
			IsImplemented: false,
		},
		Shell: {
			ScriptType:  Shell,
			ProcessName: constants.BinShellCmd,
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: true,
		},
		Bash: {
			ScriptType:  Bash,
			ProcessName: constants.BashDefaultPath,
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: true,
		},
		Perl: {
			ScriptType:  Perl,
			ProcessName: "perl",
			DefaultArguments: []string{
				"-e",
			},
			IsImplemented: true,
		},
		Python: {
			ScriptType:  Python,
			ProcessName: "python",
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: false,
		},
		Python2: {
			ScriptType:  Python2,
			ProcessName: "python2",
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: false,
		},
		Python3: {
			ScriptType:  Python3,
			ProcessName: "python3",
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: true,
		},
		CLang: {
			ScriptType:  CLang,
			ProcessName: "gcc",
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: false,
		},
		MakeScript: {
			ScriptType:  MakeScript,
			ProcessName: "python3",
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: false,
		},
		Powershell: {
			ScriptType:  Powershell,
			ProcessName: "pwsh",
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: true,
		},
		Cmd: {
			ScriptType:  Cmd,
			ProcessName: "cmd",
			DefaultArguments: []string{
				"/c",
			},
			IsImplemented: true,
		},
	}

	scriptTypeBasicEnumImpl = enumimpl.
		NewBasicByteUsingIndexedSlice(scriptTypeStringRanges[:])
)
