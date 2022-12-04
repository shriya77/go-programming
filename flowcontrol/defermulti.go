package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
