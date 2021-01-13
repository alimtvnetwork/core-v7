package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/osconsts"
)

func main() {
	fmt.Println(osconsts.IsWindows)
	fmt.Println(osconsts.IsUnixGroup)
}
