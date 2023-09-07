package pathextendinf

type Cloner interface {
	Copy() PathExtender
	CloneNew(newFileName string) PathExtender
	CloneRoot(rootPath string) PathExtender
	CloneParentDir(rootPath string) PathExtender
}
