package reqtype

import "gitlab.com/auk-go/core/constants"

func start(
	reqs []Request,
) interface{} {
	if len(reqs) == 0 {
		return nil
	}

	return (reqs)[constants.Zero]
}
