package main

import (
	"fmt"
)

func main() {
	var x interface{} = 3

	/* 変数xはinterface{}型 */
	if x == nil {
		fmt.Println("x is nil")
	} else if i, isInt := x.(int); isInt {
		fmt.Printf("x is integer : %d\n", i)
	} else if s, isString := x.(string); isString {
		fmt.Println(s)
	} else {
		fmt.Println("unsupported type")
	}

}
