package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}

// run: go run echo3.go abc def hi hello
// out: abc def hi hello\n[abc def hi hello]