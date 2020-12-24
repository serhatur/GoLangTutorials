package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonStr := `
		{
			"data": {
				"first_name":"serhat",
				"last_name": "ür"
			}
		}
	`

	var m map[string]map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}

	fmt.Println(m)

	fmt.Println("*******************")

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
