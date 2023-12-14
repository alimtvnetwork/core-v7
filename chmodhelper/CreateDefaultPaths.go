package chmodhelper

func CreateDirFilesWithRwxPermissions(
	isRemoveAllDirBeforeCreate bool,
	dirFilesWithRwxPermissions []DirFilesWithRwxPermission,
) error {
	if dirFilesWithRwxPermissions == nil || len(dirFilesWithRwxPermissions) == 0 {
		return nil
	}

	for _, pathCreate := range dirFilesWithRwxPermissions {
		err2 := CreateDirFilesWithRwxPermission(
			isRemoveAllDirBeforeCreate,
			&pathCreate)

		if err2 != nil {
			return err2
		}
	}

	return nil
}
