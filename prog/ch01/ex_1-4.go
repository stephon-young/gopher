package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := make(map[string]int)

	for _, filename := range os.Args[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex_1-4: %v\n", err)
			continue
		}

		c := make(map[string]int)
		input := bufio.NewScanner(f)
		for input.Scan() {
			c[input.Text()]++
			counts[input.Text()]++
		}

		for _, n := range c {
			if n > 1 {
				files[filename] = 1
				break
			}
		}
	}

	for k, _ := range files {
		fmt.Printf("%s\n", k)
	}
}
