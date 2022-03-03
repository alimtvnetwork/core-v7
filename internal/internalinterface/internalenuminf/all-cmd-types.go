package internalenuminf

type BaseCmdTyper interface {
	BasicEnumer
	ByteValuePlusEqualer

	FullName() string
	// HyphenName
	//
	//  "--ToNameLower()"
	HyphenName() string
	ToNameLower
	CommandName() string
	AlternativeCommandNames() map[string]bool
}

type ToNameLower interface {
	ToNameLower() string
}

type ListCmdTyper interface {
	BaseCmdTyper

	IsList() bool
	IsListJson() bool
}

type CreateCmdTyper interface {
	BaseCmdTyper

	IsCreate() bool
}

type AddCmdTyper interface {
	BaseCmdTyper

	IsAdd() bool
}

type RemoveCmdTyper interface {
	BaseCmdTyper

	IsRemove() bool
}

type ImportCmdTyper interface {
	BaseCmdTyper

	IsImport() bool
	IsImportSpecific() bool
}

type ExportCmdTyper interface {
	BaseCmdTyper

	IsExport() bool
	IsExportSpecific() bool
}

type BackupCmdTyper interface {
	BaseCmdTyper

	IsBackup() bool
	IsBackupSpecific() bool
}

type ImportExportBackupCmdTyper interface {
	BaseCmdTyper

	ImportCmdTyper
	ExportCmdTyper
	BackupCmdTyper
}

type HelpCmdTyper interface {
	BaseCmdTyper

	IsHelp() bool
}

type LogCmdTyper interface {
	BaseCmdTyper

	IsLog() bool
}

type StatusCmdTyper interface {
	BaseCmdTyper

	IsStatus() bool
}

type InstallCmdTyper interface {
	BaseCmdTyper

	IsInstall() bool
}

type ReinstallCmdTyper interface {
	BaseCmdTyper

	IsReinstall() bool
}

type CleanupCmdTyper interface {
	BaseCmdTyper

	IsCleanup() bool
}

type UninstallCmdTyper interface {
	BaseCmdTyper

	IsUninstall() bool
}

type ApplyInstallFixCmdTyper interface {
	BaseCmdTyper

	IsApplyInstallFix() bool
}

type RevertCmdTyper interface {
	BaseCmdTyper

	IsRevert() bool
}

type HistoriesCmdTyper interface {
	BaseCmdTyper

	IsHistories() bool
}

type CompressCmdTyper interface {
	BaseCmdTyper

	IsCompress() bool
}

type DecompressCmdTyper interface {
	BaseCmdTyper

	IsDecompress() bool
}

type DownloadCmdTyper interface {
	BaseCmdTyper

	IsDownload() bool
}

type DownloadDecompressCmdTyper interface {
	BaseCmdTyper

	IsDownloadDecompress() bool
}

type ChangePortCmdTyper interface {
	BaseCmdTyper

	IsChangePort() bool
}

type SwitchPortCmdTyper interface {
	BaseCmdTyper

	IsSwitchPort() bool
}

type SwitchOrChangePortCmdTyper interface {
	ChangePortCmdTyper
	SwitchPortCmdTyper
}

type WhichPortCmdTyper interface {
	BaseCmdTyper

	IsWhichPort() bool
}

type ChangeDirCmdTyper interface {
	BaseCmdTyper

	IsChangeDir() bool
}

type ApplyDefaultChmodCmdTyper interface {
	BaseCmdTyper

	IsApplyDefaultChmod() bool
}

type ApplyDefaultConfigureCmdTyper interface {
	BaseCmdTyper

	IsApplyDefaultConfig() bool
}

type SyncCmdTyper interface {
	BaseCmdTyper

	IsSync() bool
	IsSyncNow() bool
	IsSyncForce() bool
	IsInDirtyState() bool
	IsListDirtyStates() bool
	IsLastStates() bool
	IsDiscardDirtyStates() bool

	HistoriesCmdTyper
}
