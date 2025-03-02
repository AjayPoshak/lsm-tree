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

func (file *File) ReadFileByLine(key string) (string, string) {
	filePtr, err := os.Open(file.path)
	if err != nil {
		log.Fatal(err)
		return "", ""
	}
	defer filePtr.Close()
	// Instead of reading entire file, read it line by line so we don't need to read entire file into memory at once
	scanner := bufio.NewScanner(filePtr)
  value := ""
	for scanner.Scan() {
		line := scanner.Text()
    // A reasonable assumption here is about delimiter. Usually delimiters are something which won't occur in key or value so that key and value can be separated in any given string.
    // But there are a lot of edge cases for example what if key name or value contains the delimiter, it will break the parsing.  To fix this, we are wrapping
    // key and value in double quotes, and ignore everything inside the quotes while parsing.
		pair := strings.Split(line, "\"=") // Looking for "= to separate out key and value
    // After parsing the given string which looks like this ["key, "value"], we need to remove double quote at the beginning for the key, & remove quotes from beginning and end for the value.
    strippedKey := pair[0][1:]
    strippedValue := pair[1][1:len(pair[1])-1]
		if strings.TrimSpace(strippedKey) == key {
      value = strippedValue
		}
	}

	er := scanner.Err()
	if er != nil {
		log.Fatalf("Error in reading lines %v", er)
	return "", ""
	}
  return key, value
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
