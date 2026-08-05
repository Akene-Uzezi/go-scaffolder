package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	testArg := flag.String("name", "", "Test flag")

	if *testArg == "" {
		fmt.Println("place test arg")
		os.Exit(1)
	}
}
