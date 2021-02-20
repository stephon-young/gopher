package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	l := 10000

	a := make([]string, 0, l)
	for i := 0; i < l; i++ {
		a = append(a, " ")
	}

	o1 := ""
	beg := time.Now()
	for i := 0; i < l; i++ {
		o1 += a[i] + " "
	}
	end := time.Now()
	t1 := end.UnixNano() - beg.UnixNano()

	beg = time.Now()
	_ = strings.Join(a, " ")
	end = time.Now()
	t2 := end.UnixNano() - beg.UnixNano()

	fmt.Printf("t1: %v, t2: %v\n", t1, t2)
}

// run: go run ex_1-3.go abc def hi hello // l := 1000
// out: t1: 674000, t2: 12000
// run: go run ex_1-3.go abc def hi hello // l := 10000
// out: t1: 20786000, t2: 120000