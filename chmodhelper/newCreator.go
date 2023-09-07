package chmodhelper

type newCreator struct {
	RwxWrapper             newRwxWrapperCreator
	SimpleFileReaderWriter newSimpleFileReaderWriterCreator
	Attribute              newAttributeCreator
}
