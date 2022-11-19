package main

import (
	"fmt"
	"log"
	"os"
)

func repeatAudio() {
	bytes, err := os.ReadFile("C:/Users/FIJUSAU/go/src/langExGen/audio/space.mp3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(bytes))
	bytes = append(bytes, bytes...)
	fmt.Println(len(bytes))

	path := "C:/Users/FIJUSAU/go/src/langExGen/audio/"
	filePath := path + "test.mp3"
	if err := os.WriteFile(filePath, bytes, 0644); err != nil {
		log.Fatal(err)
	}
}
