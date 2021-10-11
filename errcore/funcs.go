package errcore

type (
	ErrFunc         func() error
	TaskWithErrFunc func() error
	ErrBytesFunc    func() (rawBytes []byte, err error)
	ErrStringsFunc  func() (lines []string, err error)
	ErrStringFunc   func() (line string, err error)
	ErrAnyFunc      func() (anyItem interface{}, err error)
	ErrAnyItemsFunc func() (anyItems []interface{}, err error)
	ErrInAnyFunc    func(anyItem interface{}) (err error)
)
