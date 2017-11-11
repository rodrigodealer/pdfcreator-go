package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

func CleanUp(files []string) {
	for _, element := range files {
		log.Printf("Deleted %s", element)
		os.Remove(element)
	}
}

func RandomFilename(filetype string, extension string) string {
	return fmt.Sprintf("%s-%d.%s", filetype, rand.Int(), extension)
}

func WriteStringToFile(text string, filetype string) (string, error) {
	var filename = RandomFilename(filetype, "html")
	err := ioutil.WriteFile(filename, []byte(text), 0644)
	return filename, err
}
