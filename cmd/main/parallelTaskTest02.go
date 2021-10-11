package main

import (
	"errors"
	"fmt"
	"time"

	"gitlab.com/evatix-go/core/coreasync"
)

func parallelTaskTest02() {
	finalErr := coreasync.ParallelTaskWithErrorFunctionsWaited(
		false,
		true,
		0,
		func() error {
			fmt.Println("started 1")
			go func() {
				for i := 0; i < 20; i++ {
					time.Sleep(10 * time.Millisecond)
					fmt.Println("Fun1 () - ", i)
				}
			}()
			time.Sleep(3000 * time.Millisecond)
			fmt.Println("done 1 with error")

			return errors.New("something went wrong")
		},
		func() error {
			fmt.Println("started 2")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("done 2")

			return nil
		})

	fmt.Println("Final Err")
	fmt.Println(finalErr)
}
