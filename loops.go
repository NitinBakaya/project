package main

import "fmt"

func main() {

	no := 1
	for no < 5 {
		no += 2
		no++
	}
	fmt.Println("The new number is", no)

	Test()
	TestBreak()

}
func Test() {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
func TestBreak() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}
