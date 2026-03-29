package reqtype

import "github.com/alimtvnetwork/core/constants"

func start(
	reqs []Request,
) *Request {
	r := reqs[constants.Zero]
	return &r
}
