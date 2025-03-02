package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type File struct {
	path string
}

func NewFile(path string) *File {
	// Check if file exists, if not then create
	_, error := os.Stat(path)
	if os.IsNotExist(error) {
		log.Printf("File does not exist, creating file %v", error)
		err := os.MkdirAll(strings.Split(path, "/")[0], 0755)
		if err != nil {
			log.Printf("Error creating directory %v", err)
			return nil
		}
		file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Printf("Error creating file %v", err)
			return nil
		}
		defer file.Close()
	}
	return &File{
		path: path,
	}
}

type NextLine func() string

func (file *File) ReadFileByLine(key string) (NextLine, error) {
	filePtr, err := os.Open(file.path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	scanner := bufio.NewScanner(filePtr)
	return func() string {
		if scanner.Scan() {
			return scanner.Text()
		}
		filePtr.Close()
		return ""
	}, nil
}

func (file *File) CountKeys() int {
	filePtr, err := os.Open(file.path)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	defer filePtr.Close()
	scanner := bufio.NewScanner(filePtr)
	count := 0
	for scanner.Scan() {
		scanner.Text()
		count++
	}
	return count
}

func (file *File) AppendToFile(key, value string) {
	filePtr, err := os.OpenFile(file.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer filePtr.Close()
	textToBeWritten := fmt.Sprintf("\"%s\"=\"%s\"\n", key, value)
	_, err = filePtr.WriteString(textToBeWritten)
	if err != nil {
		log.Fatalf("Error writing to file %v", err)
		return
	}

}
