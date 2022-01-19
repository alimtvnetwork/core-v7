package reflectinternal

func IsNotNull(item interface{}) bool {
	return !IsNull(item)
}
