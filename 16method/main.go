package main

import "fmt"

type User struct {
	Name    string
	Age     int
	Address string
	City    string
	State   string
	Email   string
}

func main() {

	user := User{"Varun", 25, "Lucknow", "Lucknow", "UP", "varunpratap774@gmailcom"}
	user.GetUserDetails()
	user.ResetMail("mcs20013@iiitl.ac.in")
	user.GetUserDetails()

}

func (u User) GetUserDetails() {

	fmt.Println("The following are the details of the user")
	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
	fmt.Println("Address:", u.Address)
	fmt.Println("City:", u.City)
	fmt.Println("State:", u.State)
	fmt.Println("Email:", u.Email)

}

func (u *User) ResetMail(newMail string) {
	u.Email = newMail
}
