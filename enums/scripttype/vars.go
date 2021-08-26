package scripttype

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

var (
	bashDefaultScript      = RangesMap[Bash]
	cmdDefaultScript       = RangesMap[Cmd]
	scriptTypeStringRanges = [...]string{
		Uninitialized: "Uninitialized",
		Default:       "Default",
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
		Default:       Default,
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
		Default: {
			ScriptType:  Default,
			ProcessName: "",
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
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
				constants.NonInteractivePerlFlag,
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
			IsImplemented: true,
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
				constants.NonInteractiveCmdFlag,
			},
			IsImplemented: true,
		},
	}

	scriptTypeBasicEnumImpl = enumimpl.NewBasicByteUsingIndexedSlice(
		reflectinternal.TypeName(Uninitialized),
		scriptTypeStringRanges[:])
)
