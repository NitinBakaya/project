package main

import "fmt"

func main() {

	r := rect{
		leng:  10,
		width: 20,
	}
	fmt.Println(r.area())
}
func (r rect) area() int {
	return r.leng * r.width
}

type rect struct {
	leng  int
	width int
}
