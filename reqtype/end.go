package reqtype

func end(
	reqs []Request,
) *Request {
	r := reqs[len(reqs)-1]
	return &r
}
