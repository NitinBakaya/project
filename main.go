package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func test(m messageToSend) {
	fmt.Printf("message: %v sent to: %v\n", m.message, m.phoneNumber)
	fmt.Println("================================================")
}
func main() {

	test(messageToSend{
		phoneNumber: 4323495738,
		message:     "Hello, How r u",
	})
	test(messageToSend{
		phoneNumber: 7323395938,
		message:     "Hello, meet me at bar",
	})
}
