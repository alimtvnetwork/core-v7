package pathextendinf

import "os"

type Modifier interface {
	Create() error
	CreateUsingChmod(mode os.FileMode) error
	Remove() error
	RemoveAllOnExist() error
	ApplyChmod(mode os.FileMode) error
	ApplyChmodRecursively(mode os.FileMode) error
	ApplyChown(user, group string) error
}
