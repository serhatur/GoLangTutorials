package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

//constructor kullanımı
func NewHuman() *Human {
	h := new(Human)
	return h
}

func NewHumanParams(firstName, lastName string, age int) *Human {
	h := new(Human)
	h.FirstName = firstName
	h.LastName = lastName
	h.Age = age
	return h
}

func main() {

	//v1
	x := Human{FirstName: "Serhat"}
	fmt.Println(x.FirstName)

	//v2
	x2 := new(Human)
	x2.FirstName = "Serhat ÜR"
	fmt.Println(x2.FirstName)

	//Yapıcı metot kullanımı
	x3 := NewHuman()
	x3.Age = 31
	fmt.Println(x3.Age)

	// Parametreli Yapıcı metot kullanımı
	x4 := NewHumanParams("Muhammed", "Ali", 31)
	fmt.Println(x4.FirstName)

	//veri okuma
	var data = x4.FirstName + " " + x4.LastName + " / " + strconv.Itoa(x4.Age)
	fmt.Println(data)
}
