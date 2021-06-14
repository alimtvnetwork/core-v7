package reqtype

type Status struct {
	IsSuccess  bool
	IndexMatch int
	Ranges     []Request
	Error      error
}
