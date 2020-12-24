package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Convert
	var myString = "1"
	var x = 10
	var f float32 = 2.0

	fmt.Println(myString, x, f)

	number, _ := strconv.Atoi(myString)
	fmt.Println(number)
	result := number + 2
	fmt.Println(result)

	//Casting
	var i int = 55
	var fNum float64 = float64(i)
	var u uint = uint(fNum)

	fmt.Println(i, f, u)

}
