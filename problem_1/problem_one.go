package main

import (
	"fmt"
)

func fillColors(slice1, slice2 []string) []string {
	uni := make(map[string]bool)
	for i := 0; i < len(slice1); i++ {
		uni[slice1[i]] = true
	}
	for i := 0; i < len(slice2); i++ {
		uni[slice2[i]] = true
	}
	var res []string
	for k, _ := range uni {
		res = append(res, k)
	}
	return res
}

func main() {
	slice1 := []string{"Red", "Black", "White"}
	slice2 := []string{"Black", "Yellow", "Orange"}
	res := fillColors(slice1, slice2)
	fmt.Print(res)
}
