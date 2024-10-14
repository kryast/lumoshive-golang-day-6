package main

import "fmt"

var a int
var b = 1

var (
	name string = "Lumoshive"
	age  int    = 45
)

const e float32 = 3.14

var f string

func main() {
	c := "Lumoshive"
	d := 2.56

	fmt.Println(a, b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(f)
}
