package main

import "fmt"

const thy = "TurkishAirLines"
const tai = "TAI"

func main() {
	fmt.Println(thy,tai)

	fmt.Println(FACEBOOK)
}

type Brand string

const(
	FACEBOOK Brand = "FB"
	INSTAGRAM Brand = "Insta"
)