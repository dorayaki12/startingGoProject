package main

import "fmt"

type Stringify interface {
	ToString() string
}

/* 構造体Personの定義 */
type Person struct {
	Name string
	Age  int
}

/* 構造体Carの定義 */
type Car struct {
	Number string
	Model  string
}

func (c *Car) ToStoring() string {
	return fmt.Sprintf("[%s] %s", c.Number, c.Model)
}

/*  */
func main() {

}
