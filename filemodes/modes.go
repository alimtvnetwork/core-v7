package filemodes

import "os"

// When used 0, golang treats it as octal not decimal anymore.
//goland:noinspection ALL
const (
	AllPermission                 os.FileMode = 0777
	AllExecute                    os.FileMode = 0111
	AllReadExecute                os.FileMode = 0555
	AllWriteExecute               os.FileMode = 0333
	OwnerAllReadExecuteGroupOther os.FileMode = 0755
)
