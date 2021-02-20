package main

import (
	"fmt"
	"sort"
)

func main() {
	islice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Printf("before sort: %v\n", islice)
	sort.Ints(islice)
	fmt.Printf("after sort: %v\n", islice)

	sslice := []string{"k", "j", "i", "h", "g", "f", "e", "d", "c", "b", "a"}
	fmt.Printf("before sort: %v\n", sslice)
	sort.Strings(sslice)
	fmt.Printf("after sort: %v\n", sslice)
}

// sort 自带 int/string的排序方法。
// 其余的排序需要自定义排序。
// 自定义排序需要实现如下的接口:
//  type ISort interface {
//      Len() int
//      Less(i, j int) bool
//      Swap(i, j int)
//  }
