package main

import "fmt"

func main() {
	//IF
	foo := 1
	if foo == 1 {
		println("bar")
	} else {
		println("buz")
	}
	//SWITCH
	fooSwitch := 3
	switch {
	case fooSwitch == 1:
		println("one")
	case fooSwitch == 2:
		println("two")
	case fooSwitch == 3:
		println("three")
	default:
		println("no number...")
	}

	var score float64
	fmt.Println("Enter score for your last exam")
	fmt.Scanf("%v", &score)

	switch {
	case score <= 59:
		println("E")
	case score <= 69:
		println("D")
	case score <= 79:
		println("C")
	case score <= 89:
		println("B")
	case score <= 100:
		println("A")
	default:
		println("Please enter a score <= 100")
	}

	//break & continue
	//break
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("i nin değeri", i)
		i++
	}

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Output : ", i)
	}

}
