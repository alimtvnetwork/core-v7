package chmodhelper

func CreateDirFilesWithRwxPermissionsMust(
	isRemoveAllDirBeforeCreate bool,
	dirFilesWithRwxPermissions []DirFilesWithRwxPermission,
) {
	createErr := CreateDirFilesWithRwxPermissions(
		isRemoveAllDirBeforeCreate,
		dirFilesWithRwxPermissions)

	if createErr != nil {
		panic(createErr)
	}
}
