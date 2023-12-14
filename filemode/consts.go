package filemode

import "os"

// When used 0, golang treats it as octal not decimal anymore.
// Reference : https://ss64.com/bash/chmod.html
//
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
	X100                                         os.FileMode = 0100
	X200                                         os.FileMode = 0200
	X300                                         os.FileMode = 0300
	X400                                         os.FileMode = 0400
	X500                                         os.FileMode = 0500
	X600                                         os.FileMode = 0600
	X700                                         os.FileMode = 0700
	X111                                         os.FileMode = 0111
	X222                                         os.FileMode = 0222
	X333                                         os.FileMode = 0333
	X444                                         os.FileMode = 0444
	X455                                         os.FileMode = 0455
	X466                                         os.FileMode = 0466
	X555                                         os.FileMode = 0555
	X644                                         os.FileMode = 0644
	X655                                         os.FileMode = 0655
	X666                                         os.FileMode = 0666
	X677                                         os.FileMode = 0677
	X711                                         os.FileMode = 0711
	X722                                         os.FileMode = 0722
	X744                                         os.FileMode = 0744
	X755                                         os.FileMode = 0755
	X766                                         os.FileMode = 0766
	X777                                         os.FileMode = 0777
	FileDefault                                  os.FileMode = 0644 // cannot execute by everyone OwnerCanReadWriteGroupOtherCanReadOnly
	DirDefault                                   os.FileMode = 0755 // can execute by everyone OwnerCanDoAllExecuteGroupOtherCanReadExecute
	OwnerFullAccessOnly                          os.FileMode = 0700
	OwnerGroupFullAccessOnly                     os.FileMode = 0770
	CacheFullAccess                              os.FileMode = 0777
	FullAccess                                   os.FileMode = 0777
)
