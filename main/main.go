package main

import (
	"bufio"
	"fmt"
	"log"
	"lsm-tree/lsm"
	"os"
	"strings"
)

func processCommand(input string) {
	command := strings.Split(input, " ")[0]
	if len(command) < 2 {
		log.Println("Please enter a command: get or set")
		return
	}
	key := strings.Split(input, " ")[1]
	if command == "get" {
		lsm := lsm.New()
		value := lsm.Get(key)
		log.Println(value)
	} else if command == "set" {
		value := strings.Split(input, " ")[2]
		lsm := lsm.New()
		lsm.Set(key, value)
	} else {
		log.Println("Invalid command")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter command: get or set")
	fmt.Println("Enter quit to exit")
	for {
		fmt.Print("> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading command")
			return
		}

		input := strings.TrimSpace(command)
		if input == "quit" {
			return
		}

		processCommand(input)
	}
}
