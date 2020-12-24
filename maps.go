package main

import "fmt"

func main() {

	// 1
	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[43] = "Foo"
	myMap[12] = "Bar"
	fmt.Println(myMap)

	// 2
	states := make(map[string]string)
	states["IST"] = "İstanbul"
	states["ANK"] = "Ankara"
	states["ANT"] = "Antalya"
	fmt.Println(states)

	//şehir listesinde ant anahtar adına sahip veriyi elde et
	antalya := states["ANT"]
	fmt.Println("Seçili Şehir : ", antalya)

	delete(states, "ANK")
	fmt.Println(states)

	states["SKR"] = "SAKARYA"

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	//states map nesnesinin elemanı adedince kapasiteli bir key listesi oluştur
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
}
