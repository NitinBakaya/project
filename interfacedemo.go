package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name        string
	hourlyPay   int
	hoursPerPay int
}

func (c contractor) getName() string {
	return c.name
}
func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerPay
}

type fullTime struct {
	name   string
	salary int
}

func (f fullTime) getName() string {
	return f.name
}
func (f fullTime) getSalary() int {
	return f.salary
}
func test(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("***********************************************")
}

func main() {
	test(fullTime{
		name:   "Nitin",
		salary: 150000,
	})
	test(contractor{
		name:        "Vijay",
		hourlyPay:   100,
		hoursPerPay: 9,
	})
	test(contractor{
		name:        "Raju",
		hourlyPay:   50,
		hoursPerPay: 10,
	})
}
