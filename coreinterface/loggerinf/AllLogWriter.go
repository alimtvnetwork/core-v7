package loggerinf

import "io"

type AllLogWriter interface {
	SuccessWriter() io.Writer
	InfoWriter() io.Writer
	ErrorWriter() io.Writer
	WarningWriter() io.Writer
	DebugWriter() io.Writer
	TraceWriter() io.Writer
	WriterBy(config Configurer) io.Writer
}
