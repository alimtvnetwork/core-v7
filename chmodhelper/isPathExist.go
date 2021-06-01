package chmodhelper

import "os"

func isPathExist(location string) bool {
	_, err := os.Stat(location)

	return !os.IsNotExist(err)
}
