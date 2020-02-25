package main

import (
	"fmt"
	"runtime"
)

func main() {
	var x interface{} = "mo"

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
L:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j > 1 {
				continue L

			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("ここは処理されない")
	}

	go fmt.Println("hahaha")

	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("goVersion: %s\n", runtime.Version())

	s1 := make([]int, 5)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s2 := make([]int, 5, 10)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

}
