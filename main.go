package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("hello world")
	args := os.Args[1:]
	fmt.Printf(strings.Join(args, ", "))
}
