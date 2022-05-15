package utils

import (
	"fmt"
)

func Contains(arr []string, val interface{}) bool {
	idx := 0
	for idx < len(arr) {
		fmt.Printf("%d", idx)
		if arr[idx] == val {
			return true
		}
		idx += 1
	}
	return false
}
