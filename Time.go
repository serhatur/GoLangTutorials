package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now()) //Şuanın tarihini alır
	//time.Sleep(2 * time.Second) //2sn bekletmeyi sağlar
	fmt.Println(time.Now())
	fmt.Println(time.Now().Unix())

	t := time.Date(2018, time.November, 10, 20, 0, 0, 0, time.UTC) //Belirli bir zamanı yazdırmak için
	fmt.Println(t)
	fmt.Printf("%s\n", t)

	fmt.Println("*************************************")

	now := time.Now()
	fmt.Printf("Mevcut zaman : %s\n", now)

	fmt.Println("*************************************")

	fmt.Println("Ay : ", now.Month())
	fmt.Println("Gün : ", now.Day())
	fmt.Println("Haftanın Günü : ", now.Weekday())

	fmt.Println("*************************************")

	tomorrow := now.AddDate(0, 0, 1) //tarihe zaman ekleme
	fmt.Printf("Tomorrow is %v,%v,%v,%v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	fmt.Println("*************************************")

	longFormat := "Monday,January 2,2006"
	fmt.Println("Tomorrow is ", tomorrow.Format(longFormat))
	fmt.Println("*************************************")

	shortFormat := "02/01/2006"
	fmt.Println("Tomorrow is ", tomorrow.Format(shortFormat))
	fmt.Println("*************************************")

	xTime := time.Date(1071, 10, 11, 20, 23, 0, 0, time.UTC)
	x := fmt.Println

	moment := time.Now()
	x(moment)

	x("*************************************")

	x(now.Year())
	x(now.Month())
	x(now.Day())
	x(now.Hour())
	x(now.Minute())
	x(now.Second())
	x(now.Nanosecond())
	x(now.Location())
	x(now.Weekday())

	//Tarih karşılaştırma
	x(xTime.Before(now))
	x(xTime.After(now))
	x(xTime.Equal(now))

	diff := now.Sub(xTime)
	x(diff)

	x(diff.Hours())
	x(diff.Minutes())
	x(diff.Seconds())
	x(diff.Nanoseconds())
}
