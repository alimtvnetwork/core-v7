package chmodhelper

type simpleFileWriter struct {
	CreateDir  dirCreator
	FileWriter fileWriter
}

func (it simpleFileWriter) Lock() {
	globalMutex.Lock()
}

func (it simpleFileWriter) Unlock() {
	globalMutex.Unlock()
}
