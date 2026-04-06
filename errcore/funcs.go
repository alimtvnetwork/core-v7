package errcore

type (
	ErrFunc         func() error
	ErrBytesFunc    func() (rawBytes []byte, err error)
	ErrStringsFunc  func() (lines []string, err error)
	ErrStringFunc   func() (line string, err error)
	ErrAnyFunc      func() (anyItem any, err error)
	ErrAnyItemsFunc func() (anyItems []any, err error)
	ErrInAnyFunc    func(anyItem any) (err error)
)

type TaskWithErrFunc = ErrFunc
