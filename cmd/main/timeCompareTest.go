package main

import (
	"fmt"
	"time"

	"gitlab.com/evatix-go/core/corecomparator"
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
	fmt.Println("time.Now, now = ", corecomparator.Time(time.Now(), now))
	fmt.Println("now, now = ", corecomparator.Time(now, now))
	fmt.Println("now, now2 = ", corecomparator.Time(now, now2))
	fmt.Println("now2, now2 = ", corecomparator.Time(now2, now2))
	fmt.Println("now2, now = ", corecomparator.Time(now2, now))
	fmt.Println("now2, time.now = ", corecomparator.Time(time.Now().Add(600001), now2))
}
