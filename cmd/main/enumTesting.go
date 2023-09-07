package main

import (
	"fmt"
	"log"

	"gitlab.com/auk-go/core/bytetype"
	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/core/ostype"
)

func enumTesting() {
	names, _ := corecomparator.Equal.MarshalJSON()

	fmt.Println(string(names))

	noteq := corecomparator.NotEqual

	noteq.UnmarshalJSON([]byte("3"))

	fmt.Println(noteq.Name())
	fmt.Println(noteq.NumberString())
	fmt.Println(noteq.NumberString())
	fmt.Println(bytetype.BasicEnumImpl.AppendPrependJoinValue(".", 1, 2))

	linux := ostype.Linux
	anroid := ostype.Android
	marshalledLinux, _ := linux.MarshalJSON()

	fmt.Println(marshalledLinux)

	err := anroid.UnmarshalJSON(marshalledLinux)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(anroid)

	windowsGrp := ostype.WindowsGroup
	windowsBytes, _ := windowsGrp.MarshalJSON()
	unixGrp := ostype.UnixGroup
	unixGrp.UnmarshalJSON(windowsBytes)
	fmt.Print(unixGrp)
}
