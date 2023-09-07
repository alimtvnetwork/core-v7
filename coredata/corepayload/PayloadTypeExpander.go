package corepayload

import "fmt"

type PayloadTypeExpander struct {
	CategoryStringer fmt.Stringer
	TaskTypeStringer fmt.Stringer
}
