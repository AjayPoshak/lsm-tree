package lsm

import (
	"log"
	"lsm-tree/files"
	"strings"
)

type LSM struct {
	filePath               string
	keysCountInCurrentFile int
}

func New() *LSM {
	lsmRef := &LSM{
		filePath: "data/data.txt",
	}
	lsmRef.countKeysInCurrentFile()
	return lsmRef
}

func (lsm *LSM) Get(key string) string {
	filePtr := files.NewFile(lsm.filePath)
	if filePtr == nil {
		return ""
	}
	NextLine, err := filePtr.ReadFileByLine(key)
	if err != nil {
		log.Fatal(err)
	}
	value := ""
	line := NextLine()
	for line != "" {
		// A reasonable assumption here is about delimiter. Usually delimiters are something which won't occur in key or value so that key and value can be separated in any given string.
		// But there are a lot of edge cases for example what if key name or value contains the delimiter, it will break the parsing.  To fix this, we are wrapping
		// key and value in double quotes, and ignore everything inside the quotes while parsing.
		pair := strings.Split(line, "\"=") // Looking for "= to separate out key and value
		// After parsing the given string which looks like this ["key, "value"], we need to remove double quote at the beginning for the key, & remove quotes from beginning and end for the value.
		strippedKey := pair[0][1:]
		strippedValue := pair[1][1 : len(pair[1])-1]
		if strings.TrimSpace(strippedKey) == key {
			value = strippedValue
		}
		line = NextLine()
	}

	if value == "" {
		log.Println("Key not found")
	}
	return value
}

func (lsm *LSM) Set(key, value string) {
	f := files.NewFile(lsm.filePath)
	log.Println(key, value)
	appendNewLine := value
	f.AppendToFile(key, appendNewLine)
	log.Println("Successfully written")
}

func (lsm *LSM) countKeysInCurrentFile() {
	file := files.NewFile(lsm.filePath)
	keyCount := file.CountKeys()
	lsm.keysCountInCurrentFile = keyCount
}
