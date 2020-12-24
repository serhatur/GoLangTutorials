package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2016, time.November, 10, 20, 0, 0, 0, time.UTC)

	now := time.Now()

	fmt.Println("now : ", now)
	fmt.Println(now.AddDate(1, 2, 5))
	fmt.Println(t)

	//customformat
	shortFormat := "1/2/0606"
	fmt.Println(now.Format(shortFormat))
}
