package main

import "fmt"

func pow(a [3]int) {
	/* 配列の各要素を２乗する */
	for i, v := range a {
		a[i] = v * v
	}
}

func pow_slice(a []int) {
	/* スライスの各要素を２乗する */
	for i, v := range a {
		a[i] = v * v
	}
}

func main() {
	/* ３要素の配列 */
	a := [3]int{1, 2, 3}
	pow(a)
	/* ３要素のスライス */
	b := []int{1, 2, 3}
	pow_slice(b)
	fmt.Println(a)
	fmt.Println(b)
}
