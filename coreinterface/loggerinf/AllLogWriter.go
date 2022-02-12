package loggerinf

import "io"

type AllLogWriter interface {
	InfoWriter() io.Writer
	ErrorWriter() io.Writer
	WarningWriter() io.Writer
	DebugWriter() io.Writer
	WriterBy(config Configurer) io.Writer
}
