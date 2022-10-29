package ocr

import (
	"context"
	"fmt"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"

	"spiropoulos94/ocr-server/fsutils"
)

func ExtractTextFromLocalImage(filename string) (*string, error) {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Printf("Failed to create client in order to communicate with: %v", err)
		return nil, err
	}
	defer client.Close()

	// // Sets the name of the image file to annotate.
	// filename := "samples/1copy.jpg"

	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		return nil, err
	}
	defer file.Close()

	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Printf("Failed to create image: %v", err)
		return nil, err
	}

	labels, err := client.DetectTexts(ctx, image, nil, 1)
	if err != nil {
		log.Printf("Failed to detect labels: %v", err)
		return nil, err
	}

	fmt.Println("Labels:")
	if len(labels) > 0 {
		title := "Results for " + filename + " : \n"
		fsutils.AppendToFile("results.txt", title)
	}
	for index, label := range labels {
		fmt.Println(label.Description)
		fsutils.AppendToFile("results.txt", label.Description+"\n")
		if index == len(labels)-1 {
			fsutils.AppendToFile("results.txt", "\n")
		}
	}

	return &labels[0].Description, nil

}

func ExtractTextFromRemote(imageUri string) (*string, error) {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Printf("Could not create client %v", err)
		return nil, err
	}

	image := vision.NewImageFromURI(imageUri)
	texts, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		log.Printf("Could not create new image from URI %v", err)
		return nil, err
	}

	if len(texts) == 0 {
		fmt.Println("No text found.")
		return nil, nil
	} else {
		fmt.Println("Text:")
		for _, annotation := range texts {
			fmt.Printf("%q\n", annotation.Description)
		}
	}

	return &texts[0].Description, nil
}
