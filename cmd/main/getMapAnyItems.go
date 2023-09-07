package main

import (
	"fmt"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coredata/corestr"
)

func getMapAnyItems() *coredynamic.MapAnyItems {
	fmt.Println("MapAnyItems")
	mapAnyItems := coredynamic.NewMapAnyItems(200)
	collection := corestr.New.Collection.Cap(100)
	collection.Adds("alim-1", "alim-2", "alim-3", "alim-4")
	mapAnyItems.Add("alim-something", collection)
	mapAnyItems.Add("alim-something2", collection)
	mapAnyItems.Add("alim-something3", collection.ConcatNew(1, "alim 5"))
	mapAnyItems.Add("alim-something4", collection)
	mapAnyItems.Add("alim-something5", collection)
	mapAnyItems.Add("alim-something6", collection)
	mapAnyItems.Add("alim-something7", collection)
	mapAnyItems.Add("alim-something8", collection)
	mapAnyItems.Add("alim-something9", collection)

	return mapAnyItems
}
