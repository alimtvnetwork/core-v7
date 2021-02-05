package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corestr"
)

func main() {
	// fmt.Println(osconsts.IsWindows)
	// fmt.Println(osconsts.IsUnixGroup)

	items := []string{"alim items collect 01", "alim items collect 02", "alim items collect 03"}
	items1 := []string{"alim 01"}
	items2 := []string{"alim 02"}
	items3 := []string{"alim 03"}
	items4 := []string{"alim 04"}
	items5 := []string{"alim 05"}

	collectOfCollections := corestr.NewCollectionsOfCollectionUsingStringsOfStrings(
		true,
		&items1,
		&items2,
		&items3,
		&items4,
		&items5,
	)

	collection := corestr.NewCollectionUsingStrings(&items, false)
	linkedCollection := corestr.NewLinkedCollectionsUsingCollections(collection)

	linkedCollection.AddCollection(collectOfCollections.Items()[1])
	linkedCollection.AddCollection(collectOfCollections.Items()[2])
	linkedCollection.AddCollection(collectOfCollections.Items()[3])
	linkedCollection.AddCollection(collectOfCollections.Items()[4])

	linked2 := corestr.NewLinkedCollections()
	linked2.AddStringsPtr(&items, true)
	linked2.AddCollection(collectOfCollections.Items()[1])
	linked2.AddCollection(collectOfCollections.Items()[2])
	linked2.AddCollection(collectOfCollections.Items()[3])
	linked2.AddCollection(collectOfCollections.Items()[4])

	linked2.RemoveNodeByIndex(3, 4)

	fmt.Println(linkedCollection.GetCompareSummary(linked2, "Link1", "Link2"))

	// linkedCollection.RemoveNodeByIndex(0)
	fmt.Println(linkedCollection.GetCompareSummary(linked2, "Link1", "Link2"))

}
