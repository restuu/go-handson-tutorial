package main

import (
	"flag"
	"fmt"
)

func main() {
	var str string
	flag.StringVar(&str, "name", "World", "Input your name")

	flag.Parse()

	fmt.Printf("Hello, %s\n", str)

}
