package main

import "fmt"

func main() {

	var message string //veri tipi kesin olarak biliniyorsa bu şekilde yazılabilir

	var message = "Merhaba Go!" // ilk başta tip bilinmiyorsa bu şekilde girip buna reflector un karar vermesi sağlanır

	var a, b, c int = 3, 4, 5 // tek satırda çoklu değişken tanımlama

	fmt.Println(message)

	var message = "Hello World"
	var a, b, c = 3, true, 4.5
	fmt.Println(message, a, b, c)

	var c int
	var k, o string = "abc", "xyz"

	var p = 42
	var s, b = "x,y,z", true

	//en kısa degısken tanımlama yontemı
	// : kullanılarak tanımlanması gerekmektedir
	u := 55
	v, n := "abc", true

	//eğer veri tipi kesin olarak belliyse tip belirtilmesi tavsiye edilir

	a := "Metin"
	b := 'M'     //char tipinde getirir
	c := `Metin` //alt ve 96 basılarak oluşturulur

	fmt.Println(a, b, c)

	var myFloat32 float32 = 44.
	myComplex := complex(3, 4)
	Println(myFloat32)
	Println(myComplex)

	const aciklama = "deneme"

	aciklama = "sadad" //const tanımlandıktan sonra bir daha degıstırelemez

	//tek bir yerde toplu olarak değişken kullanım örneği
	var {
		degisken1 := "Serhat"
		degisken2 := "ÜR"
	}

}
