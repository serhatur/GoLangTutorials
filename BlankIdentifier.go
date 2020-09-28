package main

import "fmt"

func main() {

	_, veri := metot()

	fmt.Println(veri)

	for _, value := range veri {
		fmt.Println(value)
	}

	var _, x, _ int = 0, 9, 0
	fmt.Println(x)
}

func metot() {
	return ""
}
