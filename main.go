package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	for {
		f, err := os.Open("./presentation.md")
		if err != nil {
			panic(err)
		}

		defer f.Close()

		fmt.Println("EJCP")
	}
}
