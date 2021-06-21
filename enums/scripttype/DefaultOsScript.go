package scripttype

import "gitlab.com/evatix-go/core/osconsts"

func DefaultOsScript() *ScriptDefault {
	if osconsts.IsWindows {
		return cmdDefaultScript
	}

	return bashDefaultScript
}
