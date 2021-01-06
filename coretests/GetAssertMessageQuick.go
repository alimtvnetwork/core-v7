package coretests

import "fmt"

func GetAssertMessageQuick(
	when,
	actual,
	expected interface{},
	counter int,
) string {
	return fmt.Sprintf("----------------------\n%d )\tWhen:%#v\n\t\tActual:%#v , Expected:%#v",
		counter,
		when,
		actual,
		expected,
	)
}
