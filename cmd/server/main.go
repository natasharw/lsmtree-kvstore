package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/natasharw/lsmtree-kvstore/pkg/storage"
)

func main() {
	fmt.Println("Welcome to the key-value store. Initialising...")
	var store storage.Storage
	store = storage.LsmTreeInit()
	fmt.Println("Ready")
	run(store)
	defer fmt.Println("Exiting key-value store")
}

func run(store storage.Storage) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter command: ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		log.Printf("Input received: %s", cmd)
		process(cmd, store)
	}
}

func process(cmd string, store storage.Storage) {
	c := strings.Fields(cmd)
	switch c[0] {
	case "get":
		if len(c[1:]) != 1 {
			fmt.Println("Incorrect command. please use \"get <yourkey> \"")
			break
		}
		keyInt, err := strconv.Atoi(c[1])
		if err != nil {
			fmt.Println("Incorrect command. <yourkey> must be of type (int)\"")
		}
		log.Printf("Processing get request")
		store.Get(keyInt)
	case "set":
		if len(c[1:]) != 2 {
			fmt.Println("Incorrect command supplied. please use \"set <yourkey> <yourvalue>\"")
			break
		}
		log.Printf("Processing set request")
		keyInt, err := strconv.Atoi(c[1])
		valInt, err := strconv.Atoi(c[2])
		if err != nil {
			fmt.Println("Incorrect command. <yourkey> and <yourvalue> must be of type (int)\"")
		}
		store.Set(keyInt, valInt)
	case "exit":
		log.Printf("Thanks for visiting the key-value store.")
		os.Exit(0)
	default:
		fmt.Println("Incorrect command supplied. hint: see README and try again")
	}
}
