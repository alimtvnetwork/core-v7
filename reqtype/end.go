package reqtype

func end(
	reqs *[]Request,
) interface{} {
	if reqs == nil || len(*reqs) == 0 {
		return nil
	}

	return (*reqs)[len(*reqs)-1]
}
