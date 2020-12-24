package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("Bu dosyaya yazıldı.\n")
	bytesWritten, err := file.Write(byteSlice)
	log.Printf("byte %d", bytesWritten)
}
