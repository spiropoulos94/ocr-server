package main

import (
	"fmt"
	"log"
	"spiropoulos94/ocr-server/ocr"
)

func main() {

	result, err := ocr.ExtractTextFromLocalImage("samples/IMG_8516.jpg")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("------ RESULT ------")
	fmt.Println(*result)

	// fsutils.AppendToFile("resultDyskolo.txt", *result)

}
