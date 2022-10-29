package fsutils

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		return false
	} else {
		return true
	}
}

func AppendToFile(filename string, stringToAppend string) {
	var file *os.File

	if !FileExists(filename) {
		_, err := os.Create(filename)

		if err != nil {
			log.Fatalf("Error while creating file: %v", err)
		}

	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()
	if err != nil {
		log.Fatalf("Failed to create client in order to communicate with gcloud: %v", err)
	}

	bytes, err := file.WriteString(stringToAppend)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
	fmt.Printf("wrote %d bytes into %s \n", bytes, filename)
}
