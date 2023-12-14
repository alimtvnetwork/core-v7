package pathinternal

import (
	"os"
)

var temp = Clean(os.TempDir())

func GetTemp() string {
	return temp
}
