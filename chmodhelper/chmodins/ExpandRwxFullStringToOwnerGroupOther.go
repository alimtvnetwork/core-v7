package chmodins

func ExpandRwxFullStringToOwnerGroupOther(rwxFullString string) (*RwxOwnerGroupOther, error) {
	err := GetRwxFullLengthError(rwxFullString)

	if err != nil {
		return nil, err
	}

	owner := rwxFullString[1:4]
	group := rwxFullString[4:7]
	other := rwxFullString[7:10]

	return NewRwxOwnerGroupOther(
		owner,
		group,
		other), nil
}
