package ocr

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

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

	textToParse := labels[0].Description

	parsedDate, err := ParseDateFromText(textToParse)
	if err != nil {
		return nil, err
	}

	return parsedDate, nil

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

	textToParse := texts[0].Description

	parsedDate, err := ParseDateFromText(textToParse)
	if err != nil {
		return nil, err
	}

	return parsedDate, nil
}

func ParseDateFromText(stringContent string) (*string, error) {

	match, err := regexp.MatchString(MatchAllValidDates, stringContent)

	if err != nil {
		log.Printf("Could not apply match string function %v", err)
		return nil, err
	}

	if match {
		r, _ := regexp.Compile(MatchAllValidDates)
		if err != nil {
			fmt.Printf("Error while compiling regex:  %v\n", err)
			return nil, err
		}

		result := r.FindString(stringContent)
		if len(result) == 0 {
			return nil, errors.New("found matches but an empty one")
		}

		return &result, nil

	} else {
		return nil, errors.New("no matches found on given text content")
	}

}
