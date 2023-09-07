package main

import (
	"fmt"
	"time"

	"gitlab.com/auk-go/core/corecmp"
)

func timeCompareTest() {
	now := time.Now()
	now21 := time.Now()
	now2 := now21.Add(time.Duration(600000))
	fmt.Println(now)
	fmt.Println(now2)
	// fmt.Println(corecomparator.Byte(1,2))
	// fmt.Println(corecomparator.Byte(2,2))
	// fmt.Println(corecomparator.Byte(3,2))
	// fmt.Println(corecomparator.Byte(3,3))
	fmt.Println("time.Now, now = ", corecmp.Time(time.Now(), now))
	fmt.Println("now, now = ", corecmp.Time(now, now))
	fmt.Println("now, now2 = ", corecmp.Time(now, now2))
	fmt.Println("now2, now2 = ", corecmp.Time(now2, now2))
	fmt.Println("now2, now = ", corecmp.Time(now2, now))
	fmt.Println("now2, time.now = ", corecmp.Time(time.Now().Add(600001), now2))
}
