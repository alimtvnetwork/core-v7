package reqtype

type ResultStatus struct {
	IsSuccess  bool
	IndexMatch int
	Ranges     []Request
	Error      error
}
