package main

import (
	"fmt"
)

func inc(p *int) {
	*p++
}

func main() {
	var i int
	p := &i
	i = 5
	fmt.Println(*p)

	inc(&i)
	inc(&i)
	inc(&i)
	inc(&i)
	inc(&i)

	fmt.Println(i)
	fmt.Println(p)
}
