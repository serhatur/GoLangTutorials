package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Kullanıcı veri oluşturma alanı v1
	fmt.Println("Kullanıcı oluşturma v1")
	user1 := &User{
		ID:        1,
		FirstName: "Serhat",
		LastName:  "ÜR",
		UserName:  "SerhatUR",
		Age:       31,
		Pay: &Payment{
			Salary: 3555,
			Bonus:  350,
		},
	}

	fmt.Println(user1)
	fmt.Println(user1.GetFullName())
	fmt.Println(user1.GetUsername())
	fmt.Println("Maaş : " + strconv.FormatFloat(user1.GetPayment(), 'g', -1, 64))

	user2 := NewUser()
	user2.FirstName = "Sibel"
	user2.LastName = "ÜR"
	user2.Age = 31
	user2.UserName = "sibosssss"

	user2.Pay = &Payment{
		Salary: 755,
		Bonus:  955,
	}
	fmt.Println(user2.GetFullName())
	fmt.Println(user2.GetUsername())
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

//Kullanıcı yapıcı metotu
func NewUser() *User {
	u := new(User)
	u.Pay = NewPayment()
	return u
}

type Payment struct {
	Salary float64
	Bonus  float64
}

//Ödemenin yapıcı metotu
func NewPayment() *Payment {
	p := new(Payment)
	return p
}

//Metotlar
func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) GetUsername() string {
	return u.UserName
}

func (u *User) GetPayment() float64 {
	pay := u.Pay.Salary + u.Pay.Bonus
	return pay
}
