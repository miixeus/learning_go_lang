package main

import (
	"fmt"
	"time"
)

func main () {
	i := 0

	for i < 10 {
		time.Sleep(time.Second)
		fmt.Println("Incrementando i")
	}

		fmt.Println(i)
		 
}