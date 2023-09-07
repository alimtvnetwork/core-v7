package enuminf

import "gitlab.com/auk-go/core/internal/internalinterface/internalenuminf"

type BaseCmdTyper interface {
	BasicEnumer
	internalenuminf.BaseCmdTyper
}

type ToNameLower interface {
	internalenuminf.ToNameLower
}

type ListCmdTyper interface {
	BasicEnumer
	internalenuminf.ListCmdTyper
}

type CreateCmdTyper interface {
	BasicEnumer
	internalenuminf.CreateCmdTyper
}

type AddCmdTyper interface {
	BasicEnumer
	internalenuminf.AddCmdTyper
}

type RemoveCmdTyper interface {
	BasicEnumer
	internalenuminf.RemoveCmdTyper
}

type ImportCmdTyper interface {
	BasicEnumer
	internalenuminf.ImportCmdTyper
}

type ExportCmdTyper interface {
	BasicEnumer
	internalenuminf.ExportCmdTyper
}

type BackupCmdTyper interface {
	BasicEnumer
	internalenuminf.BackupCmdTyper
}

type ImportExportBackupCmdTyper interface {
	BasicEnumer
	internalenuminf.ImportExportBackupCmdTyper
}

type HelpCmdTyper interface {
	BasicEnumer
	internalenuminf.HelpCmdTyper
}

type LogCmdTyper interface {
	BasicEnumer
	internalenuminf.LogCmdTyper
}

type StatusCmdTyper interface {
	BasicEnumer
	internalenuminf.StatusCmdTyper
}

type InstallCmdTyper interface {
	BasicEnumer
	internalenuminf.InstallCmdTyper
}

type ReinstallCmdTyper interface {
	BasicEnumer
	internalenuminf.ReinstallCmdTyper
}

type CleanupCmdTyper interface {
	BasicEnumer
	internalenuminf.CleanupCmdTyper
}

type UninstallCmdTyper interface {
	BasicEnumer
	internalenuminf.CleanupCmdTyper
}

type ApplyInstallFixCmdTyper interface {
	BasicEnumer
	internalenuminf.ApplyInstallFixCmdTyper
}

type RevertCmdTyper interface {
	BasicEnumer
	internalenuminf.RevertCmdTyper
}

type HistoriesCmdTyper interface {
	BasicEnumer
	internalenuminf.HistoriesCmdTyper
}

type CompressCmdTyper interface {
	BasicEnumer
	internalenuminf.CompressCmdTyper
}

type DecompressCmdTyper interface {
	BasicEnumer
	internalenuminf.DecompressCmdTyper
}

type DownloadCmdTyper interface {
	BasicEnumer
	internalenuminf.DownloadCmdTyper
}

type DownloadDecompressCmdTyper interface {
	BasicEnumer
	internalenuminf.DownloadDecompressCmdTyper
}

type ChangePortCmdTyper interface {
	BasicEnumer
	internalenuminf.ChangePortCmdTyper
}

type SwitchPortCmdTyper interface {
	BasicEnumer
	internalenuminf.SwitchPortCmdTyper
}

type SwitchOrChangePortCmdTyper interface {
	ChangePortCmdTyper
	SwitchPortCmdTyper
}

type WhichPortCmdTyper interface {
	BasicEnumer
	internalenuminf.WhichPortCmdTyper
}

type ChangeDirCmdTyper interface {
	BasicEnumer
	internalenuminf.ChangeDirCmdTyper
}

type ApplyDefaultChmodCmdTyper interface {
	BasicEnumer
	internalenuminf.ApplyDefaultChmodCmdTyper
}

type ApplyDefaultConfigureCmdTyper interface {
	BasicEnumer
	internalenuminf.ApplyDefaultConfigureCmdTyper
}

type SyncCmdTyper interface {
	BasicEnumer
	internalenuminf.SyncCmdTyper
}
