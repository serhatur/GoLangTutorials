package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Value : ", i)
	}

	//for (while yöntemi)
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}
