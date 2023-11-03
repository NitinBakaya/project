package main

import "fmt"

func main() {

	user, err := getUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	profile, err := getUserProfile(user.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	func divideByzero(x,y float64) (float64,error){
		if y == 0{
			return 0.0,error.New("no divided by zero")
		}
		return x / y, nil
	} 
}

func getUser() (user, error) {
	//user code
}
