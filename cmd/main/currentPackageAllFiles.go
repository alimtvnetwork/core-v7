package main

import (
	"fmt"
	"log"
	"path/filepath"

	"gitlab.com/auk-go/core/chmodhelper"
)

func currentPackageAllFiles() []string {
	root, _ := filepath.Abs(".")
	fmt.Println(root)

	slice, err := chmodhelper.GetRecursivePathsContinueOnError(root)

	if err != nil {
		log.Fatalln(err)
	}

	return slice
}
