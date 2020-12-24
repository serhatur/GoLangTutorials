package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("dosyam.txt")
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("ERROR", err.Error())
		os.Exit(1)
	}
}
