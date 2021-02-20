package main

import (
	"fmt"
)

func change(reversed map[string]string, data []map[string]interface{}) []map[string]interface{} {
	for i := 0; i < len(data); i++ {
		for k, _ := range data[i] {
			if vv, exist := reversed[k]; exist {
				data[i][vv] = data[i][k]
				delete(data[i], k)
			}
		}
	}

	return data
}

func main() {
	data := make([]map[string]interface{}, 0)
	for i := 0; i < 10; i++ {
		d := make(map[string]interface{}, 0)
		d["a"] = i
		d["b"] = fmt.Sprintf("%d", i)
		data = append(data, d)
	}

	fmt.Printf("before: %v\n", data)

	reversed := make(map[string]string, 0)
	reversed["a"] = "aa"
	//ßreversed["b"] = "bb"

	data2 := change(reversed, data)

	fmt.Printf("after: %v\n", data2)
}
