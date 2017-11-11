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

func WriteStringToFile(text string, filetype string) string {
	var filename = RandomFilename(filetype, "html")
	go writeToFile(filename, text)
	return filename
}

func writeToFile(file string, content string) {
	ioutil.WriteFile(file, []byte(content), 0644)
}

func Read(filename string) []byte {
	var file, err = ioutil.ReadFile(filename)
	if err != nil {
		log.Panic("Could not read pdf file", err.Error())
	}
	return file
}
