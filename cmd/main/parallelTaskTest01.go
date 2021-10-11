package main

import (
	"fmt"
	"time"

	"gitlab.com/evatix-go/core/coreasync"
)

func parallelTaskTest01() {
	coreasync.ParallelTasksWait(
		func() {
			fmt.Println("started 1")
			go func() {
				for i := 0; i < 20; i++ {
					time.Sleep(100 * time.Millisecond)
					fmt.Println("Fun1 () - ", i)
				}
			}()
			time.Sleep(5000 * time.Millisecond)
			fmt.Println("done 1")
		},
		func() {
			fmt.Println("started 2")
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("done 2")
		})
}
