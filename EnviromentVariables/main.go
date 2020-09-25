package main

import (
	"fmt"
	"os"
)

func main() {

	uName := os.Getenv("UserName")
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")

	fmt.Println("UserName : " + uName)
	fmt.Println("GOROOT : " + goRoot)
	fmt.Println("GOPATH : " + goPath)

	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }
}
