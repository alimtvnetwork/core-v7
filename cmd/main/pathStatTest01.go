package main

import (
	"fmt"

	"gitlab.com/auk-go/core/chmodhelper"
)

func pathStatTest01() {
	files := currentPackageAllFiles()
	first10 := files[:10]
	first11 := append(first10, "/something-non-exist")

	items := chmodhelper.GetPathExistStats(false, first11...)

	fmt.Println(items)
}

func pathStatTest02() {
	files := currentPackageAllFiles()
	first10 := files[:10]
	first11 := append(first10, "/something-non-exist")

	items := chmodhelper.GetPathExistStats(false, first11...)

	fmt.Println(items[0].NotExistError())
	fmt.Println(items[0].NotADirError())
	fmt.Println(items[0].NotAFileError())

	fmt.Println("----------------")
	last := len(first11) - 1
	fmt.Println(items[last].NotExistError())
	fmt.Println(items[last].NotADirError())
	fmt.Println(items[last].NotAFileError())

}
