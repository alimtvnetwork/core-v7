package reqtype

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

var (
	Ranges = [...]string{
		Invalid:                          "Invalid",
		Create:                           "CreateUsingAliasMap",
		Read:                             "Read",
		Update:                           "Update",
		Delete:                           "Delete",
		Drop:                             "Drop",
		CreateOrUpdate:                   "CreateOrUpdate",
		ExistCheck:                       "ExistCheck",
		SkipOnExist:                      "SkipOnExist",
		CreateOrSkipOnExist:              "CreateOrSkipOnExist",
		UpdateOrSkipOnNonExist:           "UpdateOrSkipOnNonExist",
		DeleteOrSkipOnNonExist:           "DeleteOrSkipOnNonExist",
		DropOrSkipOnNonExist:             "DropOrSkipOnNonExist",
		UpdateOnExist:                    "UpdateOnExist",
		DropOnExist:                      "DropOnExist",
		DropCreate:                       "DropCreate",
		Append:                           "Append",
		AppendByCompare:                  "AppendByCompare",
		AppendByCompareWhereCommentFound: "AppendByCompareWhereCommentFound",
		AppendLinesByCompare:             "AppendLinesByCompare",
		AppendLines:                      "AppendLines",
		CreateOrAppend:                   "CreateOrAppend",
		Prepend:                          "Prepend",
		CreateOrPrepend:                  "CreateOrPrepend",
		PrependLines:                     "PrependLines",
		Rename:                           "Rename",
		Change:                           "Change",
		Merge:                            "Merge",
		MergeLines:                       "MergeLines",
		GetHttp:                          "GetHttp",
		PutHttp:                          "PutHttp",
		PostHttp:                         "PostHttp",
		DeleteHttp:                       "DeleteHttp",
		PatchHttp:                        "PatchHttp",
		Touch:                            "Touch",
		Start:                            "Start",
		Stop:                             "Stop",
		Restart:                          "Restart",
		Reload:                           "Reload",
		StopSleepStart:                   "StopSleepStart",
		Suspend:                          "Suspend",
		Pause:                            "Pause",
		Resumed:                          "Resumed",
		TryRestart3Times:                 "TryRestart3Times",
		TryRestart5Times:                 "TryRestart5Times",
		TryStart3Times:                   "TryStart3Times",
		TryStart5Times:                   "TryStart5Times",
		TryStop3Times:                    "TryStop3Times",
		TryStop5Times:                    "TryStop5Times",
		InheritOnly:                      "InheritOnly",
		InheritPlusOverride:              "InheritPlusOverride",
		DynamicAction:                    "DynamicAction",
		Override:                         "Override",
		Overwrite:                        "Overwrite",
		Enforce:                          "Enforce",
	}

	httpRequests = [...]bool{
		GetHttp:    true,
		PutHttp:    true,
		PostHttp:   true,
		DeleteHttp: true,
		PatchHttp:  true,
	}

	actionRequests = [...]bool{
		Start:   true,
		Stop:    true,
		Restart: true,
		Reload:  true,
		Suspend: true,
		Pause:   true,
		Resumed: true,
	}

	RangesMap = map[string]Request{
		"Invalid":                          Invalid,
		"CreateUsingAliasMap":              Create,
		"Read":                             Read,
		"Update":                           Update,
		"Delete":                           Delete,
		"Drop":                             Drop,
		"CreateOrUpdate":                   CreateOrUpdate,
		"ExistCheck":                       ExistCheck,
		"SkipOnExist":                      SkipOnExist,
		"CreateOrSkipOnExist":              CreateOrSkipOnExist,
		"UpdateOrSkipOnNonExist":           UpdateOrSkipOnNonExist,
		"DeleteOrSkipOnNonExist":           DeleteOrSkipOnNonExist,
		"DropOrSkipOnNonExist":             DropOrSkipOnNonExist,
		"UpdateOnExist":                    UpdateOnExist,
		"DropOnExist":                      DropOnExist,
		"DropCreate":                       DropCreate,
		"Append":                           Append,
		"AppendByCompare":                  AppendByCompare,
		"AppendByCompareWhereCommentFound": AppendByCompareWhereCommentFound,
		"AppendLinesByCompare":             AppendLinesByCompare,
		"AppendLines":                      AppendLines,
		"CreateOrAppend":                   CreateOrAppend,
		"Prepend":                          Prepend,
		"CreateOrPrepend":                  CreateOrPrepend,
		"PrependLines":                     PrependLines,
		"Rename":                           Rename,
		"Change":                           Change,
		"Merge":                            Merge,
		"MergeLines":                       MergeLines,
		"GetHttp":                          GetHttp,
		"PutHttp":                          PutHttp,
		"PostHttp":                         PostHttp,
		"DeleteHttp":                       DeleteHttp,
		"PatchHttp":                        PatchHttp,
		"Touch":                            Touch,
		"Start":                            Start,
		"Stop":                             Stop,
		"Restart":                          Restart,
		"Reload":                           Reload,
		"StopSleepStart":                   StopSleepStart,
		"Suspend":                          Suspend,
		"Pause":                            Pause,
		"Resumed":                          Resumed,
		"TryRestart3Times":                 TryRestart3Times,
		"TryRestart5Times":                 TryRestart5Times,
		"TryStart3Times":                   TryStart3Times,
		"TryStart5Times":                   TryStart5Times,
		"TryStop3Times":                    TryStop3Times,
		"TryStop5Times":                    TryStop5Times,
		"InheritOnly":                      InheritOnly,
		"InheritPlusOverride":              InheritPlusOverride,
		"DynamicAction":                    DynamicAction,
		"Overwrite":                        Overwrite,
		"Override":                         Override,
		"Enforce":                          Enforce,
	}

	overrideLogicallyMap = map[Request]bool{
		Overwrite: true,
		Override:  true,
		Enforce:   true,
	}

	createMap = map[Request]bool{
		Create:              true,
		CreateOrUpdate:      true,
		CreateOrSkipOnExist: true,
		DropCreate:          true,
	}

	createUpdateMap = map[Request]bool{
		Create:                 true,
		Update:                 true,
		CreateOrUpdate:         true,
		CreateOrSkipOnExist:    true,
		UpdateOrSkipOnNonExist: true,
		UpdateOnExist:          true,
		CreateOrAppend:         true,
	}

	dropMap = map[Request]bool{
		Drop:                   true,
		DeleteOrSkipOnNonExist: true,
		DropOrSkipOnNonExist:   true,
		DropOnExist:            true,
		DropCreate:             true,
	}

	readOrEditMap = map[Request]bool{
		Read:                   true,
		Update:                 true,
		CreateOrUpdate:         true,
		CreateOrSkipOnExist:    true,
		UpdateOrSkipOnNonExist: true,
		UpdateOnExist:          true,
		Rename:                 true,
		Change:                 true,
	}

	crudMap = map[Request]bool{
		Create:                 true,
		Read:                   true,
		Update:                 true,
		Delete:                 true,
		Drop:                   true,
		CreateOrUpdate:         true,
		CreateOrSkipOnExist:    true,
		UpdateOrSkipOnNonExist: true,
		DeleteOrSkipOnNonExist: true,
		DropOrSkipOnNonExist:   true,
		UpdateOnExist:          true,
		DropOnExist:            true,
		DropCreate:             true,
	}

	editOrUpdateMap = map[Request]bool{
		Create:                 true,
		Update:                 true,
		Delete:                 true,
		CreateOrUpdate:         true,
		CreateOrSkipOnExist:    true,
		UpdateOrSkipOnNonExist: true,
		UpdateOnExist:          true,
		DropCreate:             true,
	}

	updateOrRemoveMap = map[Request]bool{
		Update:               true,
		CreateOrUpdate:       true,
		DropOrSkipOnNonExist: true,
		UpdateOnExist:        true,
		DropOnExist:          true,
		DropCreate:           true,
	}

	isExistOrSkipOnExistMap = map[Request]bool{
		ExistCheck:             true,
		SkipOnExist:            true,
		CreateOrSkipOnExist:    true,
		UpdateOrSkipOnNonExist: true,
		DeleteOrSkipOnNonExist: true,
		DropOrSkipOnNonExist:   true,
		UpdateOnExist:          true,
		DropOnExist:            true,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		reflectinternal.TypeName(Invalid),
		Ranges[:])
)
