package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/core/coreversion"
	"gitlab.com/evatix-go/core/enums/versionindexes"
)

func versionCompareTest(leftVersion, rightVersion string) corecomparator.Compare {
	fmt.Println("left, right = ", leftVersion, rightVersion)
	leftV := coreversion.New(leftVersion)
	rightV := coreversion.New(rightVersion)

	fmt.Println("   left, right = ", leftV, rightV)
	r1 := leftV.Compare(rightV)
	r2 := leftV.ComparisonValueIndexes(
		rightV,
		versionindexes.AllVersionIndexes...)
	leftVersionValues := leftV.AllVersionValues()
	rightVersionValues := rightV.AllVersionValues()

	fmt.Println("   (r1) left, right = ", r1)
	fmt.Println("   (r2) left, right = ", r2)
	fmt.Println("   (Values) left, right = ", leftVersionValues, rightVersionValues)

	r3 := corecomparator.VersionSliceInteger(
		leftVersionValues,
		rightVersionValues)

	fmt.Println("   (r3) left, right = ", r3)

	return r1
}
