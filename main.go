package main

import (
	"fmt"
	"os"
)

func doSomething(val float32) float32 {
	return val * 3 / 99
}


func main() {
	fmt.Println("Hello, World!")

	f, err := os.Open("/tmp/dat")

	if err != nil {
			panic(err))
	}

}
