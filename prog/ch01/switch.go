package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	switch_cond()
	switch_default()
	switch_no_cond()
	switch_with_fallthrough()
	switch_assign()
}

func switch_cond() {
	cond := "abc"
	switch cond {
	case "a":
		fmt.Printf("a is hit...\n")
	case "abc":
		fmt.Printf("abc is hit...\n")
	default:
		fmt.Printf("default is hit...\n")
	}
}

func switch_default() {
	cond := "abc"
	switch cond {
	default: // default语句可以在任何地方，但是最佳实践推荐在最后边。
		fmt.Printf("default is hit...\n")
	case "a":
		fmt.Printf("a is hit...\n")
	case "abc":
		fmt.Printf("abc is hit...\n")
	}
}

func switch_no_cond() {
	x := 0
	switch { // 无标签选择(tagless)，等价于 switch true
	case x > 1:
		fmt.Printf("x=%v > 1\n", x)
	case x == 1:
		fmt.Printf("x=%v == 1\n", x)
	case x < 1:
		fmt.Printf("x=%v < 1\n", x)
	}
}

func switch_with_fallthrough() {
	x := 0
	switch {
	case x > 1:
		fmt.Printf("x=%v > 1\n", x)
	case x == 1:
		fmt.Printf("x=%v == 1\n", x)
	case x < 1:
		fallthrough // 跳过当前case到下一个case.
	case x <= 0:
		fmt.Printf("x=%v <= 0\n", x)
	}
}

func switch_assign() {
	rand.Seed(time.Now().UnixNano())
	switch x := rand.Int() % 10; x {
	case 0:
		fmt.Printf("x rand 0\n")
	case 1:
		fmt.Printf("x rand 1\n")
	case 2:
		fmt.Printf("x rand 2\n")
	case 3:
		fmt.Printf("x rand 3\n")
	case 4:
		fmt.Printf("x rand 4\n")
	case 5:
		fmt.Printf("x rand 5\n")
	case 6:
		fmt.Printf("x rand 6\n")
	case 7:
		fmt.Printf("x rand 7\n")
	case 8:
		fmt.Printf("x rand 8\n")
	case 9:
		fmt.Printf("x rand 9\n")
	}
}



