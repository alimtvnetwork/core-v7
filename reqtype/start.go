package reqtype

import "gitlab.com/evatix-go/core/constants"

func start(
	reqs *[]Request,
) interface{} {
	if reqs == nil || len(*reqs) == 0 {
		return nil
	}

	return (*reqs)[constants.Zero]
}
