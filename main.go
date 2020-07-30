package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Starting: ", time.Now().Format(time.Stamp))

	for c := 0; c < 20; c++ {

		fmt.Printf("Loop: %v - %v \n", c, time.Now().Format(time.Stamp))
		time.Sleep(1 * time.Second)
	}

	fmt.Println("End: ", time.Now().Format(time.Stamp))

}
