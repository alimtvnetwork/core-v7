package filemodes

import "os"

// When used 0, golang treats it as octal not decimal anymore.
//goland:noinspection ALL
const (
	AllPermission                                os.FileMode = 0777
	AllExecute                                   os.FileMode = 0111
	AllReadExecute                               os.FileMode = 0555
	AllRead                                      os.FileMode = 0444
	AllWrite                                     os.FileMode = 0222
	AllWriteExecute                              os.FileMode = 0333
	OwnerCanDoAllExecuteGroupOtherCanReadExecute os.FileMode = 0755
	OwnerCanReadWriteGroupOtherCanReadOnly       os.FileMode = 0644
	OwnerCanDoAllGroupOtherCanReadOnly           os.FileMode = 0744
	OwnerCanDoAllGroupOtherCanReadWriteOnly      os.FileMode = 0766
	OwnerCanDoAllGroupOtherCanExecuteOnly        os.FileMode = 0711
	OwnerCanDoAllGroupOtherCanReadExecuteOnly    os.FileMode = 0755
	OwnerCanDoAllGroupOtherCanWriteOnly          os.FileMode = 0722
)
