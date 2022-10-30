package main

import "spiropoulos94/ocr-server/ocr"

func main() {

	// result, err := ocr.ExtractTextFromLocalImage("samples/IMG_8516.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result)

	// fsutils.AppendToFile("resultDyskolo.txt", *result)

	ocr.ParseDateFromText()

}
