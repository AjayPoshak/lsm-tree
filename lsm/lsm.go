package lsm

import (
	"log"
	"lsm-tree/files"
)

type LSM struct {
	filePath string
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
	key, value := filePtr.ReadFileByLine(key)
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
