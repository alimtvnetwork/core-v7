package pathinternal

import "os"

func IsPathExists(location string) bool {
	_, err := os.Stat(location)

	return err == nil || !os.IsNotExist(err)
}
