package chmodhelper

type simpleFileReaderWriterModel struct {
	ChmodDir, ChmodFile RwxWrapper
	ParentDir           string // full path to the parent dir
	FilePath            string // full path to the actual file to write to or read from.
}

type rwxWrapperModel struct {
	Chmod, // "0777"
	RwxFull string // "-rwxrwxrwx"
}
