package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}

// go run ex_1-1.go abc def hi hello
// out: /var/folders/wf/2n7mg5wx3b51w901c_9d8k_m0000gn/T/go-build197053665/b001/exe/ex_1-1
