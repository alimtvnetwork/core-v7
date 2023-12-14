package pathinternal

func JoinTemp(joiningPaths ...string) string {
	slice := make([]string, len(joiningPaths)+1)
	slice[0] = GetTemp()

	for i, curPath := range joiningPaths {
		slice[i+1] = curPath
	}

	return Join(slice...)
}
