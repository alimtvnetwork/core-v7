package taskcategory

import (
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Unknown:                      "Unknown",
		RegularTask:                  "RegularTask",
		RegularLockTask:              "RegularLockTask",
		InstructionRoot:              "InstructionRoot",
		SpecificInstructionWithType:  "SpecificInstructionWithType",
		DbTask:                       "DbTask",
		FsTask:                       "FsTask",
		PackageInstall:               "PackageInstall",
		PackageCleanup:               "PackageCleanup",
		PackageUninstall:             "PackageUninstall",
		PackageLock:                  "PackageLock",
		PackageSynchronize:           "PackageSynchronize",
		PackageDownload:              "PackageDownload",
		DatabaseCleanup:              "DatabaseCleanup",
		DatabaseMigrateAll:           "DatabaseMigrateAll",
		SpecificDatabaseMigrate:      "SpecificDatabaseMigrate",
		SoftwareUpdate:               "SoftwareUpdate",
		Reminder:                     "Reminder",
		Email:                        "Email",
		Synchronize:                  "Synchronize",
		Notification:                 "Notification",
		ErrorNotification:            "ErrorNotification",
		ServerStorageOutNotification: "ServerStorageOutNotification",
		WebServerDownNotification:    "WebServerDownNotification",
		MobileNotification:           "MobileNotification",
		EmailNotification:            "EmailNotification",
		Log:                          "Log",
		SystemCleanup:                "SystemCleanup",
		Cache:                        "Cache",
		CleanCache:                   "CleanCache",
	}

	BasicEnumImpl = enumimpl.NewBasicByteUsingIndexedSlice(
		coredynamic.TypeName(Unknown),
		Ranges[:])
)
