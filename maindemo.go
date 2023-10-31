package main

import "fmt"

func main() {
	type car struct {
		model string
		make  string
	}
	type truck struct {
		bedsize int
		car
	}
	truckOne := truck{
		bedsize: 10,
		car: car{
			make:  "bmw",
			model: "z1",
		},
	}
	fmt.Println(truckOne.bedsize)
	fmt.Println(truckOne.make)
	fmt.Println(truckOne.model)
}
